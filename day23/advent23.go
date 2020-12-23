package main

import (
	// He provat de fer servir el Ring perquè semblava l'adient https://golang.org/pkg/container/ring/
	// i quedava molt xulo.
	// Però la segona part era impossible perquè no acabava mai i he hagut de recórrer a un mapa
	// per no haver de buscar els elements de la llista (deixant els anells)
	//
	// Suposo que podia haver seguit amb els anells però ...

	"fmt"
	"strconv"
)

func stringToNumbers(input string) []int {

	result := make([]int, len(input))
	for index, value := range input {
		number, err := strconv.Atoi(string(value))
		if err != nil {
			panic("Bad input")
		}
		result[index] = number
	}
	return result
}

type number struct {
	value int
	next  *number
}

func locateDestinationKey(value int, removeds [3]int) int {

	for value > 0 {
		found := false
		for _, removed := range removeds {
			if removed == value {
				found = true
				break
			}
		}
		if !found {
			return value
		}
		value--
	}

	return -1
}

func barreja(inputs []number, vegades int) []number {

	// caché de localitzacions
	nodes := map[int]*number{}
	for index, input := range inputs {
		nodes[input.value] = &inputs[index]
	}

	current := &inputs[0]
	moves := 0
	for moves < vegades {

		// Remove
		removeds := [3]int{current.next.value, current.next.next.value, current.next.next.next.value}

		destination := locateDestinationKey(current.value-1, removeds)
		if destination == -1 {
			destination = locateDestinationKey(len(inputs), removeds)
		}
		firstRemoved := nodes[removeds[0]]
		lastRemoved := nodes[removeds[2]]
		// append removeds
		current.next = lastRemoved.next
		lastRemoved.next = nodes[destination].next
		nodes[destination].next = firstRemoved

		current = current.next

		moves++
	}
	return inputs
}

func createNumbers(inputs []int) ([]number, int) {
	size := len(inputs)

	numbers := make([]number, size)
	numbers[0] = number{inputs[0], nil}
	one := 0

	current := &numbers[0]
	for i := 1; i < size; i++ {
		numbers[i] = number{inputs[i], nil}
		if inputs[i] == 1 {
			one = i
		}
		current.next = &numbers[i]
		current = current.next
	}
	numbers[size-1].next = &numbers[0]

	return numbers, one

}

func barreja1(input string, size int, vegades int) string {

	inputs := stringToNumbers(input)

	numbers, one := createNumbers(inputs)
	numbers = barreja(numbers, vegades)

	start := numbers[one].next
	solution := ""
	for start.value != 1 {
		solution += strconv.Itoa(start.value)
		start = start.next
	}

	return solution
}

func barreja2(input string, size int, vegades int) int {

	current := stringToNumbers(input)

	inputs := make([]int, size)

	for i := 0; i < size; i++ {
		if i < len(current) {
			inputs[i] = current[i]
		} else {
			inputs[i] = i + 1
		}
	}

	numbers, one := createNumbers(inputs)
	result := barreja(numbers, vegades)

	primer := result[one].next.value
	segon := result[one].next.next.value

	return primer * segon
}

func main() {

	// numeros := "389125467"
	numeros := "157623984"

	resultat1a := barreja1(numeros, len(numeros), 100)
	fmt.Println("Part 1:", resultat1a)

	// resultat1 := anells.barreja1Ring(numeros, 100)
	// fmt.Println("Part 1 (rings):", resultat1)

	resultat2 := barreja2(numeros, 1000000, 10000000)
	fmt.Println("Part 2:", resultat2)
}
