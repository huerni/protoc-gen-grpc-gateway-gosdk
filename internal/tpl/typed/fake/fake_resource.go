package fake

import "github.com/jaronnie/protoc-gen-grpc-gateway-go/internal/vars"

type FakeResourceData struct {
	Gateways []*vars.Gateway

	IsWarpHttpResponse bool     // is warped code,data,message
	GoModule           string   // github.com/jaronnie/autosdk
	GoImportPaths      []string // pb import path [github.com/jaronnie/autosdk/pb/corev1]
	ScopeVersion       string   // corev1
	UpScopeVersion     string   // Corev1
	Resource           string   // credential
	UpResource         string   // Credential
}

var FakeResourceTpl = `
// Code generated by protoc-gen-grpc-gateway-go. DO NOT EDIT.
// versions:
//    protoc-gen-grpc-gateway-go {{getProtoToolsVersion "protoc-gen-grpc-gateway-go"}}
// type: fake_{{.Resource}}

package fake

import (
	"context"

	{{range $k, $v := .Gateways}}{{if or $v.IsStreamServer $v.IsStreamClient $v.IsSpecified}}"{{$.GoModule}}/rest"{{break}}{{end}}{{end}}

	{{range $v := .GoImportPaths}}"{{$v}}"
	{{end}}
)

var (
	{{range $k, $v := .Gateways}}FakeReturn{{$v.FuncName}} = &{{if or $v.IsStreamServer $v.IsStreamClient $v.IsSpecified}}rest.Request{}{{else}}{{$v.HttpResponseBody.RootPath}}.{{$v.HttpResponseBody.Name}}{}{{end}}
	{{end}}
)

type {{.UpResource}}Getter interface {
	{{.UpResource}}() {{.UpResource}}Interface

	Fake{{.UpResource}}Expansion
}

type {{.UpResource}}Interface interface {
	{{range $k, $v := .Gateways}}{{$v.FuncName}}({{if or $v.IsStreamServer $v.IsStreamClient $v.IsSpecified}}{{else}}ctx context.Context,{{end}} param *{{$v.ProtoRequestBody.RootPath}}.{{$v.ProtoRequestBody.Name}}) ({{if or $v.IsStreamServer $v.IsStreamClient $v.IsSpecified}}*rest.Request{{else}}*{{$v.HttpResponseBody.RootPath}}.{{$v.HttpResponseBody.Name}}{{end}}, error)
	{{end}}
}

type Fake{{.UpResource}} struct {
	Fake *Fake{{.UpScopeVersion}}
}

{{range $k, $v := .Gateways}}func (f *Fake{{$.UpResource}}) {{$v.FuncName}}({{if or $v.IsStreamServer $v.IsStreamClient $v.IsSpecified}}{{else}}ctx context.Context,{{end}}param *{{$v.ProtoRequestBody.RootPath}}.{{$v.ProtoRequestBody.Name}}) ({{if or $v.IsStreamServer $v.IsStreamClient $v.IsSpecified}}*rest.Request{{else}}*{{$v.HttpResponseBody.RootPath}}.{{$v.HttpResponseBody.Name}}{{end}}, error) {
	return FakeReturn{{$v.FuncName}}, nil
}

{{end}}
`
