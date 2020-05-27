package searchAlgorithms

import (
	"fmt"
	"math"
	"sync"
	"github.com/OneOfOne/xxhash"
)

/* 哈希查询 - 哈希表 */

const (
	expandFactor = 0.75 	// 扩容因子
)

/* 哈希表 */
type HashMap struct {
	array 			[]*keyPairs		// 哈希表数组，每个元素是一个键值对
	capacity		int				// 数组容量
	len		 		int				// 已添加键值对元素数量
	capacityMask	int				// 掩码，相当于 capacity-1, 主要用来计算数组下标
	lock 			sync.Mutex		// 并发锁
}

/* 键值对 */
type keyPairs struct {
	key 	string 		// 键
	value 	interface{}	// 值
	next 	*keyPairs	// 下一个键值对
}

/* 初始化哈希表 */
func NewHashMap(capacity int) *HashMap {
	// 默认大小为 16
	defaultCapacity := 1 << 4

	if capacity <= defaultCapacity {
		// 如果传入大小小于默认大小，那么使用默认大小 16
		capacity = defaultCapacity
	} else {
		// TODO  否则，实际大小大于 capacity 的第一个 2^k ???
		capacity = 1 << (int(math.Ceil(math.Log2(float64(capacity)))))
	}

	// 新建一个哈希表
	m := new(HashMap)
	m.array = make([]*keyPairs, capacity, capacity)

	//fmt.Println(m.array)
	m.capacity = capacity
	m.capacityMask = capacity - 1
	return m
}

/* 返回哈希表已添加的元素数量 */
func (m *HashMap) Len() int {
	return m.len
}

/* 返回Capacity */
func (m *HashMap) Capacity() int {
	return m.capacity
}

/* 计算 key 的哈希值 */
var hashAlgorithm = func(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}

/* 对键key 进行哈希求值，并计算下标 */
/* mask 容量掩码 */
func (m *HashMap) hashIndex(key string, mask int) int {
	// 求哈希
	hash := hashAlgorithm([]byte(key))
	// 求下标
	index := hash & uint64(mask)
	return int(index)
}

/* 添加键值对 */
func (m *HashMap) Put(key string, value interface{})  {
	// 并发安全
	m.lock.Lock()
	defer  m.lock.Unlock()

	// 键值对 计算的哈希表数组下标
	index := m.hashIndex(key, m.capacityMask)

	// 哈希表数组下标的元素
	elem := m.array[index]

	if elem == nil {
		// 元素为空，表示空链表，没有哈希冲突，直接复制
		m.array[index] = &keyPairs{
			key:   key,
			value: value,
		}
	} else {
		// 处理哈希冲突
		var lastPairs *keyPairs

		// 遍历链表，查看元素是否存在，存在则替换值，否则找到最后一个键值对
		for elem != nil {
			if elem.key == key {
				// 键值对存在，则更新并返回
				elem.value = value
				return
			}

			// 若不存在
			lastPairs = elem
			elem = elem.next
		}

		// 找不到键值对，将新键值对 添加到聊表尾端
		lastPairs.next = &keyPairs{
			key:   key,
			value: value,
		}
	}

	// 新的哈希表数量
	newlen := m.len + 1
	
	// 如果超出扩容因子，则需要扩容
	if float64(newlen) / float64(m.capacity) >= expandFactor {
		// 新建一个原来两倍大小的哈希表
		newM := new(HashMap)
		newM.array = make([]*keyPairs, 2 * m.capacity, 2 * m.capacity)
		newM.capacity = 2 * m.capacity
		newM.capacityMask = 2 * m.capacityMask - 1

		// 遍历老的哈希表，讲键值对重写哈希到新的哈希表
		for _, pairs := range m.array {
			for pairs != nil {
				// 直接递归put
				newM.Put(pairs.key, pairs.value)
				pairs = pairs.next
			}
		}

		// 替换老的哈希表
		m.array = newM.array
		m.capacity = newM.capacity
		m.capacityMask = newM.capacityMask
	}

	m.len = newlen

}

/* 获取键值对 */
func (m *HashMap) Get(key string) (value interface{}, ok bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	// 计算键值对下标
	index := m.hashIndex(key, m.capacityMask)

	// 哈希表数组下标元素
	elem := m.array[index]

	// 遍历链表，查看元素是否存在
	for elem != nil {
		if elem.key == key {
			return elem.value, true
		}

		elem = elem.next
	}

	return
}

/* 删除键值对 */
func (m *HashMap) Delete(key string)  {
	m.lock.Lock()
	defer m.lock.Unlock()

	// 计算哈希数组下标
	index := m.hashIndex(key, m.capacityMask)

	// 哈希数组下标元素
	elem := m.array[index]

	// 若是空链表，则直接返回
	if elem == nil {
		return
	}

	// 若链表的第一个元素就是要删除的元素
	if elem.key == key {
		// 将后驱节点键值对 连上
		m.array[index] = elem.next
		m.len = m.len - 1
		return
	}

	// 从第二个元素开始遍历链表
	nextElem := elem.next
	// 遍历，查找key
	for nextElem != nil  {
		if nextElem.key == key {
			elem.next = nextElem.next
			m.len = m.len - 1
			return
		}

		elem = nextElem
		nextElem = nextElem.next
	}
}

/* 遍历打印哈希表 */
func (m *HashMap) Range()  {
	// 并发安全
	m.lock.Lock()
	defer  m.lock.Unlock()

	for index , pairs := range m.array {
		fmt.Println("index: ", index)
		for pairs != nil {
			fmt.Printf("'%v'='%v'\n", pairs.key, pairs.value)
			pairs = pairs.next
		}
	}
}














