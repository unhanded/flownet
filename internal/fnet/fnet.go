package fnet

import (
	"fmt"

	"github.com/unhanded/flownet/pkg/flownet"
)

type Network[T any] struct {
	nodes   []flownet.Node[T]
	preLoad []flownet.Probe
}

func (f *Network[T]) Eval(r flownet.Probe) (flownet.RouteResult, error) {
	rr := &ProbeResult{NodesPassed: 0, ProbeResponses: []flownet.Response{}}

	for _, nodeId := range r.NodeIds() {
		for _, node := range f.nodes {
			if node.Id() == nodeId {
				rr.NodesPassed++
				rr.ProbeResponses = append(
					rr.ProbeResponses,
					Resistance{NodeId: nodeId, Res: node.GetResistance(r)},
				)
			}
		}
	}
	if !rr.validate() {
		return nil, fmt.Errorf("invalid route result")
	}
	return rr, nil
}

func (f *Network[T]) AddNodes(nodes ...flownet.Node[T]) error {
	for _, node := range nodes {
		f.nodes = append(f.nodes, node)
	}
	return nil
}

func (f *Network[T]) RemoveNode(nodeId string) error {
	for i, n := range f.nodes {
		if n.Id() == nodeId {
			f.nodes = append(f.nodes[:i], f.nodes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("node not found")
}

func (f *Network[T]) Nodes() []flownet.Node[T] {
	return f.nodes
}
