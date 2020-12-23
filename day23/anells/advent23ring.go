package anells

import (
	"container/ring"
	"errors"
	"fmt"
	"strconv"
)

func ringAppend(values []int) *ring.Ring {
	circle := ring.New(len(values))
	for _, value := range values {
		circle.Value = value
		circle = circle.Next()
	}
	return circle
}

func ringPrint(text string, anell *ring.Ring) {
	fmt.Print(text)
	anell.Do(func(p interface{}) {
		fmt.Print(p.(int))
	})
	fmt.Println()
}

func ringToString(anell *ring.Ring) string {
	resultat := ""
	anell.Do(func(p interface{}) {
		resultat += strconv.Itoa(p.(int))
	})
	return resultat
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
