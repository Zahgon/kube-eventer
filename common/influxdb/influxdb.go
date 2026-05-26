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
	"time"

	influxdb "github.com/influxdata/influxdb/client"
)

type InfluxdbClient interface {
	Write(influxdb.BatchPoints) (*influxdb.Response, error)
	Query(influxdb.Query) (*influxdb.Response, error)
	Ping() (time.Duration, string, error)
}

type InfluxdbConfig struct {
	User                  string
	Password              string
	Secure                bool
	Host                  string
	DbName                string
	WithFields            bool
	InsecureSsl           bool
	RetentionPolicy       string
	ClusterName           string
	DisableCounterMetrics bool
	Concurrency           int
}

func NewClient(c InfluxdbConfig) (InfluxdbClient, error) {
	_ = "STUB: not implemented"
	return *new(InfluxdbClient), nil
}

func BuildConfig(uri *url.URL) (*InfluxdbConfig, error) { _ = "STUB: not implemented"; return nil, nil }

// TODO: use more secure way to pass the password.
