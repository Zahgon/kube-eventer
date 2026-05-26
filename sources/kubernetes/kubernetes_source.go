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

package kubernetes

import (
	"net/url"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/AliyunContainerService/kube-eventer/core"
	kubeapi "k8s.io/api/core/v1"

	kubev1core "k8s.io/client-go/kubernetes/typed/core/v1"
)

const (
	// Number of object pointers. Big enough so it won't be hit anytime soon with reasonable GetNewEvents frequency.
	LocalEventsBufferSize = 100000
)

var (
	// Last time of event since unix epoch in seconds
	lastEventTimestamp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "eventer",
			Subsystem: "scraper",
			Name:      "last_time_seconds",
			Help:      "Last time of event since unix epoch in seconds.",
		})
	totalEventsNum = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "eventer",
			Subsystem: "scraper",
			Name:      "events_total_number",
			Help:      "The total number of events.",
		})
	scrapEventsDuration = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace: "eventer",
			Subsystem: "scraper",
			Name:      "duration_milliseconds",
			Help:      "Time spent scraping events in milliseconds.",
		})
)

func init() {
	prometheus.MustRegister(lastEventTimestamp)
	prometheus.MustRegister(totalEventsNum)
	prometheus.MustRegister(scrapEventsDuration)
}

// Implements core.EventSource interface.
type KubernetesEventSource struct {
	// Large local buffer, periodically read.
	localEventsBuffer chan *kubeapi.Event

	stopChannel chan struct{}

	eventClient kubev1core.EventInterface

	exportMetric bool
}

func (this *KubernetesEventSource) GetNewEvents() *core.EventBatch {
	_ = "STUB: not implemented"
	return nil
}

// Get all data from the buffer.

func (this *KubernetesEventSource) watch() {
	_ = "STUB: not implemented"
	// Outer loop, for reconnections.
	return
}

// Do not write old events.

// Inner loop, for update processing.

// Ok, buffer not full.

// Buffer full, need to drop the event.

func NewKubernetesSource(uri *url.URL, exportMetric bool) (*KubernetesEventSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
