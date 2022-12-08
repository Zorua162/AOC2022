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

func main() {
	// dat, err := os.ReadFile("./../dat")
	dat, err := os.ReadFile("./../example_dat")
	check(err)
	// fmt.Println(string(dat))

	// s := strings.Fields(string(dat[:]))
	s := strings.Split(string(dat), "\n")
	// fmt.Println(s)
	var count int = 0
	// var last int = math.MaxInt64
	// the key is the folder name and the value is a list of sub folders or files
	files := make(map[string][]string, 1)
	sizes := make(map[string]int, 1)
	var cwd string
	for _, element := range s {
		components := strings.Fields(string(element))
		if string(element[0]) == "$" {
			// fmt.Println("cmd")
			if components[1] == "cd" {
				// fmt.Println("cd")
				cwd = components[2]
			}

		} else {
			// Its a file or a folder
			files[cwd] = append(files[cwd], components[1])
			intVal, _ := strconv.Atoi(components[0])
			sizes[cwd] += intVal
		}

		// intVal, _ := strconv.Atoi(element)
		// if (intVal> last) {
		// 	count += 1
		// }
		// last = intVal
		// fmt.Println(index, element, components)
		fmt.Println(files, sizes)
		fmt.Println(count)
	}

}
