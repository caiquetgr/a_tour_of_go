package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	person := Person{
		Name: "Fred",
		Age:  40,
	}

	// Write it out!
	tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())
	err = json.NewEncoder(tmpFile).Encode(person)
	if err != nil {
		panic(err)
	}
	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}

	// Let's read it in!
	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		panic(err)
	}
	var personFromFile Person
	err = json.NewDecoder(tmpFile2).Decode(&personFromFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", personFromFile)
}
