package main

import (
	// Faig servir el Ring perquè sembla l'adient https://golang.org/pkg/container/ring/
	"container/ring"
	"errors"
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

func ringAppend(values []int) *ring.Ring {
	circle := ring.New(len(values))
	for _, value := range values {
		circle.Value = value
		circle = circle.Next()
	}
	return circle
}

func ringToString(anell *ring.Ring) string {
	resultat := ""
	anell.Do(func(p interface{}) {
		resultat += strconv.Itoa(p.(int))
	})
	return resultat
}

func ringPrint(text string, anell *ring.Ring) {
	fmt.Print(text)
	anell.Do(func(p interface{}) {
		fmt.Print(p.(int))
	})
	fmt.Println()
}

func ringLocateDestination(origen int, anell *ring.Ring) *ring.Ring {
	max := anell
	maxValue := -1
	minMax := anell
	minMaxValue := -1
	iterator := anell
	n := iterator.Len()
	for i := 0; i < n; i++ {
		valor := iterator.Value.(int)
		if valor > maxValue {
			max = iterator
			maxValue = valor
		}
		if valor < origen && valor > minMaxValue {
			minMax = iterator
			minMaxValue = valor
		}
		iterator = iterator.Next()

	}
	if minMaxValue != -1 {
		return minMax
	}
	return max
}

func ringLocateValue(searched int, anell *ring.Ring) (*ring.Ring, error) {
	iterator := anell
	n := iterator.Len()
	for i := 0; i < n; i++ {
		if iterator.Value.(int) == searched {
			return iterator, nil
		}
		iterator = iterator.Next()
	}
	return nil, errors.New("not found")
}

func barreja(numbers []int, vegades int) *ring.Ring {
	anell := ringAppend(numbers)

	moves := 0
	for moves < vegades {
		agafats := anell.Unlink(3)
		destination := ringLocateDestination(anell.Value.(int), anell)
		destination.Link(agafats)
		anell = anell.Next()
		moves++
	}
	return anell
}

func barreja1Ring(input string, vegades int) string {
	numbers := stringToNumbers(input)

	anell := barreja(numbers, vegades)
	// Localitzar el número 1
	result, err := ringLocateValue(1, anell)
	if err != nil {
		panic(err.Error())
	}
	solution := ringToString(result)

	return solution[1:]
}

func barreja1(input string, size int, vegades int) string {

	inputs := stringToNumbers(input)
	numbers := make([]number, size)
	numbers[0] = number{inputs[0], nil}
	one := 0

	current := &numbers[0]
	var value int
	for i := 1; i < size; i++ {
		numbers[i] = number{inputs[i], nil}
		if value == 1 {
			one = i
		}
		current.next = &numbers[i]
		current = current.next
	}
	numbers[size-1].next = &numbers[0]

	numbers = barrejaNoRing(numbers, vegades)

	start := numbers[one].next
	solution := ""
	for start.value != 1 {
		solution += strconv.Itoa(start.value)
		start = start.next
	}

	return solution
}

type number struct {
	value int
	next  *number
}

func locateDestination(actual *number, value int) *number {
	minMax := &number{-1, nil}
	max := &number{-1, nil}
	start := actual.next
	for start.value != value {
		currentValue := start.value
		if currentValue == value-1 {
			return start
		}
		if currentValue < value && currentValue > minMax.value {
			minMax = start
		} else if currentValue > max.value {
			max = start
		}
		start = start.next
	}
	if minMax.value != -1 {
		return minMax
	}
	return max
}

func barrejaNoRing(input []number, vegades int) []number {

	current := &input[0]
	moves := 0
	for moves < vegades {
		if moves%100000 == 0 {
			fmt.Println("... ", moves)
		}
		// Remove
		firstRemoved := current.next
		lastRemoved := firstRemoved.next.next
		current.next = lastRemoved.next
		// append removeds
		destination := locateDestination(current, current.value)
		afterdestination := destination.next
		destination.next = firstRemoved
		lastRemoved.next = afterdestination
		// step
		current = current.next

		moves++
	}
	return input
}

func barreja2(input string, quantitat int, vegades int) int {
	inputs := stringToNumbers(input)

	numbers := make([]number, quantitat)

	numbers[0] = number{inputs[0], nil}
	one := 0

	current := &numbers[0]
	var value int
	for i := 1; i < quantitat; i++ {
		if i < len(inputs) {
			value = inputs[i]
		} else {
			value = i
		}
		numbers[i] = number{value, nil}
		if value == 1 {
			one = i
		}
		current.next = &numbers[i]
		current = current.next
	}
	numbers[quantitat-1].next = &numbers[0]

	result := barrejaNoRing(numbers, vegades)

	primer := result[one].next.value
	segon := result[one].next.next.value

	return primer * segon
}

func main() {

	// numeros := "389125467"
	numeros := "157623984"

	resultat1a := barreja1(numeros, len(numeros), 100)
	fmt.Println("Part 1:", resultat1a)
	resultat1 := barreja1Ring(numeros, 100)
	fmt.Println("Part 1:", resultat1)

	resultat2 := barreja2(numeros, 1000000, 10000000)
	fmt.Println("Part 2:", resultat2)
}
