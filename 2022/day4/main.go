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
	sum2 := 0
	for _, line := range data {
		pairs := strings.Split(strings.ReplaceAll(line, "-", " "), ",")
		if len(pairs) != 2 {
			break
		}

		str1 := strings.Split(pairs[0], " ")
		str2 := strings.Split(pairs[1], " ")

		nums1 := toNums(str1)
		nums2 := toNums(str2)
		if contains(nums1, nums2) {
			sum++
			sum2++
		} else if overlaps(nums1, nums2) {
			sum2++
		}

	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

func toNums(str []string) []int {

	var nums []int

	for _, n := range str {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)

	}

	return nums
}

func contains(nums1, nums2 []int) bool {
	return condContains(nums1, nums2) || condContains(nums2, nums1)
}
func condContains(nums1, nums2 []int) bool {
	return (nums1[0] <= nums2[0] && nums1[1] >= nums2[1]) || (nums2[0] <= nums1[0] && nums2[1] >= nums1[1])

}
func overlaps(nums1, nums2 []int) bool {
	return condOverlaps(nums1, nums2) || condOverlaps(nums2, nums1)
}

func condOverlaps(nums1, nums2 []int) bool {

	return ((nums1[0] >= nums2[0] && nums1[0] <= nums2[1]) || (nums1[1] >= nums2[0] && nums1[1] <= nums2[1]))
}
