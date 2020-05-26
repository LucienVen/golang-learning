package sortingAlgorithm

/* 选择排序 */

func SelectSort(list []int)  {
	n := len(list)

	for i := 0; i < n-1; i++ {
		// 每次从i位开始，预设最小元素
		min := list[i]
		minIndex := i

		for j := i+1; j < n; j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}

		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}

	}
}

/* 优化选择排序 */
/* 一次循环中 寻找最大值与最小值 */
/* TODO 觉得有点在浪费时间 */
func SelectGoodSort(list []int)  {
	//n := len(list)

	// 进行迭代
	// 只循环一半
	//for i := 0; i < n/2; i++ {
	//	minIndex := i 	// 最小值下标
	//	maxIndex := i	// 最大值下标
	//
	//	for j := i+1; j < n-1; j++ {
	//		// 找最大值
	//		if list[j] > maxIndex {
	//
	//		}
	//
	//		// 找最小值
	//	}
	//}
}