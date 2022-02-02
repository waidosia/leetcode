package main

import (
	"fmt"
)

func bubbleSort(nums []int) {
	// 常规写法
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i+1; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func new_bubbleSort(nums []int) {
	hasChange := true
	for i, n := 0, len(nums)-1; i < n && hasChange; i++ {
		hasChange = false
		for j := 0; j < n-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				hasChange = true
			}
		}
	}
}
func main() {
	nums := []int{1, 2, 7, 9, 5, 8}
	bubbleSort(nums)
	fmt.Println(nums)

	new_bubbleSort(nums)
	fmt.Println(nums)

}
