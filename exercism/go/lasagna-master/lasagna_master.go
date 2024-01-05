package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return 2 * len(layers)
	}

	return time * len(layers)
}

func Quantities(layers []string) (noodlesGrams int, SauceLiters float64) {
	const GramsPerNoodle = 50
	const LitersPerSauce = 0.2

	for _, l := range layers {
		switch l {
		case "noodles":
			noodlesGrams += GramsPerNoodle
		case "sauce":
			SauceLiters += LitersPerSauce
		}
	}
	return
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return myList
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	scaleFactor := float64(portions) / 2
	for i, q := range quantities {
		scaledQuantities[i] = q * scaleFactor
	}

	return scaledQuantities
}
