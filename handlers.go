package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(sensorUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := sensorData{}
	var body []byte
	body, err = ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Body.Close()
	renderTemplate(w, selectImage(&data))
}
