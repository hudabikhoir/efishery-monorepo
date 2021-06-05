package commodity

import (
	"efishery/business/commodity"
	"efishery/util"
)

//RepositoryFactory Will return business.commodity.Repository based on resource rest api
func RepositoryFactory(dbCon *util.DatabaseConnection) commodity.Repository {
	commodityRepo := NewRESTAPIRepository()

	return commodityRepo
}
