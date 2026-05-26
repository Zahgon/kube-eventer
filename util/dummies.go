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

package util

import (
	"sync"
	"time"

	"github.com/AliyunContainerService/kube-eventer/core"
)

type DummySink struct {
	name        string
	mutex       sync.Mutex
	exportCount int
	stopped     bool
	latency     time.Duration
}

func (this *DummySink) Name() string { _ = "STUB: not implemented"; return "" }

func (this *DummySink) ExportEvents(*core.EventBatch) { _ = "STUB: not implemented"; return }

func (this *DummySink) Stop() { _ = "STUB: not implemented"; return }

func (this *DummySink) IsStopped() bool { _ = "STUB: not implemented"; return false }

func (this *DummySink) GetExportCount() int { _ = "STUB: not implemented"; return 0 }

func NewDummySink(name string, latency time.Duration) *DummySink {
	_ = "STUB: not implemented"
	return nil
}

type DummyEventSource struct {
	eventBatch *core.EventBatch
}

func (this *DummyEventSource) GetNewEvents() *core.EventBatch {
	_ = "STUB: not implemented"
	return nil
}

func NewDummySource(eventBatch *core.EventBatch) *DummyEventSource {
	_ = "STUB: not implemented"
	return nil
}
