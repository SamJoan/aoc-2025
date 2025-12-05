package main

import (
	"fmt"
	"os"
	"bufio"
)

func blankRow(length int) []bool {
	row := []bool{}
	for i := 0; i < length + 2; i++ {
		row = append(row, false)
	}

	return row
}

func parse_inputs(filename string) ([][]bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	diagram := [][]bool{}

	lineLen := -1
	for scanner.Scan() {
		line := scanner.Text()

		if len(diagram) == 0 {
			lineLen = len(line)
			diagram = append(diagram, blankRow(lineLen))
		}

		row := []bool{}

		row = append(row, false)
		for _, c := range line {
			val := false
			if string(c) == "@" {
				val = true
			}

			row = append(row, val)
		}
		row = append(row, false)

		diagram = append(diagram, row)		
	}

	diagram = append(diagram, blankRow(lineLen))

	return diagram, nil
}

func countTp(array [][]bool, x int, y int) int {
	result := []bool{array[y-1][x-1], array[y-1][x+1], array[y+1][x-1], array[y+1][x+1], 
		array[y+1][x], array[y-1][x], array[y][x-1], array[y][x+1]}

	count := 0
	for _, toiletPaper := range result {
		if toiletPaper {
			count += 1
		}
	}

	return count
}

func solve_part1(diagram [][]bool) int {
	count := 0
	for y, line := range diagram {
		if y == 0 || y == len(diagram) - 1 {
			continue
		}

		for x := range diagram[y] {
			if x == 0 || x == len(line) - 1 {
				continue
			}

			if diagram[y][x] == true {
				surroundedTpCount := countTp(diagram, x, y)
				if surroundedTpCount < 4 {
					count += 1
				}
			}
		}
	}

	return count;
}

func duplicate_diagram(in [][]bool) [][]bool {
	out := [][]bool{}
	for _, line := range in {
		row := []bool{}
		for _, c := range line {
			row = append(row, c)
		}

		out = append(out, row)
	}

	return out
}

func solve_part2(diagram [][]bool) int {
	count := 0

	newDiagram := duplicate_diagram(diagram)

	for y, line := range diagram {
		if y == 0 || y == len(diagram) - 1 {
			continue
		}

		for x := range diagram[y] {
			if x == 0 || x == len(line) - 1 {
				continue
			}

			if diagram[y][x] == true {
				surroundedTpCount := countTp(diagram, x, y)
				if surroundedTpCount < 4 {
					newDiagram[y][x] = false
					count += 1
				}
			}
		}
	}

	if count > 0 {
		count += solve_part2(newDiagram)
	}

	return count;
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
