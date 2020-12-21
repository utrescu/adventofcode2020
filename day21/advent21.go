package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
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

type fatalCombination struct {
	alergic    string
	ingredient string
}

func searchAlergens(ingredientCount map[string]int, alergens map[string]alergen) (int, string) {
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

	part2 := make([]fatalCombination, 0)

	for len(alergens) > 0 {
		actualIngredient := ""
		actualAlergia := ""
		for _, alergia := range alergens {
			if len(alergia.ingredients) == 1 {
				for actualIngredient = range alergia.ingredients {
				}
				actualAlergia = alergia.name
				part2 = append(part2, fatalCombination{
					alergia.name,
					actualIngredient})
				break
			}
		}
		delete(alergens, actualAlergia)
		for _, purge := range alergens {
			delete(purge.ingredients, actualIngredient)
		}
	}

	sort.Slice(part2, func(i, j int) bool {
		return part2[i].alergic < part2[j].alergic
	})

	resultat := make([]string, 0)
	for _, cosa := range part2 {
		resultat = append(resultat, cosa.ingredient)
	}
	return part1, strings.Join(resultat, ",")
}

func main() {
	sumaIngredients, alergens, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1, resultat2 := searchAlergens(sumaIngredients, alergens)
	fmt.Println("Part 1: ", correctes1)
	fmt.Println("Part 2:", resultat2)

}
