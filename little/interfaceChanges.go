package main

import (
	"errors"
	"fmt"
)

// interface 断言使用

func main() {
	fmt.Println("Hello, world")

	var param interface{}

	param = "off"

	value, ok := param.(string)
	if !ok {
		panic("assert err")
	}

	status := ActStatus(value)
	if ok, err := status.CheckParam(); !ok {
		panic(err)
	}

	fmt.Printf("value:%v, type:%T", status, status)
}

// 活动状态
type ActStatus string

const (
	ActStatus_On  ActStatus = "on"
	ActStatus_Off ActStatus = "off"
)

func (param ActStatus) CheckParam() (bool, error) {
	if param != ActStatus_On && param != ActStatus_Off {
		return false, errors.New(fmt.Sprintf("ActStatus check err:%s", param))
	}

	return true, nil
}
