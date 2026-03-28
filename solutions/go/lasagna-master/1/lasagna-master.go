package lasagna


// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, average int) int {
    if (average == 0) {
        average = 2
    }
	
    numlayers := len(layers)

    return numlayers * average
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    var noodles int = 0
    var sauce float64 = 0.0


    for _, item := range layers {
        switch item {
            case "sauce":
        		sauce += 0.2
            case "noodles":
        		noodles += 50
        }
    }

    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) []string {
    modifiedList := myList
    lastItem := friendsList[len(friendsList) - 1]
    modifiedList = append(modifiedList, lastItem)
    return modifiedList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, number int) []float64 {
    var toscale float64
    toscale =  float64(number) / float64(2)
    
    scaledquantities := make([]float64, len(quantities))
    copy(scaledquantities, quantities)
                             

    for i := 0; i < len(scaledquantities); i++ {
        scaledquantities[i] *= toscale
    }


    return scaledquantities
}