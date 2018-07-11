package ynab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"strconv"

	"bmvs.io/ynab/api"
)

const apiEndpoint = "https://api.youneedabudget.com/v1"

// NewClient facilitates the creation of a new client instance
func NewClient(accessToken string) *client {
	c := &client{
		accessToken: accessToken,
		client:      http.DefaultClient,
	}
	return c
}

// client API
type client struct {
	accessToken string
	client      *http.Client
}

// GET sends a GET request to the YNAB API
func (c *client) GET(url string, responseModel interface{}) error {
	return c.do(http.MethodGet, url, responseModel)
}

// do sends a request to the YNAB API
func (c *client) do(method, url string, responseModel interface{}) error {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", apiEndpoint, url), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		response := struct {
			Error *api.Error `json:"error"`
		}{}

		if err := json.Unmarshal(body, &response); err != nil {
			// returns a forged *api.Error fore ease of use
			// because either the response body is empty or the response is
			// non compliant with YNAB's API specification
			// https://api.youneedabudget.com/#errors
			apiError := &api.Error{
				ID:     strconv.Itoa(res.StatusCode),
				Name:   "unknown_api_error",
				Detail: "Unknown API error",
			}
			return apiError
		}

		return response.Error
	}

	if err := json.Unmarshal(body, &responseModel); err != nil {
		return err
	}

	return nil
}
