package auth

import (
	validator "github.com/go-playground/validator/v10"
)

//Repository ingoing port for auth
type Repository interface {
	//FetchMeByToken is a function to get auth by token param
	FetchMeByToken(Token string) (*User, error)
}

//Service outgoing port for auth
type Service interface {
	Validate(Param string) (*User, error)
}

//=============== The implementation of those interface put below =======================

type service struct {
	repository Repository
	validate   *validator.Validate
}

//NewService Construct auth service object
func NewService(repository Repository) Service {
	return &service{
		repository,
		validator.New(),
	}
}

//GetCommodities Get all commodities by given tag, return zero array if not match
func (s *service) Validate(Param string) (*User, error) {
	authRes, err := s.repository.FetchMeByToken(Param)
	if err != nil {
		return nil, err
	}

	return authRes, err
}
