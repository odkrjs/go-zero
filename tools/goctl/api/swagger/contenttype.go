package swagger

import (
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func consumesFromTypeOrDef(method string, tp spec.Type) []string {
	if strings.EqualFold(method, http.MethodGet) {
		return []string{}
	}
	if tp == nil {
		return []string{}
	}
	structType, ok := tp.(spec.DefineStruct)
	if !ok {
		return []string{}
	}
	if typeContainsTag(structType, tagJson) {
		return []string{applicationJson}
	}
	return []string{applicationForm}
}
