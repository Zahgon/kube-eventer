package filters

import (
	v1 "k8s.io/api/core/v1"
)

// All filter interface
type Filter interface {
	Filter(event *v1.Event) (matched bool)
}

func GetValues(o []string) []string { _ = "STUB: not implemented"; return nil }
