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
	// dat, err := os.ReadFile("./../dat")
	dat, err := os.ReadFile("./../example_dat")
	check(err)
	// fmt.Println(string(dat))

	s := strings.Fields(string(dat[:]))
	// fmt.Println(s)
	var count int = 0
	// var last int = math.MaxInt64
	for index, element := range s {
		// intVal, _ := strconv.Atoi(element)
		// if (intVal> last) {
		// 	count += 1
		// }
		// last = intVal
		fmt.Println(index, element)
	}
	fmt.Println(count)

}
