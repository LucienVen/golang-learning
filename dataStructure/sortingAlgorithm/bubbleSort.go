package sortingAlgorithm

/* 冒泡排序 */

func BubbleSort(list []int)  {
	n := len(list)

	// 在一轮中没有交换过
	didSwap := false

	// 进行 n-1 轮迭代
	for i := n - 1; i > 0; i-- {
		// 交换
		for j := 0; j < i; j++ {
			// 从小到大排序
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}

			// 从大到小
			// if list[j] < list[j+1]
		}

		//fmt.Println("didSwap: ", didSwap)

		// 判断该轮有无进行交换
		if !didSwap {
			return
		} else {
			didSwap = false
		}

	}
}