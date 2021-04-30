package list

import "fmt"

type Node struct {
	Data	interface{}
	Next	*Node
}

func NewNode (v interface{}) *Node {
	node := new(Node)
	node.Data = v
	node.Next = nil
	return node
}

type List struct{
	Header	*Node
	//footer	*Node	
}

func NewList() *List{
	list := new(List)
	list.Header = nil
	return list
}

func (l *List) Insert(v interface{}) {
	node := NewNode(v)

	node.Next = l.Header
	l.Header = node
}

func (l List) Show() {
	if l.Header == nil{
		fmt.Println(l.Header)
	}else{
		for l.Header != nil {
			fmt.Println(l.Header.Data)
			l.Header = l.Header.Next
		}			
	}

}

func (l List) Search(dest interface{}) []interface{} {
	idx := new(Node)
	var out []interface{}
	for idx = l.Header; idx != nil; idx = idx.Next{
		if dest == idx.Data{
			out = append(out, idx.Data)
			//return out
		}
	}
	return out
}

func (l *List) Remove (data interface{}) {
	var current, prev *Node
	var found bool
	current = l.Header
	prev = nil
	found = false

	for current != nil && !found {
		found = (current.Data == data)
		if !found {
			prev = current
			current = current.Next
		}
	}
	if current != nil {
		if current == l.Header {
			l.Header = current.Next
		}else {
			prev.Next = current.Next
		}
	}
}