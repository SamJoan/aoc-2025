package main

import (
	"fmt"
	"os"
	"bufio"
)

type Range struct {
	start int64
	end int64
}

func parse_inputs(filename string) ([]Range, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// rotations := []Rotation{}
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}

	return nil, nil
}

func solve_part1(rotations []Range) int {
	return 0
}

func main() {
	items, err := parse_inputs(os.Args[1])
	if err != nil {
		panic(err)
	}

	part1 := solve_part1(items)
	
	fmt.Printf("Result %d\n", part1)
}
