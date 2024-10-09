package routers

import (
	"auto_reconcile_service_v2/internal/app/routers/wallet_gateway"
	"net/http"

	"auto_reconcile_service_v2/internal/app/routers/bank_statement"
	"auto_reconcile_service_v2/internal/app/routers/execute_workflow"
	"auto_reconcile_service_v2/internal/app/routers/export_file"
	"auto_reconcile_service_v2/internal/app/routers/reconcile"
	"auto_reconcile_service_v2/internal/app/routers/statistic"
	"auto_reconcile_service_v2/internal/domain/service"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo, s *service.Service) {
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	v2 := e.Group("v2")

	statistic.NewStatisticRouter(v2, s)
	execute_workflow.NewRerunWorkflow(v2, s)
	bank_statement.NewBankStatementRouter(v2, s)
	reconcile.NewReconcileRouter(v2, s)
	export_file.NewExportFileRouter(v2, s)
	wallet_gateway.NewWalletGatewayRouter(v2, s)
}
