package main

import "fmt"

func main() {
	var nums = []int{2, 1, 0, 3, 0, 1, 0, 4, 5, 6, 0}
	removeElement(nums, 0)
}

func removeElement(nums []int, val int) int {
	length := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[length] = nums[i]
			length++
		}
	}
	fmt.Println(nums)
	return length
}
