package main

import (
	"errors"
	"fmt"
)

type Node struct {
	item string
	next *Node //data type, pointer to a node, node as struct
}

type linkedList struct {
	head *Node
	size int
}

// *linkedList is a receiver
// constructer for struct use brace {}
// appointing the addNode function to the p (pointer to a linked list)
func (p *linkedList) addNode(name string) error { //interface, addNode is a function that linked to p linklst
	newNode := &Node{
		item: name,
		next: nil,
	}

	if p.head == nil { //empty list at the start //head is a pointer to a node
		p.head = newNode // assign to the next node
		//linkup the data , head doesnt have next, it doesnt have another compartmnet
		//
	} else { // if head is not empty
		currentNode := p.head

		//tranversal
		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = newNode // need a temp pointer, cannot use p
	}
	p.size++ // post increment
	return nil

}

func (p *linkedList) removeNode(index int) (string, error) {
	if p.head == nil {
		return "", errors.New("List is empty.") // EMPTY OR NOT
	}
	if index < 1 || index > p.size { // out of range
		return "", errors.New("Invalid index position.")
	}

	var item string // declare a temp variable

	if index == 1 { // removal at the 1st location

		//fmt.Println("Node to be removed is front node") // check for
		item = p.head.item
		p.head = p.head.next

	} else {
		currentNode := p.head //
		prevNode := p.head    // currentNode := p.head (same output)
		for i := 1; i <= index-1; i++ {
			prevNode = currentNode
			currentNode = currentNode.next // read currentNode.NEXT = curentNode
		}
		item = currentNode.item          //prevNode.next.item is the same output
		prevNode.next = currentNode.next // destination to the pointer prevNode, right to left like computer
	}
	p.size--
	return item, nil
}

func (p *linkedList) printAllNodes() error {
	if p.head == nil {

		fmt.Println("There are no nodes/elements in the list")
		return nil // we can customize, error message
	}
	currentNode := p.head
	fmt.Printf("%v\n", currentNode.item)
	// traverse
	// for loop below to print to the next element
	for currentNode.next != nil { // not equals to nil, it will jump to next

		currentNode = currentNode.next
		fmt.Printf("%v\n", currentNode.item)
	}

	return nil
}

func main() {
	// picked up the first node in the memory, thus creating memory/pointer
	myList := &linkedList{nil, 0} // creating an instance // when there's no node, head is grounded line 12

	fmt.Printf("Adding nodes to linked list....\n")
	myList.addNode("Mary")
	myList.addNode("Tom")
	myList.addNode("Dan")
	fmt.Print("Printing all the nodes...\n")
	// check for many names in the list, for size struct
	//
	myList.printAllNodes()
	fmt.Println("Remove one node...")
	item, _ := myList.removeNode(2)
	fmt.Println(item)

	fmt.Println("Printing all nodes of linked list...")
	myList.printAllNodes()

}
