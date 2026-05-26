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

package mongo

import (
	"net/url"
	"sync"
	"time"

	"github.com/AliyunContainerService/kube-eventer/core"
	"go.mongodb.org/mongo-driver/mongo"
	kube_api "k8s.io/api/core/v1"
)

type mongoSink struct {
	client  *mongo.Client
	closeDB func()
	sync.RWMutex
}

type mongoSinkPoint struct {
	Count                    int32     `bson:"count,omitempty"`
	Namespace                string    `bson:"namespace,omitempty"`
	Kind                     string    `bson:"kind,omitempty"`
	Name                     string    `bson:"name,omitempty"`
	Type                     string    `bson:"type,omitempty"`
	Reason                   string    `bson:"reason,omitempty"`
	Message                  string    `bson:"message,omitempty"`
	EventID                  string    `bson:"event_id,omitempty"`
	FirstOccurrenceTimestamp time.Time `bson:"first_occurrence_time,omitempty"`
	LastOccurrenceTimestamp  time.Time `bson:"last_occurrence_time,omitempty"`
}

func (m *mongoSink) Name() string { _ = "STUB: not implemented"; return "" }

func (m *mongoSink) saveData(sinkData *mongoSinkPoint) error { _ = "STUB: not implemented"; return nil }

func eventToPoint(event *kube_api.Event) (*mongoSinkPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Part of k8s resources FirstOccurrenceTimestamp/LastOccurrenceTimestamp is nil

func (m *mongoSink) ExportEvents(eventBatch *core.EventBatch) { _ = "STUB: not implemented"; return }

func (m *mongoSink) Stop() { _ = "STUB: not implemented"; return }

func CreateMongoSink(uri *url.URL) (core.EventSink, error) {
	_ = "STUB: not implemented"
	return *new(core.EventSink), nil
}
