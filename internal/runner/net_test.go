package runner_test

import (
	"encoding/json"
	"testing"

	"github.com/unhanded/flownet/internal/runner"
	"github.com/unhanded/flownet/pkg/flownet"
	"github.com/unhanded/luid/pkg/luid"
)

func TestFNet(t *testing.T) {
	fnt := runner.New[DummyData]()
	n := GetTestNodes()
	fnt.AddNodes(n...)

	r := RouteImplementation{ids: []string{n[0].Id(), n[1].Id()}}

	result, err := fnt.Eval(&r)
	if err != nil {
		t.Errorf("Error: %s", err)
	} else if result == nil {
		t.Errorf("Result is nil")
	}
	if len(result.RelativeFlow()) != 2 {
		t.Errorf("Expected 2 relative flow responses, got %d", len(result.RelativeFlow()))
	}
	b, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("Result:\n%s", b)

	flow := result.RelativeFlow()
	b, _ = json.MarshalIndent(flow, "", "  ")
	t.Logf("Flow:\n%s", b)

	data := n[0].Data()
	b, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("Data:\n%s", b)
}

func TestFNetError(t *testing.T) {
	fnt := runner.New[DummyData]()
	n := GetTestNodes()
	fnt.AddNodes(n...)

	r := RouteImplementation{ids: []string{}}

	_, err := fnt.Eval(&r)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestFNetRemove(t *testing.T) {
	fnt := runner.New[DummyData]()
	n := GetTestNodes()
	fnt.AddNodes(n...)
	if len(fnt.Nodes()) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(fnt.Nodes()))
	}
	fnt.RemoveNode(n[0].Id())

	if len(fnt.Nodes()) != 1 {
		t.Errorf("Expected 1 nodes after delete, got %d", len(fnt.Nodes()))
	}
}

// Convenience function to make a couple of test node.
// Uses the example implementation below.
func GetTestNodes() []flownet.Node[DummyData] {
	n1 := &SimpleNodeImplementation{id: luid.New().String(), val: 100}
	n2 := &SimpleNodeImplementation{id: luid.New().String(), val: 180}
	return []flownet.Node[DummyData]{n1, n2}
}

// SimpleNodeImplementation is a highly simplified implementation of a node.
type SimpleNodeImplementation struct {
	id  string
	val float64
}

func (n *SimpleNodeImplementation) Id() string {
	if n.id == "" {
		n.id = luid.New().String()
	}
	return n.id
}

func (n *SimpleNodeImplementation) GetResistance(r flownet.Probe) float64 {
	attr := r.Attributes()
	// If we are doing a smarter node, the attributes are the parameters that we use.
	// But for now, we are just using a static value, so let's discard it.
	_ = attr
	// And simply return val.
	return n.val
}

// Name is also mandatory, but we are not using it in this example.
func (n *SimpleNodeImplementation) Name() string {
	return "NodeImplementation"
}

func (n *SimpleNodeImplementation) Data() DummyData {
	return DummyData{VeryInterestingField: "superFascinatingValue"}
}

// RouteImplementation is a highly simplified implementation of a route structure.
// The route is just a sequence of node ids and attributes where applicable.
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

func (r *RouteImplementation) Attributes() flownet.Attributes {
	return map[string]uint32{"A": 1}
}

type DummyData struct {
	VeryInterestingField string `json:"veryInterestingField"`
}
