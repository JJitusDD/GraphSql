package export_file

import (
	"auto_reconcile_service_v2/internal/domain/service"
	"github.com/labstack/echo/v4"
)

func NewExportFileRouter(g *echo.Group, s *service.Service) {
	g.POST("/export-file-va", s.UploadFileVAF88)
	g.POST("/export-file-settlement", s.UploadFileSettlementF88)
}
