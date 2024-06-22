package ifnet

// FNet is the central interface for flownet.
type FNet interface {
	// AddNodes adds nodes to the network, as one might expect.
	AddNodes(...Node) error
	// Eval evaluates a route through the network.
	Eval(Route) (RouteResult, error)
	// Nodes() returns all nodes in the network.
	Nodes() []Node
	// RemoveNode discards a node in the network by node id.
	RemoveNode(string) error
}

type RouteResult interface {
	TimeoutResponses() []Response
	RelativeFlow() []Response
}

type Response interface {
	Id() string
	Value() float64
}

type Attributes map[string]uint32

type Node interface {
	Name() string
	Id() string
	GetTimeoutDuration(Route) float64
}

type Route interface {
	NodeIds() []string
	Attributes() Attributes
}
