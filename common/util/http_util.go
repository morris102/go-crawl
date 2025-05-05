package util

import (
	"io"
	"net/http"
)

type HttpClient struct{}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (inst *HttpClient) Do(url, method string, body io.Reader) (*http.Response, error) {
	if method == "" {
		method = "GET"
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return &http.Response{StatusCode: http.StatusBadRequest}, err
	}

	httpClient := http.Client{
		Transport: &http.Transport{},
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, err
}
