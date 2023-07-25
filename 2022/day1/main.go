package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	d, _ := os.ReadFile("input.txt")

	data := strings.Split(string(d), "\n")
	sum := 0
	max := []int{0, 0, 0}
	for _, line := range data {
		fmt.Println(line)
		if line == "" {
			for i, n := range max {
				if sum > n {
					max[i] = sum
					break
				}
			}
			sum = 0
			continue
		}

		n, _ := strconv.Atoi(line)

		sum += n
	}
	sum = 0
	for _, n := range max {
		sum += n
	}
	fmt.Println(sum)
}
