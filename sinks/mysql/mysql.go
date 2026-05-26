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

package mysql

import (
	"net/url"
	"sync"

	mysql_common "github.com/AliyunContainerService/kube-eventer/common/mysql"
	"github.com/AliyunContainerService/kube-eventer/core"
	kube_api "k8s.io/api/core/v1"
)

// SaveDataFunc is a pluggable function to enforce limits on the object
type SaveDataFunc func(sinkData []interface{}) error

type mysqlSink struct {
	mysqlSvc  *mysql_common.MysqlService
	saveData  SaveDataFunc
	flushData func() error
	closeDB   func() error
	sync.RWMutex
	uri *url.URL
}

const (
	// Maximum number of mysql Points to be sent in one batch.
	maxSendBatchSize = 1
)

func (sink *mysqlSink) createDatabase() error { _ = "STUB: not implemented"; return nil }

// Generate point value for event
func getEventValue(event *kube_api.Event) (string, error) {
	_ = "STUB: not implemented"
	// TODO: check whether indenting is required.
	return "", nil
}

func eventToPoint(event *kube_api.Event) (*mysql_common.MysqlKubeEventPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sink *mysqlSink) ExportEvents(eventBatch *core.EventBatch) { _ = "STUB: not implemented"; return }

func (sink *mysqlSink) Name() string { _ = "STUB: not implemented"; return "" }

func (sink *mysqlSink) Stop() { _ = "STUB: not implemented"; return }

// Returns a thread-safe implementation of core.EventSink for InfluxDB.
func CreateMysqlSink(uri *url.URL) (core.EventSink, error) {
	_ = "STUB: not implemented"
	return *new(core.EventSink), nil
}
