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

package influxdb

import (
	"net/url"
	"sync"

	influxdb_common "github.com/AliyunContainerService/kube-eventer/common/influxdb"
	"github.com/AliyunContainerService/kube-eventer/core"
	influxdb "github.com/influxdata/influxdb/client"
	kube_api "k8s.io/api/core/v1"
)

type influxdbSink struct {
	client influxdb_common.InfluxdbClient
	sync.RWMutex
	c        influxdb_common.InfluxdbConfig
	dbExists bool
}

const (
	eventMeasurementName = "log/events"
	// Event special tags
	eventUID = "uid"
	// Value Field name
	valueField = "value"
	// Event special tags
	dbNotFoundError = "database not found"

	// Maximum number of influxdb Points to be sent in one batch.
	maxSendBatchSize = 1000
)

func (sink *influxdbSink) resetConnection() { _ = "STUB: not implemented"; return }

// Generate point value for event
func getEventValue(event *kube_api.Event) (string, error) {
	_ = "STUB: not implemented"
	// TODO: check whether indenting is required.
	return "", nil
}

func eventToPointWithFields(event *kube_api.Event) (*influxdb.Point, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func eventToPoint(event *kube_api.Event) (*influxdb.Point, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sink *influxdbSink) ExportEvents(eventBatch *core.EventBatch) {
	_ = "STUB: not implemented"
	return
}

func (sink *influxdbSink) sendData(dataPoints []influxdb.Point) { _ = "STUB: not implemented"; return }

func (sink *influxdbSink) Name() string { _ = "STUB: not implemented"; return "" }

func (sink *influxdbSink) Stop() {
	_ = "STUB: not implemented"
	// nothing needs to be done.
	return
}

func (sink *influxdbSink) createDatabase() error { _ = "STUB: not implemented"; return nil }

// We want to return error only if it is not "already exists" error.

func (sink *influxdbSink) createRetentionPolicy() error { _ = "STUB: not implemented"; return nil }

// We want to return error only if it is not "already exists" error.

// Returns a thread-safe implementation of core.EventSink for InfluxDB.
func newSink(c influxdb_common.InfluxdbConfig) core.EventSink {
	_ = "STUB: not implemented"
	return *new(core.EventSink)
}

// can be nil

func CreateInfluxdbSink(uri *url.URL) (core.EventSink, error) {
	_ = "STUB: not implemented"
	return *new(core.EventSink), nil
}
