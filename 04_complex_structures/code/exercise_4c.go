package main

import "fmt"

func average(scores [5]float64) float64 {
	total := 0.0
	for _, num := range scores {
		total += num
	}
	return total / float64(len(scores))
}

var initialPets map[string]string = map[string]string{
	"fido":     "dog",
	"penelope": "horse",
	"nancy":    "cat",
}

func doesPetExist(petName string) bool {
	_, ok := initialPets[petName]
	return ok
}

var initialGroceries = []string{"beans", "lemons", "chicken", "fruit"}

func addGroceryToList(newGrocery ...string) []string {
	foods := initialGroceries
	for _, g := range newGrocery {
		foods = append(foods, g)
	}
	return foods
}

func main() {
	scores := [5]float64{2, 7, 8, 1, 9}
	fmt.Println(average(scores))

	pet := "fido"
	petExists := doesPetExist(pet)
	fmt.Println(petExists)

	groceryList := addGroceryToList("beets", "chocolate")
	fmt.Println(groceryList)
}
