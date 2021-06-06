package commodity

import (
	"bufio"
	"bytes"
	"efishery/business/commodity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
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
	req, err := http.NewRequest("GET", "https://free.currconv.com/api/v7/convert/", nil)

	if err != nil {
		return 0, err
	}

	req.AddCookie(&http.Cookie{Name: "c", Value: "ccc"})
	q := req.URL.Query()
	q.Add("apiKey", "7b69cf393e72fc95d3ab")
	q.Add("compact", "ultra")
	q.Add("q", "USD_IDR")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	// you can set turn on/off cache body with set DumpResponse parameter to true or false
	body, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return 0, err
	}

	// wrap the cached response
	r := bufio.NewReader(bytes.NewReader(body))

	// ReadResponse by default assumes the request for the response was a "GET" requested
	// If you want the method to be different, you must pass an http.Request to ReadResponse (instead of nil)
	resp, err = http.ReadResponse(r, nil)
	if err != nil {
		return 0, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var responseObject PriceConvertResponse
	json.Unmarshal(responseData, &responseObject)
	return responseObject.PriceUSD, nil
}
