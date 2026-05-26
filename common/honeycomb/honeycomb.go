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
	"io"
	"net/http"
	"net/url"
	"time"
)

type config struct {
	APIHost  string
	Dataset  string
	WriteKey string
}

func BuildConfig(uri *url.URL) (*config, error) { _ = "STUB: not implemented"; return nil, nil }

type Client interface {
	SendBatch(batch Batch) error
}

type HoneycombClient struct {
	config     config
	httpClient http.Client
}

func NewClient(uri *url.URL) (*HoneycombClient, error) { _ = "STUB: not implemented"; return nil, nil }

type BatchPoint struct {
	Data      interface{}
	Timestamp time.Time
}

type Batch []*BatchPoint

func (c *HoneycombClient) SendBatch(batch Batch) error { _ = "STUB: not implemented"; return nil }

// Nothing to send

func (c *HoneycombClient) makeRequest(body io.Reader) error { _ = "STUB: not implemented"; return nil }
