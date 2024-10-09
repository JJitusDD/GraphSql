package statistic

import (
	"auto_reconcile_service_v2/internal/domain/service"

	"github.com/labstack/echo/v4"
)

func NewStatisticRouter(g *echo.Group, s *service.Service) {
	g = g.Group("/statistic")
	g.POST("/statement-reconciliation-report", s.GetStatementReconciliationReport)
	g.GET("/download-file-report/:file_path", s.DownloadStatementReconciliationReport)
	g.POST("/download-file-report", s.DownloadStatementReconciliationReportV1)
}
