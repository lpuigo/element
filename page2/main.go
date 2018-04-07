package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/element/model"
	"github.com/lpuig/element/model/node"
	"github.com/lpuig/element/component"
)

//go:generate bash ./makejs.sh

func main() {
	go aNodeTree()
}

/////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////////////////

type DemoElement struct {
	*js.Object

	Nodes     []*node.HoursNode `js:"nodes"`
	NodeProps js.M         `js:"nodeProps"`

	Selected *node.HoursNode `js:"selected"`

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

func (de *DemoElement) HandleNodeClick(n *node.HoursNode) {
	de.Selected = n
}

func createNodeTree() []*node.HoursNode {
	hours := []float64{1,2,3,4,5,6,7,8,9,0,1,2,3,4,5,6,7,8,9,0}
	res := make([]*node.HoursNode, 0)

	hn10 := node.NewHoursNode("Node 1-1", hours)
	hn11 := node.NewHoursNode("Node 1-2", hours)
	hn1 := node.NewHoursNode("Node 1", nil)
	hn1.AddChild(hn10)
	hn1.AddChild(hn11)
	res = append(res, hn1)
	res = append(res, node.NewHoursNode("Node 2", hours))
	res = append(res, node.NewHoursNode("Node 3", hours))
	println(res)
	return res
}

func aNodeTree() {
	de := NewDemoElement()
	component.NewColoredTableRowComp()

	hvue.NewVM(
		hvue.El("#app"),
		hvue.DataS(de),
		hvue.MethodsOf(de),
	)

	js.Global.Set("de", de)
}

////////////////////////////////////////////////////////////////////////////////////
