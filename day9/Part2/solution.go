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
}

func (head Head) getLocation() Location {
	return head.location
}

type Tail struct {
	location           Location
	previous_locations []Location
}

func (tail Tail) getLocation() Location {
	return tail.location
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

func move_tail_right(object Object, tail Tail) Tail {

	if object.getLocation().x-2 == tail.getLocation().x {
		tail.location.x++

		if object.getLocation().y != tail.location.y {
			tail.location.y = object.getLocation().y
		}

		tail = tail.add_location()
	}

	return tail
}

func move_right(head Head, tail_list []Tail, amount int) (Head, []Tail) {
	for i := 0; i < amount; i++ {
		head.location.x++
		tail_list[0] = move_tail_right(head, tail_list[0])
		for i := range tail_list[1:] {
			tail_list[i+1] = move_tail_right(tail_list[i], tail_list[i+1])
		}
		// fmt.Println(head, tail_list)
		print_data(head, tail_list)
	}

	return head, tail_list
}

func move_tail_left(object Object, tail Tail) Tail {
	if object.getLocation().x+2 == tail.location.x {
		tail.location.x--

		if object.getLocation().y != tail.location.y {
			tail.location.y = object.getLocation().y
		}
		tail = tail.add_location()
	}
	return tail
}

func move_left(head Head, tail_list []Tail, amount int) (Head, []Tail) {
	for i := 0; i < amount; i++ {
		head.location.x--
		tail_list[0] = move_tail_left(head, tail_list[0])
		for i := range tail_list[1:] {
			tail_list[i+1] = move_tail_left(tail_list[i], tail_list[i+1])
		}
		print_data(head, tail_list)
	}
	return head, tail_list
}

func move_tail_up(object Object, tail Tail) Tail {
	if object.getLocation().y-2 == tail.location.y {
		tail.location.y++

		if object.getLocation().x != tail.location.x {
			tail.location.x = object.getLocation().x
		}

		tail = tail.add_location()
	}
	if tail.location.x == object.getLocation().x-3 {
		tail.location.x = object.getLocation().x - 1
		tail.location.y = object.getLocation().y
	}
	return tail
}

func move_up(head Head, tail_list []Tail, amount int) (Head, []Tail) {
	for i := 0; i < amount; i++ {
		head.location.y++
		tail_list[0] = move_tail_up(head, tail_list[0])
		for i := range tail_list[1:] {
			tail_list[i+1] = move_tail_up(tail_list[i], tail_list[i+1])
		}
		print_data(head, tail_list)
	}
	return head, tail_list
}

func move_tail_down(object Object, tail Tail) Tail {
	if object.getLocation().y+2 == tail.location.y {
		tail.location.y--
		if object.getLocation().x != tail.location.x {
			tail.location.x = object.getLocation().x
		}
		tail = tail.add_location()
	}
	return tail
}

func move_down(head Head, tail_list []Tail, amount int) (Head, []Tail) {
	for i := 0; i < amount; i++ {
		head.location.y--
		tail_list[0] = move_tail_down(head, tail_list[0])
		for i := range tail_list[1:] {
			tail_list[i+1] = move_tail_down(tail_list[i], tail_list[i+1])
		}
		print_data(head, tail_list)
	}

	return head, tail_list
}

func move_head(head Head, tail_list []Tail, direction string, amount string) (Head, []Tail) {
	amount_int, _ := strconv.Atoi(amount)
	if direction == "R" {
		head, tail_list = move_right(head, tail_list, amount_int)
	}
	if direction == "L" {
		head, tail_list = move_left(head, tail_list, amount_int)
	}
	if direction == "U" {
		head, tail_list = move_up(head, tail_list, amount_int)
	}
	if direction == "D" {
		head, tail_list = move_down(head, tail_list, amount_int)
	}
	return head, tail_list
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

func add_tails(locations []Location, tails []Tail) []Location {
	for _, tail := range tails {
		for _, location := range tail.previous_locations {
			locations = append(locations, location)
		}
	}
	return locations
}

func print_board_from_visited(head Head, tail Tail, tail_list []Tail, f *os.File) {
	visited := tail.previous_locations
	visited_and_head := append(visited, head.location)
	visited_and_head = add_tails(visited_and_head, tail_list)
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

	for i, tail := range tail_list {
		board = set_pos(board, translate_pos(min_x, tail.location.x), translate_pos(min_y, tail.location.y), strconv.Itoa(i+1))
	}

	for _, pos := range visited {
		board = set_pos(board, translate_pos(min_x, pos.x), translate_pos(min_y, pos.y), "#")
	}

	board = set_pos(board, translate_pos(min_x, 0), translate_pos(min_y, 0), "S")

	board = set_pos(board, translate_pos(min_x, tail.location.x), translate_pos(min_y, tail.location.y), "T")

	board = set_pos(board, translate_pos(min_x, head.location.x), translate_pos(min_y, head.location.y), "H")


	for _, line := range board {
		// fmt.Println(line)
		f.WriteString(strings.Join(line, "") + "\n")
	}

}

func print_data(head Head, tail_list []Tail) {
	var out_str string
	out_str += head.location.to_string()
	for _, tail := range tail_list {
		out_str += "; "
		out_str += tail.location.to_string()
	}
	fmt.Println(out_str)
}

func main() {
	// dat, err := os.ReadFile("./../smol_dat")
	// dat, err := os.ReadFile("./../dat")
	dat, err := os.ReadFile("./../test_dat")
	// dat, err := os.ReadFile("./../example_dat")
	// dat, err := os.ReadFile("./../med_exa_dat")

	check(err)
	// fmt.Println(string(dat))

	s := strings.Fields(string(dat[:]))
	// fmt.Println(s)
	var count int = 0

	head := Head{Location{0, 0}}
	tail_list := make([]Tail, 0)
	num_tails := 9
	for i := 0; i < num_tails; i++ {
		tail := Tail{Location{0, 0}, make([]Location, 0)}
		tail.previous_locations = append(tail.previous_locations, tail.location)
		tail_list = append(tail_list, tail)
	}

	// var last int = math.MaxInt64
	var last_tail Tail


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
		head, tail_list = move_head(head, tail_list, s[i], s[i+1])
		// fmt.Println(head, tail_list)
		// print_data(head, tail_list)
		last_tail = tail_list[len(tail_list)-1]
		print_board_from_visited(head, last_tail, tail_list, f)
	}
	count = len(last_tail.previous_locations)
	fmt.Println(count)

}
