package main

type ProcessInputOutputCombination struct {
	Input  []ProcessResource
	Output []ProcessResource
}

type ProcessResource struct {
	Id       string
	Quantity int
	Price    int
}

func balancer(resourcesParameters []ResourceParameter, productionBuildingsParameters []ProductionBuildingParameters) []ProductionOutput {
	// For all production buildings, check all possibilities
	// Check all inputs and output for all times and its prices

	var output []ProductionOutput
	var productionOutput ProductionOutput
	var productionProcessOutput ProductionProcessOutput
	//var productionProcessResourceOutput []ProductionResourceOutput

	for _, productionBuilding := range productionBuildingsParameters {
		productionOutput.Id = productionBuilding.Id
		productionOutput.Processes = make([]ProductionProcessOutput, 0)

		for _, process := range productionBuilding.Processes {
			productionProcessOutput.Id = process.Id
			productionProcessOutput.Input = make([]ProductionResourceOutput, 0)
			productionProcessOutput.Output = make([]ProductionResourceOutput, 0)

			//combinations := getInputOutputCombinations(process)

			for time := process.MinTime; time < process.MaxTime; time++ {
				productionProcessOutput.Time = time

				// First, get all combinations of inputs and outputs with quantities and prices

				// Calculate price of inputs, calculate price of outputs and compute
			}

			productionOutput.Processes = append(productionOutput.Processes, productionProcessOutput)
		}

		output = append(output, productionOutput)
	}

	return output
}

func getInputOutputCombinations(process ProductionProcessesParameters) []ProcessInputOutputCombination {
	combinations := make([]ProcessInputOutputCombination, 0)

	// Primero obtener para cada input y output sus combinaciones particulares
	// Si un input tiene una cantidad entre 1 y 3 y un precio entre 1 y 5, las combinaciones son
	// 1/1, 1/2, 1/3, 1/4, 1/5, 2/1, 2/

	//inputCombinations := getResourcesCombinations(process.Input)
	//outputCombinations := getResourcesCombinations(process.Output)

	//if len(inputCombinations) != 0 {
	/* for _, input := range inputCombinations {
		for _, output := range outputCombinations {
			combinations = append(combinations, ProcessInputOutputCombination{Input: make([]ProcessResource, 0), Output: outputCombinations})

		}
	} */

	//} else {
	//7	combinations = append(combinations, ProcessInputOutputCombination{Input: make([]ProcessResource, 0), Output: outputCombinations})
	//}

	return combinations
}

// Obtener todas las combinaciones para una lista de recursos
// La función es recursiva: recorre los recursos para obtener sus combinaciones
// las combina con las combinaciones de los recursos anteriores y pasa la lista de recursos restantes
func getResourcesCombinations(combinations [][]ProcessResource, resources []ProductionResourceParameters) [][]ProcessResource {
	// Combinations es una lista de combinaciones entre recursos
	if len(resources) > 0 {
		resourceCombinations := getResourceCombinations(resources[0])

		combinationsWithResource := make([][]ProcessResource, 0)

		if len(combinations) == 0 {
			for _, resource := range resourceCombinations {
				combinationsWithResource = append(combinationsWithResource, []ProcessResource{resource})
			}
		} else {
			for _, combination := range combinations {
				// Imaginemos que la combinacion tiene agua y semillas
				// Al añadir electricidad, para cada combinación de agua, semillas y electricidad debe añadir una línea
				for _, resource := range resourceCombinations {
					combinationsWithResource = append(combinationsWithResource, append(combination, resource))
				}
			}
		}
		combinations = getResourcesCombinations(combinationsWithResource, resources[1:])
	}

	return combinations
}

// Obtener todas las combinaciones de cantidades y precios de un recurso
func getResourceCombinations(resource ProductionResourceParameters) []ProcessResource {
	resourcePrice := getResourceParameters(resource.Id)

	resourceCombinations := make([]ProcessResource, 0)

	for quantity := resource.Min; quantity <= resource.Max; quantity++ {
		for price := resourcePrice.Min; price <= resourcePrice.Max; price++ {
			resourceCombinations = append(resourceCombinations, ProcessResource{Id: resource.Id, Quantity: quantity, Price: price})
		}
	}

	return resourceCombinations
}
