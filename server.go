package main

import (
	"flag"
	"log"
	"net/http"
	"runtime"
)

func init() {
	flag.StringVar(&port, "port", "8080", "HTTP listen port")
	SetConfig()
}

func main() {
	flag.Parse()

	log.Print("Welcome to klocontrol!")
	log.Printf("GOMAXPROCS: %v", runtime.GOMAXPROCS(-1))
	log.Printf("Port: %v", port)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
