// Copyright 2017 Google Inc. All Rights Reserved.
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

package riemann

import (
	"net/url"
	"sync"
	"time"

	riemannCommon "github.com/AliyunContainerService/kube-eventer/common/riemann"
	"github.com/AliyunContainerService/kube-eventer/core"
	kube_api "k8s.io/api/core/v1"
)

// contains the riemann client, the riemann configuration, and a RWMutex
type RiemannSink struct {
	client riemanngo.Client
	config riemannCommon.RiemannConfig
	sync.RWMutex
}

// creates a Riemann sink. Returns a riemannSink
func CreateRiemannSink(uri *url.URL) (core.EventSink, error) {
	_ = "STUB: not implemented"
	return *new(core.EventSink), nil
}

// Return a user-friendly string describing the sink
func (sink *RiemannSink) Name() string { _ = "STUB: not implemented"; return "" }

func (sink *RiemannSink) Stop() { _ = "STUB: not implemented"; return }

func getEventState(event *kube_api.Event) string { _ = "STUB: not implemented"; return "" }

func appendEvent(events []riemanngo.Event, sink *RiemannSink, event *kube_api.Event, timestamp time.Time) []riemanngo.Event {
	_ = "STUB: not implemented"
	return nil
}

// client will reconnect later

func (sink *RiemannSink) ExportEvents(eventBatch *core.EventBatch) {
	_ = "STUB: not implemented"
	return
}

// the client could be nil here, so we reconnect

// creates an event and add it to dataEvent

// client will reconnect later
