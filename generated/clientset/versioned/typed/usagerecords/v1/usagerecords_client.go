/*
Copyright (c) 2023 SUSE LLC

This program is free software; you can redistribute it and/or
modify it under the terms of version 3 of the GNU General Public License as
published by the Free Software Foundation.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.   See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, contact SUSE LLC.

To contact SUSE about this file by physical or electronic mail,
you may find current contact information at www.suse.com
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"net/http"

	v1 "github.com/SUSE-Enceladus/csp-rancher-usage-operator/api/usagerecords/v1"
	"github.com/SUSE-Enceladus/csp-rancher-usage-operator/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type UsagerecordsV1Interface interface {
	RESTClient() rest.Interface
	ProductUsagesGetter
}

// UsagerecordsV1Client is used to interact with features provided by the usagerecords.suse.com group.
type UsagerecordsV1Client struct {
	restClient rest.Interface
}

func (c *UsagerecordsV1Client) ProductUsages() ProductUsageInterface {
	return newProductUsages(c)
}

// NewForConfig creates a new UsagerecordsV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*UsagerecordsV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new UsagerecordsV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*UsagerecordsV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &UsagerecordsV1Client{client}, nil
}

// NewForConfigOrDie creates a new UsagerecordsV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *UsagerecordsV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new UsagerecordsV1Client for the given RESTClient.
func New(c rest.Interface) *UsagerecordsV1Client {
	return &UsagerecordsV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *UsagerecordsV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
