package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

func main() {
	httpPort := flag.String("port", "8080", "HTTP listen port")
	flag.Parse()

	if v := os.Getenv("sensor_url"); v != "" {
		sensorUrl = v
	}
	if v := os.Getenv("light_threshold"); v != "" {
		if newValue, err := strconv.Atoi(v); err == nil {
			lightThreshold = newValue
		}
	}
	log.Printf("GOMAXPROCS: %v", runtime.GOMAXPROCS(-1))
	log.Printf("Port: %v", *httpPort)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(":"+*httpPort, nil))
}

func selectImage(data *Data) string {
	log.Printf("data.Light: %v", data.Light)
	if data.Light > lightThreshold {
		return "toilet_full.svg"
	}
	return "toilet_empty.svg"
}

func renderTemplate(w http.ResponseWriter, image string) {
	err := index.Execute(w, image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
