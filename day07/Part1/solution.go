package main

import (
	"fmt"
	"os"
	"sort"

	// "sort"
	"strconv"
	"strings"
	// "math"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type File struct {
	filename string
	folder bool
	contains_files []string
	size int
}

func Contains(list []string, x string) bool {
	for _, item := range list {
		if item == x {
			return true
		}
	}
	return false
}

func search_file_tree(start string, files map[string][]string, sizes map[string]int, total_sizes map[string]int, already_searched []string) (map[string]int, []string) {
	total_sizes[start] = sizes[start]
	// Deal with possible infinite loops that somehow exist
	if Contains(already_searched, start) {
		return total_sizes, already_searched
	}
	already_searched = append(already_searched, start)
	for _, elem := range files[start] {
		fmt.Println(elem)
		_, exists := files[elem]
		if exists {
			total_sizes, already_searched = search_file_tree(elem, files, sizes, total_sizes, already_searched)
			total_sizes[start] += sizes[elem]
		}
	}

	return total_sizes, already_searched
}

func main() {
	dat, err := os.ReadFile("./../dat")
	// dat, err := os.ReadFile("./../example_dat")
	check(err)
	// fmt.Println(string(dat))

	// s := strings.Fields(string(dat[:]))
	s := strings.Split(string(dat), "\n")
	// fmt.Println(s)
	// var count int = 0
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
		// fmt.Println(files, sizes)
		// fmt.Println(count)
	}
	fmt.Println(sizes, files)
	total_sizes := make(map[string]int)
	var already_searched []string
	fmt.Println("Starting to search tree")
	total_sizes, already_searched = search_file_tree("/", files, sizes, total_sizes, already_searched)
	fmt.Println(total_sizes, already_searched)

	keys := make([]string, 0, len(total_sizes))

	for key := range total_sizes {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return total_sizes[keys[i]] < total_sizes[keys[j]]
	})

	fmt.Println(keys)

	count := 0

	for _, key := range keys {
		fmt.Println(key)
		value := total_sizes[key]
		if value > 100000 {
			continue
		}
		count += value
	}
	fmt.Println(count)

	// 748557

}
