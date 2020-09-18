package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"

	"tor-ent/common"
)

var cache map[string]string

func init() {
	cache = make(map[string]string)
}

func getFromCache(name string) string {
	onion, ok := cache[name]
	if !ok {
		return ""
	}
	return onion
}

func addToCache(e common.Entry) {
	cache[e.Name] = e.Address
}

func getAddressFromServer(name string) (string, error) {
	e := common.Entry{Name: name}
	url := fmt.Sprintf("%s/get", *serverAddress)
	resp, err := sendRequest(http.MethodGet, url, e)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(b, &e)
	if err != nil {
		return "", err
	}
	return e.Address, nil
}

func getOnion(name string) (string, error) {
	onion := getFromCache(name)
	if onion != "" {
		return onion, nil
	}
	onion, err := getAddressFromServer(name)
	if err != nil {
		return "", err
	}
	e := common.Entry{Name: name, Address: onion}
	addToCache(e)
	return onion, nil
}