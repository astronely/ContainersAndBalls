package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in *bufio.Reader
var out *bufio.Writer

// solveWithSort solve "Containers with balls" problem with sort arrays
func solveWithSort() string {
	var n, x int
	fmt.Fscan(in, &n)

	// Arrays with const size & capacity
	sumX, sumY := make([]int, n), make([]int, n)

	// Input data
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &x)
			sumX[i] += x
			sumY[j] += x
		}
	}

	// Calculate result
	sort.Ints(sumX)
	sort.Ints(sumY)

	for i := 0; i < n; i++ {
		if sumX[i] != sumY[i] {
			return "no"
		}
	}
	return "yes"
}

// solveWithMap solve "Containers with balls" problem with map, without sort
func solveWithMap() string {
	var n, x int
	fmt.Fscan(in, &n)

	// Arrays with const size & capacity
	sumX, sumY := make([]int, n), make([]int, n)

	// Input data
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &x)
			sumX[i] += x
			sumY[j] += x
		}
	}

	// Calculate result
	ballsSumCounter := make(map[int]int)
	for i := 0; i < n; i++ {
		ballsSumCounter[sumX[i]]++
		ballsSumCounter[sumY[i]]--
	}

	for _, val := range ballsSumCounter {
		if val != 0 {
			return "no"
		}
	}

	return "yes"
}

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	//fmt.Fprint(out, solveWithSort())
	fmt.Fprint(out, solveWithMap())
}
