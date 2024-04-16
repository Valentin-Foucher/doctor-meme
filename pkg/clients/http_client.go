package clients

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Valentin-Foucher/doctor-meme/pkg/utils"
)

type HttpClient struct {
	baseUrl    string
	maxRetries int
	c          http.Client
}

func getRequest(url, method string, body io.Reader, headers map[string]string, queryParameters map[string][]string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	if _, found := headers["Content-Type"]; !found {
		headers["Content-Type"] = "application/json"
	}

	for k, v := range headers {
		request.Header.Add(k, v)
	}
	for k, v := range queryParameters {
		for _, e := range v {
			request.Form.Add(k, e)
		}
	}
	return request, nil
}

func (client *HttpClient) request(method, path string, statusCodes []int, body io.Reader, headers map[string]string, queryParameters map[string][]string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", client.baseUrl, path)
	request, err := getRequest(url, method, body, headers, queryParameters)

	if err != nil {
		return nil, err
	}

	tries_count := 0
	var response *http.Response
	for {
		response, err := client.c.Do(request)

		if !utils.SliceContains(statusCodes, response.StatusCode) {
			return nil, fmt.Errorf("unexpected status code (%d): %s", response.StatusCode, response.Body)
		}
		if err == nil {
			break
		}
		if tries_count > client.maxRetries {
			return nil, err
		}
		tries_count += 1
	}

	return response, nil
}

func (client *HttpClient) Get(path string, statusCodes []int, headers map[string]string, queryParameters map[string][]string) (*http.Response, error) {
	return client.request("GET", path, statusCodes, nil, headers, queryParameters)
}

func (client *HttpClient) Post(path string, statusCodes []int, body io.Reader, headers map[string]string, queryParameters map[string][]string) (*http.Response, error) {
	return client.request("POST", path, statusCodes, body, headers, queryParameters)
}

func (client *HttpClient) Put(path string, statusCodes []int, body io.Reader, headers map[string]string, queryParameters map[string][]string) (*http.Response, error) {
	return client.request("PUT", path, statusCodes, body, headers, queryParameters)
}

func (client *HttpClient) Delete(path string, statusCodes []int, headers map[string]string, queryParameters map[string][]string) (*http.Response, error) {
	return client.request("DELETE", path, statusCodes, nil, headers, queryParameters)
}
