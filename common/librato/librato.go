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

package librato

import (
	"net/http"
	"net/url"
)

type Measurement struct {
	Name  string            `json:"name,omitempty"`
	Value float64           `json:"value,omitempty"`
	Tags  map[string]string `json:"tags,omitempty"`
	Time  int64             `json:"time,omitempty"`
}

type request struct {
	Tags         map[string]string `json:"tags,omitempty"`
	Measurements []Measurement     `json:"measurements,omitempty"`
}

type Client interface {
	Write([]Measurement) error
}

type LibratoClient struct {
	httpClient *http.Client
	config     LibratoConfig
}

func (c *LibratoClient) Write(measurements []Measurement) error {
	_ = "STUB: not implemented"
	return nil
}

type LibratoConfig struct {
	Username string
	Token    string
	API      string
	Prefix   string
	Tags     map[string]string
}

func NewClient(c LibratoConfig) *LibratoClient { _ = "STUB: not implemented"; return nil }

func BuildConfig(uri *url.URL) (*LibratoConfig, error) { _ = "STUB: not implemented"; return nil, nil }

// TODO: use more secure way to pass the password.
