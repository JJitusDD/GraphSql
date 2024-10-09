package execute_workflow

import (
	"auto_reconcile_service_v2/internal/domain/service"

	"github.com/labstack/echo/v4"
)

func NewRerunWorkflow(g *echo.Group, s *service.Service) {
	g.POST("/:workflow_id/reconciliation-process", s.RerunReconciliationProcess)

	g.POST("/:workflow_id/bank-statement", s.RerunBankStatement)

	g.POST("/:workflow_id/cancel-workflow", s.CancelWorkflowByService)
}
