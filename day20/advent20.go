package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// COSTATS de les peces
const COSTATS = 4

// TOP nom del costat
const TOP = 0

// LEFT nom del costat
const LEFT = 1

// RIGHT nom del costat
const RIGHT = 2

// BOTTOM nom del costat
const BOTTOM = 3

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
				lines[id] = casella{id, getWalls(caselles), caselles}
			} else {

				caracters := make([]string, 0)
				for _, c := range line {
					caracters = append(caracters, string(c))
				}
				caselles = append(caselles, caracters)
			}
		}
	}

	lines[id] = casella{id, getWalls(caselles), caselles}

	return lines, scanner.Err()
}

type casella struct {
	id       int
	costats  []string
	caselles [][]string
}

func rotate(c [][]string) [][]string {

	columnallarg := len(c)
	newContent := make([][]string, columnallarg)

	for y := 0; y < columnallarg; y++ {
		newContent[y] = make([]string, len(c[y]))
		filallarg := len(c[y])
		for x := 0; x < filallarg; x++ {
			newContent[y][x] = c[x][filallarg-1-y]
		}
	}
	return newContent
}

func reverseSlice(c []string) []string {
	llarg := len(c)
	next := make([]string, llarg)
	for i, v := range c {
		next[llarg-1-i] = v
	}
	return next
}

func getWalls(caselles [][]string) []string {
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
	return []string{
		top,
		left,
		right,
		botom,
	}

}

func (c casella) rotateTile() casella {
	novesCaselles := rotate(c.caselles)

	return casella{
		id:       c.id,
		costats:  getWalls(novesCaselles),
		caselles: novesCaselles,
	}
}

func (c casella) reverseTile() casella {
	novesCaselles := make([][]string, 0)
	for _, linia := range c.caselles {
		novesCaselles = append(novesCaselles, reverseSlice(linia))
	}
	return casella{
		id:       c.id,
		costats:  getWalls(novesCaselles),
		caselles: novesCaselles,
	}
}

func (c casella) getTile() []string {
	files := make([]string, 0)

	for i, fila := range c.caselles {
		linia := ""
		if i != 0 || i != len(c.caselles)-1 {
			for j, valor := range fila {
				if j != 0 && j != len(fila)-1 {
					linia += valor
				}
			}
		}
		files = append(files, linia)
	}
	return files
}

func (c casella) contains(paret string) bool {
	for _, costat := range c.costats {
		if costat == paret || reverse(costat) == paret {
			return true
		}
	}
	return false
}

func (c casella) containsOther(other casella) bool {
	for _, paret := range c.costats {
		if other.contains(paret) || other.contains(reverse(paret)) {
			return true
		}
	}

	return false
}

func (c casella) otherFitsOn(other casella) (int, error) {
	for index, paret := range c.costats {
		if other.contains(paret) {
			return index % COSTATS, nil
		}
	}
	return -1, errors.New("No quadren")
}

// --- PART 1
func processaCaselles(caselles map[int]casella) ([]int, map[int][]casella) {

	ids := make(map[int][]casella)

	for id, actual := range caselles {
		for id2, other := range caselles {
			if id != id2 {

				if other.containsOther(actual) {
					_, ok := ids[id]
					if ok {
						ids[id] = append(ids[id], other)
					} else {
						ids[id] = make([]casella, 1)
						ids[id][0] = other
					}
				}
			}
		}
	}

	resultat := make([]int, 0)
	for k, v := range ids {
		if len(v) == 2 {
			resultat = append(resultat, k)
		}
	}

	return resultat, ids

}

func searchMonster(caselles map[int]casella, casellesAdjacents map[int][]casella) int {

	// Localitzar top, left (0,0)
	topleft := -1
	for actual, adjacents := range casellesAdjacents {
		dret := caselles[actual].costats[RIGHT]
		baix := caselles[actual].costats[BOTTOM]
		if len(adjacents) == 2 {

			if adjacents[0].contains(dret) && adjacents[1].contains(baix) ||
				adjacents[0].contains(baix) && adjacents[1].contains(dret) {
				topleft = actual
			}
		}
	}
	fmt.Println(topleft)
	// Per cada adjacent posar-les on toca (rota, flip)
	// mapa := composeMap(resultat)
	// comprovar amb expressió regular?

	/// Monster
	// #
	// #    ##    ##    ###
	//  #  #  #  #  #  #

	return 0
}

func main() {
	linies, err := readLines("inputtest")
	if err != nil {
		panic("File read failed")
	}

	resultat, adjacents := processaCaselles(linies)
	correctes1 := 1
	for _, value := range resultat {
		correctes1 *= value
	}
	fmt.Println("Part 1: ", correctes1)

	correctes2 := searchMonster(linies, adjacents)
	fmt.Println("Part 2: ", correctes2)

}
