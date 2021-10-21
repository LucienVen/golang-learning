// @Author : a1234
// @Time : 2021/10/21 11:14 上午 
// @Describe :  剑指 Offer 11. 旋转数组的最小数字

// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/

package saber

func MinArray(numbers []int) int {
	l := 0
	r := len(numbers) - 1
	mid := l + (r - l) / 2

	if numbers[l] < numbers[r] {
		// 单调递增
		return numbers[l]
	}

	for l <= r {
		if numbers[l] < numbers[r] {
			return numbers[l]
		}

		// 若最左小于mid元素，则最左到mid是严格递增的，那么最小元素必定在mid之后
		if numbers[l] < numbers[mid] {
			l = mid + 1
			mid = (l + r) / 2

		} else if numbers[l] > numbers[mid] {
			// 若最左大于mid元素，则最小元素必定在最左到mid之间(不包括最左，因为最左已经大于mid)
			r = mid
			l++
			mid = (l + r) / 2
		} else {
			//若二者相等，则最小元素必定在l+1到r之间，因为l和mid相等，故可以去除
			l++
			mid = (l + r) / 2
		}

	}
	return numbers[mid]
}
