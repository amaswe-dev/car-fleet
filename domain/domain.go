package domain

import (
	"encoding/json"
	"io/ioutil"
)

type Vehicle struct {
	Id    string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

type Metrics struct {
	Brand   string         `json:"brand"`
	Overall int            `json:"overall"`
	Models  map[string]int `json:"models"`
}

func ParseInput(file string) ([]Vehicle, error) {
	f, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	var v []Vehicle

	err = json.Unmarshal(f, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
