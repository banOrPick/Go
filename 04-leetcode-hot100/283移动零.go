package main

import "fmt"

func main() {
	var nums = []int{2, 1, 0, 3, 0, 1, 0, 4, 5, 6, 0}
	moveZeroes(nums)
}

//func moveZeroes(nums []int) {
//	zero := 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == zero {
//			for j := i; j < len(nums); j++ {
//				if nums[j] != zero {
//					temp := nums[j]
//					nums[j] = zero
//					nums[i] = temp
//					break
//				}
//			}
//		}
//	}
//	fmt.Println(nums)
//}

func moveZeroes(nums []int) {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != 0 {
			temp := nums[slowIndex]
			nums[slowIndex] = nums[fastIndex]
			nums[fastIndex] = temp
			slowIndex++
		}
	}
	fmt.Println(nums)
}
