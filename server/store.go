package main

import (
	"errors"
	"fmt"
	"tor-ent/common"
)

var (
	store map[string]string
)

func init() {
	store = make(map[string]string)
}

func addToStore(e common.Entry) (error) {
	onion, _ := getFromStore(e.Name)
	if onion != "" {
		return errors.New("name already taken")
	}
	store[e.Name] = e.Address
	fmt.Println(store)
	return nil
}

func getFromStore(name string) (string, error) {
	onion, ok := store[name]
	if !ok {
		return "", errors.New("name does not exist")
	}
	return onion, nil
}