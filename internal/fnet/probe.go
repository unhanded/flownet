package fnet

import "github.com/unhanded/flownet/pkg/flownet"

// ProbeResult is the result of evaluating a route.
type ProbeResult struct {
	// NodesPassed is the number of nodes that was passed while evaluating.
	NodesPassed int64
	// Responses is the responses produced by checking the nodes along the route.
	ProbeResponses []flownet.Response
}

// Validate checks if the RouteResult is valid, used internally.
func (rr *ProbeResult) validate() bool {
	if rr.NodesPassed != int64(len(rr.ProbeResponses)) {
		return false
	}
	if rr.NodesPassed == 0 {
		return false
	}
	return true
}

func (rr *ProbeResult) Responses() []flownet.Response {
	return rr.ProbeResponses
}

func (rr *ProbeResult) RelativeFlow() []flownet.Response {
	slowest := rr.probeSlowest()
	var flows []flownet.Response
	for _, r := range rr.ProbeResponses {
		flows = append(
			flows,
			&RelativeResistance{
				NodeId: r.Id(),
				Flow:   slowest / r.Value(),
			},
		)
	}
	return flows
}

func (rr *ProbeResult) probeSlowest() float64 {
	var max float64 = 0
	for _, r := range rr.ProbeResponses {
		if r.Value() > max {
			max = r.Value()
		}
	}
	return max
}
