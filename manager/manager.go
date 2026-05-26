// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package manager

import (
	"time"

	"github.com/AliyunContainerService/kube-eventer/core"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Last time of eventer housekeep since unix epoch in seconds
	lastHousekeepTimestamp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "eventer",
			Subsystem: "manager",
			Name:      "last_time_seconds",
			Help:      "Last time of eventer housekeep since unix epoch in seconds.",
		})

	// Time of latest scrape operation
	LatestScrapeTime = time.Now()
)

func init() {
	prometheus.MustRegister(lastHousekeepTimestamp)
}

type Manager interface {
	Start()
	Stop()
}

type realManager struct {
	source    core.EventSource
	sink      core.EventSink
	frequency time.Duration
	stopChan  chan struct{}
}

func NewManager(source core.EventSource, sink core.EventSink, frequency time.Duration) (Manager, error) {
	_ = "STUB: not implemented"
	return *new(Manager), nil
}

func (rm *realManager) Start() { _ = "STUB: not implemented"; return }

func (rm *realManager) Stop() { _ = "STUB: not implemented"; return }

func (rm *realManager) Housekeep() {
	_ = "STUB: not implemented"

	// Try to invoke housekeep at fixed time.
	return
}

func (rm *realManager) housekeep() { _ = "STUB: not implemented"; return }

// No parallelism. Assumes that the events are pushed to Heapster. Add parallelism
// when this stops to be true.
