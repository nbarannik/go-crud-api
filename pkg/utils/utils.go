package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, i interface{}) error {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(requestBody, i)
	return err
}
