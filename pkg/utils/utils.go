package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, i interface{}) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(requestBody, i)
	if err != nil {
		panic(err)
	}
}
