package chance
import "math/rand"
//import "time"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
    //rand.Seed(time.Now().UnixNano())
    n := rand.Intn(20) + 1
    return n
    
	panic("Please implement the RollADie function")
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
    //rand.Seed(time.Now().UnixNano())
    f := rand.Float64()*12
    return f
	panic("Please implement the GenerateWandEnergy function")
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
    var Animals = []string {"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog",}
    rand.Shuffle(len(Animals), func(i, j int) {
	Animals[i], Animals[j] = Animals[j], Animals[i]
})
    return Animals
}
