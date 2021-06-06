package auth

import (
	"efishery/business/auth"
)

//RepositoryFactory Will return business.auth.Repository based on resource rest api
func RepositoryFactory() auth.Repository {
	authRepo := NewRESTAPIRepository()

	return authRepo
}
