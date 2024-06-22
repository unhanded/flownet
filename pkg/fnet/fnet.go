package fnet

import (
	"github.com/unhanded/flownet/internal/fnet"
)

type FNet *fnet.FRouter

func New() FNet {
	return &fnet.FRouter{}
}
