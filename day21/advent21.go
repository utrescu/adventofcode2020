package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (map[string]int, map[string]alergen, error) {
	re := regexp.MustCompile(`(.+) \(contains (.+)\)`)
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	suma := map[string]int{}
	alergens := map[string]alergen{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		currentIngredients := make(map[string]bool, 0)
		separaCamps := re.FindStringSubmatch(line)
		for _, ingredientName := range strings.Fields(separaCamps[1]) {
			suma[ingredientName]++
			currentIngredients[ingredientName] = false
		}

		alergensList := strings.Replace(separaCamps[2], ",", "", -1)
		for _, alergenName := range strings.Fields(alergensList) {
			value, ok := alergens[alergenName]
			if !ok {
				newIngredients := map[string]bool{}
				for k := range currentIngredients {
					newIngredients[k] = false
				}
				alergens[alergenName] = alergen{alergenName, newIngredients}
			} else {

				for nom := range value.ingredients {
					if _, ok := currentIngredients[nom]; !ok {
						delete(value.ingredients, nom)
					}
				}

			}
		}

	}
	return suma, alergens, scanner.Err()
}

type alergen struct {
	name        string
	ingredients map[string]bool
}

func (a alergen) print() {
	fmt.Println("Alergen", a.name, a.ingredients)
}

func main() {
	sumaIngredients, alergens, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	for _, alergen := range alergens {
		alergen.print()
	}

	correctes1 := searchAlergens(sumaIngredients, alergens)
	fmt.Println("Part 1: ", correctes1)

}

func searchAlergens(ingredientCount map[string]int, alergens map[string]alergen) int {
	part1 := 0

	for ingredient, count := range ingredientCount {
		found := false
		for _, alergen := range alergens {
			if _, ok := alergen.ingredients[ingredient]; ok {
				found = true
				break
			}
		}
		if !found {
			part1 += count
		}
	}

	return part1
}
