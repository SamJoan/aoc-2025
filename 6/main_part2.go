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

func get_lines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}

func get_next_chunk_size(lines []string, offset int) int {
	remainingLastLine := lines[len(lines) - 1][offset + 1:]
	result := -1
	for nb, char := range remainingLastLine {
		if string(char) != " " {
			result = nb
			break
		}
	}

	if result == -1 {
		result = len(remainingLastLine) + 1
	}

	return result
}

func split_chunks(lines []string) [][]string {
	chunks := [][]string{}
	for _, line := range lines {
		lineArr := []string{}

		remaining := line
		originalLen := len(remaining)
		for {
			offset := originalLen - len(remaining)
			chunkSize := get_next_chunk_size(lines, offset)

			chunk := remaining[:chunkSize]
			lineArr = append(lineArr, chunk)
			if len(remaining) - chunkSize == 0 {
				break
			}
			remaining = remaining[chunkSize + 1:]
		}

		chunks = append(chunks, lineArr)
	}

	return chunks
}

func parse_inputs(filename string) ([]Problem, error) {
	lines, err := get_lines(filename)
	if err != nil {
		return nil, err
	}

	chunks := split_chunks(lines)

	problems := []Problem{}
	for x := 0; x <= len(chunks[0]) - 1; x++ {
		numbersRaw := []string{}
		problem := Problem{}
		for y := 0; y <= len(chunks) - 1; y++ {
			last := y + 1 == len(chunks)
			col := chunks[y][x]
			if last {
				problem.operator = strings.TrimSpace(col)
			} else {
				numbersRaw = append(numbersRaw, col)
			}
		}

		problem.numbers = parseCephalopod(numbersRaw)
		problems = append(problems, problem)
	}

	return problems, nil
}

func parseCephalopod(numbersRaw []string) []int64 {
	fmt.Println(numbersRaw)
	numbers := []int64{}
	digitSize := len(numbersRaw[0]) - 1
	for i := digitSize; i >= 0; i-- {
		numberStr := ""
		for _, digit := range numbersRaw {
			char := string(digit[i])
			if char != " " {
				numberStr += char
			}
		}

		nb, err := atoi_64(numberStr); if err != nil { panic(err) }
		numbers = append(numbers, nb)
	}

	return numbers
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
	
	fmt.Printf("Result PART 2 %d\n", part1)
}
