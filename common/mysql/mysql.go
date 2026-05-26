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
	"database/sql"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DEFAULT_TABLE = "kube_event"
)

type MysqlService struct {
	db    *sql.DB
	table string
	dsn   string
}

type MysqlKubeEventPoint struct {
	Namespace                string
	Kind                     string
	Name                     string
	Type                     string
	Reason                   string
	Message                  string
	EventID                  string
	FirstOccurrenceTimestamp string
	LastOccurrenceTimestamp  string
}

func (mySvc MysqlService) SaveData(sinkData []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Prepare statement for inserting data

func (mySvc MysqlService) FlushData() error { _ = "STUB: not implemented"; return nil }

func (mySvc MysqlService) CreateDatabase(name string) error { _ = "STUB: not implemented"; return nil }

func (mySvc MysqlService) CloseDB() error { _ = "STUB: not implemented"; return nil }

func NewMysqlClient(uri *url.URL) (*MysqlService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Open doesn't open a connection. Validate DSN data:
