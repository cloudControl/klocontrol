package main

import (
	"fmt"
	"log"
	"net/http"
)

func selectImage(data *Data) string {
	msg := fmt.Sprintf("Light: %d", data.Light)
	if data.Light > lightThreshold {
		log.Printf("%s - Occupied", msg)
		return "toilet_full.svg"
	}
	log.Printf("%s - Free", msg)
	return "toilet_empty.svg"
}

func renderTemplate(w http.ResponseWriter, image string) {
	err := index.Execute(w, image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
