package learn_go_pzn_golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("sample_output.json")

	customer := Customer{
		FirstName: "Iqbal",
		LastName:  "Atma",
		Age:       25,
	}

	encoder := json.NewEncoder(writer)
	encoder.Encode(customer)
}
