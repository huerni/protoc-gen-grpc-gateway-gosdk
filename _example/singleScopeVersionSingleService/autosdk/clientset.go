// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package autosdk

import (
	"github.com/jaronnie/autosdk/rest"
	"github.com/jaronnie/autosdk/typed"
	corev1 "github.com/jaronnie/autosdk/typed/corev1"
)

type Interface interface {
	Direct() typed.DirectInterface

	Corev1() corev1.Corev1Interface
}

type Clientset struct {
	// direct client to request
	direct *typed.DirectClient

	corev1 *corev1.Corev1Client
}

func (x *Clientset) Direct() typed.DirectInterface {
	return x.direct
}

func (x *Clientset) Corev1() corev1.Corev1Interface {
	return x.corev1
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
	cs.corev1, err = corev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return &cs, nil
}
