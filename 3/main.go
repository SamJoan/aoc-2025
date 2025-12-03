package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func parse_inputs(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	banks := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		banks = append(banks, line)
	}

	return banks, nil
}

func index(bank string) map[int][]int {
	i := make(map[int][]int)

	for pos, char := range bank {
		val, err := strconv.Atoi(string(char)); if err != nil { panic(err) }
		i[val] = append(i[val], pos)
	}

	return i
}

func solve(bank map[int][]int, bankSize int, desiredLen int) int {
	solution := []int{}
	lastPos := -1
	for i := 0; i < desiredLen; i++ {
		foundSuitableDigit := false
		for highestNum := 9; highestNum >= 1; highestNum-- {
			positions := bank[highestNum]
			if len(positions) > 0 {
				for _, pos := range positions {
					lastPossiblePos := bankSize - (desiredLen - len(solution)) 
					// fmt.Printf("nb: %d, pos: %d, lastPos: %d, lastPossiblePos: %d, bankSize: %d\n", highestNum, pos, lastPos, lastPossiblePos, bankSize)
					if pos > lastPos && pos <= lastPossiblePos {
						solution = append(solution, highestNum)
						foundSuitableDigit = true
						lastPos = pos
						break
					}
				}
			}

			if foundSuitableDigit {
				break
			}
		}
	}

	solutionStr := ""
	for _, i := range solution {
		solutionStr += strconv.Itoa(i)
	}
	
	numericValue, err := strconv.Atoi(solutionStr); if err != nil {panic(err)}

	return numericValue
}

func solve_part1(banks []string) int {
	total := 0
	for _, bank := range banks {
		bankIndex := index(bank)
		solution := solve(bankIndex, len(bank), 2)
		total += solution
	}

	return total
}

func solve_part2(banks []string) int {
	total := 0
	for _, bank := range banks {
		bankIndex := index(bank)
		solution := solve(bankIndex, len(bank), 12)
		total += solution
	}

	return total
}

func main() {
	items, err := parse_inputs(os.Args[1])
	if err != nil {
		panic(err)
	}

	part1 := solve_part1(items)
	part2 := solve_part2(items)
	
	fmt.Printf("Result %d\n", part1)
	fmt.Printf("Result %d\n", part2)
}
