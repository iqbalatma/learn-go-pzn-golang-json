package learn_go_pzn_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Hobbies   []string
	Address   []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Iqbal",
		LastName:  "Atma",
		Age:       25,
		Hobbies:   []string{"Naik Gunung", "Ngoding"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
