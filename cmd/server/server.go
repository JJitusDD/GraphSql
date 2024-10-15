package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"project-test/configs"
	"project-test/graph"
	"project-test/internal/domain/facade"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const defaultPort = "8080"

type App struct {
	Config *configs.Config
	Echo   *echo.Echo
	Facade *facade.AutoReconcileFacade
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e := echo.New()
	confg, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	f := facade.NewAutoReconcileFacade(confg)
	e.Logger.SetLevel(log.INFO)

	if confg.ENV != "PRODUCTION" {
		f.Logger.WithFields(logrus.Fields{
			"config": confg,
		}).Info("Logging config")
	}

	// setup Validator
	validate := validator.New()
	e.Validator = &CustomValidator{validator: validate}

	// error handler
	e.HTTPErrorHandler = CustomHTTPErrorHandler(f.Logger)
	// Recover middleware
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
	e.POST("/query", func(c echo.Context) error {
		// Handle GraphQL request
		srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	app := &App{
		Config: confg,
		Echo:   e,
		Facade: f,
	}

	go func() {
		if err := app.Echo.Start(":" + app.Config.Port); err != nil {
			app.Echo.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

func CustomHTTPErrorHandler(l *logrus.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			l.WithError(err).WithFields(logrus.Fields{
				"trace_id": c.Get("trace_id"),
			}).Error("echo http error")

			res := configs.Response{
				Meta: configs.Meta{
					Code:    he.Code,
					Message: he.Message,
				},
				Data: nil,
			}
			if he.Internal != nil {
				res.Meta.Msg = he.Internal.Error()
			}
			return
		} else if validationErrs, ok := err.(validator.ValidationErrors); ok {
			var messages error
			for _, validationErr := range validationErrs {
				// custom handling of specific validation errors
				switch validationErr.Tag() {
				case "required":
					messages = fmt.Errorf("%s is required", validationErr.Field())
				case "email":
					messages = fmt.Errorf("%s must be a valid email", validationErr.Field())
				case "gte":
					messages = fmt.Errorf("%s must be greater than or equal to %s", validationErr.Field(), validationErr.Param())
				case "lte":
					messages = fmt.Errorf("%s must be less than or equal to %s", validationErr.Field(), validationErr.Param())
				default:
					messages = fmt.Errorf("some thing went wrong!")
				}
			}

			l.WithError(err).WithFields(logrus.Fields{
				"trace_id": c.Get("trace_id"),
			}).Error("validator error")

			// Return the validation error messages as a JSON response_definition.Response
			c.JSON(http.StatusBadRequest, configs.Response{
				Meta: configs.Meta{
					Code:    http.StatusBadRequest,
					Message: nil,
					Msg:     messages.Error(),
				},
				Data: nil,
			})
			return
		}
		// If it's not a validation error, return the default error response
		c.Echo().DefaultHTTPErrorHandler(err, c)

	}
}
