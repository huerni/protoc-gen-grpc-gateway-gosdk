// Code generated by protoc-gen-grpc-gateway-go. DO NOT EDIT.
// versions:
//    protoc-gen-grpc-gateway-go 1.7.1-next-bad862
// type: fake_clientset

package fake

import (
	"github.com/jaronnie/autosdk/typed"
	corev1 "github.com/jaronnie/autosdk/typed/corev1"
	fakecorev1 "github.com/jaronnie/autosdk/typed/corev1/fake"
	"github.com/jaronnie/autosdk/typed/fake"
)

type Clientset struct{}

func (f *Clientset) Direct() typed.DirectInterface {
	return &fake.FakeDirect{}
}

func (f *Clientset) Corev1() corev1.Corev1Interface {
	return &fakecorev1.FakeCorev1{}
}
