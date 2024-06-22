package fnet

import (
	"github.com/unhanded/flownet/internal/fnet"
	"github.com/unhanded/flownet/internal/ifnet"
)

func New() ifnet.FNet {
	return &fnet.FNetImpl{}
}
