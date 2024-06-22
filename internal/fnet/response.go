package fnet

// FTimeoutResponse is a struct that represents a single response of a node.
type FTimeoutResponse struct {
	NodeId          string
	TimeoutDuration float64
}

func (r FTimeoutResponse) Id() string {
	return r.NodeId
}

func (r FTimeoutResponse) Value() float64 {
	return r.TimeoutDuration
}

// FRelativeFlowResponse is a struct that represents the flow value of a node relative to the entire route.
type FRelativeFlowResponse struct {
	NodeId string
	Flow   float64
}

func (r FRelativeFlowResponse) Id() string {
	return r.NodeId
}

func (r FRelativeFlowResponse) Value() float64 {
	return r.Flow
}
