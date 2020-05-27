package main

/* 哈希表 */

import (
	"errors"
	"fmt"
	s "go.learning/dataStructure/searchAlgorithms"
)

func main()  {
	// 新建一个哈希表
	hashMap := s.NewHashMap(16)
	//fmt.Printf("cap: %d, len: %d\n", hashMap.Capacity(), hashMap.Len())
	//panic("end")
	// 存放 35 个值
	for i := 0; i < 35; i++ {

		key := fmt.Sprintf("%d", i)
		value := fmt.Sprintf("v%d", i)

		hashMap.Put(key, value)
		//fmt.Println(key, value)
	}

	// 打印哈希表信息
	fmt.Printf("cap: %d, len: %d\n", hashMap.Capacity(), hashMap.Len())

	// 打印全部键值对
	//hashMap.Range()

	key := "4"
	value, ok := hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}

	// 删除
	hashMap.Delete(key)
	fmt.Printf("cap: %d, len: %d\n", hashMap.Capacity(), hashMap.Len())

	value, ok = hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}

	fmt.Println("=================================")
	//test2()
	
}

func test()  {
	i := 2
	if i > 1 {
		fmt.Println(" start i: ",i)
		i, err := doDivision(i, 2)
		if err != nil {
			panic(err)
		}
		fmt.Println(i)
	}
	fmt.Println(i)
}

func doDivision(x, y int) (int, error) {
	fmt.Println("x: ", x, " y: ", y)
	if y == 0 {
		return 0, errors.New("input is invalid")
	}
	return x / y, nil
}

type person struct {
	name   string
	age    byte
	isDead bool
}

func test1()  {
	p1 := person{name: "zzy", age: 100}
	p2 := person{name: "dj", age: 99}
	p3 := person{name: "px", age: 20}
	people := []person{p1, p2, p3}
	whoIsDead(people)
	for _, p := range people {
		if p.isDead {
			fmt.Println("who is dead?", p.name)
		}
	}
}

func whoIsDead(people []person) {
	for _, p := range people {
		if p.age < 50 {
			p.isDead = true
		}
	}
}

func test2()  {
	p1 := person{name: "zzy", age: 100}
	p2 := p1
	p1.name = "changed"
	fmt.Println(p2.name)
}