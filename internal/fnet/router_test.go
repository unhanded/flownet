package fnet_test

import (
	"encoding/json"
	"testing"

	"github.com/unhanded/flownet/internal/fnet"
	"github.com/unhanded/luid/pkg/luid"
)

func TestFRouter(t *testing.T) {
	rtr := fnet.FRouter{}
	n := TstNodes()
	rtr.AddNodes(n...)

	r := RouteImplementation{ids: []string{n[0].Id(), n[1].Id()}}

	result, err := rtr.Eval(&r)
	if err != nil {
		t.Errorf("Error: %s", err)
	} else if result == nil {
		t.Errorf("Result is nil")
	}

	if result.NodesPassed != 2 {
		t.Errorf("NodesPassed is not 1")
	}
	b, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("Result:\n%s", b)

	flow := result.RelativeFlow()
	b, _ = json.MarshalIndent(flow, "", "  ")
	t.Logf("Flow:\n%s", b)
}

func TstNodes() []fnet.FNode {
	n1 := &NodeImplementation{id: luid.New().String(), val: 100}
	n2 := &NodeImplementation{id: luid.New().String(), val: 180}
	return []fnet.FNode{n1, n2}
}

type NodeImplementation struct {
	id  string
	val float64
}

func (n *NodeImplementation) Id() string {
	if n.id == "" {
		n.id = luid.New().String()
	}
	return n.id
}

func (n *NodeImplementation) GetTimeoutDuration(fnet.FRoute) float64 {
	return n.val
}

func (n *NodeImplementation) Name() string {
	return "NodeImplementation"
}

type RouteImplementation struct {
	ids []string
}

func (r *RouteImplementation) NodeIds() []string {
	ids := []string{}
	for _, id := range r.ids {
		ids = append(ids, id)
	}
	return ids
}

func (r *RouteImplementation) Attributes() fnet.Attributes {
	return map[string]uint32{"A": 1}
}
