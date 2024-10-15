package facade

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"project-test/configs"
	"project-test/internal/domain/repository"
	"project-test/internal/infrastructure/persistence"
	"project-test/pkg/agent/mongodb"
	"project-test/pkg/agent/postgres"
	"project-test/pkg/agent/redis"
	"project-test/pkg/loggers"
)

type Facade struct {
	Logger *logrus.Logger
	Config *configs.Config
	Redis  repository.IRedis
	Mongo  *repository.MongoRepoImpl
	User   repository.IUser
}

func NewFacade(config *configs.Config) *Facade {
	log := loggers.NewLogger()

	pg := postgres.NewPostgres(config.Postgress)
	dbMongo := mongodb.NewMongoDbConnection(config.MongoURI)

	redis, err := redis.NewRedis(config, log)
	if err != nil {
		fmt.Print("NewRedis err: %s", err.Error())
	}
	f := &Facade{
		Logger: log,
		Config: config,
		Redis:  persistence.NewRedis(redis),
		Mongo:  repository.NewMongoDB(dbMongo),
		User:   persistence.NewUser(pg, log),
	}

	return f
}
