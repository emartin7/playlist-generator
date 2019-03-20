package clients

import (
	"log"
	"net/http"
	"playlist-generator/models"
)

func Get(request models.HttpRequest) (*http.Response, error) {
	requestType := "GET"
	return execute(requestType, request)
}

func Post(request models.HttpRequest) (*http.Response, error) {
	requestType := "POST"
	return execute(requestType, request)
}

func execute(requestType string, request models.HttpRequest) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(requestType, request.Path, request.Body)
	setDefaultHeaders(req)
	setCustomHeaders(request.Headers, req)
	SetQueryParams(request.QueryParams, req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return client.Do(req)
}

func setDefaultHeaders(request *http.Request) {
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
}

func setCustomHeaders(headerMap map[string]string, request *http.Request) {
	for key, value := range headerMap {
		request.Header.Set(key, value)
	}
}

func SetQueryParams(queryParamMap map[string]string, request *http.Request) {
	if queryParamMap != nil {
		for key, value := range queryParamMap {
			if value != "" && value != "0" {
				q := request.URL.Query()
				q.Add(key, value)
				request.URL.RawQuery = q.Encode()
			}
		}
	}
}

func GetFullRequest(requestType string, request models.HttpRequest) (*http.Request, error) {
	req, err := http.NewRequest(requestType, request.Path, request.Body)
	setDefaultHeaders(req)
	setCustomHeaders(request.Headers, req)
	SetQueryParams(request.QueryParams, req)
	return req, err
}
