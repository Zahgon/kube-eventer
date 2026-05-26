package filters

import (
	"reflect"

	v1 "k8s.io/api/core/v1"
)

type GenericFilter struct {
	field  string
	keys   []string
	regexp bool
}

func IsZero(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

func (gf *GenericFilter) Filter(event *v1.Event) (matched bool) {
	_ = "STUB: not implemented"
	return false
}

// enable regexp

// Generic Filter
func NewGenericFilter(field string, keys []string, regexp bool) *GenericFilter {
	_ = "STUB: not implemented"
	return nil
}
