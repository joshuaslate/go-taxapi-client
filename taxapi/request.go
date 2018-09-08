package taxapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func (c *Client) makeRequest(resourcePath, method string, parameters [][]string, result interface{}) (statusCode int, err error) {
	req, err := http.NewRequest(http.MethodGet, rootURI+"v"+version+resourcePath, nil)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	req.Header.Add("Accept", "application/json")

	if len(parameters) > 0 {
		qs := url.Values{}
		for _, param := range parameters {
			qs.Add(param[0], param[1])
		}

		req.URL.RawQuery = qs.Encode()
	}

	// Make the request to TaxAPI.io
	if !time.Now().After(c.nextRequestAt) {
		fmt.Println("here?")
		time.Sleep(c.nextRequestAt.Sub(time.Now()))
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		if res != nil {
			return res.StatusCode, err
		}

		return http.StatusInternalServerError, err
	}

	// Close the request when done
	defer res.Body.Close()

	// Make sure we got a 200
	if res.StatusCode != 200 {
		return res.StatusCode, errors.New(res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(result)
	return res.StatusCode, err
}
