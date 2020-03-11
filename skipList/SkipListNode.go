package skipList



type SkipListNode struct {
	level []*SkipListNode
	score float64
	index uint64
}

func newSkipListNode(score float64, index uint64, level int) *SkipListNode {
	return &SkipListNode{
		score: score,
		index: index,
		level: make([]*SkipListNode, level, level),
	}
}

// 返回node的index
func (node SkipListNode) Index() uint64{
	return node.index
}

// 返回node的value
func (node SkipListNode) Value() float64{
	return node.score
}