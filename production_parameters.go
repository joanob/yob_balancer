package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ProductionBuildingParameters struct {
	Id        string                          `json:"id"`
	Processes []ProductionProcessesParameters `json:"processes"`
}

type ProductionProcessesParameters struct {
	Id      string                         `json:"id"`
	MinTime int                            `json:"minTime"`
	MaxTime int                            `json:"maxTime"`
	Input   []ProductionResourceParameters `json:"input"`
	Output  []ProductionResourceParameters `json:"output"`
}

type ProductionResourceParameters struct {
	Id  string `json:"id"`
	Min int    `json:"min"`
	Max int    `json:"max"`
}

type JSONProductionBuildingsParameters struct {
	ProductionBuildingParameters []ProductionBuildingParameters `json:"productionBuildingParameters"`
}

func LoadProductionBuildingsParameters() []ProductionBuildingParameters {
	file, err := os.Open("production_parameters.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var jsonProductionBuildingParameters JSONProductionBuildingsParameters

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonProductionBuildingParameters)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonProductionBuildingParameters.ProductionBuildingParameters)

	return jsonProductionBuildingParameters.ProductionBuildingParameters
}
