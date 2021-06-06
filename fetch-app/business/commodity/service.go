package commodity

import (
	"fmt"
	"strconv"

	golinq "github.com/ahmetb/go-linq/v3"

	validator "github.com/go-playground/validator/v10"
)

//Repository ingoing port for commodity
type Repository interface {
	//FetchCommodities is a function to fetch commodities from resource data
	FetchCommodities() ([]Commodity, error)

	//FetchPriceConverter is a function to fetch price converter from resource data
	FetchPriceConverter() (float64, error)
}

//Service outgoing port for commodity
type Service interface {
	GetCommodities() ([]Commodity, error)
	GetReportCommodities() ([]CommodityReport, error)
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
		// only showing used data
		if val.UUID != "" {
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
	}
	return commodities, err
}

func (s *service) GetReportCommodities() ([]CommodityReport, error) {
	commoditiesRes, err := s.repository.FetchCommodities()
	var reports []CommodityReport

	// Use golinq to create query from struct
	var query []golinq.Group
	golinq.From(commoditiesRes).GroupByT(
		func(p Commodity) string { return p.Province },
		func(p Commodity) string { return p.Size },
	).OrderByT(
		func(g golinq.Group) string { return g.Key.(string) },
	).ToSlice(&query)

	// parse query to report data
	reports = parseGroupToReport(query)

	return reports, err
}

// parseGroupToReport function to parse query group to report data
func parseGroupToReport(query []golinq.Group) []CommodityReport {
	var reports []CommodityReport
	var report CommodityReport

	for _, comGroup := range query {
		// sorted size
		size := parseListGroupToFloat(comGroup.Group)

		max := golinq.From(size).Max()
		maxInt, _ := max.(float64)

		min := golinq.From(size).Min()
		minInt, _ := min.(float64)

		average := golinq.From(size).Average()

		median := len(size) / 2

		report.Province = fmt.Sprintf("%v", comGroup.Key)
		report.Max = maxInt
		report.Min = minInt
		report.Average = average
		report.Median = median

		reports = append(reports, report)
	}

	return reports
}

// parseListGroupToFloat function to parse interface to float and then sort it!
func parseListGroupToFloat(listGroup []interface{}) []float64 {
	var szFloat []float64
	var sizeSorted []float64
	for _, val := range listGroup {
		sizeStr := val.(string)
		sizeFloat, _ := strconv.ParseFloat(sizeStr, 64)
		szFloat = append(szFloat, sizeFloat)
	}
	golinq.From(szFloat).
		Sort(
			func(i interface{}, j interface{}) bool { return i.(float64) < j.(float64) },
		).
		ToSlice(&sizeSorted)

	return sizeSorted
}
