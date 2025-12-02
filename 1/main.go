package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

type Rotation struct {
	direction string
	amount int
}

func mod(a, b int) int {
    return (a % b + b) % b
}

func parse_inputs(filename string) ([]Rotation, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	rotations := []Rotation{}
	for scanner.Scan() {
		line := scanner.Text()

		direction := line[:1]
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		rotation := Rotation{direction: direction, amount: amount}
		rotations = append(rotations, rotation)
	}

	return rotations, nil
}

func solve_part1(rotations []Rotation) int {
	value := 50
	zeroCount := 0
	for _, rotation := range rotations {
		if rotation.direction == "L" {
			value -= rotation.amount
		} else {
			value += rotation.amount
		}
		value := mod(value, 100)
		if value == 0 {
			zeroCount += 1
		}
	}

	return zeroCount
}

func solve_part2(rotations []Rotation) int {
	value := 50
	zeroCount := 0
	for _, rotation := range rotations {
		if rotation.direction == "L" {
			if value == 0 {
				value = 100
			}

			value -= rotation.amount
			if value <= 0 {
				shouldAdd := ((value * -1) / 100) + 1
				zeroCount += shouldAdd
			}
		} else {
			value += rotation.amount
			if value >= 100 {
				shouldAdd := (value / 100)
				zeroCount += shouldAdd
			}
		}

		value = mod(value, 100)
	}

	return zeroCount
}

func main() {
	rotations, err := parse_inputs(os.Args[1])
	if err != nil {
		panic(err)
	}

	part1 := solve_part1(rotations)
	part2 := solve_part2(rotations)
	
	fmt.Printf("Result part 1: %d\n", part1)
	fmt.Printf("Result part 2: %d\n", part2)
}
