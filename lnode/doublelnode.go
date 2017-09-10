package lnode

type Node struct {
	Data     int
	PrevNode *Node
	NextNode *Node
}
