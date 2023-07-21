// Code generated by protoc-gen-grpc-gateway-gosdk. DO NOT EDIT.
// versions:
//    protoc-gen-grpc-gateway-gosdk 1.8.1-next-f3e4de
// type: clientset

package autosdk

import (
	"github.com/jaronnie/autosdk/rest"
	"github.com/jaronnie/autosdk/typed"
	av1 "github.com/jaronnie/autosdk/typed/av1"
	bv1 "github.com/jaronnie/autosdk/typed/bv1"
)

type Interface interface {
	Direct() typed.DirectInterface

	Av1() av1.Av1Interface
	Bv1() bv1.Bv1Interface
}

type Clientset struct {
	// direct client to request
	direct *typed.DirectClient

	av1 *av1.Av1Client
	bv1 *bv1.Bv1Client
}

func (x *Clientset) Direct() typed.DirectInterface {
	return x.direct
}

func (x *Clientset) Av1() av1.Av1Interface {
	return x.av1
}

func (x *Clientset) Bv1() bv1.Bv1Interface {
	return x.bv1
}

func NewClientWithOptions(ops ...rest.Opt) (*Clientset, error) {
	c := &rest.RESTClient{}
	for _, op := range ops {
		if err := op(c); err != nil {
			return nil, err
		}
	}
	configShallowCopy := *c
	var cs Clientset
	var err error
	cs.direct, err = typed.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.av1, err = av1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.bv1, err = bv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return &cs, nil
}