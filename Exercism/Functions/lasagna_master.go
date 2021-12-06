package lasagna

func PreparationTime(layers []string, avg_prep int) int {
	if avg_prep == 0 {
		avg_prep = 2
	}
	return avg_prep * len(layers)
}

func Quantities(layers []string) (noodles int, sauce float64) {
	// Common pattern matching for slice.
	for _, v := range layers {
		switch {
		case v == "noodles":
			noodles += 50
		case v == "sauce":
			sauce += 0.2
		}
	}
	return

}

func AddSecretIngredient(friendList []string, ownList []string) []string {
	ownList = append(ownList, friendList[len(friendList)-1])
	return ownList
}

func ScaleRecipe(slice []float64, amounts int) []float64 {
	scaledQuantities := make([]float64, len(slice)) // make another slice with the same length as original slice.
	for k, v := range slice {
		scaledQuantities[k] = v * float64(amounts) * 0.5 //for every element in scaledQuantities slice, multiply it by the amount times 0.5
	}
	return scaledQuantities
}
