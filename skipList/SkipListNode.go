package skipList



type SkipListNode struct {
	level []skipListLevel
	backward *SkipListNode
	score float64
}

type skipListLevel struct {
	forward *SkipListNode
	span int
}

func NewSkipListNode(score float64) *SkipListNode{
	node := &SkipListNode{}
	node.level := [32]*skipListLevel

}
