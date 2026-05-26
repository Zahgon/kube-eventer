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

package flags

import (
	"net/url"
)

type Uri struct {
	Key string
	Val url.URL
}

func (u *Uri) String() string { _ = "STUB: not implemented"; return "" }

func (u *Uri) Set(value string) error { _ = "STUB: not implemented"; return nil }

type Uris []Uri

func (us *Uris) String() string { _ = "STUB: not implemented"; return "" }

func (us *Uris) Set(value string) error { _ = "STUB: not implemented"; return nil }

func (us *Uris) Type() string { _ = "STUB: not implemented"; return "" }
