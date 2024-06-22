package fnet

import (
	"fmt"

	"github.com/unhanded/flownet/internal/ifnet"
)

type FNetImpl struct {
	nodes []ifnet.Node
}

func (f *FNetImpl) Eval(r ifnet.Route) (ifnet.RouteResult, error) {
	rr := &FRouteResult{NodesPassed: 0, Responses: []ifnet.Response{}}

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

func (f *FNetImpl) AddNodes(nodes ...ifnet.Node) error {
	for _, node := range nodes {
		f.nodes = append(f.nodes, node)
	}
	return nil
}

func (f *FNetImpl) RemoveNode(nodeId string) error {
	for i, n := range f.nodes {
		if n.Id() == nodeId {
			f.nodes = append(f.nodes[:i], f.nodes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("node not found")
}

func (f *FNetImpl) Nodes() []ifnet.Node {
	return f.nodes
}
