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

func solve_part1(banks []string) int {
	total := 0
	for _, bank := range banks {
		bankLen := len(bank)
		firstValue := -1
		firstValuePos := -1
		for targetNumber := 9; targetNumber >= 1; targetNumber-- {
			// -1 for array offset, -1 for the fact that we can't have the first char be at the last spot.
			for pos:= 0; pos <= bankLen - 2; pos++ {
				valueInBank, err := strconv.Atoi(string(bank[pos])); if err != nil { panic(err) }
				if valueInBank == targetNumber {
					firstValue = valueInBank
					firstValuePos = pos
					break
				}
			}

			if firstValue != -1 {
				break
			}
		}

		secondValue := -1
		for targetNumber := 9; targetNumber >= 1; targetNumber-- {
			for pos := bankLen - 1; pos > firstValuePos; pos-- {
				valueInBank, err := strconv.Atoi(string(bank[pos])); if err != nil { panic(err) }
				if valueInBank == targetNumber {
					secondValue = valueInBank
					break
				}
			}

			if secondValue != -1 {
				break
			}
		}

		firstValueStr := strconv.Itoa(firstValue)
		secondValueStr := strconv.Itoa(secondValue)

		jolts, err := strconv.Atoi(firstValueStr + secondValueStr); if err != nil { panic(err) }

		total += jolts
	}

	return total
}

func main() {
	items, err := parse_inputs(os.Args[1])
	if err != nil {
		panic(err)
	}

	part1 := solve_part1(items)
	
	fmt.Printf("Result %d\n", part1)
}
