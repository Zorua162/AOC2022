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

type Object interface {
	getLocation() Location
}

type Head struct {
	location Location
	tail     *Tail
}

func (head Head) getLocation() Location {
	return head.location
}

func (tail Tail) getLocation() Location {
	return tail.location
}

type Tail struct {
	location           Location
	previous_locations []Location
	tail               *Tail
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

func move_to_object(object Object, tail Tail) Tail {
	// Object can be a head or a tail

	// If the head is ever two steps directly up, down, left, or right from the tail,
	// the tail must also move one step in that direction so it remains close enough.
	// Otherwise, if the head and tail aren't touching and aren't in the same row or
	// column, the tail always moves one step diagonally to keep up:

	// Check if two steps directly up or down and move to it
	x_diff := object.getLocation().x - tail.location.x
	y_diff := object.getLocation().y - tail.location.y

	// Horizontal direct move
	// if abs(x_diff) == 2 && y_diff == 0 {
	// 	// Directly left or right
	// 	if x_diff > 0 {
	// 		// Positive so the tail is to the left and it needs to go right one
	// 		tail.location.x++
	// 	} else {
	// 		tail.location.x--
	// 	}
	// }
	// // Vertical direct move
	// if abs(y_diff) == 2 && x_diff == 0 {
	// 	if y_diff > 0 {
	// 		tail.location.y++
	// 	} else {
	// 		tail.location.y--
	// 	}
	// }

	// Figure out the direction that it needs to move diagonally
	var y_add int
	if y_diff > 0 && abs(y_diff) == 2 {
		y_add = 1
	} else {
		y_add = -1
	}
	var x_add int
	if x_diff > 0 && abs(x_diff) == 2 {
		x_add = 1
	} else {
		x_add = -1
	}

	tail.location.x = tail.location.x + x_add
	tail.location.y = tail.location.y + y_add

	if tail.tail != nil {
		tails_tail := move_to_object(tail, *tail.tail)
		tail.tail = &tails_tail
	}
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

func move_head(head Head, direction string) Head {
	fmt.Println(direction)
	if direction == "R" {
		head.location.x++
		fmt.Println("Moved right")
	}
	if direction == "L" {
		head.location.x--
	}
	if direction == "U" {
		head.location.y++
	}
	if direction == "D" {
		head.location.y--
	}
	return head
}

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

func (loc Location) to_string() string {
	out_str := ""
	out_str += strconv.Itoa(loc.x) + ", " + strconv.Itoa(loc.y)
	return out_str
}

func add_tails(locations []Location, tail *Tail) []Location {
	if tail.tail != nil {
		locations = add_tails(locations, tail.tail)
	}

	for _, location := range tail.previous_locations {
		locations = append(locations, location)
	}
	return locations
}

func add_tail_pos(board [][]string, tail *Tail, val int) [][]string {
	current_tail := *tail
	board = set_pos(board, current_tail.location.x, current_tail.location.y, strconv.Itoa(val))
	if tail.tail != nil {
		val--
		board = add_tail_pos(board, tail.tail, val)
	}
	return board
}

func print_board_from_visited(head Head, tail Tail, f *os.File) *os.File {
	visited := tail.previous_locations
	visited_and_head := append(visited, head.location)
	visited_and_head = add_tails(visited_and_head, &tail)
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

	board = add_tail_pos(board, &tail, 9)

	for _, pos := range visited {
		board = set_pos(board, translate_pos(min_x, pos.x), translate_pos(min_y, pos.y), "#")
	}
	board = set_pos(board, translate_pos(min_x, 0), translate_pos(min_y, 0), "S")

	// board = set_pos(board, translate_pos(min_x, tail.location.x), translate_pos(min_y, tail.location.y), "T")

	// for i, tail := range tail_list {
	// 	board = set_pos(board, translate_pos(min_x, tail.location.x), translate_pos(min_y, tail.location.y), strconv.Itoa(i+1))
	// }

	board = set_pos(board, translate_pos(min_x, head.location.x), translate_pos(min_y, head.location.y), "H")

	for _, line := range board {
		// fmt.Println(line)
		f.WriteString(strings.Join(line, "") + "\n")
	}
	f.WriteString("--------\n\n")
	return f

}

func main() {
	dat, err := os.ReadFile("./../smol_dat")
	// dat, err := os.ReadFile("./../dat")
	// dat, err := os.ReadFile("./../test_dat")
	// dat, err := os.ReadFile("./../example_dat")
	check(err)
	// fmt.Println(string(dat))

	s := strings.Fields(string(dat[:]))
	// fmt.Println(s)
	var count int = 0

	end_tail := Tail{Location{0, 0}, make([]Location, 0), nil}

	num_tails := 9
	var tail Tail
	for i := 0; i < num_tails-1; i++ {
		tail = Tail{Location{0, 0}, make([]Location, 0), &end_tail}
		tail.previous_locations = append(tail.previous_locations, tail.location)
	}
	head := Head{Location{0, 0}, &tail}

	f, err := os.Create("out.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	for i := 0; i < len(s); i += 2 {
		// intVal, _ := strconv.Atoi(element)
		// if (intVal> last) {
		// 	count += 1
		// }
		// last = intVal
		fmt.Println(s[i], s[i+1])
		amount, _ := strconv.Atoi(s[i+1])
		f.WriteString("\n" + s[i] + s[i+1] + "\n")
		for j := 0; j < amount; j++ {
			head = move_head(head, s[i])
			fmt.Println(head)
			tail = move_to_object(head, tail)
			print_board_from_visited(head, tail, f)
			fmt.Println(head, tail)
		}
	}
	count = len(tail.previous_locations)
	fmt.Println(count)

}
