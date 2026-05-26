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

package logsink

import (
	"github.com/AliyunContainerService/kube-eventer/core"
)

type LogSink struct {
}

func (this *LogSink) Name() string { _ = "STUB: not implemented"; return "" }

func (this *LogSink) Stop() {
	_ = "STUB: not implemented"
	// Do nothing.
	return
}

func batchToString(batch *core.EventBatch) string { _ = "STUB: not implemented"; return "" }

func (this *LogSink) ExportEvents(batch *core.EventBatch) { _ = "STUB: not implemented"; return }

func CreateLogSink() (*LogSink, error) { _ = "STUB: not implemented"; return nil, nil }
