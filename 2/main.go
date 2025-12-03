package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Range struct {
	start int64
	end int64
}

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

func parse_inputs(filename string) ([]Range, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	ranges := []Range{}
	for scanner.Scan() {
		line := scanner.Text()
		for _, rangeString := range strings.Split(line, ",") {
			splat := strings.Split(rangeString, "-")
			if len(splat) != 2 {
				return nil, fmt.Errorf("What?")
			}

			start, err := atoi_64(splat[0])
			if err != nil {
				return nil, err
			}

			end, err := atoi_64(splat[1])
			if err != nil {
				return nil, err
			}

			r := Range{start: start, end: end}
			ranges = append(ranges, r)
		}
	}

	return ranges, nil
}

func solve_part1(ranges []Range) int64 {
	var invalidIdSum int64 = 0
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			nbString := itoa_64(i)

			stringLen := len(nbString)
			if stringLen % 2 != 0 {
				continue
			}

			halfLen := stringLen / 2

			firstHalf := nbString[:halfLen]
			secondHalf := nbString[halfLen:]

			if firstHalf == secondHalf {
				invalidIdSum += i
			}
		}
	}

	return invalidIdSum
}

func solve_part2(ranges []Range) int64 {
	var invalidIdSum int64 = 0
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			nbString := itoa_64(i)
			stringLen := len(nbString)
			halfLen := stringLen / 2
			
			isInvalid := false
			for partSize := 1; partSize <= halfLen; partSize++ {
				if stringLen % partSize != 0 {
					continue
				}

				// Split into equal sized parts of a certain size
				parts := []string{}
				for j := 0; j < stringLen / partSize; j++ {
					startIndex := j*partSize
					endIndex := (j+1)*partSize
					parts = append(parts, nbString[startIndex:endIndex])
				}

				allEqual := true
				for _, elem := range parts {
					if elem != parts[0] {
						allEqual = false
						break
					}
				}

				if allEqual {
					isInvalid = true
					break
				}
			}

			if isInvalid {
				invalidIdSum += i
			}
		}
	}

	return invalidIdSum
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
