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

package kafka

import (
	"net/url"
	"sync"
	"time"

	kafka_common "github.com/AliyunContainerService/kube-eventer/common/kafka"
	event_core "github.com/AliyunContainerService/kube-eventer/core"
	kube_api "k8s.io/api/core/v1"
)

type KafkaSinkPoint struct {
	EventValue     interface{}
	EventTimestamp time.Time
	EventTags      map[string]string
}

type kafkaSink struct {
	kafka_common.KafkaClient
	sync.RWMutex
}

func getEventValue(event *kube_api.Event) (string, error) {
	_ = "STUB: not implemented"
	// TODO: check whether indenting is required.
	return "", nil
}

func eventToPoint(event *kube_api.Event) (*KafkaSinkPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sink *kafkaSink) ExportEvents(eventBatch *event_core.EventBatch) {
	_ = "STUB: not implemented"
	return
}

func NewKafkaSink(uri *url.URL) (event_core.EventSink, error) {
	_ = "STUB: not implemented"
	return *new(event_core.EventSink), nil
}
