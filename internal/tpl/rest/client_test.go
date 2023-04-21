package rest

import (
	"fmt"
	"testing"

	"github.com/jaronnie/protoc-gen-grpc-gateway-go/utilx"
)

func TestGenerateRestClien(t *testing.T) {
	b, err := utilx.ParseTemplate(nil, []byte(ClientTpl))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
