package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Object interface {
	getSize() int
	getName() string
}

type File struct {
	name string
	size int
}

type Directory struct {
	name       string
	size       int
	contains   []Object
	parent_dir *Directory
}

func (f File) getSize() int {
	return f.size
}

func (d Directory) getSize() int {
	return d.size
}

func (f File) getName() string {
	return f.name
}

func (d Directory) getName() string {
	return d.name
}

func replace_current(parent Directory, current Directory) Directory {

	for i, dir := range parent.contains {
		if dir.getName() == current.getName() {
			parent.contains[i] = current
			return parent
		}
	}
	parent.contains = append(parent.contains, current)
	return parent
}

func (dir Directory) calculate_size() Directory {
	for _, file := range dir.contains {
		dir.size += file.getSize()
	}
	return dir
}

func change_dir(current_dir Directory, new_location string, small_dirs []Directory) (Directory, []Directory) {
	if new_location == ".." {
		fmt.Println("Changing to parent dir")
		fmt.Println(current_dir)
		current_dir = current_dir.calculate_size()

		if current_dir.size <= 100000 {
			small_dirs = append(small_dirs, current_dir)
		}

		parent_dir := *current_dir.parent_dir
		parent_dir = replace_current(parent_dir, current_dir)

		return parent_dir, small_dirs
	}
	new_dir := Directory{new_location, 0, make([]Object, 0), &current_dir}
	current_dir.add_dir(new_dir)
	return new_dir, small_dirs
}

func go_to_top(dir Directory, small_dirs []Directory) Directory {
	for dir.getName() != "/" {
		fmt.Println("going to top" + dir.getName())
		dir, small_dirs = change_dir(dir, "..", small_dirs)
	}
	return dir
}

func (current_dir Directory) add_file(name string, size int) Directory {
	new_file := File{name, size}
	current_dir.contains = append(current_dir.contains, new_file)
	return current_dir
}

func (dir Directory) add_dir(add_dir Directory) Directory {
	dir.contains = append(dir.contains, add_dir)
	return dir
}

func main() {
	dat, err := os.ReadFile("./../dat")
	// dat, err := os.ReadFile("./../example_dat")
	check(err)

	s := strings.Split(string(dat), "\n")

	current_dir := Directory{"home", 0, make([]Object, 0), nil}
	small_dirs := make([]Directory, 0)

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
				current_dir, small_dirs = change_dir(current_dir, components[2], small_dirs)
			}

		} else {
			// Its a file or a folder
			// files[cwd] = append(files[cwd], components[1])
			if components[0] != "dir" {
				intVal, _ := strconv.Atoi(components[0])
				current_dir = current_dir.add_file(components[1], intVal)
				fmt.Println(current_dir)
				// sizes[cwd] += intVal
			}
		}
	}
	fmt.Println(current_dir)
	// fmt.Println(*current_dir.parent_dir)
	// fmt.Println(*current_dir.parent_dir)

	current_dir = go_to_top(current_dir, small_dirs)
	current_dir = current_dir.calculate_size()
	fmt.Println(current_dir)

	total := 0
	for _, dir := range small_dirs {
		total += dir.size
		fmt.Println(dir.getName() + " " + strconv.Itoa(dir.getSize()))
	}
	fmt.Println(total)

	// 1148367 too low
	// 1534236 guess too low
	// 1148367

}
