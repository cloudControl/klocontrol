package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(sensorUrl)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data Data
	if err = data.Read(res.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Body.Close()

	err = index.Execute(w, data.SelectImage())
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
