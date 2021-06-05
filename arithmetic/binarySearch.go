package main

/* 二分查找 - 查找一个数在数组中下标的位置 */

import (
	"fmt"
	"runtime/debug"
)

func main() {
	//list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	//list = alg.BubbleSort(list)


	nums := []int{1, 2, 3, 4, 5, 6, 7}

	aimIndex := BinarySearchRecursive(6, nums, 0, len(nums))
	fmt.Println("BinarySearchRecursive: ", aimIndex)


	nums2 := []int{1, 2, 3, 4, 5, 6, 7}
	aimIndex2 := BinarySearch(6, nums2)

	fmt.Println("BinarySearch: ", aimIndex2)

	debug.PrintStack()



}

func BinarySearch(target int, nums []int) int {
	left := 0
	right := len(nums)
	for left < right  {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
			continue
		}

		if nums[mid] < target {
			left = mid + 1
			continue
		}


	}

	return -1
}


/* 递归实现 */
func BinarySearchRecursive(target int, nums []int, left int, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	if target == nums[mid] {
		return mid
	}

	if nums[mid] > target {
		return BinarySearchRecursive(target, nums, left, mid+1)
	}

	if nums[mid] < target {
		return BinarySearchRecursive(target, nums, mid+1, right)
	}

	return -1
}