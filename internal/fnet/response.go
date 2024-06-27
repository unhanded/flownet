package fnet

// Resistance is a struct that represents a single response of a node.
type Resistance struct {
	NodeId string
	Res    float64
}

func (r Resistance) Id() string {
	return r.NodeId
}

func (r Resistance) Value() float64 {
	return r.Res
}

// RelativeResistance is a struct that represents the flow value of a node relative to the entire route.
type RelativeResistance struct {
	NodeId string
	Flow   float64
}

func (r RelativeResistance) Id() string {
	return r.NodeId
}

func (r RelativeResistance) Value() float64 {
	return r.Flow
}
