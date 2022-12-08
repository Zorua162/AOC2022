package main

import (
	"fmt"
	"os"
	"strings"
	// "strconv"
	// "math"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	/*
		A, B, C Opponent plays rock, paper or scissors
		X, Y, Z Player plays rock, paper or scissors
		X = L, Y = D, Z = W
	*/

	choice := map[string]map[string]string{"A": {"X": "Z", "Y": "X", "Z": "Y"},
		"B": {"X": "X", "Y": "Y", "Z": "Z"},
		"C": {"X": "Y", "Y": "Z", "Z": "X"}}

	// score_played := map[string]int{"X": 1, "Y": 2, "Z": 3}
	// score_result := map[string]int{"L": 0, "D": 3, "W": 6}
	output := map[string]map[string]int{"A": {}, "B": {}, "C": {}}
	output["A"]["X"] = 4
	output["A"]["Y"] = 8
	output["A"]["Z"] = 3
	output["B"]["X"] = 1
	output["B"]["Y"] = 5
	output["B"]["Z"] = 9
	output["C"]["X"] = 7
	output["C"]["Y"] = 2
	output["C"]["Z"] = 6

	fmt.Println(output)

	// fmt.Println(scores)

	dat, err := os.ReadFile("./../dat")
	// dat, err := os.ReadFile("./../example_dat")
	check(err)
	// fmt.Println(string(dat))

	s := strings.Fields(string(dat[:]))
	// fmt.Println(s)
	var count int = 0
	// var last int = math.MaxInt64
	// for index, element := range s {
	for i := 0; i < len(s)-1; i = i + 2 {

		s[i+1] = choice[s[i]][s[i+1]]

		// count += score_played[s[i+1]]
		count += output[s[i]][s[i+1]]
		// intVal, _ := strconv.Atoi(element)
		// if (intVal> last) {
		// 	count += 1
		// }
		// last = intVal
		fmt.Println(s[i], s[i+1])
	}

	fmt.Println(count)

}
