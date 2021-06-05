package modules

import (
	"efishery/api"
	"efishery/api/common"
	"efishery/util"

	commodityCtrlV1 "efishery/api/v1/commodity"
	commodityBussiness "efishery/business/commodity"
	commodityRepo "efishery/modules/repository/commodity"

	echo "github.com/labstack/echo"
)

//SetErrorHandler - set error response
func SetErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// error message must be known RC value
		errResp := common.NewInternalServerErrorResponse()
		c.JSON(500, errResp)
	}
}

//RegisterController - register the controller
func RegisterController(dbCon *util.DatabaseConnection) api.Controller {

	//initiate commodity
	commodityPermitRepo := commodityRepo.RepositoryFactory(dbCon)
	commodityPermitService := commodityBussiness.NewService(commodityPermitRepo)
	commodityPermitControllerV1 := commodityCtrlV1.NewController(commodityPermitService)

	//lets put the controller together
	controllers := api.Controller{
		CommodityController: commodityPermitControllerV1,
	}

	return controllers
}
