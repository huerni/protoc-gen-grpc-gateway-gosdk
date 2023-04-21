// Code generated by protoc-gen-grpc-gateway-go. DO NOT EDIT.
// versions:
//    protoc-gen-grpc-gateway-go 1.7.1-next-bad862
// type: fake_credential

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

	FakeCredentialExpansion
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
