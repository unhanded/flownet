package fnet

import "github.com/unhanded/flownet/internal/ifnet"

// FRouteResult is the result of evaluating a route.
type FRouteResult struct {
	// NodesPassed is the number of nodes that was passed while evaluating.
	NodesPassed int64
	// Responses is the responses produced by checking the nodes along the route.
	Responses []ifnet.Response
}

// Validate checks if the RouteResult is valid, used internally.
func (rr *FRouteResult) validate() bool {
	if rr.NodesPassed != int64(len(rr.Responses)) {
		return false
	}
	return true
}

func (rr *FRouteResult) TimeoutResponses() []ifnet.Response {
	return rr.Responses
}

func (rr *FRouteResult) RelativeFlow() []ifnet.Response {
	fastest := rr.routeFastest()
	var flows []ifnet.Response
	for _, r := range rr.Responses {
		flows = append(flows, &FRelativeFlowResponse{NodeId: r.Id(), Flow: fastest / r.Value()})
	}
	return flows
}

func (rr *FRouteResult) routeFastest() float64 {
	var min float64 = 99999999
	for _, r := range rr.Responses {
		if r.Value() < min {
			min = r.Value()
		}
	}
	return min
}
