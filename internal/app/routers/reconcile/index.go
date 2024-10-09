package reconcile

import (
	"auto_reconcile_service_v2/internal/domain/service"

	"github.com/labstack/echo/v4"
)

func NewReconcileRouter(g *echo.Group, s *service.Service) {
	g.POST("/update-reconcile-status", s.UpdateReconcileStatus)
	g.GET("/job-update-reconcile-wallet", s.JobUpdateReconcileStatusWallet)
	g.POST("/decode-napas-file", s.DecodeNapasFile) // only use this endpoint in sandbox
	g.POST("/reconcile-bankstatement-orders", s.ReconcileBankStatementOrders)
}
