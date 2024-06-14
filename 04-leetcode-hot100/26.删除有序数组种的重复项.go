package main

import "fmt"

func main() {
	var nums = []int{0, 1, 1, 3, 5, 5, 6, 6, 7, 7, 8}
	result := removeDuplicates(nums)
	fmt.Println(result)
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	fmt.Println(nums)
	return slow

}
