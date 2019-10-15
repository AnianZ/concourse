package tracing

import (
	"go.opentelemetry.io/api/core"
	"go.opentelemetry.io/api/key"
)

type Attrs map[string]string

// keyValueSlice converts our internal representation of kv pairs to the tracing
// SDK's kv representation.
//
func keyValueSlice(attrs Attrs) []core.KeyValue {
	res := make([]core.KeyValue, len(attrs))

	idx := 0
	for k, v := range attrs {
		res[idx] = key.New(k).String(v)
		idx++
	}

	return res
}
