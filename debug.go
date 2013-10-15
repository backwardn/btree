package btree

import (
	"fmt"
)

//debug func
func (n *Node) PrintChildrens() {
	for i := range n.Childrens {
		fmt.Println("Node", n.GetId(), "Child", n.Childrens[i])
	}
}

func (n *Node) PrintKeys() {
	for i := range n.Keys {
		fmt.Println("Node", n.GetId(), "Key", string(n.Keys[i]))
	}
}

func (l *Leaf) PrintKeys() {
	for i := range l.Keys {
		fmt.Println("Leaf", l.GetId(), "Key", string(l.Keys[i]))
	}
}

func (t *Btree) PrintInfo() {
	fmt.Println("Root", t.GetRoot())
	fmt.Println("IndexCursor", t.GetIndexCursor())
	fmt.Println("LeafCount", *t.LeafCount)
	fmt.Println("NodeCount", *t.NodeCount)
}

func (t *Btree) PrintTree() {
	fmt.Println("-----------Tree-------------")
	for i := 0; i < int(t.GetIndexCursor()); i++ {
		if node, ok := t.nodes[i].(*Node); ok {
			node.PrintKeys()
			node.PrintChildrens()
		}
		if leaf, ok := t.nodes[i].(*Leaf); ok {
			leaf.PrintKeys()
		}
		fmt.Println("AA")
	}
}
