package flownet

// FNet is the central interface for flownet.
type FNet[T any] interface {
	// AddNodes adds nodes to the network, as one might expect.
	AddNodes(...Node[T]) error
	// Eval evaluates a route through the network.
	Eval(Probe) (RouteResult, error)
	// Nodes() returns all nodes in the network.
	Nodes() []Node[T]
	// RemoveNode discards a node in the network by node id.
	RemoveNode(string) error
}

type RouteResult interface {
	Responses() []Response
	RelativeFlow() []Response
}

type Response interface {
	Id() string
	Value() float64
}

type Attributes map[string]uint32

type Node[T any] interface {
	Name() string
	Id() string
	GetResistance(Probe) float64
	Data() T
}

type Probe interface {
	NodeIds() []string
	Attributes() Attributes
}
