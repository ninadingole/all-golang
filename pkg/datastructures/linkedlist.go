package datastructures

type Node struct {
	value int
	next  *Node
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Next() *Node {
	return n.next
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) First() *Node {
	return ll.head
}

func (ll *LinkedList) Last() *Node {
	return ll.tail
}

func (ll *LinkedList) Push(val int) {
	node := &Node{value: val}

	if ll.head == nil {
		ll.head = node
	} else {
		ll.tail.next = node
	}

	ll.tail = node
}
