package node


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
		//hn.Hours[i] += h // Gopherjs style :(
		o := hn.Object.Get("hours")
		o.SetIndex(i, o.Index(i).Float()+h)
	}
}