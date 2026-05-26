// Copyright 2015 Google Inc. All Rights Reserved.
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
package elasticsearch

import (
	"net/http"
	"net/url"
	"time"
)

const (
	ESIndex       = "heapster"
	ESClusterName = "default"
)

type UnsupportedVersion struct{}

func (UnsupportedVersion) Error() string { _ = "STUB: not implemented"; return "" }

type elasticWrapper interface {
	IndexExists(indices ...string) (bool, error)
	CreateIndex(name string, mapping string) (bool, error)
	AddAlias(index string, alias string) (bool, error)
	HasAlias(index string, alias string) (bool, error)
	AddBulkReq(index, typeName string, data interface{}) error
	ErrorStats() int64
	FlushBulk() error
}

type ElasticConfig struct {
	Url         []string
	User        string
	Secret      string
	MaxRetries  *int
	HealthCheck *bool
	Timeout     *time.Duration
	HttpClient  *http.Client
	Sniff       *bool
}

type ElasticSearchService struct {
	EsClient     elasticWrapper
	baseIndex    string
	ClusterName  string
	UseNamespace bool
}

func (esSvc *ElasticSearchService) Index(date time.Time, namespace string) string {
	_ = "STUB: not implemented"
	return ""
}

func (esSvc *ElasticSearchService) IndexAlias(typeName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (esSvc *ElasticSearchService) FlushData() error { _ = "STUB: not implemented"; return nil }

func (esSvc *ElasticSearchService) ErrorStats() int64 { _ = "STUB: not implemented"; return 0 }

// SaveDataIntoES save metrics and events to ES by using ES client
func (esSvc *ElasticSearchService) SaveData(date time.Time, typeName string, namespace string, sinkData []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Use the IndexExists service to check if a specified index exists.

// Create a new index.

// CreateElasticSearchConfig creates an ElasticSearch configuration struct
// which contains an ElasticSearch client for later use
func CreateElasticSearchService(uri *url.URL) (*ElasticSearchService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// set the index for es,the default value is "heapster"

// Set the URL endpoints of the ES's nodes. Notice that when sniffing is
// enabled, these URLs are used to initially sniff the cluster on startup.

// If the ES cluster needs authentication, the username and secret
// should be set in sink config.Else, set the Authenticate flag to false
