package learn_go_pzn_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "Iqbal",
		LastName:  "Atma",
		Age:       25,
		Hobbies:   []string{"Naik Gunung", "Ngoding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Iqbal","LastName":"Atma","Age":25,"Hobbies":["Naik Gunung","Ngoding"]}`
	jsonBytes := []byte(jsonString)

	custom := &Customer{}

	err := json.Unmarshal(jsonBytes, custom)
	if err != nil {
		panic(err)
	}

	fmt.Println(custom)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Iqbal",
		Address: []Address{
			{
				Street:     "Street",
				Country:    "Indonesia",
				PostalCode: "1123",
			},
			{
				Street:     "X",
				Country:    "Indonesia",
				PostalCode: "1123xx",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayOnlyJSONArray(t *testing.T) {
	jsonString := `[{"Street":"Street","Country":"Indonesia","PostalCode":"1123"},{"Street":"X","Country":"Indonesia","PostalCode":"1123xx"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	errr := json.Unmarshal(jsonBytes, addresses)
	if errr != nil {
		panic(errr)
	}

	fmt.Println(addresses)
}
