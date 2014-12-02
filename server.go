package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

var (
	sensorUrl      = "https://api.spark.io/v1/devices/53ff69066667574832581667/light?access_token=b77ca0da2879ebfa132bd6e8861c09b4a0adc33f"
	lightThreshold = 1000
	templates      = template.Must(template.ParseFiles("templates/index.html"))
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
	http.HandleFunc("/", reqHandler)

	log.Fatal(http.ListenAndServe(":"+*httpPort, nil))
}

type sensorData struct {
	Light int `json:"result"`
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
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

func selectImage(data *sensorData) string {
	log.Printf("data.Light: %v", data.Light)
	if data.Light > lightThreshold {
		return "toilet_full.svg"
	}
	return "toilet_empty.svg"
}

func renderTemplate(w http.ResponseWriter, image string) {
	err := templates.Execute(w, image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
