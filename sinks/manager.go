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

package sinks

import (
	"time"

	"github.com/AliyunContainerService/kube-eventer/core"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	DefaultSinkExportEventsTimeout = 20 * time.Second
	DefaultSinkStopTimeout         = 60 * time.Second
)

var (
	// Time spent exporting events to sink in milliseconds.
	exporterDuration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: "eventer",
			Subsystem: "exporter",
			Name:      "duration_milliseconds",
			Help:      "Time spent exporting events to sink in milliseconds.",
		},
		[]string{"exporter"},
	)
)

func init() {
	prometheus.MustRegister(exporterDuration)
}

type sinkHolder struct {
	sink              core.EventSink
	eventBatchChannel chan *core.EventBatch
	stopChannel       chan bool
}

// Sink Manager - a special sink that distributes data to other sinks. It pushes data
// only to these sinks that completed their previous exports. Data that could not be
// pushed in the defined time is dropped and not retried.
type sinkManager struct {
	sinkHolders         []sinkHolder
	exportEventsTimeout time.Duration
	// Should be larger than exportEventsTimeout, although it is not a hard requirement.
	stopTimeout time.Duration
}

func NewEventSinkManager(sinks []core.EventSink, exportEventsTimeout, stopTimeout time.Duration) (core.EventSink, error) {
	_ = "STUB: not implemented"
	return *new(core.EventSink), nil
}

// ExportEvents Guarantees that the export will complete in exportEventsTimeout.
func (this *sinkManager) ExportEvents(data *core.EventBatch) { _ = "STUB: not implemented"; return }

// everything ok

// Wait for all pushes to complete or timeout.

func (this *sinkManager) Name() string { _ = "STUB: not implemented"; return "" }

func (this *sinkManager) Stop() { _ = "STUB: not implemented"; return }

// everything ok

func export(s core.EventSink, data *core.EventBatch) { _ = "STUB: not implemented"; return }
