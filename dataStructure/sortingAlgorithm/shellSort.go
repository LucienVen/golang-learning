package sortingAlgorithm

/* 希尔排序 */
/* 使用步长的插入排序 */

func ShellSort(list []int)  {
	n := len(list)


	// 每次减半，直到步长等于1
	for step := n/2 ; step >= 1; step /= 2 {

		// 开始插入排序
		for i:=step; i < n; i = i + step {

			for j := i - step; j >= 0; j -= step {
				if list[j + step] < list[j] {
					list[j], list[j + step] = list[j + step], list[j]
					continue
				}
				break
			}
		}

	}
}
