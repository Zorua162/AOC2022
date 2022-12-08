package main

import (
	"fmt"
	"os"
	// "strings"
	// "strconv"
	// "math"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func contains_duplicates(input string) bool {
	char_count_map := make(map[string]int)
	for _, char := range input {
		val, inMap := char_count_map[string(char)]
		if inMap {
			val += 1
		} else {
			char_count_map[string(char)] = 1
		}

		if val >= 2 {
			return true
		}
	}
	fmt.Println(char_count_map)
	return false
}

func main() {
	dat, err := os.ReadFile("./../dat")
	// dat, err := os.ReadFile("./../example_dat")
	check(err)
	// fmt.Println(string(dat))

	// s := strings.Fields(string(dat[:]))
	s := string(dat)
	fmt.Println(s)
	msg_len := 14
	var count int = (msg_len - 1)
	// var last int = math.MaxInt64
	for i := 0; i < len(s)-(msg_len-1); i = i + 1 {
		// intVal, _ := strconv.Atoi(element)
		// if (intVal> last) {
		count += 1
		// }
		// last = intVal
		// fmt.Println(string(s[i]), string(s[i+1]), string(s[i+2]), string(s[i+3]))
		if !contains_duplicates(s[i : i+msg_len]) {
			fmt.Println(count)
			break
		}
	}
	fmt.Println(count)

}
