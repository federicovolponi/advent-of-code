package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"aoc/shared"	
)

func extractBank(bank string) []int {
	res := make([]int, len(bank))
	for i := 0; i < len(bank); i++ {
		res[i] = int(bank[i] - '0')
	}
	return res
}

func findMaxBatteries(bank []int) (int, int) {
	first := 0
	second := 1

	for i := 1; i < len(bank); i++ {
		if bank[i-1] > bank[first] {
			first = i - 1
		}
		if bank[i] > bank[second] {
			second = i
		}
		if i+1 < len(bank) && bank[second] > bank[first] {
			first = second
			second++
		}
	}
	return bank[first], bank[second]
}

func sliceToNumber(digits []int) (int, error) {
	n := 0
	for _, d := range digits {
		if d < 0 || d > 9 {
			return 0, fmt.Errorf("invalid digit %d; must be 0-9", d)
		}
		n = n*10 + d
	}
	return n, nil
}

func findMaxBatteriesStack(bank []int) int {
	stack := ds.NewStack()
	length := len(bank)
	
	for i, battery := range bank {
		if (stack.Len() == 0 || stack.Tail() >= battery) && stack.Len() < 12 {
			stack.Push(battery)
		} else {
			for stack.Len() > 0 && battery > stack.Tail() && length - i > 12 - stack.Len() {
				stack.Pop()
			}

			if stack.Len() < 12 {
				stack.Push(battery)
			}
		}
	}
	res, _ := sliceToNumber(stack.S())
	return res
}

func findMaxBatteriesPt2BF(bank []int) int {
/*		0  1  2  3  4  5  6  7  8  9		
2350
Input: [1, 2, 3, 6, 5, 8, 7, 8, 9, 8] bank 5
Stack: [] -> while greater than last in stack pop and then append
-> I need to make sure to pop at max unil I have space left
[8, 7] I can pop only if len(bank) - currIdx > 5 - len(stack)

while stack:
	if stack empty or stack[-1] <= currNum:
		stack.append(currNum)
	else:
		while currNum > stack[-1] and len(bank) - currIdx > 12 - len(stack):
			stack.pop()
		stack.append(currNum)
*/
	res := make([]int, 0)
	var curr string

	var helper func(int, string)
	helper = func(idx int, curr string) {
		if idx >= len(bank) || len(curr) == 12 {
			// Skip when the batteries set is not complete
			if len(curr) != 12 {
				return
			}
			num, err := strconv.Atoi(curr)
			if err != nil {
				log.Fatal(err)
			}
			res = append(res, num)
			return
		}
		helper(idx + 1, curr)
		curr += strconv.Itoa(bank[idx])
		helper(idx + 1, curr)
		curr = curr[:len(curr) - 1]
	}

	helper(0, curr)
	return slices.Max(res)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	resPt1 := 0
	resPt2 := 0
	for scanner.Scan() {
		bank := extractBank(scanner.Text())
		first, second := findMaxBatteries(bank)
		maxJoltage, err := strconv.Atoi(fmt.Sprintf("%d%d", first, second))
		if err != nil {
			log.Fatal(err)
		}
		
		resPt2 += findMaxBatteriesStack(bank)
		resPt1 += maxJoltage
	}

	fmt.Println("Part 1: The output joltage is ", resPt1)
	fmt.Println("Part 2: The output joltage is ", resPt2)
}
