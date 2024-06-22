package fnet

// FRouteResult is the result of evaluating a route.
type FRouteResult struct {
	// NodesPassed is the number of nodes that was passed while evaluating.
	NodesPassed int64
	// Responses is the responses produced by checking the nodes along the route.
	Responses []FTimeoutResponse
}

// Validate checks if the RouteResult is valid, used internally.
func (rr *FRouteResult) validate() bool {
	if rr.NodesPassed != int64(len(rr.Responses)) {
		return false
	}
	return true
}

func (rr *FRouteResult) TimeoutResponses() []FTimeoutResponse {
	return rr.Responses
}

func (rr *FRouteResult) RelativeFlow() []FRelativeFlow {
	fastest := rr.routeFastest()
	var flows []FRelativeFlow
	for _, r := range rr.Responses {
		flows = append(flows, FRelativeFlow{NodeId: r.NodeId, Flow: fastest / r.TimeoutDuration})
	}
	return flows
}

func (rr *FRouteResult) routeFastest() float64 {
	var min float64 = 99999999
	for _, r := range rr.Responses {
		if r.TimeoutDuration < min {
			min = r.TimeoutDuration
		}
	}
	return min
}

type FRoute interface {
	NodeIds() []string
	Attributes() Attributes
}
