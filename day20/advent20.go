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

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (map[int]casella, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make(map[int]casella)
	var (
		id       int
		caselles [][]string
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Tile ") {
			// nova casella
			line2 := strings.Replace(line, "Tile ", "", -1)

			caselles = make([][]string, 0)
			id, _ = stringToInt(line2[:len(line2)-1])
		} else {
			if line == "" {
				// Calcular costats
				top := ""
				botom := ""
				left := ""
				right := ""
				for column := 0; column < len(caselles[0]); column++ {
					top = top + caselles[0][column]
					botom = botom + caselles[len(caselles)-1][column]
				}
				for row := 0; row < len(caselles); row++ {
					left = left + caselles[row][0]
					right = right + caselles[row][len(caselles)-1]
				}
				costats := make([]string, 0)
				costats = append(costats, top)
				costats = append(costats, left)
				costats = append(costats, right)
				costats = append(costats, botom)

				costats = append(costats, reverse(top))
				costats = append(costats, reverse(left))
				costats = append(costats, reverse(right))
				costats = append(costats, reverse(botom))

				lines[id] = casella{costats, caselles}
			} else {

				caracters := make([]string, 0)
				for _, c := range line {
					caracters = append(caracters, string(c))
				}
				caselles = append(caselles, caracters)
			}
		}
	}
	top := ""
	botom := ""
	left := ""
	right := ""
	for column := 0; column < len(caselles[0]); column++ {
		top = top + caselles[0][column]
		botom = botom + caselles[len(caselles)-1][column]
	}
	for row := 0; row < len(caselles); row++ {
		left = left + caselles[row][0]
		right = right + caselles[row][len(caselles)-1]
	}
	costats := make([]string, 0)
	costats = append(costats, top)
	costats = append(costats, left)
	costats = append(costats, right)
	costats = append(costats, botom)

	costats = append(costats, reverse(top))
	costats = append(costats, reverse(left))
	costats = append(costats, reverse(right))
	costats = append(costats, reverse(botom))

	lines[id] = casella{costats, caselles}

	return lines, scanner.Err()
}

type casella struct {
	costats []string
	casella [][]string
}

func (c casella) contains(paret string) bool {
	for _, costat := range c.costats {
		if costat == paret {
			return true
		}
	}
	return false
}

func (c casella) containsOther(other casella) bool {
	for _, paret := range c.costats {
		if other.contains(paret) {
			return true
		}
	}
	return false
}

func anyMatches2(matches []int) bool {
	for _, value := range matches {
		if value == 2 {
			return true
		}
	}
	return false
}

func processaCaselles(caselles map[int]casella) int {

	ids := make(map[int]int)

	for id, actual := range caselles {
		if id == 2311 {
			fmt.Println("2311")
		}
		for id2, other := range caselles {
			if id != id2 {

				if other.containsOther(actual) {
					count, ok := ids[id]
					if ok {
						ids[id] = count + 1
					} else {
						ids[id] = 1
					}
				}
			}
		}
	}

	resultat := 1
	for k, v := range ids {
		if v == 2 {
			resultat = resultat * k
		}
	}

	return resultat

}

func main() {
	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1 := processaCaselles(linies)
	fmt.Println("Part 1: ", correctes1)

}
