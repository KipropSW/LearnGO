package main

import (
	"fmt"
)

func main() {
	ns := GetNutrionalScore(NutrionalData{
		Energy:              EnergyFromKcal(0),
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(2),
		Sodium:              SodiumMiligram(500),
		Fruits:              FruitsPercent(60),
		Fibre:               FibreGram(7),
		Protein:             ProteinGram(4),
	}, Food)

	fmt.Printf("Nutritional Score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
