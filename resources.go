package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Resource struct {
	Id string `json:"id"`
}

type JSONResources struct {
	Resources []Resource `json:"resources"`
}

func LoadResources() []Resource {
	file, err := os.Open("resources.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var jsonResources JSONResources

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonResources)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonResources.Resources)

	return jsonResources.Resources
}
