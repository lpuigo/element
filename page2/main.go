package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/element/model"
	"github.com/lpuig/element/model/node"
)

//go:generate bash ./makejs.sh

func main() {
	go aNodeTree()
}

/////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////////////////

type DemoElement struct {
	*js.Object

	Nodes     []*node.Node `js:"nodes"`
	NodeProps js.M         `js:"nodeProps"`

	Selected *node.Node `js:"selected"`

	VM *hvue.VM `js:"vm"`
}

func NewDemoElement() *DemoElement {
	de := &DemoElement{Object: model.O()}
	de.Nodes = createNodeTree()
	de.NodeProps = js.M{
		"children": "children",
		"label":    "label",
	}
	de.Selected = nil
	return de
}

func (de *DemoElement) HandleNodeClick(n *node.Node) {
	de.Selected = n
}

func createNodeTree() []*node.Node {
	res := []*node.Node{}
	res = append(res, node.NewNode("Node 1",
		[]*node.Node{
			node.NewNode("Node 1-1", nil),
			node.NewNode("Node 1-2", nil),
		}))
	res = append(res, node.NewNode("Node 2", nil))
	res = append(res, node.NewNode("Node 3", nil))
	return res
}

func aNodeTree() {
	de := NewDemoElement()

	hvue.NewVM(
		hvue.El("#app"),
		hvue.DataS(de),
		hvue.MethodsOf(de),
	)

	js.Global.Set("de", de)
}

////////////////////////////////////////////////////////////////////////////////////
