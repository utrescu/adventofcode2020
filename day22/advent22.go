package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	player := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Player") {
			if len(player) != 0 {
				lines = append(lines, player)
				player = make([]int, 0)
			}
		} else {
			value, err := stringToInt(line)
			if err == nil {
				player = append(player, value)
			}
		}
	}
	lines = append(lines, player)
	return lines, scanner.Err()
}

func sumaPunts(cartes []int) int {
	resultat := 0
	numCartes := len(cartes)

	for i, valor := range cartes {
		resultat += valor * (numCartes - i)
	}
	return resultat
}

// Part 1 ----

func playCombat(cartes [][]int) int {

	var winner []int

	for {
		player0, cartes0 := cartes[0][0], cartes[0][1:]
		player1, cartes1 := cartes[1][0], cartes[1][1:]

		if player0 > player1 {
			cartes0 = append(cartes0, player0, player1)
			if len(cartes1) == 0 {
				winner = cartes0
				break
			}
		} else {
			cartes1 = append(cartes1, player1, player0)
			if len(cartes0) == 0 {
				winner = cartes1
				break
			}
		}
		cartes[0] = cartes0
		cartes[1] = cartes1
	}

	return sumaPunts(winner)
}

func recursive(cartes [][]int, num int) int {
	jugadaExistent := false

	jugadesAnteriors := make([][][]int, 0)

	for {

		// Mirar si ja teniem aquesta jugada anteriorment
		if jugadaExistent {

			return 0
		}

		// Afegir jugada a anteriors (les he de copiar perquè es passa per referència)
		anterior0 := make([]int, len(cartes[0]))
		copy(anterior0, cartes[0])
		anterior1 := make([]int, len(cartes[1]))
		copy(anterior1, cartes[0])
		jugadesAnteriors = append(jugadesAnteriors, [][]int{anterior0, anterior1})

		player0, cartes0 := cartes[0][0], cartes[0][1:]
		player1, cartes1 := cartes[1][0], cartes[1][1:]

		if len(cartes0) >= player0 && len(cartes1) >= player1 {
			// TODO: subgame
			subcartes0 := make([]int, player0)
			copy(subcartes0, cartes0)
			subcartes1 := make([]int, player1)
			copy(subcartes1, cartes1)

			guanyador := recursive([][]int{subcartes0, subcartes1}, num+1)

			if guanyador == 0 {
				cartes0 = append(cartes0, player0, player1)
				if len(cartes1) == 0 {
					return guanyador
				}
			} else {
				cartes1 = append(cartes1, player1, player0)
				if len(cartes0) == 0 {
					return guanyador
				}
			}

		} else {
			// Normal game
			if player0 > player1 {
				cartes0 = append(cartes0, player0, player1)
				if len(cartes1) == 0 {
					return 0
				}
			} else {
				cartes1 = append(cartes1, player1, player0)
				if len(cartes0) == 0 {
					return 1
				}
			}
		}
		cartes[0] = cartes0
		cartes[1] = cartes1
	}

}

func playCombatRecursive(cartes [][]int) int {
	winner := cartes[recursive(cartes, 0)]
	return sumaPunts(winner)
}

func main() {
	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1 := playCombat(linies)
	fmt.Println("Part 1: ", correctes1)

	correctes2 := playCombatRecursive(linies)
	fmt.Println("Part 2", correctes2)
}
