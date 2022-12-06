// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package fake

import (
	"context"

	"github.com/jaronnie/autosdk/pb/corev1"
)

var (
	FakeReturnInitCredential   = &corev1.Credential{}
	FakeReturnUpdateCredential = &corev1.Credential{}
)

type CredentialGetter interface {
	Credential() CredentialInterface
}

type CredentialInterface interface {
	InitCredential(ctx context.Context, param *corev1.Credential) (*corev1.Credential, error)
	UpdateCredential(ctx context.Context, param *corev1.Credential) (*corev1.Credential, error)
}

type FakeCredential struct {
	Fake *FakeCorev1
}

func (f *FakeCredential) InitCredential(ctx context.Context, param *corev1.Credential) (*corev1.Credential, error) {
	return FakeReturnInitCredential, nil
}

func (f *FakeCredential) UpdateCredential(ctx context.Context, param *corev1.Credential) (*corev1.Credential, error) {
	return FakeReturnUpdateCredential, nil
}
