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

func GetMap(item interface{} )(map[string]interface{}, error) {
	var itemMap map[string]interface{}
	tempByte, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(tempByte, &itemMap)
	if err != nil {
		return nil, err
	}
	return itemMap, err
}