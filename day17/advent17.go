package main

import (
	"bufio"
	"fmt"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (energycube, error) {
	result := energycube{posicions: map[point]rune{}}

	file, err := os.Open(path)
	if err != nil {
		return energycube{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, caracter := range line {
			result.posicions[point{x, y, 0, 0}] = caracter
		}
		result.maxY = y
		result.maxX = len(line)
		y++
	}
	return result, scanner.Err()
}

type point struct {
	x, y, z, w int
}

type energycube struct {
	es4d      bool
	posicions map[point]rune

	minX int
	maxX int

	minY int
	maxY int

	minZ int
	maxZ int

	minW int
	maxW int
}

func (d *energycube) obtenirCasella(x, y, z, w int) rune {
	if t, ok := d.posicions[point{x, y, z, w}]; ok {
		return t
	}

	return '.'
}

func (d *energycube) casellaActiva(x int, y int, z int, w int) int {

	if d.obtenirCasella(x, y, z, w) == '#' {
		return 1
	}
	return 0
}

func (d *energycube) comptaVeins(x, y, z, w int) int {
	suma := 0

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if d.es4d {
					// Part 2
					for dw := -1; dw <= 1; dw++ {
						if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
							continue
						}
						suma += d.casellaActiva(x+dx, y+dy, z+dz, w+dw)
					}
				} else {
					// Part 1
					if dx == 0 && dy == 0 && dz == 0 {
						continue
					}
					suma += d.casellaActiva(x+dx, y+dy, z+dz, 0)
				}
			}
		}
	}

	return suma
}

func (d *energycube) modificaCasella(x int, y int, z int, w int, numVeins int, valorActual rune) {
	switch {
	case valorActual == '#' && (numVeins == 2 || numVeins == 3):
		d.posicions[point{x, y, z, w}] = '#'
	case valorActual == '.' && numVeins == 3:
		d.posicions[point{x, y, z, w}] = '#'
	default:
		d.posicions[point{x, y, z, w}] = '.'
	}
}

func (d *energycube) step() {
	newWorld := &energycube{posicions: map[point]rune{}}

	for x := d.minX - 1; x <= d.maxX+1; x++ {
		for y := d.minY - 1; y <= d.maxY+1; y++ {
			for z := d.minZ - 1; z <= d.maxZ+1; z++ {
				// Part 2
				if d.es4d {
					for w := d.minW - 1; w <= d.maxW+1; w++ {
						newWorld.modificaCasella(x, y, z, w, d.comptaVeins(x, y, z, w), d.obtenirCasella(x, y, z, w))
					}
				} else {
					// Part 1
					newWorld.modificaCasella(x, y, z, 0, d.comptaVeins(x, y, z, 0), d.obtenirCasella(x, y, z, 0))
				}
			}
		}
	}

	// Restaurar canvis
	d.posicions = newWorld.posicions
	d.minX--
	d.minY--
	d.minZ--
	d.maxX++
	d.maxY++
	d.maxZ++

	d.minW--
	d.maxW++
}

func (d *energycube) comptaActives() int {
	count := 0
	for _, t := range d.posicions {
		if t == '#' {
			count++
		}
	}

	return count
}

func processaMapa(world energycube) int {
	pas := 0
	for pas < 6 {
		world.step()
		pas++
	}

	return world.comptaActives()
}

func main() {
	world, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	world.es4d = false
	correctes1 := processaMapa(world)
	fmt.Println("Part 1: ", correctes1)

	world.es4d = true
	correctes2 := processaMapa(world)
	fmt.Println("Part 2: ", correctes2)

}
