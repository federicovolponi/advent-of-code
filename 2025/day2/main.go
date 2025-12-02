package main

import (
	//"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func isInvalidProductPt1(product string) bool {
	length := len(product)
	if length % 2 != 0 {
		return false
	}
	
	mid := length / 2
	for i := range mid {
		if product[i] != product[mid + i] {
			return false
		}
	}
	return true
}

func isInvalidProductPt2(product string) bool {
	length := len(product)

	for s := 1; s < (length / 2) + 1; s++ {
		chunk := product[:s]
		comp := strings.Repeat(chunk, length / s)
		if comp == product {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	res_pt1 := 0
	res_pt2 := 0
	for rawRange := range strings.SplitSeq(string(file), ",") {
		bounds := strings.Split(rawRange, "-")

		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(strings.TrimRight(bounds[1], "\n"))
		if err != nil {
			log.Fatal(err)
		}

		for product := start; product <= end; product++ {
			if isInvalidProductPt1(strconv.Itoa(product)) {
				res_pt1 += product
			}
			if isInvalidProductPt2(strconv.Itoa(product)) {
				res_pt2 += product
			}
		}
	}
	fmt.Println("1st Part: Sum of invalid products is ", res_pt1)
	fmt.Println("2st Part: Sum of invalid products is ", res_pt2)

}
