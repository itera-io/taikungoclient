/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the AlertLabels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertLabels{}

// AlertLabels struct for AlertLabels
type AlertLabels struct {
	Alertname *string `json:"alertname,omitempty"`
	Condition *string `json:"condition,omitempty"`
	Container *string `json:"container,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Instance *string `json:"instance,omitempty"`
	Job *string `json:"job,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Node *string `json:"node,omitempty"`
	Pod *string `json:"pod,omitempty"`
	Service *string `json:"service,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status *string `json:"status,omitempty"`
	Daemonset *string `json:"daemonset,omitempty"`
}

// NewAlertLabels instantiates a new AlertLabels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertLabels() *AlertLabels {
	this := AlertLabels{}
	return &this
}

// NewAlertLabelsWithDefaults instantiates a new AlertLabels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertLabelsWithDefaults() *AlertLabels {
	this := AlertLabels{}
	return &this
}

// GetAlertname returns the Alertname field value if set, zero value otherwise.
func (o *AlertLabels) GetAlertname() string {
	if o == nil || IsNil(o.Alertname) {
		var ret string
		return ret
	}
	return *o.Alertname
}

// GetAlertnameOk returns a tuple with the Alertname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertLabels) GetAlertnameOk() (*string, bool) {
	if o == nil || IsNil(o.Alertname) {
		return nil, false
	}
	return o.Alertname, true
}

// HasAlertname returns a boolean if a field has been set.
func (o *AlertLabels) HasAlertname() bool {
	if o != nil && !IsNil(o.Alertname) {
		return true
	}

	return false
}

// SetAlertname gets a reference to the given string and assigns it to the Alertname field.
func (o *AlertLabels) SetAlertname(v string) {
	o.Alertname = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *AlertLabels) GetCondition() string {
	if o == nil || IsNil(o.Condition) {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertLabels) GetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *AlertLabels) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *AlertLabels) SetCondition(v string) {
	o.Condition = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *AlertLabels) GetContainer() string {
	if o == nil || IsNil(o.Container) {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertLabels) GetContainerOk() (*string, bool) {
	if o == nil || IsNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *AlertLabels) HasContainer() bool {
	if o != nil && !IsNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *AlertLabels) SetContainer(v string) {
	o.Container = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *AlertLabels) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertLabels) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *AlertLabels) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *AlertLabels) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *AlertLabels) GetInstance() string {
	if o == nil || IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertLabels) GetInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *AlertLabels) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *AlertLabels) SetInstance(v string) {
	o.Instance = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *AlertLabels) GetJob() string {
	if o == nil || IsNil(o.Job) {
		var ret string
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertLabels) GetJobOk() (*string, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *AlertLabels) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given string and assigns it to the Job field.
func (o *AlertLabels) SetJob(v string) {
	o.Job = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *AlertLabels) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertLabels) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *AlertLabels) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *AlertLabels) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *AlertLabels) GetNode() string {
	if o == nil || IsNil(o.Node) {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertLabels) GetNodeOk() (*string, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *AlertLabels) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *AlertLabels) SetNode(v string) {
	o.Node = &v
}

// GetPod returns the Pod field value if set, zero value otherwise.
func (o *AlertLabels) GetPod() string {
	if o == nil || IsNil(o.Pod) {
		var ret string
		return ret
	}
	return *o.Pod
}

// GetPodOk returns a tuple with the Pod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertLabels) GetPodOk() (*string, bool) {
	if o == nil || IsNil(o.Pod) {
		return nil, false
	}
	return o.Pod, true
}

// HasPod returns a boolean if a field has been set.
func (o *AlertLabels) HasPod() bool {
	if o != nil && !IsNil(o.Pod) {
		return true
	}

	return false
}

// SetPod gets a reference to the given string and assigns it to the Pod field.
func (o *AlertLabels) SetPod(v string) {
	o.Pod = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *AlertLabels) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertLabels) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *AlertLabels) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *AlertLabels) SetService(v string) {
	o.Service = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *AlertLabels) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertLabels) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *AlertLabels) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *AlertLabels) SetSeverity(v string) {
	o.Severity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlertLabels) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertLabels) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlertLabels) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlertLabels) SetStatus(v string) {
	o.Status = &v
}

// GetDaemonset returns the Daemonset field value if set, zero value otherwise.
func (o *AlertLabels) GetDaemonset() string {
	if o == nil || IsNil(o.Daemonset) {
		var ret string
		return ret
	}
	return *o.Daemonset
}

// GetDaemonsetOk returns a tuple with the Daemonset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertLabels) GetDaemonsetOk() (*string, bool) {
	if o == nil || IsNil(o.Daemonset) {
		return nil, false
	}
	return o.Daemonset, true
}

// HasDaemonset returns a boolean if a field has been set.
func (o *AlertLabels) HasDaemonset() bool {
	if o != nil && !IsNil(o.Daemonset) {
		return true
	}

	return false
}

// SetDaemonset gets a reference to the given string and assigns it to the Daemonset field.
func (o *AlertLabels) SetDaemonset(v string) {
	o.Daemonset = &v
}

func (o AlertLabels) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertLabels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alertname) {
		toSerialize["alertname"] = o.Alertname
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.Container) {
		toSerialize["container"] = o.Container
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.Pod) {
		toSerialize["pod"] = o.Pod
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Daemonset) {
		toSerialize["daemonset"] = o.Daemonset
	}
	return toSerialize, nil
}

type NullableAlertLabels struct {
	value *AlertLabels
	isSet bool
}

func (v NullableAlertLabels) Get() *AlertLabels {
	return v.value
}

func (v *NullableAlertLabels) Set(val *AlertLabels) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertLabels) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertLabels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertLabels(val *AlertLabels) *NullableAlertLabels {
	return &NullableAlertLabels{value: val, isSet: true}
}

func (v NullableAlertLabels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertLabels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


