package utils

import (
	"github.com/rs/xid"
	"net/http"
	"encoding/json"
)

func GenerateId() string {
	return xid.New().String()
}

func ParseJSON(r *http.Request, item interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(item)
	if err != nil {
		return err
	}

	return nil
}