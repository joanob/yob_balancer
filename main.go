package main

func main() {
	//resources := LoadResources()
	//productionBuilidings := LoadProductionBuildings()
	resourcesParameters := LoadResourcesParameters()
	productionBuildingsParameters := LoadProductionBuildingsParameters()

	productionOutput := balancer(resourcesParameters, productionBuildingsParameters)

	SaveOutput(productionOutput)
}
