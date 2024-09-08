package learn_go_pzn_golang_json

import (
	"encoding/json"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Iqbal","LastName":"Atma","Age":25}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		return
	}
}
