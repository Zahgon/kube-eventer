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

package elasticsearch

import (
	"net/url"
	"sync"
	"time"

	esCommon "github.com/AliyunContainerService/kube-eventer/common/elasticsearch"
	event_core "github.com/AliyunContainerService/kube-eventer/core"
	"github.com/prometheus/client_golang/prometheus"
	kube_api "k8s.io/api/core/v1"
)

const (
	typeName = "events"
)

// SaveDataFunc is a pluggable function to enforce limits on the object
type SaveDataFunc func(date time.Time, namespace string, sinkData []interface{}) error

type elasticSearchSink struct {
	esSvc     esCommon.ElasticSearchService
	saveData  SaveDataFunc
	flushData func() error
	sync.RWMutex
	errorRate prometheus.Gauge
}

type EsSinkPoint struct {
	Count                    interface{}
	Metadata                 interface{}
	InvolvedObject           interface{}
	Source                   interface{}
	FirstOccurrenceTimestamp time.Time
	LastOccurrenceTimestamp  time.Time
	Message                  string
	Reason                   string
	Type                     string
	EventTags                map[string]string
}

func eventToPoint(event *kube_api.Event, clusterName string) (*EsSinkPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Part of k8s resources FirstOccurrenceTimestamp/LastOccurrenceTimestamp is nil

func (sink *elasticSearchSink) ExportEvents(eventBatch *event_core.EventBatch) {
	_ = "STUB: not implemented"
	return
}

func (sink *elasticSearchSink) Name() string { _ = "STUB: not implemented"; return "" }

func (sink *elasticSearchSink) Stop() {
	_ = "STUB: not implemented"
	// nothing needs to be done.
	return
}

func NewElasticSearchSink(uri *url.URL) (event_core.EventSink, error) {
	_ = "STUB: not implemented"
	return *new(event_core.EventSink), nil
}
