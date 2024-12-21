package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ResourceParameter struct {
	Id  string `json:"id"`
	Min int    `json:"min"`
	Max int    `json:"max"`
}

type JSONResourcesParameters struct {
	ResourcesParameters []ResourceParameter `json:"resourcesParameters"`
}

func LoadResourcesParameters() []ResourceParameter {
	file, err := os.Open("resources_parameters.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var jsonResourcesParameters JSONResourcesParameters

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonResourcesParameters)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonResourcesParameters.ResourcesParameters)

	return jsonResourcesParameters.ResourcesParameters
}
