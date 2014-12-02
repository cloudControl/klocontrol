package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
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
