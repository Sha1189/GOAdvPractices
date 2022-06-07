package main

import (
	"errors"
	"fmt"
)

type binaryNode struct {
	item  string
	left  *binaryNode
	right *binaryNode
}

type BST struct {
	root *binaryNode
}

// wrapper function, (t **binaryNode, item string)declaration
func (bst *BST) insertNode(t **binaryNode, item string) error { // recursion, a pointer to a pointer **binaryNode
	if *t == nil { // if nil can sit down there
		newNode := &binaryNode{
			item:  item,
			left:  nil,
			right: nil,
		}
		*t = newNode // t is now pointing to the left, so you need to def and assigned to new node
		return nil   // *t is the root/ return nil to go up the tree
		// pointer need to be linked to the new node as it is nil so create new node

	} //base case

	// below divide an conquer if root is not nil
	if item < (*t).item { // you need to dereference only 1 layer
		bst.insertNode(&(*t).left, item) // i call the insertNode repetitively until root is found
		// dereference t, (*t) : deferencing the pointer to access the data/ . includes 1 layer of derefencing
	} else {
		bst.insertNode(&(*t).right, item)
	}
	return nil
}

// those with exact value will go to the right
// unless you write <= it will go to the left

// anything we want to change, we want to update the original position.
//  // if not we are using the copy of struct, if we want the original struct, we need to use pointer
// pointer of a pointer to receive it

func (bst *BST) insert(item string) error { // this a copy of the string you parsing in

	bst.insertNode(&bst.root, item)
	return nil
}

func (bst *BST) inOrderTraversal(t *binaryNode) error { //keep going to the left if nil retun to t.left print item,
	// go to right, then finish no more code, then retun nil, which then it will go up
	if t != nil {
		bst.inOrderTraversal(t.left)
		fmt.Println(t.item)
		bst.inOrderTraversal(t.right) // calling root right
	}
	return nil //go to the above layer of left
}

//traversal
func (bst *BST) inOrder() {
	bst.inOrderTraversal(bst.root) // no need ampersand, bcos you are not manipulating the data
}

func (bst *BST) preOrderTraversal(t *binaryNode) error {
	if t != nil {
		fmt.Println(t.item)
		bst.preOrderTraversal(t.left)
		bst.preOrderTraversal(t.right)
	}
	return nil
}

func (bst *BST) preOrder() {
	bst.preOrderTraversal(bst.root)
}

func (bst *BST) postOrderTraversal(t *binaryNode) error {
	if t != nil {
		bst.postOrderTraversal(t.left)
		bst.postOrderTraversal(t.right)
		fmt.Println(t.item)
	}
	return nil
}

func (bst *BST) postOrder() {
	bst.postOrderTraversal(bst.root)
}

func (bst *BST) findSucessor(t *binaryNode) string {
	for t.right != nil { // find node on extreme right
		t = t.right
	}
	return t.item
}

func (bst *BST) removeNode(t **binaryNode, item string) (*binaryNode, error) {
	if *t == nil { // check if root is nil
		return nil, errors.New("Tree is empty")
	} else if item < (*t).item {
		(*t).left, _ = bst.removeNode(&(*t).left, item) //get attached to the left of the tree
	} else if item > (*t).item {
		(*t).right, _ = bst.removeNode(&(*t).right, item)
	} else { // 2nd case of 1 child
		if (*t).left == nil {
			return (*t).right, nil
		} else if (*t).right == nil {
			return (*t).left, nil
		} else { // 3rd case of 2 children
			(*t).item = bst.findSucessor((*t).left)
			(*t).left, _ = bst.removeNode(&(*t).left, item)
		}
	}
	return *t, nil
}

func (bst *BST) remove(item string) {
	bst.root, _ = bst.removeNode(&bst.root, item)
}

func main() {

	bst := &BST{nil}
	bst.insert("Mary")
	bst.insert("Tom")
	bst.insert("Lucas")
	fmt.Println("Inorder")
	bst.inOrder()
	fmt.Println("Preorder")
	bst.preOrder()
	fmt.Println("POSTorder")
	bst.postOrder()
	bst.remove("Tom")
}
