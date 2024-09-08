package learn_go_pzn_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Macbook Pro","image_url":"https://www.apple.com/img/jpg"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result["id"])
}
