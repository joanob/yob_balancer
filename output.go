package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ProductionOutput struct {
	Id        string                    `json:"id"`
	Processes []ProductionProcessOutput `json:"processes"`
}

type ProductionProcessOutput struct {
	Id            string                     `json:"id"`
	Time          int                        `json:"time"`
	Input         []ProductionResourceOutput `json:"input"`
	Output        []ProductionResourceOutput `json:"output"`
	ProfitPerHour float64                    `json:"profitPerHour"`
}

type ProductionResourceOutput struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type JSONOutput struct {
	Production []ProductionOutput `json:"production"`
}

func SaveOutput(production []ProductionOutput) {
	now := time.Now()
	filename := fmt.Sprintf("%v%v%v_%v%v%v.json", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var output JSONOutput
	output.Production = production

	decoder := json.NewEncoder(file)
	err = decoder.Encode(&output)
	if err != nil {
		panic(err)
	}
}
