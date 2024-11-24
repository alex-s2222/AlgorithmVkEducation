package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (linkedList *LinkedList)addNewHead(el int) {
	node := &Node{
		data: el,
		next: nil,
	}

	// если список был пуст
	if linkedList.head == nil {
		linkedList.tail = node
	} else {
		// прежний head сдвигаем на один узел назад
		// созданный узел в качестве next ссылается на head
		node.next = linkedList.head
	}
	// записываем новый узел в качестве head
	linkedList.head = node

}

func (linkedList *LinkedList) addNewTail(el int){
	node := &Node{
		data: el,
		next: nil,
	}
	// если список был пуст 
	if (linkedList.tail == nil){
		linkedList.head = node
	}else{
		linkedList.tail.next = node
	}
	// записываем новый узел в качестве tail
	linkedList.tail = node
}

// если нет указателя на конец списка
// приводит к сложности O(n)
func (linkedList *LinkedList) addNewTailNotTail(el int){
	node := &Node{
		data: el,
		next: nil,
	}

	currentNode := linkedList.head
	for currentNode.next != nil{
		currentNode = currentNode.next
	}
	currentNode.next = node
}


func (linketList *LinkedList) insert(after int, el int){
	search := linketList.head
	for search != nil {
		if search.data == after{
			break 
		}
		// переходим к следующему элементу
		search = search.next
	}
	// если нашли элемент
	if search != nil{
		node := &Node{
			data: el,
			next: nil,
		}
		// если элемент после которого нужно вставить явл последним
		// то теперь новый элемент явл последний
		if search == linketList.tail {
			linketList.tail = node
		}
		// узел который был для search следующим
		// теперь является next для нового узла
		node.next = search.next
		search.next = node
	}
}

func (linketList *LinkedList) search(el int) (node *Node){
	current := linketList.head

	for current != nil{
		if current.data == el{
			return current
		}
		current = current.next
	}
	return nil

}


func main() {
	list := &LinkedList{}
	list.addNewHead(3)
	fmt.Println(list.tail, list.head)
	list.addNewTail(4)
	fmt.Println(list.tail, list.head)
	list.insert(4, 6)
	fmt.Println(list.tail, list.head)
	fmt.Println(list.search(4))

}