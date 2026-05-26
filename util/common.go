package util

import (
	"time"

	v1 "k8s.io/api/core/v1"
)

func GetLastEventTimestamp(event *v1.Event) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
