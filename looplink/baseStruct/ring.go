package baseStruct

// 循环链表
type Ring struct {
	next 	*Ring			// 后驱节点
	prev 	*Ring			// 前驱节点
	Value 	interface{}
}

// 初始化 空的循环链表，前驱后驱都是指向自己
func (r *Ring) Init() *Ring {
	r.next = r
	r.prev = r
	return r
}

//func main()  {
//	r := new(Ring)
//	//r.Init()
//	fmt.Printf("type: %T, value: %v\n",r ,r)
//	// type: *main.Ring, value: &{<nil> <nil> <nil>}
//}

// 创建指定大小为N的循环链表, 值全为空
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}

	// new 的作用是根据传入的类型分配一片内存空间并返回指向这片内存空间的指针
	r := new(Ring)
	// 复制了r
	// p 设置为当前node
	p := r

	for i := 1; i < n; i++ {
		// 当前node 指向新的node，并设置好prev，next指向
		p.next = &Ring{prev: p}
		// 把新node设为当前node
		p = p.next
	}

	// 处理循环链表 收尾链接
	p.next = r
	r.prev = p

	return r
}

// 获取 前一个节点
func (r *Ring) Prev() *Ring {
	if r == nil {
		return r.Init()
	}

	return r.prev
}

// 获取 后一个节点
func (r *Ring) Next() *Ring {
	if r == nil {
		return r.Init()
	}

	return r.next
}

// 获取第N个节点
func (r *Ring) Move(n int) *Ring {
	if r == nil {
		return r.Init()
	}

	switch {
	// 若 n 小于 0 ，则向前遍历
	case n < 0:
		for ; n < 0; n ++ {
			r = r.prev
		}
	// 若 n 大于 0 ，则向后遍历
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}

	return r

}

// 添加节点
// 往节点A，链接一个节点，并且返回之前节点A的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		// 处理prev, next指向
		//p := s.Prev()

		r.next = s
		s.prev = r
		//TODO 此处有点疑惑，我认为是 n.prev = s，还是说两者相等？
		//n.prev = p
		//p.next = n

		n.prev = s
		s.next = n
	}

	return n
}

/* TODO 复制 */
func (r *Ring) Link1(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}






















