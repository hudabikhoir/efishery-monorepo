package response

import (
	"efishery/business/commodity"
)

type GetCommodityReportResponse struct {
	Province string  `json:"province"`
	Min      float64 `json:"min"`
	Max      float64 `json:"max"`
	Median   int     `json:"median"`
	Average  float64 `json:"average"`
}

//GetCommoditieResponse Get commoditie by tag response payload
type GetCommoditiesReportResponse struct {
	Code    string                        `json:"code"`
	Message string                        `json:"message"`
	Data    []*GetCommodityReportResponse `json:"data"`
}

//NewGetCommodityReportResponse construct GetCommodityReportResponse
func NewGetCommodityReportResponse(commodity commodity.CommodityReport) *GetCommodityReportResponse {
	var commodityResponse GetCommodityReportResponse
	commodityResponse.Province = commodity.Province
	commodityResponse.Min = commodity.Min
	commodityResponse.Max = commodity.Max
	commodityResponse.Median = commodity.Median
	commodityResponse.Average = commodity.Average

	return &commodityResponse
}

//NewGetCommoditieResponse construct GetCommoditieResponse
func NewGetCommoditiesReportResponse(commodities []commodity.CommodityReport) *GetCommoditiesReportResponse {
	var commoditiesResponses []*GetCommodityReportResponse
	commoditiesResponses = make([]*GetCommodityReportResponse, 0)

	for _, commodities := range commodities {
		commoditiesResponses = append(commoditiesResponses, NewGetCommodityReportResponse(commodities))
	}

	return &GetCommoditiesReportResponse{
		"00",
		"Success",
		commoditiesResponses,
	}
}
