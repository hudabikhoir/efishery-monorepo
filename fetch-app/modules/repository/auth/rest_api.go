package auth

import (
	"bufio"
	"bytes"
	"efishery/business/auth"
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
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    DetailResponse `json:"data"`
}
type DetailResponse struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Role  int    `json:"role"`
}

//NewRESTAPIRepository fetch data from external datasource
func NewRESTAPIRepository() *HTTPRepository {
	return &HTTPRepository{}
}

//FetchPriceConverter Find auth based on given ID. Its return nil if not found
func (repo *HTTPRepository) FetchMeByToken(Token string) (*auth.User, error) {
	var user *auth.User
	var bearer = "Bearer " + Token
	req, err := http.NewRequest("GET", "http://127.0.0.1:8089/me", nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "c", Value: "ccc"})
	fmt.Println(req.URL.String())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// you can set turn on/off cache body with set DumpResponse parameter to true or false
	body, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}

	// wrap the cached response
	r := bufio.NewReader(bytes.NewReader(body))

	// ReadResponse by default assumes the request for the response was a "GET" requested
	// If you want the method to be different, you must pass an http.Request to ReadResponse (instead of nil)
	resp, err = http.ReadResponse(r, nil)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("responseData", string(responseData))
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	user = &auth.User{
		Role:  responseObject.Data.Role,
		Phone: responseObject.Data.Phone,
		Name:  responseObject.Data.Name,
	}

	return user, nil
}
