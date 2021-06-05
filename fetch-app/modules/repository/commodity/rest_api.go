package commodity

import (
	"efishery/business/commodity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HTTPRepository struct {
}

// A Response struct to map the Entire Response
type Response struct {
	UUID      string `json:"uuid"`
	Commodity string `json:"komoditas"`
	Province  string `json:"area_provinsi"`
	City      string `json:"area_kota"`
	Size      string `json:"size"`
	Price     string `json:"price"`
	ParsedAt  string `json:"tgl_parsed"`
	Timestamp string `json:"timestamp"`
}

// A response struct to map price convert response
type PriceConvertResponse struct {
	PriceUSD float64 `json:"USD_IDR"`
}

//NewRESTAPIRepository fetch data from external datasource
func NewRESTAPIRepository() *HTTPRepository {
	return &HTTPRepository{}
}

//FetchCommodities Find commodity based on given ID. Its return nil if not found
func (repo *HTTPRepository) FetchCommodities() ([]commodity.Commodity, error) {
	var commodities []commodity.Commodity
	var commodity commodity.Commodity
	response, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list/")

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject []Response
	json.Unmarshal(responseData, &responseObject)

	for _, val := range responseObject {
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

	return commodities, nil
}

//FetchPriceConverter Find commodity based on given ID. Its return nil if not found
func (repo *HTTPRepository) FetchPriceConverter() (float64, error) {
	response, err := http.Get("https://free.currconv.com/api/v7/convert?q=USD_IDR&compact=ultra&apiKey=7b69cf393e72fc95d3ab")

	fmt.Println(response, err)
	if err != nil {
		return 0, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	fmt.Println(string(responseData))
	var responseObject PriceConvertResponse
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.PriceUSD)

	return responseObject.PriceUSD, nil
}
