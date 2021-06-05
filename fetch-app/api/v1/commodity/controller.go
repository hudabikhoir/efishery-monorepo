package commodity

import (
	"efishery/api/common"
	"efishery/api/v1/commodity/response"
	commodityBusiness "efishery/business/commodity"
	"net/http"

	v10 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

//Controller Get commodity API controller
type Controller struct {
	service   commodityBusiness.Service
	validator *v10.Validate
}

//NewController Construct commodity API controller
func NewController(service commodityBusiness.Service) *Controller {
	return &Controller{
		service,
		v10.New(),
	}
}

//GetAllCommodities Get commodity by ID echo handler
func (controller *Controller) GetAllCommodities(c echo.Context) error {
	commoditys, err := controller.service.GetCommodities()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewGetCommoditieResponse(commoditys)
	return c.JSON(http.StatusOK, response)
}

//GetReportCommodities Get commodity by ID echo handler
func (controller *Controller) GetReportCommodities(c echo.Context) error {
	commodities, err := controller.service.GetReportCommodities()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewGetCommoditiesReportResponse(commodities)
	return c.JSON(http.StatusOK, response)
}
