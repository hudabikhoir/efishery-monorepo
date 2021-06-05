package commodity

import (
	"fmt"
	"strconv"

	validator "github.com/go-playground/validator/v10"
)

//Repository ingoing port for commodity
type Repository interface {
	//FetchCommodities is a function to fetch commodities from resource data
	FetchCommodities() ([]Commodity, error)

	FetchPriceConverter() (float64, error)
}

//Service outgoing port for commodity
type Service interface {
	GetCommodities() ([]Commodity, error)
}

//=============== The implementation of those interface put below =======================

type service struct {
	repository Repository
	validate   *validator.Validate
}

//NewService Construct commodity service object
func NewService(repository Repository) Service {
	return &service{
		repository,
		validator.New(),
	}
}

//GetCommodities Get all commodities by given tag, return zero array if not match
func (s *service) GetCommodities() ([]Commodity, error) {
	var commodities []Commodity
	var commodity Commodity
	var err error

	commoditiesRes, err := s.repository.FetchCommodities()
	if err != nil {
		return nil, err
	}

	convertPrice, err := s.repository.FetchPriceConverter()
	if err != nil {
		return nil, err
	}

	for _, val := range commoditiesRes {
		price, _ := strconv.ParseFloat(val.Price, 64)
		commodity.ConvertPrice = fmt.Sprintf("%.2f", (price / convertPrice))
		commodity.UUID = val.UUID
		commodity.Commodity = val.Commodity
		commodity.Province = val.Province
		commodity.City = val.City
		commodity.Size = val.Size
		commodity.Price = val.Price
		commodity.ParsedAt = val.ParsedAt
		commodity.Timestamp = val.Timestamp

		commodities = append(commodities, commodity)
	}
	return commodities, err
}
