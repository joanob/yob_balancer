package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ProductionBuilding struct {
	Id        string                `json:"id"`
	Processes []ProductionProcesses `json:"processes"`
}

type ProductionProcesses struct {
	Id     string               `json:"id"`
	Input  []ProductionResource `json:"input"`
	Output []ProductionResource `json:"output"`
}

type ProductionResource struct {
	Id string `json:"id"`
}

type JSONProductionBuildings struct {
	ProductionBuildings []ProductionBuilding `json:"productionBuildings"`
}

func LoadProductionBuildings() []ProductionBuilding {
	file, err := os.Open("production.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var jsonProductionBuildings JSONProductionBuildings

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonProductionBuildings)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonProductionBuildings.ProductionBuildings)

	return jsonProductionBuildings.ProductionBuildings
}
