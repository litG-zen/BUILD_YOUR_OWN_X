package main

import "fmt"

/*
A linked list is a fundamental data structure in computer science. It mainly allows efficient insertion
and deletion operations compared to arrays.

 Like arrays, it is also used to implement other data structures like stack, queue and deque
*/

// We will be creating a simple linked list using Struct and pointers for revision, before we delve into BlockChain.

type LinkedNode struct {
	data int // We will be taking a simple int-value as data for the linked list.
	next *LinkedNode
}

// Functionalities for LinkedList
// Node initiation
// Node addition
// Node search by data
// Node deletion
// Node traversal
// Node reversal

func InitNode(data int) *LinkedNode {
	return &LinkedNode{
		data: data,
		next: nil, // Explicitly setting next to nil (optional, since it's default)
	}
}

func (l *LinkedNode) AddNode(data int) {
	new_node := InitNode(data)
	l.next = new_node
}

func main() {
	first_node := InitNode(220396)
	fmt.Printf("First node data :%v", first_node)

	fmt.Println("\n New node addition")
	first_node.AddNode(160896)
	fmt.Printf("\n Updated LinkedList : %v %v", first_node.data, first_node.next)

}
