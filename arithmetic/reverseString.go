package main

import "fmt"

/* 字符串反转 */

func main()  {

	toDoString := "ABCDEFG"

	//fmt.Println(len(toDoString))

	temp := []rune(toDoString)

	for i, j := 0, len(toDoString)-1; j > i ; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
		//fmt.Println(i)
		//fmt.Println(j)
		//fmt.Println("----")
	}

	fmt.Println(string(temp))
	
}