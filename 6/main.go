package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"
)

func atoi_64(numberString string) (int64, error) {
	i, err := strconv.ParseInt(numberString, 10, 64)
	if err != nil {
		return -1, err
	}

	return i, nil
}

func itoa_64(nb int64) string {
	return strconv.FormatInt(nb, 10)
}

type Problem struct {
	numbers []int64
	operator string
}

func parse_inputs(filename string) ([]Problem, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		lineArr := []string{}
		for _, col := range strings.Split(line, " ") {
			if col != "" {
				lineArr = append(lineArr, col)
			}
		}

		lines = append(lines, lineArr)
	}

	problems := []Problem{}
	for x := 0; x <= len(lines[0]) - 1; x++ {
		numbers := []int64{}
		problem := Problem{}
		for y := 0; y <= len(lines) - 1; y++ {
			last := y + 1 == len(lines)
			col := lines[y][x]
			if last {
				problem.operator = col
			} else {
				nb, err := atoi_64(col); if err != nil { panic(err) }
				numbers = append(numbers, nb)
			}
		}

		problem.numbers = numbers
		problems = append(problems, problem)
	}

	return problems, nil
}

func (p Problem) solve() int64 {
	var total int64 = -1
	for _, nb := range p.numbers {
		if total == -1 {
			total = nb
			continue
		}

		if p.operator  == "*" {
			total *= nb
		} else if p.operator == "+" {
			total += nb
		} else {
			panic("what?")
		}
	}

	return total
}

func solve_part1(problems []Problem) int64 {
	var total int64 = 0
	for _, problem := range problems {
		result := problem.solve()
		fmt.Printf("%v: %d\n", problem, result)
		total += result
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
