package main

import (
	"errors"
	"fmt"
)

type node struct {
	item string
	next *node
}

type queue struct {
	front *node
	back  *node
	size  int
}

func (p *queue) enqueue(name string) error {
	//creating new node
	newNode := &node{
		priority: prty,
		item: name,
		next: nil,
	}
	if p.front == nil {
		p.front = newNode
	} else {
		if p.front.priority >prty {

			
		}
		p.back.next = newNode
	}
	p.back = newNode
	p.size++
	return nil
}

func (p *queue) printAllNodes() (int, error) { //include count

	var count int = 0

	if p.front == nil {
		return 0, errors.New("Queue is empty.")
	}
	currentNode := p.front
	for currentNode != nil {
		fmt.Println(currentNode.item)
		currentNode = currentNode.next
		count++
	}
	return count, nil
}

func main() {
	myQueue := &queue{nil, nil, 0}
	fmt.Println("Enqueueing...")
	myQueue.enqueue("Mary")
	myQueue.enqueue("Tom")
	myQueue.enqueue("Luca")
	fmt.Println("Showing all nodes of queue...")
	count, _ := myQueue.printAllNodes()
	fmt.Printf("%+v names are in the queue", count)
}
