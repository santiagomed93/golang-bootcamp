package services

import (
	"encoding/json"
	"net/http"
)

func getJSON(url string, target interface{}) (interface{}, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode == http.StatusNotFound {
		return target, nil
	}
	err = json.NewDecoder(r.Body).Decode(&target)
	if err != nil {
		return target, err
	}
	return target, nil
}
