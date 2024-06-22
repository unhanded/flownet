package fnet

// FTimeoutResponse is a struct that represents a single response of a node.
type FTimeoutResponse struct {
	NodeId          string
	TimeoutDuration float64
}

// FRelativeFlow is a struct that represents the flow value of a node relative to the entire route.
type FRelativeFlow struct {
	NodeId string
	Flow   float64
}
