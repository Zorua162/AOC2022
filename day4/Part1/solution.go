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

func to_list_of_numbers(elf []string) []int {
	var elf_list []int
	start, _ := strconv.Atoi(elf[0])
	end, _ := strconv.Atoi(elf[1])
	for i := start; i < end+1; i = i + 1 {
		elf_list = append(elf_list, i)
	}
	return elf_list

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

func int_list_contains_int(list []int, input int) bool {
	for _, comp := range list {
		if comp == input {
			return true
		}
	}
	return false
}

func find_duplicate_numbers(list1 []int, list2 []int) []int {
	var out_ints []int

	for _, comp_int := range list1 {
		if int_list_contains_int(list2, comp_int) {
			out_ints = append(out_ints, comp_int)
		}
	}
	return out_ints
}

func main() {
	// dat, err := os.ReadFile("./../extra_senario")
	dat, err := os.ReadFile("./../dat")
	// dat, err := os.ReadFile("./../example_dat")
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
		var elfs []string = strings.Split(element, ",")
		elf1 := strings.Split(elfs[0], "-")
		elf2 := strings.Split(elfs[1], "-")
		elf1_list := to_list_of_numbers(elf1)
		elf2_list := to_list_of_numbers(elf2)
		// elf2 := strings.Split(elfs[1], "-")
		// Determine if one is completely in the other
		fmt.Println(elf1_list, elf2_list)

		// dupes1 := find_duplicate_letters(strings.Join(elf1_list, ""), strings.Join(elf2_list, ""))
		// dupes2 := find_duplicate_letters(strings.Join(elf2_list, ""), strings.Join(elf1_list, ""))
		// fmt.Println(dupes1, len(dupes1), dupes2, len(dupes2), len(elf1_list), len(elf2_list))
		// if len(dupes1) == len(elf1_list) {
		// 	count += 1
		// }
		// if len(dupes1) == len(elf2_list) {
		// 	count += 1
		// }

		dupes := find_duplicate_numbers(elf1_list, elf2_list)
		fmt.Println(dupes)
		// dupes = find_duplicate_numbers(elf2_list, elf1_list)
		// fmt.Println(dupes)
		if len(dupes) == len(elf1_list) {
			count += 1
		} else if len(dupes) == len(elf2_list) {
			count += 1
		}

	}
	fmt.Println(count)

	// input := []string{"22", "24"}
	// fmt.Println(to_list_of_numbers(input))

	// 55
	// 60

}
