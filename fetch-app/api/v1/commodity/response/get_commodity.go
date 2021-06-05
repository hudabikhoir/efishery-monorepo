package response

import (
	"efishery/business/commodity"
)

//GetCommodityByIDResponse Get commodity by ID response payload
type GetCommodityByIDResponse struct {
	UUID      string `json:"uuid"`
	Commodity string `json:"commodity"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Size      string `json:"size"`
	IDR       string `json:"idr"`
	USD       string `json:"usd"`
	ParsedAt  string `json:"parsed_at"`
	Timestamp string `json:"timestamp"`
}

//GetCommoditieResponse Get commoditie by tag response payload
type GetCommoditieResponse struct {
	Commodities []*GetCommodityByIDResponse `json:"commodities"`
}

//NewGetCommodityByIDResponse construct GetCommodityByIDResponse
func NewGetCommodityByIDResponse(commodity commodity.Commodity) *GetCommodityByIDResponse {
	var commodityResponse GetCommodityByIDResponse
	commodityResponse.UUID = commodity.UUID
	commodityResponse.Commodity = commodity.Commodity
	commodityResponse.Province = commodity.Province
	commodityResponse.City = commodity.City
	commodityResponse.Size = commodity.Size
	commodityResponse.IDR = commodity.Price
	commodityResponse.USD = commodity.ConvertPrice
	commodityResponse.ParsedAt = commodity.ParsedAt
	commodityResponse.Timestamp = commodity.Timestamp

	return &commodityResponse
}

//NewGetCommoditieResponse construct GetCommoditieResponse
func NewGetCommoditieResponse(commodities []commodity.Commodity) *GetCommoditieResponse {
	var commoditiesResponses []*GetCommodityByIDResponse
	commoditiesResponses = make([]*GetCommodityByIDResponse, 0)

	for _, commodities := range commodities {
		commoditiesResponses = append(commoditiesResponses, NewGetCommodityByIDResponse(commodities))
	}

	return &GetCommoditieResponse{
		commoditiesResponses,
	}
}
