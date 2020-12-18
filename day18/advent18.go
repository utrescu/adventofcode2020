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

type operation struct {
	number   int
	operator string
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

	}
	return lines, scanner.Err()
}

func opera(numero1 int, numero2 int, operador string) int {
	switch {
	case operador == "+":
		return numero1 + numero2
	case operador == "*":
		return numero1 * numero2
	default:
		return 0
	}
}

func operaLinia(linia string) string {
	var re = regexp.MustCompile(`(?m)(\d+) (\D) (\d+)`)
	match := re.FindStringSubmatch(linia)
	for len(match) > 1 {
		numero1, _ := stringToInt(match[1])
		operador := match[2]
		numero2, _ := stringToInt(match[3])
		resultat := opera(numero1, numero2, operador)
		linia = strings.Replace(linia, match[0], strconv.Itoa(resultat), 1)
		match = re.FindStringSubmatch(linia)
	}

	return linia
}

func operaSumes(linia string) string {
	var re = regexp.MustCompile(`(?m)(\d+) (\+) (\d+)`)

	match := re.FindStringSubmatch(linia)
	for len(match) > 1 {
		numero1, _ := stringToInt(match[1])
		operador := match[2]
		numero2, _ := stringToInt(match[3])
		resultat := opera(numero1, numero2, operador)
		linia = strings.Replace(linia, match[0], strconv.Itoa(resultat), 1)
		match = re.FindStringSubmatch(linia)
	}

	return linia
}

func calcula(linia string, advanced bool) int {

	// Elimina parèntesis
	var re = regexp.MustCompile(`(?m)(\([^\(\)]+\))`)

	match := re.FindStringSubmatch(linia)
	for len(match) > 1 {
		contingut := match[1]
		numero := calcula(contingut[1:len(contingut)-1], advanced)
		linia = strings.Replace(linia, contingut, strconv.Itoa(numero), 1)
		match = re.FindStringSubmatch(linia)
	}

	if advanced {
		// Elimina sumes
		linia = operaSumes(linia)
	}

	// Calcula
	resultat, _ := stringToInt(operaLinia(linia))
	return resultat
}

func sumaOperacions(linies []string, advanced bool) int {

	suma := 0
	for _, linia := range linies {
		suma += calcula(linia, advanced)
	}
	return suma
}

func main() {
	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1 := sumaOperacions(linies, false)
	fmt.Println("Part 1: ", correctes1)

	correctes2 := sumaOperacions(linies, true)
	fmt.Println("Part 2: ", correctes2)
}
