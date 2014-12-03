package main

import (
	"html/template"
	"os"
	"strconv"
)

var (
	sensorUrl      = "https://api.spark.io/v1/devices/53ff69066667574832581667/light?access_token=b77ca0da2879ebfa132bd6e8861c09b4a0adc33f"
	lightThreshold = 1000
	index          = template.Must(template.ParseFiles("templates/index.html"))

	port string
)

func SetConfig() {
	if v := os.Getenv("sensor_url"); v != "" {
		sensorUrl = v
	}
	if v := os.Getenv("light_threshold"); v != "" {
		if newValue, err := strconv.Atoi(v); err == nil {
			lightThreshold = newValue
		}
	}
}
