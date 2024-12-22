package main

func main() {
	//resources := LoadResources()
	//productionBuilidings := LoadProductionBuildings()
	LoadResourcesParameters()
	productionBuildingsParameters := LoadProductionBuildingsParameters()

	productionOutput := balancer(productionBuildingsParameters)

	SaveOutput(productionOutput)
}
