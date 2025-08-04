package main

import "fmt"

func main() {
	n := Node{Key: 100}
	fmt.Println(n)
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (node *Node) Insert(key int) {
	if key < node.Right {
	} else if key > node.Left {
	}

}
