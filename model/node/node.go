package node

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lpuig/element/model"
)

var id int

type Node struct {
	*js.Object

	Id       int     `js:"id"`
	Label    string  `js:"label"`
	Children []*Node `js:"children"`
}

func NewNode(label string, children []*Node) *Node {
	n := &Node{Object: model.O()}
	n.Id = id
	n.Label = label
	n.Children = children
	id++
	return n
}
