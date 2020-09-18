package main

import (
	"net/http"
	"encoding/json"
	"bytes"
)

var client *http.Client

func init() {
	client = &http.Client{}
}

func sendRequest(method, url string, body interface{}) (*http.Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}