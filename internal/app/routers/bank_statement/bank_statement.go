package bank_statement

import (
	"auto_reconcile_service_v2/internal/domain/service"

	"github.com/labstack/echo/v4"
)

func NewBankStatementRouter(g *echo.Group, s *service.Service) {
	g = g.Group("/bank-statement")

	g.POST("/update-bank-statement-status", s.UpdateBankStatementStatus)
	g.POST("/update-bank-statement-data", s.UpdateBankStatementData)
	g.POST("/update-bank-statement", s.UpdateBankStatement)
}
