package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var p Person

	data := `{"name": "Fred", "age": 40}
	{"name": "Mary", "age": 21}
	{"name": "Pat", "age": 30}`

	dec := json.NewDecoder(strings.NewReader(data))
	personList := make([]Person, 0, 3)

	for dec.More() {
		err := dec.Decode(&p)
		if err != nil {
			return
		}
		fmt.Println(p)
		personList = append(personList, p)
	}

	var outputJson bytes.Buffer
	enc := json.NewEncoder(&outputJson)

	for _, input := range personList {
		err := enc.Encode(input)
		if err != nil {
			return
		}
	}
	fmt.Println(outputJson.String())
}
