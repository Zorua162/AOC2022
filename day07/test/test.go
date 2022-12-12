package main

import (
	"fmt"
)

type File struct {
	name   string
	parent *File
	contains []string
}

func (file File) change_dir(name string) File {
	new_file := File{name, &file}
	return new_file

}

func main() {
	fmt.Println("test")
	top := File{"top_file", nil}
	bottom := File{"bottom_file", &top}
	fmt.Println(top)
	fmt.Println(bottom)
	fmt.Println(bottom.parent.name)

	current_dir := bottom

	list := [4]string{"a", "b", "c", "d"}

	for _, v := range list {
		current_dir = current_dir.change_dir(v)
		fmt.Println(current_dir)
	}
}
