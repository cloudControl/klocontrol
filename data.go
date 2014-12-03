package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type Data struct {
	Light int `json:"result"`
}

func (data *Data) Read(response io.Reader) error {
	body, err := ioutil.ReadAll(response)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	return nil
}

func (data *Data) SelectImage() string {
	msg := fmt.Sprintf("Light: %d", data.Light)
	if data.Light > lightThreshold {
		log.Printf("%s - Occupied", msg)
		return "toilet_full.svg"
	}
	log.Printf("%s - Free", msg)
	return "toilet_empty.svg"
}
