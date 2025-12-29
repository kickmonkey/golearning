package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, n int) (int){
    if n == 0 {
        return len(layers) * 2
    }
    return len(layers) * n
    panic("")
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string)(int, float64){
    noods := 0
    saus := 0
    for _, v :=range layers{
        if v == "noodles"{
            noods ++
        }
        if v == "sauce"{
            saus++
        }
    }
    return noods * 50, float64(saus) * 0.2
    panic("")
    
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string){
    myList[len(myList)-1] = friendList[len(friendList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64{
    factor := float64(portions)/2.0
    result := make([]float64, len(quantities))
    for i, v := range quantities{
        result[i] = v*factor
    }
    return result
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
