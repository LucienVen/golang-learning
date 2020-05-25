package baseStruct

import (
	"fmt"
	"sync"
	"unsafe"
)

/* 不可重复集合 Set */

type Set struct {
	m 		map[int]struct{}		// 字典实现
	len 	int						// 集合的大小
	sync.RWMutex					// 读写锁
}


/* 集合初始化 */
func NewSet(cap int) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{
		m:    temp,
	}
}

/* 添加元素 */
func (s *Set) Add(item int)  {
	// 加写锁
	s.Lock()
	defer s.Unlock()

	// 字典实现
	// 字典键，是不可重复的，利用字典键来存放item，可自动去重
	// struct{}{} 占用字节为0
	s.m[item] = struct{}{}
	s.len = len(s.m)
}

/* 删除元素 */
func (s *Set) Remove(item int)  {
	s.Lock()
	defer s.Unlock()

	if s.len == 0 {
		panic("Set empty")
	}

	delete(s.m, item)	// 实际上，从字典删除该值
	s.len = len(s.m)	// 重新计算集合大小
}

/* 查看元素是否在集合中 */
func (s *Set) Has(item int) bool {
	// 加读锁 （读锁可以多个）
	s.RLock()
	s.RUnlock()

	_, ok := s.m[item]
	return ok
}

/* 查看集合大小 */
func (s *Set) Len() int {
	return s.len
}

/* 查看集合是否为空 */
func (s *Set) IsEmpty() bool  {
	return s.len == 0
}

/* 清除所有集合元素 */
func (s *Set) Clear()  {
	// 加读锁
	s.Lock()
	s.Unlock()

	s.m = map[int]struct{}{}
	s.len = 0
}

/* 将集合转换为列表（切片） */
func (s *Set) List() []int {
	// 读锁
	s.RLock()
	s.RUnlock()

	list := make([]int, 0, s.len)

	for item, _ := range s.m {
		list = append(list, item)
	}

	return list
}

/* 为什么使用struct{} */
func Other()  {
	a := struct {}{}
	b := struct {}{}

	if a == b {
		fmt.Printf("a == b right! %p\n", &a)
	}

	fmt.Println("unsafe.Sizeof result: ", unsafe.Sizeof(a))
}