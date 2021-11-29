// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/client/access_profiles"
	"github.com/itera-io/taikungoclient/client/admin"
	"github.com/itera-io/taikungoclient/client/alerting_integrations"
	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/client/auth"
	"github.com/itera-io/taikungoclient/client/aws"
	"github.com/itera-io/taikungoclient/client/azure"
	"github.com/itera-io/taikungoclient/client/backup"
	"github.com/itera-io/taikungoclient/client/billing"
	"github.com/itera-io/taikungoclient/client/checker"
	"github.com/itera-io/taikungoclient/client/cloud_credentials"
	"github.com/itera-io/taikungoclient/client/common"
	"github.com/itera-io/taikungoclient/client/cron_job"
	"github.com/itera-io/taikungoclient/client/doc"
	"github.com/itera-io/taikungoclient/client/flavors"
	"github.com/itera-io/taikungoclient/client/images"
	"github.com/itera-io/taikungoclient/client/invoices"
	"github.com/itera-io/taikungoclient/client/keycloak"
	"github.com/itera-io/taikungoclient/client/kube_config"
	"github.com/itera-io/taikungoclient/client/kube_config_role"
	"github.com/itera-io/taikungoclient/client/kubernetes"
	"github.com/itera-io/taikungoclient/client/kubernetes_profiles"
	"github.com/itera-io/taikungoclient/client/kubespray"
	"github.com/itera-io/taikungoclient/client/notifications"
	"github.com/itera-io/taikungoclient/client/opa_profiles"
	"github.com/itera-io/taikungoclient/client/openstack"
	"github.com/itera-io/taikungoclient/client/ops_credentials"
	"github.com/itera-io/taikungoclient/client/organization_subscriptions"
	"github.com/itera-io/taikungoclient/client/organizations"
	"github.com/itera-io/taikungoclient/client/partner"
	"github.com/itera-io/taikungoclient/client/payment"
	"github.com/itera-io/taikungoclient/client/pre_defined_queries"
	"github.com/itera-io/taikungoclient/client/project_quotas"
	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/itera-io/taikungoclient/client/prometheus"
	"github.com/itera-io/taikungoclient/client/request"
	"github.com/itera-io/taikungoclient/client/s3_credentials"
	"github.com/itera-io/taikungoclient/client/search"
	"github.com/itera-io/taikungoclient/client/security_group"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/itera-io/taikungoclient/client/showback"
	"github.com/itera-io/taikungoclient/client/slack"
	"github.com/itera-io/taikungoclient/client/ssh_users"
	"github.com/itera-io/taikungoclient/client/stand_alone"
	"github.com/itera-io/taikungoclient/client/stand_alone_profile"
	"github.com/itera-io/taikungoclient/client/stand_alone_vm_disks"
	"github.com/itera-io/taikungoclient/client/subscription"
	"github.com/itera-io/taikungoclient/client/ticket"
	"github.com/itera-io/taikungoclient/client/user_projects"
	"github.com/itera-io/taikungoclient/client/users"
)

// Default taikungoclient HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api-staging.taikun.dev"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new taikungoclient HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Taikungoclient {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new taikungoclient HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Taikungoclient {
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

// New creates a new taikungoclient client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Taikungoclient {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Taikungoclient)
	cli.Transport = transport
	cli.AccessProfiles = access_profiles.New(transport, formats)
	cli.Admin = admin.New(transport, formats)
	cli.AlertingIntegrations = alerting_integrations.New(transport, formats)
	cli.AlertingProfiles = alerting_profiles.New(transport, formats)
	cli.Auth = auth.New(transport, formats)
	cli.Aws = aws.New(transport, formats)
	cli.Azure = azure.New(transport, formats)
	cli.Backup = backup.New(transport, formats)
	cli.Billing = billing.New(transport, formats)
	cli.Checker = checker.New(transport, formats)
	cli.CloudCredentials = cloud_credentials.New(transport, formats)
	cli.Common = common.New(transport, formats)
	cli.CronJob = cron_job.New(transport, formats)
	cli.Doc = doc.New(transport, formats)
	cli.Flavors = flavors.New(transport, formats)
	cli.Images = images.New(transport, formats)
	cli.Invoices = invoices.New(transport, formats)
	cli.Keycloak = keycloak.New(transport, formats)
	cli.KubeConfig = kube_config.New(transport, formats)
	cli.KubeConfigRole = kube_config_role.New(transport, formats)
	cli.Kubernetes = kubernetes.New(transport, formats)
	cli.KubernetesProfiles = kubernetes_profiles.New(transport, formats)
	cli.Kubespray = kubespray.New(transport, formats)
	cli.Notifications = notifications.New(transport, formats)
	cli.OpaProfiles = opa_profiles.New(transport, formats)
	cli.Openstack = openstack.New(transport, formats)
	cli.OpsCredentials = ops_credentials.New(transport, formats)
	cli.OrganizationSubscriptions = organization_subscriptions.New(transport, formats)
	cli.Organizations = organizations.New(transport, formats)
	cli.Partner = partner.New(transport, formats)
	cli.Payment = payment.New(transport, formats)
	cli.PreDefinedQueries = pre_defined_queries.New(transport, formats)
	cli.ProjectQuotas = project_quotas.New(transport, formats)
	cli.Projects = projects.New(transport, formats)
	cli.Prometheus = prometheus.New(transport, formats)
	cli.Request = request.New(transport, formats)
	cli.S3Credentials = s3_credentials.New(transport, formats)
	cli.Search = search.New(transport, formats)
	cli.SecurityGroup = security_group.New(transport, formats)
	cli.Servers = servers.New(transport, formats)
	cli.Showback = showback.New(transport, formats)
	cli.Slack = slack.New(transport, formats)
	cli.SSHUsers = ssh_users.New(transport, formats)
	cli.StandAlone = stand_alone.New(transport, formats)
	cli.StandAloneProfile = stand_alone_profile.New(transport, formats)
	cli.StandAloneVMDisks = stand_alone_vm_disks.New(transport, formats)
	cli.Subscription = subscription.New(transport, formats)
	cli.Ticket = ticket.New(transport, formats)
	cli.UserProjects = user_projects.New(transport, formats)
	cli.Users = users.New(transport, formats)
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

// Taikungoclient is a client for taikungoclient
type Taikungoclient struct {
	AccessProfiles access_profiles.ClientService

	Admin admin.ClientService

	AlertingIntegrations alerting_integrations.ClientService

	AlertingProfiles alerting_profiles.ClientService

	Auth auth.ClientService

	Aws aws.ClientService

	Azure azure.ClientService

	Backup backup.ClientService

	Billing billing.ClientService

	Checker checker.ClientService

	CloudCredentials cloud_credentials.ClientService

	Common common.ClientService

	CronJob cron_job.ClientService

	Doc doc.ClientService

	Flavors flavors.ClientService

	Images images.ClientService

	Invoices invoices.ClientService

	Keycloak keycloak.ClientService

	KubeConfig kube_config.ClientService

	KubeConfigRole kube_config_role.ClientService

	Kubernetes kubernetes.ClientService

	KubernetesProfiles kubernetes_profiles.ClientService

	Kubespray kubespray.ClientService

	Notifications notifications.ClientService

	OpaProfiles opa_profiles.ClientService

	Openstack openstack.ClientService

	OpsCredentials ops_credentials.ClientService

	OrganizationSubscriptions organization_subscriptions.ClientService

	Organizations organizations.ClientService

	Partner partner.ClientService

	Payment payment.ClientService

	PreDefinedQueries pre_defined_queries.ClientService

	ProjectQuotas project_quotas.ClientService

	Projects projects.ClientService

	Prometheus prometheus.ClientService

	Request request.ClientService

	S3Credentials s3_credentials.ClientService

	Search search.ClientService

	SecurityGroup security_group.ClientService

	Servers servers.ClientService

	Showback showback.ClientService

	Slack slack.ClientService

	SSHUsers ssh_users.ClientService

	StandAlone stand_alone.ClientService

	StandAloneProfile stand_alone_profile.ClientService

	StandAloneVMDisks stand_alone_vm_disks.ClientService

	Subscription subscription.ClientService

	Ticket ticket.ClientService

	UserProjects user_projects.ClientService

	Users users.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Taikungoclient) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AccessProfiles.SetTransport(transport)
	c.Admin.SetTransport(transport)
	c.AlertingIntegrations.SetTransport(transport)
	c.AlertingProfiles.SetTransport(transport)
	c.Auth.SetTransport(transport)
	c.Aws.SetTransport(transport)
	c.Azure.SetTransport(transport)
	c.Backup.SetTransport(transport)
	c.Billing.SetTransport(transport)
	c.Checker.SetTransport(transport)
	c.CloudCredentials.SetTransport(transport)
	c.Common.SetTransport(transport)
	c.CronJob.SetTransport(transport)
	c.Doc.SetTransport(transport)
	c.Flavors.SetTransport(transport)
	c.Images.SetTransport(transport)
	c.Invoices.SetTransport(transport)
	c.Keycloak.SetTransport(transport)
	c.KubeConfig.SetTransport(transport)
	c.KubeConfigRole.SetTransport(transport)
	c.Kubernetes.SetTransport(transport)
	c.KubernetesProfiles.SetTransport(transport)
	c.Kubespray.SetTransport(transport)
	c.Notifications.SetTransport(transport)
	c.OpaProfiles.SetTransport(transport)
	c.Openstack.SetTransport(transport)
	c.OpsCredentials.SetTransport(transport)
	c.OrganizationSubscriptions.SetTransport(transport)
	c.Organizations.SetTransport(transport)
	c.Partner.SetTransport(transport)
	c.Payment.SetTransport(transport)
	c.PreDefinedQueries.SetTransport(transport)
	c.ProjectQuotas.SetTransport(transport)
	c.Projects.SetTransport(transport)
	c.Prometheus.SetTransport(transport)
	c.Request.SetTransport(transport)
	c.S3Credentials.SetTransport(transport)
	c.Search.SetTransport(transport)
	c.SecurityGroup.SetTransport(transport)
	c.Servers.SetTransport(transport)
	c.Showback.SetTransport(transport)
	c.Slack.SetTransport(transport)
	c.SSHUsers.SetTransport(transport)
	c.StandAlone.SetTransport(transport)
	c.StandAloneProfile.SetTransport(transport)
	c.StandAloneVMDisks.SetTransport(transport)
	c.Subscription.SetTransport(transport)
	c.Ticket.SetTransport(transport)
	c.UserProjects.SetTransport(transport)
	c.Users.SetTransport(transport)
}
