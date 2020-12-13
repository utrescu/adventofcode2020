package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func stringToInt(str string) (uint64, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.ParseUint(nonFractionalPart[0], 10, 64)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (uint64, []autobus, error) {
	file, err := os.Open(path)
	if err != nil {
		return ^uint64(0), nil, err
	}
	defer file.Close()

	var lines []autobus
	scanner := bufio.NewScanner(file)

	var timestamp uint64
	i := 0
	for scanner.Scan() {
		if i == 0 {
			timestamp, _ = stringToInt(scanner.Text())
			i++
		} else {
			for index, text := range strings.Split(scanner.Text(), ",") {
				number, err := stringToInt(text)
				if err == nil {
					lines = append(lines, autobus{number, uint64(index)})
				}
			}
		}
	}
	return timestamp, lines, scanner.Err()
}

type autobus struct {
	id     uint64
	offset uint64
}

func firstBus(timestamp uint64, busos []autobus) uint64 {

	bestBus := ^uint64(0)
	timeWait := ^uint64(0)

	for _, bus := range busos {
		nextAt := ((timestamp / bus.id) + 1) * bus.id
		wait := nextAt - timestamp
		if wait < timeWait {
			bestBus = bus.id
			timeWait = wait
		}
	}
	return bestBus * timeWait
}

func firstBus2(busos []autobus) uint64 {

	sort.Slice(busos, func(i, j int) bool {
		return busos[i].id > busos[j].id
	})

	timestamp := uint64(0)
	timestampstep := uint64(1)
	for _, bus := range busos {
		for (timestamp+bus.offset)%bus.id != 0 {
			timestamp += timestampstep
		}
		timestampstep *= bus.id
	}
	return timestamp

}

func main() {
	timestamp, busos, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1 := firstBus(timestamp, busos)
	fmt.Println("Part 1: ", correctes1)

	correctes2 := firstBus2(busos)
	fmt.Println("Part 2: ", correctes2)
}
