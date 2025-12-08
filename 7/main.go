package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func parse_inputs(filename string) ([][]bool, []int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil,  err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	m := [][]bool{}
	var coords []int = nil

	y := 0
	for scanner.Scan() {
		line := scanner.Text()

		lineInt := []bool{}
		for x, c := range line {
			char := string(c)
			if char == "^" {
				lineInt = append(lineInt, true)
			} else {
				lineInt = append(lineInt, false)
				if char == "S" {
					coords = []int{x, y}
				}
			}
		}

		m = append(m, lineInt)

		y++
	}

	return m, coords, nil
}

func solve_part1(m [][]bool, startCoords []int) int {
	queue := [][]int{}
	queue = append(queue, startCoords)

	alreadyVisited := map[string]bool{}

	splitTimes := 0
	for len(queue) > 0 {
		coords := queue[0]
		queue = queue[1:]

		x, y := coords[0], coords[1]
		key := fmt.Sprintf("%d,%d", x, y)
		if alreadyVisited[key] {
			continue
		}

		alreadyVisited[key] = true

		if y >= len(m) - 1 {
			continue
		}

		elem := m[y][x]
		if elem == false {
			queue = append(queue, []int{x, y + 1})
		} else {
			splitTimes += 1
			queue = append(queue, []int{x + 1, y})
			queue = append(queue, []int{x - 1, y})
		}
	}

	return splitTimes
}

func k(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func p(key string) (int, int, string) {
	splat := strings.Split(key, ",")

	x, err := strconv.Atoi(splat[0]); if err != nil { panic(err) }
	y, err := strconv.Atoi(splat[1]); if err != nil { panic(err) }

	return x, y, splat[2]
}

func solve_part2(m *[][]bool, cache *map[string]int, x int, y int) int {
	cache_key := k(x, y)
	val, ok := (*cache)[cache_key] 
	if ok {
		return val
	}
	
	if y > len(*m) - 1 {
		return 1
	}

	result := 0
	if (*m)[y][x] {
		result += solve_part2(m, cache, x + 1, y)
		result += solve_part2(m, cache, x - 1, y)
	} else {
		result += solve_part2(m, cache, x, y + 1)
	}

	(*cache)[cache_key] = result

	return result
}

func main() {
	m, coords, err := parse_inputs(os.Args[1])

	if err != nil {
		panic(err)
	}

	cache := map[string]int{}

	part1 := solve_part1(m, coords)
	part2 := solve_part2(&m, &cache, coords[0], coords[1])
	
	fmt.Printf("Result %d\n", part1)
	fmt.Printf("Result %d\n", part2)
}
