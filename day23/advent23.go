package main

import (
	// Faig servir el Ring perquè sembla l'adient https://golang.org/pkg/container/ring/
	"container/ring"
	"errors"
	"fmt"
	"strconv"
)

func stringToNumbersRing(input string) *ring.Ring {

	circle := ring.New(len(input))
	for _, value := range input {
		number, err := strconv.Atoi(string(value))
		if err != nil {
			panic("Bad input")
		}
		circle.Value = number
		circle = circle.Next()
	}
	return circle
}

func findDestination(current int, numbers []int) (int, int) {
	max := 0
	maxIndex := -1
	proper := -1
	properIndex := -1
	for i, candidate := range numbers {
		if candidate > max {
			max = candidate
			maxIndex = i
		}
		if candidate < current {
			if candidate > proper {
				proper = candidate
				properIndex = i
			}
		}
	}

	if proper == -1 {
		return maxIndex, max
	}
	return properIndex, proper
}

func ringString(anell *ring.Ring) string {
	resultat := ""
	anell.Do(func(p interface{}) {
		resultat += strconv.Itoa(p.(int))
	})
	return resultat
}
func printRing(text string, anell *ring.Ring) {
	fmt.Print(text)
	anell.Do(func(p interface{}) {
		fmt.Print(p.(int))
	})
	fmt.Println()
}

func locateDestination(origen int, anell *ring.Ring) *ring.Ring {
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

func locateValueInRing(searched int, anell *ring.Ring) (*ring.Ring, error) {
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

func barreja(input string, vegades int) string {
	anell := stringToNumbersRing(input)

	moves := 0

	for moves < vegades {
		agafats := anell.Unlink(3)
		destination := locateDestination(anell.Value.(int), anell)
		destination.Link(agafats)
		anell = anell.Next()
		moves++
	}

	// Localitzar el número 1
	result, err := locateValueInRing(1, anell)
	if err != nil {
		panic(err.Error())
	}
	solution := ringString(result)

	return solution[1:]
}

func main() {

	// numeros := "389125467"
	numeros := "157623984"
	resultat1 := barreja(numeros, 100)
	fmt.Println("Part 1:", resultat1)

}
