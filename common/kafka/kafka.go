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
	"crypto/tls"
	"net/url"
	"time"

	kafka "github.com/Shopify/sarama"
)

const (
	brokerClientID         = "kafka-sink"
	brokerDialTimeout      = 10 * time.Second
	brokerDialRetryLimit   = 1
	brokerDialRetryWait    = 0
	brokerLeaderRetryLimit = 1
	brokerLeaderRetryWait  = 0
	metricsTopic           = "heapster-metrics"
	eventsTopic            = "heapster-events"
)

const (
	TimeSeriesTopic = "timeseriestopic"
	EventsTopic     = "eventstopic"
)

type KafkaClient interface {
	Name() string
	Stop()
	ProduceKafkaMessage(msgData interface{}) error
}

type kafkaSink struct {
	producer  kafka.SyncProducer
	dataTopic string
}

func (sink *kafkaSink) ProduceKafkaMessage(msgData interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (sink *kafkaSink) Name() string { _ = "STUB: not implemented"; return "" }

func (sink *kafkaSink) Stop() { _ = "STUB: not implemented"; return }

func getTopic(opts map[string][]string, topicType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getCompression(opts url.Values) (kafka.CompressionCodec, error) {
	_ = "STUB: not implemented"
	return *new(kafka.CompressionCodec), nil
}

func getTlsConfiguration(opts url.Values) (*tls.Config, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func getSASLConfiguration(opts url.Values) (string, string, bool, error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

func getOptionsWithoutSecrets(values url.Values) string { _ = "STUB: not implemented"; return "" }

func NewKafkaClient(uri *url.URL, topicType string) (KafkaClient, error) {
	_ = "STUB: not implemented"
	return *new(KafkaClient), nil
}

//structure the config of broker

// set up producer of kafka server.
