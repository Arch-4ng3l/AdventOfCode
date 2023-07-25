package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	d, _ := os.ReadFile("input.txt")
	data := strings.Split(string(d), "\n")
	chars := make(map[rune]int)
	chars2 := make(map[rune]int)
	chars2temp := make(map[rune]int)
	count := 0
	sum := 0
	sum2 := 0
	for _, line := range data {

		mid := len(line) / 2
		for _, c := range line[:mid] {

			chars[c]++
		}

		for _, c := range line[mid:] {
			if _, ok := chars[c]; ok {
				if 'a' <= int(c) && int(c) <= 'z' {
					sum += int(c) - 'a' + 1
				} else {
					sum += int(c) - 'A' + 27
				}

				break
			}
		}

		chars = make(map[rune]int)

		// PART 2

		count++
		if count == 3 {
			for _, c := range line {
				if _, ok := chars2[c]; ok {
					if 'a' <= int(c) && int(c) <= 'z' {
						sum2 += int(c) - 'a' + 1
					} else {
						sum2 += int(c) - 'A' + 27
					}
					break

				}
			}
			count = 0
			chars2 = make(map[rune]int)
		} else if count == 1 {
			for _, c := range line {
				chars2temp[c]++
			}
		} else {
			for _, c := range line {
				if _, ok := chars2temp[c]; ok {
					chars2[c]++
				}
			}
			chars2temp = make(map[rune]int)
		}

	}

	fmt.Println(sum)
	fmt.Println(sum2)
}
