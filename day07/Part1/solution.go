package main

import (
	"fmt"
	"os"

	// "sort"

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
	filename       string
	folder         bool
	contains_files []File
	parent_dir     *File
	size           int
}

func change_dir(current_dir File, change_to string) File {
	if change_to == ".." {
		// Parent directory of current_dir is now the current_dir
		fmt.Println(current_dir)
		parent_dir := *current_dir.parent_dir
		return parent_dir
	}

	// We're creating a new file instead of this current folder
	new_file := File{change_to, true, make([]File, 0), &current_dir, 0}
	current_dir.contains_files = append(current_dir.contains_files, new_file)
	return new_file

}

func add_file(current_dir File, file_name string, file_size int) File {
	new_file := File{file_name, false, make([]File, 0), &current_dir, file_size}
	current_dir.contains_files = append(current_dir.contains_files, new_file)
	current_dir.size += file_size
	return current_dir
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

func return_to_top(current_folder File) File {
	for current_folder.parent_dir != nil {
		current_folder = change_dir(current_folder, "..")
	}
	// we are now in home folder, so return to /

	return current_folder
}

func main() {
	// dat, err := os.ReadFile("./../dat")
	dat, err := os.ReadFile("./../example_dat")
	check(err)
	// fmt.Println(string(dat))

	// s := strings.Fields(string(dat[:]))
	s := strings.Split(string(dat), "\n")
	// fmt.Println(s)
	// var count int = 0
	// var last int = math.MaxInt64
	// the key is the folder name and the value is a list of sub folders or files
	// files := make(map[string][]string, 1)
	// sizes := make(map[string]int, 1)
	// root_dir := File{
	// filename: "/",
	// folder: true,
	// contains_files: make([]string, 0),
	// size: 0}

	// var current_dir File
	current_dir := File{"/", true, make([]File, 0), nil, 0}

	for _, element := range s {
		fmt.Println("---")
		fmt.Println(element)
		components := strings.Fields(string(element))
		if string(element[0]) == "$" {
			// fmt.Println("cmd")
			if components[1] == "cd" {
				// fmt.Println("cd")
				// cwd = components[2]
				fmt.Println(components[2])
				current_dir = change_dir(current_dir, components[2])
			}

		} else {
			// Its a file or a folder
			// files[cwd] = append(files[cwd], components[1])
			intVal, _ := strconv.Atoi(components[0])
			current_dir = add_file(current_dir, components[1], intVal)
			fmt.Println(current_dir)
			// sizes[cwd] += intVal
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

	fmt.Println(current_dir)

	current_dir = return_to_top(current_dir)
	fmt.Println(current_dir)

	// fmt.Println(sizes, files)
	// total_sizes := make(map[string]int)
	// var already_searched []string
	// fmt.Println("Starting to search tree")
	// total_sizes, already_searched = search_file_tree("/", files, sizes, total_sizes, already_searched)
	// fmt.Println(total_sizes, already_searched)

	// keys := make([]string, 0, len(total_sizes))

	// for key := range total_sizes {
	// 	keys = append(keys, key)
	// }

	// sort.SliceStable(keys, func(i, j int) bool {
	// 	return total_sizes[keys[i]] < total_sizes[keys[j]]
	// })

	// fmt.Println(keys)

	// count := 0

	// for _, key := range keys {
	// 	fmt.Println(key)
	// 	value := total_sizes[key]
	// 	if value > 100000 {
	// 		continue
	// 	}
	// 	count += value
	// }
	// fmt.Println(count)

	// 748557

}
