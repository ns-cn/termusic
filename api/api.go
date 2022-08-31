package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func Get[V any](url string, data *V) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}
	return nil
}
