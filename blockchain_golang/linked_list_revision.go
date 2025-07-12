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

func main() {
	fmt.Println("Hello! this is first commit for Linkedlist.")
}
