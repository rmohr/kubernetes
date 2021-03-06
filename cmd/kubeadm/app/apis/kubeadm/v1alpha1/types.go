/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MasterConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	API               API        `json:"api"`
	Discovery         Discovery  `json:"discovery"`
	Etcd              Etcd       `json:"etcd"`
	Networking        Networking `json:"networking"`
	KubernetesVersion string     `json:"kubernetesVersion"`
	CloudProvider     string     `json:"cloudProvider"`
	AuthorizationMode string     `json:"authorizationMode"`

	// SelfHosted enables an alpha deployment type where the apiserver, scheduler, and
	// controller manager are managed by Kubernetes itself. This option is likely to
	// become the default in the future.
	SelfHosted bool `json:"selfHosted"`

	APIServerExtraArgs         map[string]string `json:"apiServerExtraArgs"`
	ControllerManagerExtraArgs map[string]string `json:"controllerManagerExtraArgs"`
	SchedulerExtraArgs         map[string]string `json:"schedulerExtraArgs"`

	// APIServerCertSANs sets extra Subject Alternative Names for the API Server signing cert
	APIServerCertSANs []string `json:"apiServerCertSANs"`
	// CertificatesDir specifies where to store or look for all required certificates
	CertificatesDir string `json:"certificatesDir"`
}

type API struct {
	// AdvertiseAddress sets the address for the API server to advertise.
	AdvertiseAddress string `json:"advertiseAddress"`
	// BindPort sets the secure port for the API Server to bind to
	BindPort int32 `json:"bindPort"`
}

type Discovery struct {
	HTTPS *HTTPSDiscovery `json:"https"`
	File  *FileDiscovery  `json:"file"`
	Token *TokenDiscovery `json:"token"`
}

type HTTPSDiscovery struct {
	URL string `json:"url"`
}

type FileDiscovery struct {
	Path string `json:"path"`
}

type TokenDiscovery struct {
	ID        string   `json:"id"`
	Secret    string   `json:"secret"`
	Addresses []string `json:"addresses"`
}

type Networking struct {
	ServiceSubnet string `json:"serviceSubnet"`
	PodSubnet     string `json:"podSubnet"`
	DNSDomain     string `json:"dnsDomain"`
}

type Etcd struct {
	Endpoints []string `json:"endpoints"`
	CAFile    string   `json:"caFile"`
	CertFile  string   `json:"certFile"`
	KeyFile   string   `json:"keyFile"`
}

type NodeConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	CACertPath               string   `json:"caCertPath"`
	DiscoveryFile            string   `json:"discoveryFile"`
	DiscoveryToken           string   `json:"discoveryToken"`
	DiscoveryTokenAPIServers []string `json:"discoveryTokenAPIServers"`
	TLSBootstrapToken        string   `json:"tlsBootstrapToken"`
	Token                    string   `json:"token"`
}

// ClusterInfo TODO add description
type ClusterInfo struct {
	metav1.TypeMeta `json:",inline"`
	// TODO(phase1+) this may become simply `api.Config`
	CertificateAuthorities []string `json:"certificateAuthorities"`
	Endpoints              []string `json:"endpoints"`
}
