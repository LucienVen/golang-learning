package alg

/* 冒泡排序 */

func BubbleSort(arr []int) []int {
	if arr == nil || len(arr) == 1 {
		return arr
	}

	n := len(arr)

	//for i := 0; i < lenArr ; i++ {
	//	for j := i; j < lenArr ; j++ {
	//		if arr[j] > arr[j+1] {
	//			arr[j], arr[j+1] = arr[j+1], arr[j]
	//		}
	//	}
	//}

	for i := n - 1; i > 0 ; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

