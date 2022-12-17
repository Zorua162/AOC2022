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
		if location == tail.location {
			return tail
		}
	}
	tail.previous_locations = append(tail.previous_locations, tail.location)
	return tail
}

func move_tail_right(head Head, tail Tail) Tail {
	if head.location == tail.location {
		return tail
	}
	fmt.Println(head.location.x-2, tail.location.x)
	if head.location.x-2 == tail.location.x {
		tail.location.x++
		tail = tail.add_location()
	}
	if head.location.x+1 == tail.location.x {
		// Tail is ahead of the head so the tail doesn't move
		return tail
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

func move_head(head Head, tail Tail, direction string, amount string) (Head, Tail) {
	amount_int, _ := strconv.Atoi(amount)
	if direction == "R" {
		head, tail = move_right(head, tail, amount_int)
	}
	if direction == "L" {
		head, tail = move_left(head, tail, amount_int)
	}
	return head, tail

}

func main() {
	// dat, err := os.ReadFile("./../dat")
	dat, err := os.ReadFile("./../example_dat")
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
	}
	fmt.Println(count)

}
