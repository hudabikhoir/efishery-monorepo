package api

import (
	"efishery/api/v1/commodity"

	"github.com/labstack/echo"
)

// Controller to define controller that we use
type Controller struct {
	CommodityController *commodity.Controller
}

//RegisterPath Registera V1 API path
func RegisterPath(e *echo.Echo, ctrl Controller) {
	//commodity
	commodityV1 := e.Group("/v1/commodities")
	commodityV1.GET("", ctrl.CommodityController.GetAllCommodities)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
