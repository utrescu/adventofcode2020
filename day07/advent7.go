package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

type bag struct {
	name string
	qty  int
}

func splitContents(content string, separator string) []bag {
	var re = regexp.MustCompile(`(?m)^(\d) (.+) bags?`)

	var bags []bag
	stringBags := strings.Split(content, separator)

	for _, stringBag := range stringBags {
		match := re.FindStringSubmatch(stringBag)
		if len(match) > 0 {
			qty, _ := stringToInt(match[1])
			newBag := bag{
				name: match[2],
				qty:  qty,
			}
			bags = append(bags, newBag)
		} else {
			bags = append(bags, bag{name: "", qty: 0})
		}
	}

	return bags
}

func getColorName(contents string) string {
	var re = regexp.MustCompile(`(?m)^(.+) bags?`)
	return re.FindStringSubmatch(contents)[1]
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (map[string][]bag, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines = make(map[string][]bag)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linia := scanner.Text()
		principal := strings.Split(linia, " contain ")
		name := getColorName(principal[0])

		lines[name] = splitContents(principal[1], ", ")
	}
	return lines, scanner.Err()
}

// Part 1
func teShiny(coloret bag, bags map[string][]bag) bool {
	if coloret.name == "shiny gold" {
		return true
	}
	if coloret.qty == 0 {
		return false
	}

	for _, innerColorets := range bags[coloret.name] {
		if teShiny(innerColorets, bags) {
			return true
		}
	}
	return false
}

func searchShinyGoldBags(tots map[string][]bag) (int, int) {
	suma := 0
	for name, colorets := range tots {
		ok := false
		if name != "shiny gold" {
			for _, coloret := range colorets {

				ok = teShiny(coloret, tots)
				if ok {
					suma = suma + 1
					break
				}
			}
		}
	}

	return suma, 0
}

func main() {
	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1, correctes2 := searchShinyGoldBags(linies)

	fmt.Println("Part 1: ", correctes1)
	fmt.Println("Part 2: ", correctes2)
}
