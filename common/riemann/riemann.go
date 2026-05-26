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

package riemann

import (
	"net/url"
	"sync"
)

// Used to store the Riemann configuration specified in the Heapster cli
type RiemannConfig struct {
	Host      string
	Ttl       float32
	State     string
	Tags      []string
	BatchSize int
}

// contains the riemann client, the riemann configuration, and a RWMutex
type RiemannSink struct {
	Client riemanngo.Client
	Config RiemannConfig
	sync.RWMutex
}

// creates a Riemann sink. Returns a riemannSink
func CreateRiemannSink(uri *url.URL) (*RiemannSink, error) {
	_ = "STUB: not implemented"
	// Default configuration
	return nil, nil
}

// check host

// check ttl

// check batch size

// check state

// check tags

// Warn but return the sink => the client in the sink can be nil

// Receives a sink, connect the riemann client.
func GetRiemannClient(config RiemannConfig) (riemanngo.Client, error) {
	_ = "STUB: not implemented"
	return *new(riemanngo.Client), nil
}

// 5 seconds timeout

// Send Events to Riemann using the client from the sink.
func SendData(client riemanngo.Client, events []riemanngo.Event) error {
	_ = "STUB: not implemented"
	// do nothing if we are not connected
	return nil
}
