// Copyright 2014 Google Inc. All Rights Reserved.
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

package kubernetes

import (
	"net/url"

	"k8s.io/client-go/kubernetes"
	kube_rest "k8s.io/client-go/rest"
	kubeClientCmd "k8s.io/client-go/tools/clientcmd"
)

const (
	APIVersion = "v1"

	defaultKubeletPort        = 10255
	defaultKubeletHttps       = false
	defaultUseServiceAccount  = false
	defaultServiceAccountFile = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	defaultInClusterConfig    = true
)

var KubernetesClientSingleton kubernetes.Interface = nil

func getConfigOverrides(uri *url.URL) (*kubeClientCmd.ConfigOverrides, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetKubeClientConfig(uri *url.URL) (*kube_rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use the current context in kubeconfig

// Load structured kubeconfig data from the given path.

// Flatten the loaded data to a particular restclient.Config based on the current context.

// If a readable service account token exists, then use it

func GetKubernetesClient(uri *url.URL) (client kubernetes.Interface, err error) {
	_ = "STUB: not implemented"
	return *new(kubernetes.Interface), nil
}
