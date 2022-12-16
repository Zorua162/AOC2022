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

func check_edge(x, y int, array [][]int, visible [][]bool, checked [][]bool) ([][]int, [][]bool, [][]bool) {
	loc_h := array[y][x]
	fmt.Println(loc_h)
	is_visible := false
	// Compare each side to see if this is shorter
	// Check each isn't outside of array first and if it is outside then mark as visible

	// left
	if x == 0 {
		is_visible = true
	}

	// right
	if x == len(array[0])-1 {
		is_visible = true
	}

	// top
	if y == 0 {
		is_visible = true
	}

	// bottom
	if y == len(array)-1 {
		is_visible = true
	}

	visible[y][x] = is_visible
	if is_visible {
		checked[y][x] = true
	}
	return array, visible, checked
}

func check_left(x, y int, array [][]int, visible [][]bool) ([][]int, [][]bool) {
	for i := x - 1; i >= 0; i -= 1 {
		fmt.Println(i)
		if array[y][i] >= array[y][x] {
			return array, visible
		}
	}
	visible[y][x] = true
	return array, visible
}

func check_right(x, y int, array [][]int, visible [][]bool) ([][]int, [][]bool) {
	for i := x + 1; i < len(array[0]); i += 1 {
		fmt.Println(i)
		if array[y][i] >= array[y][x] {
			return array, visible
		}
	}
	visible[y][x] = true
	return array, visible
}

func check_top(x, y int, array [][]int, visible [][]bool) ([][]int, [][]bool) {
	for i := y - 1; i >= 0; i -= 1 {
		fmt.Println(i)
		if array[i][x] >= array[y][x] {
			return array, visible
		}
	}
	visible[y][x] = true
	return array, visible
}

func check_bottom(x, y int, array [][]int, visible [][]bool) ([][]int, [][]bool) {
	for i := y + 1; i < len(array[0]); i += 1 {
		fmt.Println(i)
		if array[i][x] >= array[y][x] {
			return array, visible
		}
	}
	visible[y][x] = true
	return array, visible
}

func check_location(x, y int, array [][]int, visible [][]bool, checked [][]bool) ([][]int, [][]bool, [][]bool) {
	// Skip this location if it is already checked i.e it is already visible
	if checked[y][x] {
		return array, visible, checked
	}

	array, visible = check_left(x, y, array, visible)

	array, visible = check_right(x, y, array, visible)

	array, visible = check_top(x, y, array, visible)

	array, visible = check_bottom(x, y, array, visible)

	// Look to the left of it
	checked[y][x] = true

	return array, visible, checked
}

func main() {
	dat, err := os.ReadFile("./../dat")
	// dat, err := os.ReadFile("./../example_dat")
	check(err)
	// fmt.Println(string(dat))

	s := strings.Fields(string(dat[:]))
	// fmt.Println(s)
	var count int = 0
	// var last int = math.MaxInt64
	array := [][]int{}
	visible := [][]bool{}
	checked := [][]bool{}
	for index, element := range s {
		array = append(array, make([]int, 0))
		visible = append(visible, make([]bool, 0))
		checked = append(checked, make([]bool, 0))
		// intVal, _ := strconv.Atoi(element)
		// if (intVal> last) {
		// 	count += 1
		// }
		// last = intVal
		for _, char := range element {
			intVal, _ := strconv.Atoi(string(char))
			array[index] = append(array[index], intVal)
			visible[index] = append(visible[index], false)
			checked[index] = append(checked[index], false)
		}
	}

	// Get all the edges
	for i, elem := range array {
		for j := range elem {
			array, visible, checked = check_edge(j, i, array, visible, checked)
		}
	}

	// Check each location, treat it individually and look at each position
	for i, elem := range array {
		for j := range elem {
			array, visible, checked = check_location(j, i, array, visible, checked)
		}
	}

	// Count up all non visible elements
	for _, elem := range visible {
		for _, vis := range elem {
			if vis {
				count += 1
			}

		}
	}

	fmt.Println(visible)
	fmt.Println(checked)
	fmt.Println(array)

	fmt.Println("output")
	fmt.Println(count)

}
