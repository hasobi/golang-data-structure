package main

import "fmt"

type Node struct {
 Data int
 Next *Node
}

func (n *Node) AddNode(data int) {
 newNode := Node{data, nil}
 iter := n
 for iter.Next != nil {
  iter = iter.Next
 }
 iter.Next = &newNode
}

func (n *Node) PrintNode() {
 iter := n
 for iter != nil {
  fmt.Println(iter.Data)
  iter = iter.Next
 }
}

func main() {
 newNode := Node{10, nil}
 newNode.AddNode(20)
 newNode.AddNode(30)
 newNode.AddNode(40)
 newNode.PrintNode()
}
