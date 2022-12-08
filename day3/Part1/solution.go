package main

import (
	"fmt"
	"os"
	"strings"
	// "math"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func find_duplicate_letters(list1 string, list2 string) string {
	var out_runes []rune

	for _, rune_comp := range list2 {
		if strings.Contains(list1, string(rune_comp)) {
			out_runes = append(out_runes, rune_comp)
		}
	}

	return string(out_runes)
}

func get_letter_priority(char rune) int {
	var val int
	if char < 96 {
		// Its a caps
		val = int(char - 38)
		// fmt.Println(string(char), strconv.Itoa(val))
	} else {
		val = int(char - 96)
		// fmt.Println(string(char), strconv.Itoa(val))
	}
	return val
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

	for i := 0; i < len(s)-1; i = i + 3 {
		// fmt.Println(s[i])
		// fmt.Println(s[i + 1])
		// fmt.Println(s[i + 2])
		// fmt.Println("----")
		dupes := find_duplicate_letters(s[i], s[i+1])
		dupes = find_duplicate_letters(dupes, s[i+2])
		// fmt.Println(string(dupes[0]))
		count += get_letter_priority(rune(dupes[0]))

	}

	fmt.Println(count)

	// fmt.Println([]rune("A")[0] - 38)
	// fmt.Println(([]rune("a")[0]) - 96)

}
