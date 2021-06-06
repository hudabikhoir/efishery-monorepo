package commodity

import (
	"efishery/business/commodity"
)

//RepositoryFactory Will return business.commodity.Repository based on resource rest api
func RepositoryFactory() commodity.Repository {
	commodityRepo := NewRESTAPIRepository()

	return commodityRepo
}
