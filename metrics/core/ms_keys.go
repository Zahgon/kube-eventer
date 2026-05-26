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

package core

// MetricsSet keys inside of DataBatch. The structure of the returned string is
// an implementation detail and no component should rely on it as it may change
// anytime. It it only guaranteed that it is unique for the unique combination of
// passed parameters.

func PodContainerKey(namespace, podName, containerName string) string {
	_ = "STUB: not implemented"
	return ""
}

func PodKey(namespace, podName string) string { _ = "STUB: not implemented"; return "" }

func NamespaceKey(namespace string) string { _ = "STUB: not implemented"; return "" }

func NodeKey(node string) string { _ = "STUB: not implemented"; return "" }

func NodeContainerKey(node, container string) string { _ = "STUB: not implemented"; return "" }

func ClusterKey() string { _ = "STUB: not implemented"; return "" }
