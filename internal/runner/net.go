package runner

import (
	"github.com/unhanded/flownet/internal/fnet"
	"github.com/unhanded/flownet/pkg/flownet"
)

func New[T any]() flownet.FNet[T] {
	return &fnet.Network[T]{}
}
