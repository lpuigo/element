package node

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lpuig/element/model"
)

/*
type HoursNode struct {
	*Node

	Hours []float64 `js:"hours"`
}

func NewHoursNode(label string, hours []float64) *HoursNode {
	n := &HoursNode{Node: NewNode(label, nil)}
	n.Hours = hours
	return n
}

func (hn *HoursNode) AddChild(c *HoursNode) {
	hn.Node.AddChild(c)

	if len(hn.Hours) == 0 {
		hn.Hours = make([]float64, len(c.Hours))
	}

	for i, h := range c.Hours {
		hn.Hours[i] += h
	}
}
*/


var nhid int

type HoursNode struct {
	*js.Object

	Id       int     `js:"id"`
	Label    string  `js:"label"`
	Children []*HoursNode `js:"children"`

	Hours []float64 `js:"hours"`
}

func NewHoursNode(label string, hours []float64) *HoursNode {
	hn := &HoursNode{Object: model.O()}
	hn.Id = nhid
	hn.Label = label
	hn.Children = []*HoursNode{}
	hn.Hours = hours
	nhid++
	return hn
}

func (hn *HoursNode) AddChild(c *HoursNode) {
	hn.Children = append(hn.Children, c)
	if len(hn.Hours) == 0 {
		hn.Hours = make([]float64, len(c.Hours))
	}

	for i, h := range c.Hours {
		//hn.Hours[i] += h // Gopherjs style :(
		o := hn.Object.Get("hours")
		o.SetIndex(i, o.Index(i).Float()+h)
	}
}
