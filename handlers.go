package main

import (
	"net/http"
	"log"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(sensorUrl)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data Data
	err = data.Read(res.Body)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Body.Close()
	renderTemplate(w, selectImage(&data))
}
