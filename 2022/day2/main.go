package main

import (
	"fmt"
	"os"
	"strings"
)

// rock	    1 -> 2 -> 3
// paper	2 -> 3 -> 1
// scissor  3 -> 1 -> 2

const (
	A = 1
	B = 2
	C = 3

	DRAW = 3
	WIN  = 6
)

func score1(p1, p2 int) int {

	lose := p2 + 1
	if lose > 3 {
		lose -= 3
	}
	win := p2 + 2
	if win > 3 {
		win -= 3
	}

	if p1 == win {
		return p2 + WIN
	}
	if p1 == lose {
		return p2
	}

	return p1 + DRAW
}
func score2(p1, p2 int) int {
	switch p2 {

	case 1:
		score := p1 + 2
		if score > 3 {
			score -= 3
		}
		return score
	case 2:
		return p1 + DRAW
	case 3:
		score := p1 + 1
		if score > 3 {
			score -= 3
		}
		return score + WIN
	}

	return 0
}

func main() {

	d, _ := os.ReadFile("input.txt")
	data := strings.Split(string(d), "\n")
	sum1 := 0
	sum2 := 0
	for _, line := range data {
		if len(line) == 0 {
			break
		}
		p1, p2 := line[0]-64, line[2]-87
		sum1 += score1(int(p1), int(p2))
		sum2 += score2(int(p1), int(p2))

	}
	fmt.Println(sum1)
	fmt.Println(sum2)

}
