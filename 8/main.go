package main

import (
	"fmt"
	"math"
	"slices"
	"os"
	"maps"
	"bufio"
	"strings"
	"strconv"
	"sort"
)

type Coord struct {
	x float64
	y float64
	z float64

	edges *[]Coord
}

type Circuit struct {
	startingNode Coord
	size int
}

func parse_inputs(filename string) ([]Coord, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	coords := []Coord{}
	for scanner.Scan() {
		line := scanner.Text()
		coord := Coord{x: -1, y: -1, z: -1, edges: &[]Coord{}}
		for i, val := range strings.Split(line, ",") {
			intval, err := strconv.ParseFloat(val, 64); if err != nil { panic(err) }
			if i == 0 {
				coord.x = intval
			} else if i == 1 {
				coord.y = intval
			} else if i == 2 {
				coord.z = intval
			} else {
				panic("what?")
			}
		}

		if coord.x == -1 || coord.y == -1 || coord.z == -1 {
			panic("what?")
		}

		coords = append(coords, coord)
	}

	return coords, nil
}

func absoluteDistance(a Coord, b Coord) float64 {
	return math.Sqrt( math.Pow(a.x - b.x, 2) + math.Pow(a.y - b.y, 2) + math.Pow(a.z - b.z, 2))
}

func by_distance(coords []Coord) map[float64][][]Coord {
	// Migrate to https://pkg.go.dev/sort
	byDistance := map[float64][][]Coord{}
	for _, a := range coords {
		for _, b := range coords {
			if a == b {
				continue
			}

			distance := absoluteDistance(a, b)
			containsOpposite := slices.ContainsFunc(byDistance[distance], func(coord []Coord) bool {
				sliceA := coord[0]
				sliceB := coord[1]

				return sliceA == b && sliceB == a
			})

			if !containsOpposite {
				byDistance[distance] = append(byDistance[distance], []Coord{a, b})
			}
		}
	}

	return byDistance
}

func get_ten_closest(coords []Coord) [][]Coord {
	distanceMap := by_distance(coords)
	tenClosest := [][]Coord{}
	keysSeq := maps.Keys(distanceMap)
	keys := []float64{}
	for key := range keysSeq {
		keys = append(keys, key)
	}

	slices.Sort(keys)
	for len(tenClosest) < 1000 {
	// for len(tenClosest) < 10 {
		closest := keys[0]

		fmt.Println(closest, distanceMap[closest][0])

		tenClosest = append(tenClosest, distanceMap[closest][0])
		keys = keys[1:]
	}

	fmt.Println()

	return tenClosest
}

func isolateCircuits(nodes map[Coord]bool) []Circuit {
	already := map[Coord]bool{}

	circuits := []Circuit{}

	for n := range maps.Keys(nodes) {
		if already[n] {
			continue
		} else {
			already[n] = true
		}

		circuit := Circuit{startingNode: n, size: 0}

		next := []Coord{n}
		for len(next) > 0 {
			cur := next[0]
			next = next[1:]

			circuit.size++

			for _, edge := range *cur.edges {
				if !already[edge] {
					next = append(next, edge)
					already[edge] = true
				}
			}
		}

		circuits = append(circuits, circuit)
	}

	return circuits
}

type Circuits []Circuit
type BySize struct{ Circuits }
func (s BySize) Less(i, j int) bool { return s.Circuits[i].size < s.Circuits[j].size }
func (s Circuits) Len() int      { return len(s) }
func (s Circuits) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func solve_part1(coords []Coord) int {
	tenClosest := get_ten_closest(coords)
	nodes := map[Coord]bool{}
	for _, coordPair := range tenClosest {
		a, b := coordPair[0], coordPair[1]
		*a.edges = append(*a.edges, b)
		*b.edges = append(*b.edges, a)

		nodes[a] = true
		nodes[b] = true
	}

	circuits := isolateCircuits(nodes)
	sort.Sort(sort.Reverse(BySize{circuits}))

	result := 1
	for i, c := range circuits {
		result *= c.size

		if i == 2 {
			break
		}
	}

	return result
}

func main() {
	items, err := parse_inputs(os.Args[1])
	if err != nil {
		panic(err)
	}

	part1 := solve_part1(items)
	
	fmt.Printf("Result %d\n", part1)
}
