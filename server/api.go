package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"tor-ent/common"
)


func addEntry(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	var e common.Entry
	err = json.Unmarshal(b, &e)
	fmt.Println(e)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	err = addToStore(e)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("added"))
}

func getEntry(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	var e common.Entry
	err = json.Unmarshal(b, &e)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	onion, err := getFromStore(e.Name)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	e.Address = onion
	b, err = json.Marshal(e)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}