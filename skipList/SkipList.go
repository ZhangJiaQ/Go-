package skipList

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

type SkipList struct {
	header *SkipListNode   // 指向跳跃表的头结点
	tail *SkipListNode     // 指向跳跃表的尾节点
	level int              // 不包含头结点的最高层数
	length int             // 不包含头结点的跳跃表的长度
}


func SkipListCreate() *SkipList{
	skiplist := &SkipList{}
	skiplist.level = 0
	skiplist.length = 0
	skiplist.header = new(SkipListNode{})
	return skiplist
}