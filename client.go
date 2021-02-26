package update_api_1c

import (
	"encoding/json"
	"github.com/monaco-io/request"
	"github.com/monaco-io/request/response"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"

	"github.com/khorevaa/logos"
)

var log = logos.New("github.com/v8platform/update-api-1c")

const (
	baseURL   = "https://update-api.1c.ru"
	userAgent = "1C+Enterprise/8.3"
)

type Client struct {
	BaseURL string

	Username string
	Password string
}

func NewClient(username, password string) *Client {

	return &Client{
		baseURL,
		username,
		password,
	}

}

type apiRequest struct {
	path   string
	method string
	data   interface{}
}

func (c *Client) doFileRequest(fileRequestUrl string) (io.ReadCloser, error) {

	req, err := http.NewRequest("GET", fileRequestUrl, nil)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Content-Type", "application/json")

	cj, _ := cookiejar.New(nil)
	httpClient := &http.Client{
		Jar: cj,
	}

	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusBadRequest, http.StatusNotFound:

		var err RequestError

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr.Error())
			return nil, readErr
		}

		jsonErr := json.Unmarshal(body, &err)
		if jsonErr != nil {
			log.Fatal(jsonErr.Error())
			return nil, jsonErr
		}
		return nil, err
	}

	return res.Body, nil

}

func (c *Client) doRequest(req apiRequest) (*response.Sugar, error) {

	client := request.Client{
		URL:    c.BaseURL + req.path,
		Method: req.method,
		JSON:   req.data,
		Header: map[string]string{
			"Content-Type": "application/json",
			"User-Agent":   userAgent,
		},
	}

	resp := client.Send()
	httpRes := resp.Response()

	switch httpRes.StatusCode {
	case http.StatusBadRequest, http.StatusNotFound:
		var err RequestError
		resp.Scan(&err)
		return nil, err
	}

	if !resp.OK() {
		log.Error("Error while doRequest: " + resp.Error().Error())
		return nil, resp.Error()
	}

	return resp, nil

}
