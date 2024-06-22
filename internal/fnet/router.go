package fnet

import (
	"fmt"
)

type FRouter struct {
	nodes []FNode
}

func (f *FRouter) Eval(r FRoute) (*FRouteResult, error) {
	rr := &FRouteResult{NodesPassed: 0, Responses: []FTimeoutResponse{}}

	for _, nodeId := range r.NodeIds() {
		for _, node := range f.nodes {
			if node.Id() == nodeId {
				rr.NodesPassed++
				rr.Responses = append(
					rr.Responses,
					FTimeoutResponse{NodeId: nodeId, TimeoutDuration: node.GetTimeoutDuration(r)},
				)
			}
		}
	}
	if !rr.validate() {
		return nil, fmt.Errorf("invalid route result")
	}
	return rr, nil
}

func (f *FRouter) AddNodes(nodes ...FNode) error {
	for _, node := range nodes {
		f.nodes = append(f.nodes, node)
	}
	return nil
}

func (f *FRouter) RemoveNode(node FNode) error {
	for i, n := range f.nodes {
		if n.Id() == node.Id() {
			f.nodes = append(f.nodes[:i], f.nodes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("node not found")
}

func (f *FRouter) Nodes() []FNode {
	return f.nodes
}
