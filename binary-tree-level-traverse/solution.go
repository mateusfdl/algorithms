package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func buildBinaryTree() *BinaryTree {
	return &BinaryTree{
		Root: &Node{
			Value: 1,
			Left: &Node{
				Value: 2,
				Left:  &Node{Value: 4},
				Right: &Node{Value: 5},
			},
			Right: &Node{
				Value: 3,
				Left:  &Node{Value: 6},
				Right: &Node{Value: 7},
			},
		},
	}
}

func main() {
	tree := buildBinaryTree()

	levelTraverse(tree.Root)
}

type Queue []*Node

func (q *Queue) Pop() {
	if len((*q)) < 1 {
		return
	}

	nq := make([]*Node, len((*q))-1)

	copy(nq, (*q)[1:])

	*q = nq
}

func (q *Queue) Push(n *Node) {
	nq := make([]*Node, len((*q))+1)

	copy(nq, (*q))

	nq = append(nq, n)

	*q = nq
}

func levelTraverse(root *Node) {
	q := make(Queue, 1)

	for {
		if len(q) < 1 {
			break
		}
	}
}
