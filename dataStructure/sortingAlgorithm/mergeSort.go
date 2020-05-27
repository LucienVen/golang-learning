package sortingAlgorithm

/* 归并排序 */

/* 自顶向下 */
func MergeSort(array []int, begin int, end int)  {
	
	// 元素数量大于1 才进入递归
	if end - begin > 1 {

		// 将数组一分为二
		mid := begin + (end - begin + 1)/2

		MergeSort(array, begin, mid)

		MergeSort(array, mid, end)
	}
	
}

/* 自底向上 */

