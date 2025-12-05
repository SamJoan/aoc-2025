package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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


type Range struct {
	start int64
	end int64
}

type InventorySystem struct {
	ranges []Range
	items []int64
}

func parse_inputs(filename string) (*InventorySystem, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	ranges := []Range{}
	items := []int64{}

	parsingRanges := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			splat := strings.Split(line, "-")
			start, err := atoi_64(splat[0]); if err != nil {panic(err)}
			end, err := atoi_64(splat[1]); if err != nil {panic(err)}
			r := Range{start: start, end: end}

			ranges = append(ranges, r)
		} else {
			number, err := atoi_64(line); if err != nil {panic(err)}
			items = append(items, number)
		}
	}

	inventorySystem := InventorySystem{ranges: ranges, items: items}

	return &inventorySystem, nil
}

func (is InventorySystem) isFresh(i int64) bool {
	for _, r := range is.ranges {
		if i >= r.start && i <= r.end {
			return true
		}
	}

	return false
}

func solve_part1(is *InventorySystem) int {
	nbFresh := 0
	for _, item := range is.items {
		if is.isFresh(item) {
			nbFresh += 1
		}
	}

	return nbFresh
}

func (is InventorySystem) mergeIntervals() []Range {
	sort.Slice(is.ranges, func(i, j int) bool {
		return is.ranges[i].start < is.ranges[j].start
	})

	newRanges := []Range{}
	curr := Range{start: -1, end: -1}
	for _, r := range is.ranges {
		if curr.start == -1 {
			curr = r
			continue
		}

		if curr.end < r.start {
			newRanges = append(newRanges, curr)
			curr = r
		} else {
			if curr.end < r.end {
				curr.end = r.end
			}
		}
	}

	newRanges = append(newRanges, curr)

	return newRanges
}

func solve_part2(is *InventorySystem) int64 {
	newRanges := is.mergeIntervals()

	var total int64 = 0
	for _, r := range newRanges {
		total += r.end - r.start + 1
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
