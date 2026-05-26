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

package honeycomb

import (
	"net/url"
	"sync"

	honeycomb_common "github.com/AliyunContainerService/kube-eventer/common/honeycomb"
	event_core "github.com/AliyunContainerService/kube-eventer/core"
	kube_api "k8s.io/api/core/v1"
)

type honeycombSink struct {
	client honeycomb_common.Client
	sync.Mutex
}

type exportedData struct {
	Namespace       string `json:"namespace"`
	Kind            string `json:"kind"`
	Name            string `json:"name"`
	SubObject       string `json:"subobject"`
	SourceComponent string `json:"source.component"`
	SourceHost      string `json:"source.host"`
	Count           int32  `json:"count"`
	Type            string `json:"type"`
	Reason          string `json:"reason"`
	Message         string `json:"message"`
}

func getExportedData(e *kube_api.Event) *exportedData { _ = "STUB: not implemented"; return nil }

func (sink *honeycombSink) ExportEvents(eventBatch *event_core.EventBatch) {
	_ = "STUB: not implemented"
	return
}

func (sink *honeycombSink) Stop() { _ = "STUB: not implemented"; return }

func (sink *honeycombSink) Name() string { _ = "STUB: not implemented"; return "" }

func NewHoneycombSink(uri *url.URL) (event_core.EventSink, error) {
	_ = "STUB: not implemented"
	return *new(event_core.EventSink), nil
}
