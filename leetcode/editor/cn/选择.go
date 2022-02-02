package main

import "fmt"

func selectionSort(nums []int) {

	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}

}

func main() {
	nums := []int{1, 2, 7, 9, 5, 8}
	selectionSort(nums)
	fmt.Println(nums)
}
