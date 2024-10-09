package wallet_gateway

import (
	"auto_reconcile_service_v2/internal/domain/service"
	"github.com/labstack/echo/v4"
)

func NewWalletGatewayRouter(g *echo.Group, s *service.Service) {
	g = g.Group("/wallet-gateway")
	g.POST("/unlink", s.UnlinkBankWalletGateway)
}
