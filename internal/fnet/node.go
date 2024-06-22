package fnet

type FNode interface {
	Name() string
	Id() string
	GetTimeoutDuration(FRoute) float64
}
