// Copyright 2018 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package sls

import (
	"net/url"

	"github.com/AliyunContainerService/kube-eventer/core"
	"github.com/AliyunContainerService/kube-eventer/sinks/utils"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	sls_producer "github.com/aliyun/aliyun-log-go-sdk/producer"
	v1 "k8s.io/api/core/v1"
)

const (
	slsSinkName        = "SLSSink"
	eventId            = "eventId"
	podEvent           = "Pod"
	eventLevel         = "level"
	SLSDefaultEndpoint = "log.aliyuncs.com"
	SLSUserAgent       = "ack-kube-eventer"
)

/*
 * Usage:
 * --sink=sls:https://sls.aliyuncs.com?logStore=[your_log_store]&project=[your_project_name]&label=<key,value>
 */
type SLSSink struct {
	Config   *Config
	Project  string
	LogStore string

	// current sls producer
	Producer *sls_producer.Producer

	// current ak_info with expiration time
	AkInfo *utils.AKInfo
}

// Config can be specific
type Config struct {
	project         string
	logStore        string
	topic           string
	regionId        string
	internal        bool
	accessKeyId     string
	accessKeySecret string
	label           map[string]string
}

func (s *SLSSink) Name() string { _ = "STUB: not implemented"; return "" }

func (s *SLSSink) ExportEvents(batch *core.EventBatch) { _ = "STUB: not implemented"; return }

func (s *SLSSink) Stop() {
	_ = "STUB: not implemented"
	// safe close producer: close after all data is sent
	return
}

// get a sls producer.
// if akInfo expiration, recreate a new producer.
func (s *SLSSink) getProducer() *sls_producer.Producer { _ = "STUB: not implemented"; return nil }

// if akInfo expiration, recreate a new producer.

// 1. stop the old producer

// 2. create a new producer

// 3. start the new producer

func eventToContents(event *v1.Event, labels map[string]string) []*sls.LogContent {
	_ = "STUB: not implemented"
	return nil
}

// deep copy

// NewSLSSink returns new SLSSink
func NewSLSSink(uri *url.URL) (*SLSSink, error) { _ = "STUB: not implemented"; return nil, nil }

// parseConfig create config from uri
func parseConfig(uri *url.URL) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func parseLabels(labelsStrs []string) map[string]string { _ = "STUB: not implemented"; return nil }

// newProducer create producer with config and new akInfo.
func newProducer(c *Config) (*sls_producer.Producer, *utils.AKInfo, error) {
	_ = "STUB: not implemented"
	// get region from env
	return nil, nil, nil
}

// region from client

// region from meta data

// get ak info

// construct sls producer config

// refer doc: https://help.aliyun.com/zh/sls/developer-reference/endpoints
func getSLSEndpoint(region string, internal bool) string { _ = "STUB: not implemented"; return "" }

// vpc network

// public network

// callback, use it to implement the sls_producer.Callback interface
// to obtain the result of each send,
// because the producer sends requests to the server asynchronously.
type callback struct {
}

func (c callback) Success(result *sls_producer.Result) { _ = "STUB: not implemented"; return }

func (c callback) Fail(result *sls_producer.Result) { _ = "STUB: not implemented"; return }
