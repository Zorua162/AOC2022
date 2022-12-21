package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Location struct {
	x int
	y int
}

type Head struct {
	location Location
}

type Tail struct {
	location           Location
	previous_locations []Location
}

func (tail Tail) add_location() Tail {
	for _, location := range tail.previous_locations {
		// if location == tail.location {
		if location.x == tail.location.x && location.y == tail.location.y {
			return tail
		}
	}
	tail.previous_locations = append(tail.previous_locations, tail.location)
	return tail
}

func move_tail_right(head Head, tail Tail) Tail {

	if head.location.x-2 == tail.location.x {
		tail.location.x++

		if head.location.y != tail.location.y {
			tail.location.y = head.location.y
		}

		tail = tail.add_location()
	}

	return tail
}

func move_right(head Head, tail Tail, amount int) (Head, Tail) {
	for i := 0; i < amount; i++ {
		head.location.x++
		tail = move_tail_right(head, tail)
		fmt.Println(head, tail)
	}

	return head, tail
}

func move_tail_left(head Head, tail Tail) Tail {
	if head.location.x+2 == tail.location.x {
		tail.location.x--

		if head.location.y != tail.location.y {
			tail.location.y = head.location.y
		}

		tail = tail.add_location()
	}
	return tail
}

func move_left(head Head, tail Tail, amount int) (Head, Tail) {
	for i := 0; i < amount; i++ {
		head.location.x--

		tail = move_tail_left(head, tail)
	}
	return head, tail
}

func move_tail_up(head Head, tail Tail) Tail {
	if head.location.y-2 == tail.location.y {
		tail.location.y++

		if head.location.x != tail.location.x {
			tail.location.x = head.location.x
		}

		tail = tail.add_location()
	}
	return tail
}

func move_up(head Head, tail Tail, amount int) (Head, Tail) {
	for i := 0; i < amount; i++ {
		head.location.y++
		tail = move_tail_up(head, tail)
	}
	return head, tail
}

func move_tail_down(head Head, tail Tail) Tail {
	if head.location.y+2 == tail.location.y {
		tail.location.y--
		if head.location.x != tail.location.x {
			tail.location.x = head.location.x
		}
		tail = tail.add_location()
	}
	return tail
}

func move_down(head Head, tail Tail, amount int) (Head, Tail) {
	for i := 0; i < amount; i++ {
		head.location.y--
		tail = move_tail_down(head, tail)
	}

	return head, tail
}

func move_head(head Head, tail Tail, direction string, amount string) (Head, Tail) {
	amount_int, _ := strconv.Atoi(amount)
	if direction == "R" {
		head, tail = move_right(head, tail, amount_int)
	}
	if direction == "L" {
		head, tail = move_left(head, tail, amount_int)
	}
	if direction == "U" {
		head, tail = move_up(head, tail, amount_int)
	}
	if direction == "D" {
		head, tail = move_down(head, tail, amount_int)
	}
	return head, tail
}

// func create_board(size_x, size_y int) [][]string {
// 	board := make([][]string, size_y)
// 	for i := range board {
// 		board[i] = make([]string, size_x)
// 		for j := range board[i] {
// 			board[i][j] = "."
// 		}
// 	}
//
// 	return board
// }
//
// func set_pos_old(board [][]string, x, y int, set_to string, shift_x, shift_y int) [][]string {
// 	fmt.Println(len(board)-1-(y+shift_y), x+shift_x)
// 	board[len(board)-2-(y+shift_y)][x+shift_x+1] = set_to
// 	return board
// }
//
// func set_head(board [][]string, head Head, shift_x, shift_y int) [][]string {
// 	set_pos(board, head.location.x, head.location.y, "H", shift_x, shift_y)
// 	return board
// }
//
// func set_tail(board [][]string, tail Tail, shift_x, shift_y int) [][]string {
// 	set_pos(board, tail.location.x, tail.location.y, "T", shift_x, shift_y)
// 	return board
// }
//
// func compare(head, tail int) (int, int) {
// 	if head > tail {
// 		return head, tail
// 	}
// 	return tail, head
// }
//
// func get_current_size(head Head, tail Tail) (int, int, int, int) {
// 	var x int
// 	var y int
// 	shift_x := 0
// 	shift_y := 0
//
// 	maxx, minx := compare(head.location.x, tail.location.x)
// 	maxy, miny := compare(head.location.x, tail.location.x)
//
// 	// if its small then make it 5 instead
// 	if maxx < 5 {
// 		x = 4
// 	}
// 	if maxy < 5 {
// 		y = 4
// 	}
//
// 	if minx < 0 {
// 		shift_x = minx
// 	}
//
// 	if miny < 0 {
// 		shift_y = miny
// 	}
//
// 	// Add one to fix for size to index
// 	x++
// 	y++
//
// 	return x, y, shift_x, shift_y
// }
//
// func print_board(head Head, tail Tail) {
// 	// size_x, size_y, shift_x, shift_y := get_current_size(head, tail)
//
// 	// board := create_board(size_x+shift_x+1, size_y+shift_y+1)
//
// 	// for _, loc := range tail.previous_locations {
// 	// 	set_pos(board, loc.x, loc.y, "#", shift_x, shift_y)
// 	// }
// 	// set_pos(board, shift_x, shift_y, "S", shift_x, shift_y)
//
// 	// board = set_tail(board, tail, shift_x, shift_y)
// 	// board = set_head(board, head, shift_x, shift_y)
//
// 	// fmt.Println(board)
// 	// for _, line := range board {
// 	// 	fmt.Println(line)
// 	// }
// 	fmt.Println("---")
// }

func get_exteme_vals(vals []int) (int, int) {
	min := math.MaxInt32
	max := -math.MaxInt32

	for _, val := range vals {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	return min, max
}

func get_extreme_x(locations []Location) (int, int) {
	var values []int
	for _, loc := range locations {
		values = append(values, loc.x)
	}
	return get_exteme_vals(values)
}

func get_extreme_y(locations []Location) (int, int) {
	var values []int
	for _, loc := range locations {
		values = append(values, loc.y)
	}
	return get_exteme_vals(values)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func translate_pos(min, pos int) int {
	// if +ve or 0 then add abs(min_x)
	// if -ve then abs(min_x) + pos
	if min < 0 {
		if pos < 0 {
			return abs(min) + pos
		}
		return pos + abs(min)
	}
	return pos - min
}

func set_pos(board [][]string, x, y int, val string) [][]string {
	board[len(board)-y-1][x] = val
	return board
}

func get_range(min, max int) int {
	if min < 0 {
		return max + abs(min) + 1 // + 1 is for 0, 0
	}
	return max - min + 1
}

func print_board_from_visited(head Head, tail Tail) {
	visited := tail.previous_locations
	visited_and_head := append(visited, head.location)
	min_x, max_x := get_extreme_x(visited_and_head)
	min_y, max_y := get_extreme_y(visited_and_head)
	// Create the board based on the difference between mins and maxes
	range_x := get_range(min_x, max_x)
	range_y := get_range(min_y, max_y)

	fmt.Println(min_x, max_x)
	board := make([][]string, 0)

	for i := 0; i < range_y; i++ {
		board = append(board, make([]string, 0))
		for j := 0; j < range_x; j++ {
			board[i] = append(board[i], ".")
		}
	}

	for _, pos := range visited {
		board = set_pos(board, translate_pos(min_x, pos.x), translate_pos(min_y, pos.y), "#")
	}

	board = set_pos(board, translate_pos(min_x, 0), translate_pos(min_y, 0), "S")

	board = set_pos(board, translate_pos(min_x, tail.location.x), translate_pos(min_y, tail.location.y), "T")

	board = set_pos(board, translate_pos(min_x, head.location.x), translate_pos(min_y, head.location.y), "H")

	f, err := os.Create("out.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	for _, line := range board {
		// fmt.Println(line)
		f.WriteString(strings.Join(line, "") + "\n")
	}

}

func main() {
	// dat, err := os.ReadFile("./../smol_dat")
	dat, err := os.ReadFile("./../dat")
	// dat, err := os.ReadFile("./../test_dat")
	// dat, err := os.ReadFile("./../example_dat")
	check(err)
	// fmt.Println(string(dat))

	s := strings.Fields(string(dat[:]))
	// fmt.Println(s)
	var count int = 0

	head := Head{Location{0, 0}}
	tail := Tail{Location{0, 0}, make([]Location, 0)}
	tail.previous_locations = append(tail.previous_locations, tail.location)
	// var last int = math.MaxInt64
	for i := 0; i < len(s); i += 2 {
		// intVal, _ := strconv.Atoi(element)
		// if (intVal> last) {
		// 	count += 1
		// }
		// last = intVal
		fmt.Println(s[i], s[i+1])
		head, tail = move_head(head, tail, s[i], s[i+1])
		fmt.Println(head, tail)
		print_board_from_visited(head, tail)
	}
	count = len(tail.previous_locations)
	fmt.Println(count)

	// 5079 too low

}
