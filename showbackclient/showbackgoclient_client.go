// Code generated by go-swagger; DO NOT EDIT.

package showbackclient

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/showbackclient/projects"
	"github.com/itera-io/taikungoclient/showbackclient/showback_credentials"
	"github.com/itera-io/taikungoclient/showbackclient/showback_rules"
	"github.com/itera-io/taikungoclient/showbackclient/showback_summaries"
)

// Default showbackgoclient HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.taikun.dev"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new showbackgoclient HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Showbackgoclient {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new showbackgoclient HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Showbackgoclient {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)

	// add support for "application/problem+json" media type
	transport.Consumers["application/problem+json"] = runtime.JSONConsumer()
	transport.Producers["application/problem+json"] = runtime.JSONProducer()

	// add support for "application/*+json" media type
	transport.Consumers["application/*+json"] = runtime.JSONConsumer()
	transport.Producers["application/*+json"] = runtime.JSONProducer()

	// add support for "application/json-patch+json" media type
	transport.Consumers["application/json-patch+json"] = runtime.JSONConsumer()
	transport.Producers["application/json-patch+json"] = runtime.JSONProducer()

	// add support for "text/json" media type
	transport.Consumers["text/json"] = runtime.JSONConsumer()
	transport.Producers["text/json"] = runtime.JSONProducer()

	return New(transport, formats)
}

// New creates a new showbackgoclient client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Showbackgoclient {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Showbackgoclient)
	cli.Transport = transport
	cli.Projects = projects.New(transport, formats)
	cli.ShowbackCredentials = showback_credentials.New(transport, formats)
	cli.ShowbackRules = showback_rules.New(transport, formats)
	cli.ShowbackSummaries = showback_summaries.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Showbackgoclient is a client for showbackgoclient
type Showbackgoclient struct {
	Projects projects.ClientService

	ShowbackCredentials showback_credentials.ClientService

	ShowbackRules showback_rules.ClientService

	ShowbackSummaries showback_summaries.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Showbackgoclient) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Projects.SetTransport(transport)
	c.ShowbackCredentials.SetTransport(transport)
	c.ShowbackRules.SetTransport(transport)
	c.ShowbackSummaries.SetTransport(transport)
}
