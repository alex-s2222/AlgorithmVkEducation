package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (linketList *LinkedList) appendFont(data int){
	node := &Node{
		data: data,
		next: nil,
		prev: nil,
	}
	if linketList.head == nil{
		linketList.head = node
		return
	}

	node.next = linketList.head
	linketList.head.prev = node
	linketList.head = node
}

func (linketList *LinkedList) appendBack(data int){
	node := &Node{
		data: data,
		next: nil,
		prev: nil,
	}
	if linketList.head == nil{
		linketList.head = node
		return
	}

	cur_node := linketList.head
	for cur_node.next != nil{
		cur_node = cur_node.next
	}

	cur_node.next = node
	node.prev = cur_node
}

func main(){
	list := &LinkedList{}
	list.appendFont(3)
	fmt.Println(list.head)
	list.appendBack(4)
	fmt.Println(list.head)
}