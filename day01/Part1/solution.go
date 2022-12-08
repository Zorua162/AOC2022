package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Max(s []int) int {
	current_max := math.MinInt64
	for _, element := range s {
		if element > current_max {
			current_max = element
		}
	}
	return current_max
}

func Sum(s []int) int {
	total := 0
	for _, element := range s {
		total += element
	}
	return total
}

func main() {
	dat, err := os.ReadFile("./../dat")
	// dat, err := os.ReadFile("./../example_dat")
	check(err)
	// fmt.Println(string(dat))

	// s := strings.Fields(string(dat[:]))
	// fmt.Println(s)
	// var last int = math.MaxInt64

	var elf_amounts []int
	var current_amount int
	splitData := strings.Split(string(dat), "\n")
	for _, element := range splitData {
		intVal, _ := strconv.Atoi(strings.TrimSpace(element))
		current_amount += intVal
		if strings.TrimSpace(string(element)) == "" {
			elf_amounts = append(elf_amounts, current_amount)
			current_amount = 0
		}
	}
	// fmt.Println(elf_amounts)
	// fmt.Println(Max(elf_amounts))
	sort.Ints(elf_amounts)
	fmt.Println(elf_amounts)
	top_three := elf_amounts[len(elf_amounts)-3:]
	fmt.Println(top_three)
	fmt.Println(Sum(top_three))

}
