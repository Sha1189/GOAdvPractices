package main

import (
	"errors"
	"fmt"
)

type node struct {
	item     string
	next     *node
	priority int
}

type queue struct {
	front *node
	back  *node
	size  int
}

func (p *queue) enqueue(name string, prty int) error {
	//creating new node
	newNode := &node{
		priority: prty, //new pq
		item:     name,
		next:     nil,
	}
	if p.front == nil {
		p.front = newNode
		p.back = newNode

	} else { // below code, it will check the priority of the string

		if p.front.priority > prty { // incoming higher priority than front, new pq, do not lost track assessing the fron tnode
			newNode.next = p.front
			p.front = newNode

		} else {

			currentNode := p.front

			for currentNode.next != nil && currentNode.next.priority <= prty {
				currentNode = currentNode.next
			}
			if currentNode.next == nil {

				p.back = newNode
			}
			newNode.next = currentNode.next
			currentNode.next = newNode
		}
	}
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
	myQueue.enqueue("Mary", 6)
	myQueue.enqueue("Tom", 4)
	myQueue.enqueue("Luca", 9)
	fmt.Println("Showing all nodes of queue...")
	count, _ := myQueue.printAllNodes()
	fmt.Printf("%+v names are in the queue", count)
}
