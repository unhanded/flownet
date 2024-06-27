package core

import (
	"github.com/unhanded/flownet/internal/runner"
	"github.com/unhanded/flownet/pkg/flownet"
)

func New[T any]() flownet.FNet[T] {
	return runner.New[T]()
}
