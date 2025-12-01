package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func mod(a int, b int) int {
	return (a%b + b) % b
}

func firstProblem(n_zeros *int, dial int, rotation int, orientation byte) int {
	var offset int
	if orientation == 'L' {
		offset = dial - rotation
	} else {
		offset = dial + rotation
	}

	if dial == 0 {
		*n_zeros++
	}
	return offset
}

func secondProblem(n_zeros *int, dial int, rotation int, orientation byte) int {
    var offset int

    if orientation == 'R' {
        k := (100 - dial) % 100
        if k == 0 {
            k = 100
        }
        if k <= rotation {
            *n_zeros += 1 + (rotation-k)/100
        }

        offset = dial + rotation
    } else {
        k := dial % 100
        if k == 0 {
            k = 100
        }
        if k <= rotation {
            *n_zeros += 1 + (rotation-k)/100
        }

        offset = dial - rotation
    }
    return offset
}

func parseInput(line string) (orientation byte, rotation int) {
	rotation, err := strconv.Atoi(line[1:])
	if err != nil {
		log.Fatal(err)
	}
	orientation = line[0]
	return

}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	dial := 50
	n_zeros_first := 0
	n_zeros_second := 0
	for scanner.Scan() {
		line := scanner.Text()
		orientation, rotation := parseInput(line)

		offset_first := firstProblem(&n_zeros_first, dial, rotation, orientation)
		offset_second := secondProblem(&n_zeros_second, dial, rotation, orientation)
		if offset_first != offset_second {
			panic("The two offset should be equal!")
		}
		dial = mod(offset_first, 100)
	}

	fmt.Printf("First solution: %d\nSecond solution: %d\n", n_zeros_first, n_zeros_second)
}
