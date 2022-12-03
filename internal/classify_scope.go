package internal

import (
	"github.com/jaronnie/protoc-gen-go-httpsdk/internal/vars"
	"github.com/jaronnie/protoc-gen-go-httpsdk/utilx"
)

// getAllScopeVersionsMap {"corev1":"CoreV1", "corev2":"CoreV2"}
func getAllScopeVersionsMap(scopeResourceGws vars.ScopeResourceGateway) map[string]string {
	scopeVersionMap := make(map[string]string, 0)

	for k, _ := range scopeResourceGws {
		scopeVersionMap[string(k)] = utilx.FirstUpper(string(k))
	}
	return scopeVersionMap
}
