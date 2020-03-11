package skipList

import (
	"sync"
	"math"
)

const FLOAT_MAX = math.MaxFloat64
const FLOAT_MIN = math.SmallestNonzeroFloat64

type SkipList struct {
	head *SkipListNode   // 指向跳跃表的头结点
	tail *SkipListNode     // 指向跳跃表的尾节点
	level int              // 跳跃表的最高层数
	length int             // 跳跃表的长度
	mutex  sync.RWMutex    // 读写锁
}


func newSkipList(level int) *SkipList {

	// create header node
	head := newSkipListNode(FLOAT_MIN, 0,level  )

	// create tail node
	var tail *SkipListNode

	for i:= 0; i< len(head.level); i++ {
		head.level[i] = tail
	}

	return &SkipList{
		head: head,
		tail: tail,
		level: level,
		length: 0,
	}

}




func (s *SkipList) insert(index uint64, value float64) {
	// write lock
	s.mutex.Lock()
	defer s.mutex.Unlock()

}



