// Code generated by protoc-gen-grpc-gateway-gosdk. DO NOT EDIT.
//    protoc-gen-grpc-gateway-gosdk 1.8.1-next-f3e4de
// type: fake_bv1_client

package fake

import (
	"github.com/jaronnie/autosdk/rest"
	"github.com/jaronnie/autosdk/typed/bv1"
)

type FakeBv1 struct{}

func (f *FakeBv1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (f *FakeBv1) User() bv1.UserInterface {
	return &FakeUser{Fake: f}
}