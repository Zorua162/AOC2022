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

// func RemoveIndex(s []string, index int) []string{
//     ret := make([]string, 0)
//     ret = append(ret, s[:index]...)
//     return append(ret, s[index+1:]...)
// }

func move_crates(amount int, from int, to int, storage [][]string) [][]string {
	// Move the given amount from index (from/to-1) to given location

	for i := 0; i < amount; i = i + 1 {
		storage[to-1] = append(storage[to-1], storage[from-1][len(storage[from-1])])
		storage[from-1] = storage[from-1][:len(storage[from-1])-1]
	}
	return storage
}

func main() {
	// fmt.Println(string(dat))
	// fmt.Println(string(dat[54:]))
	// fmt.Println(string(dat[:]))

	dat, err := os.ReadFile("./../example_dat")

	example_input := make([][]string, 3)
	example_input[0] = []string{"Z", "N"}
	example_input[1] = []string{"M", "C", "D"}
	example_input[2] = []string{"P"}

	storage := example_input

	/*

		    [G]         [P]         [M]
		    [V]     [M] [W] [S]     [Q]
		    [N]     [N] [G] [H]     [T] [F]
		    [J]     [W] [V] [Q] [W] [F] [P]
		[C] [H]     [T] [T] [G] [B] [Z] [B]
		[S] [W] [S] [L] [F] [B] [P] [C] [H]
		[G] [M] [Q] [S] [Z] [T] [J] [D] [S]
		[B] [T] [M] [B] [J] [C] [T] [G] [N]

	*/

	// dat, err := os.ReadFile("./../dat")
	// input := [][]string{{"B", "G", "S", "C"},
	// 	{"T", "M", "W", "H", "J", "N", "V", "G"},
	// 	{"M", "Q", "S"},
	// 	{"B", "S", "L", "T", "W", "N", "M"},
	// 	{"J", "Z", "F", "T", "V", "G", "W", "P"},
	// 	{"C", "T", "B", "G", "Q", "H", "S"},
	// 	{"T", "J", "P", "B", "W"},
	// 	{"G", "D", "C", "Z", "F", "T", "Q", "M"},
	// 	{"N", "S", "H", "B", "P", "F"}}

	// storage := input

	// s := strings.Fields(string(dat[:]))
	check(err)
	s := strings.Split(string(dat[:]), "\n")

	for _, element := range s {
		fmt.Println(element)
	}

	for _, element := range s[10:] {
		// intVal, _ := strconv.Atoi(element)
		// if (intVal> last) {
		// 	count += 1
		// }
		// last = intVal
		// fmt.Println(index, element)
		split_element := strings.Fields(element)
		// fmt.Println(split_element)
		amount, _ := strconv.Atoi(string(split_element[1]))
		from, _ := strconv.Atoi(string(split_element[3]))
		to, _ := strconv.Atoi(string(split_element[5]))

		fmt.Println(amount, from, to)

		storage = move_crates(amount, from, to, storage)

		// CGSMPSWMF
	}
	fmt.Println(storage)

}
