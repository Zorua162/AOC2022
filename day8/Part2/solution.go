package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	// "math"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func check_left(x, y int, array [][]int) int {
	score := 0
	for i := x - 1; i >= 0; i -= 1 {
		score += 1
		if array[y][i] >= array[y][x] {
			return score
		}
	}
	return score
}

func check_right(x, y int, array [][]int) int {
	score := 0
	for i := x + 1; i < len(array[0]); i += 1 {
		score += 1
		if array[y][i] >= array[y][x] {
			return score
		}
	}
	return score
}

func check_top(x, y int, array [][]int) int {
	score := 0
	for i := y - 1; i >= 0; i -= 1 {
		score += 1
		if array[i][x] >= array[y][x] {
			return score
		}
	}
	return score
}

func check_bottom(x, y int, array [][]int) int {
	score := 0
	for i := y + 1; i < len(array[0]); i += 1 {
		score += 1
		if array[i][x] >= array[y][x] {
			return score
		}
	}
	return score
}

func get_score(x, y int, array [][]int) int {

	score := check_left(x, y, array)
	score *= check_right(x, y, array)
	score *= check_top(x, y, array)
	score *= check_bottom(x, y, array)

	return score
}

func main() {
	dat, err := os.ReadFile("./../dat")
	// dat, err := os.ReadFile("./../example_dat")
	check(err)
	// fmt.Println(string(dat))

	s := strings.Fields(string(dat[:]))
	// fmt.Println(s)
	// var last int = math.MaxInt64
	array := [][]int{}
	scores := [][]int{}
	for index, element := range s {
		array = append(array, make([]int, 0))
		scores = append(scores, make([]int, 0))
		// intVal, _ := strconv.Atoi(element)
		// if (intVal> last) {
		// 	count += 1
		// }
		// last = intVal
		for _, char := range element {
			intVal, _ := strconv.Atoi(string(char))
			array[index] = append(array[index], intVal)
			scores[index] = append(scores[index], 0)
		}
	}

	// Check each location, treat it individually and look at each position
	for i, elem := range array {
		for j := range elem {
			score := get_score(j, i, array)
			scores[i][j] = score
		}
	}

	largest := 0
	// Count up all non visible elements
	for _, elem := range scores {
		for _, score := range elem {
			if score > largest {
				largest = score
			}

		}
	}

	fmt.Println(array)

	fmt.Println("output")
	fmt.Println(largest)
	// 126 too low

}
