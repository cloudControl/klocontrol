package main

import (
	"net/http"
)

func renderTemplate(w http.ResponseWriter, image string) {
	err := index.Execute(w, image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
