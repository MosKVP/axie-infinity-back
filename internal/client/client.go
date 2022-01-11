package client

import (
	"axie-infinity-back/internal/log"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 10,
}

func Send(url string, method string, header map[string]string, reqBody interface{}, resBody interface{}) (*http.Response, error) {
	var req *http.Request
	var err error

	// Create Request
	if reqBody != nil {
		var b []byte
		log.Logger.Debugf("Request Body: %+v", reqBody)
		b, err = json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
		body := bytes.NewBuffer(b)
		req, err = http.NewRequest(method, url, body)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, err
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	// Send Request
	httpRes, err := client.Do(req)
	if err != nil {
		return httpRes, err
	}
	defer httpRes.Body.Close()

	// Parse Response
	httpResBody, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return httpRes, err
	}
	log.Logger.Debugf("Response Body: %s", httpResBody)

	err = json.Unmarshal(httpResBody, resBody)
	if err != nil {
		return httpRes, err
	}
	return httpRes, nil
}
