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
	"bytes"
	"fmt"
)

// checks if the KubernetesResourcesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesResourcesDto{}

// KubernetesResourcesDto struct for KubernetesResourcesDto
type KubernetesResourcesDto struct {
	Pods int64 `json:"pods"`
	Crds int64 `json:"crds"`
	Helmreleases int64 `json:"helmreleases"`
	DaemonSets int64 `json:"daemonSets"`
	Deployments int64 `json:"deployments"`
	StatefulSets int64 `json:"statefulSets"`
	Jobs int64 `json:"jobs"`
	CronJobs int64 `json:"cronJobs"`
	ConfigMaps int64 `json:"configMaps"`
	Namespaces int64 `json:"namespaces"`
	Nodes int64 `json:"nodes"`
	Pvcs int64 `json:"pvcs"`
	Secrets int64 `json:"secrets"`
	Services int64 `json:"services"`
	Ingresses int64 `json:"ingresses"`
	NetworkPolicies int64 `json:"networkPolicies"`
	Pdbs int64 `json:"pdbs"`
	StorageClasses int64 `json:"storageClasses"`
}

type _KubernetesResourcesDto KubernetesResourcesDto

// NewKubernetesResourcesDto instantiates a new KubernetesResourcesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesResourcesDto(pods int64, crds int64, helmreleases int64, daemonSets int64, deployments int64, statefulSets int64, jobs int64, cronJobs int64, configMaps int64, namespaces int64, nodes int64, pvcs int64, secrets int64, services int64, ingresses int64, networkPolicies int64, pdbs int64, storageClasses int64) *KubernetesResourcesDto {
	this := KubernetesResourcesDto{}
	this.Pods = pods
	this.Crds = crds
	this.Helmreleases = helmreleases
	this.DaemonSets = daemonSets
	this.Deployments = deployments
	this.StatefulSets = statefulSets
	this.Jobs = jobs
	this.CronJobs = cronJobs
	this.ConfigMaps = configMaps
	this.Namespaces = namespaces
	this.Nodes = nodes
	this.Pvcs = pvcs
	this.Secrets = secrets
	this.Services = services
	this.Ingresses = ingresses
	this.NetworkPolicies = networkPolicies
	this.Pdbs = pdbs
	this.StorageClasses = storageClasses
	return &this
}

// NewKubernetesResourcesDtoWithDefaults instantiates a new KubernetesResourcesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesResourcesDtoWithDefaults() *KubernetesResourcesDto {
	this := KubernetesResourcesDto{}
	return &this
}

// GetPods returns the Pods field value
func (o *KubernetesResourcesDto) GetPods() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Pods
}

// GetPodsOk returns a tuple with the Pods field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetPodsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pods, true
}

// SetPods sets field value
func (o *KubernetesResourcesDto) SetPods(v int64) {
	o.Pods = v
}

// GetCrds returns the Crds field value
func (o *KubernetesResourcesDto) GetCrds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Crds
}

// GetCrdsOk returns a tuple with the Crds field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetCrdsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Crds, true
}

// SetCrds sets field value
func (o *KubernetesResourcesDto) SetCrds(v int64) {
	o.Crds = v
}

// GetHelmreleases returns the Helmreleases field value
func (o *KubernetesResourcesDto) GetHelmreleases() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Helmreleases
}

// GetHelmreleasesOk returns a tuple with the Helmreleases field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetHelmreleasesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Helmreleases, true
}

// SetHelmreleases sets field value
func (o *KubernetesResourcesDto) SetHelmreleases(v int64) {
	o.Helmreleases = v
}

// GetDaemonSets returns the DaemonSets field value
func (o *KubernetesResourcesDto) GetDaemonSets() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DaemonSets
}

// GetDaemonSetsOk returns a tuple with the DaemonSets field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetDaemonSetsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaemonSets, true
}

// SetDaemonSets sets field value
func (o *KubernetesResourcesDto) SetDaemonSets(v int64) {
	o.DaemonSets = v
}

// GetDeployments returns the Deployments field value
func (o *KubernetesResourcesDto) GetDeployments() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetDeploymentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deployments, true
}

// SetDeployments sets field value
func (o *KubernetesResourcesDto) SetDeployments(v int64) {
	o.Deployments = v
}

// GetStatefulSets returns the StatefulSets field value
func (o *KubernetesResourcesDto) GetStatefulSets() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StatefulSets
}

// GetStatefulSetsOk returns a tuple with the StatefulSets field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetStatefulSetsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatefulSets, true
}

// SetStatefulSets sets field value
func (o *KubernetesResourcesDto) SetStatefulSets(v int64) {
	o.StatefulSets = v
}

// GetJobs returns the Jobs field value
func (o *KubernetesResourcesDto) GetJobs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetJobsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jobs, true
}

// SetJobs sets field value
func (o *KubernetesResourcesDto) SetJobs(v int64) {
	o.Jobs = v
}

// GetCronJobs returns the CronJobs field value
func (o *KubernetesResourcesDto) GetCronJobs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CronJobs
}

// GetCronJobsOk returns a tuple with the CronJobs field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetCronJobsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CronJobs, true
}

// SetCronJobs sets field value
func (o *KubernetesResourcesDto) SetCronJobs(v int64) {
	o.CronJobs = v
}

// GetConfigMaps returns the ConfigMaps field value
func (o *KubernetesResourcesDto) GetConfigMaps() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ConfigMaps
}

// GetConfigMapsOk returns a tuple with the ConfigMaps field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetConfigMapsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigMaps, true
}

// SetConfigMaps sets field value
func (o *KubernetesResourcesDto) SetConfigMaps(v int64) {
	o.ConfigMaps = v
}

// GetNamespaces returns the Namespaces field value
func (o *KubernetesResourcesDto) GetNamespaces() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetNamespacesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespaces, true
}

// SetNamespaces sets field value
func (o *KubernetesResourcesDto) SetNamespaces(v int64) {
	o.Namespaces = v
}

// GetNodes returns the Nodes field value
func (o *KubernetesResourcesDto) GetNodes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetNodesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value
func (o *KubernetesResourcesDto) SetNodes(v int64) {
	o.Nodes = v
}

// GetPvcs returns the Pvcs field value
func (o *KubernetesResourcesDto) GetPvcs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Pvcs
}

// GetPvcsOk returns a tuple with the Pvcs field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetPvcsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pvcs, true
}

// SetPvcs sets field value
func (o *KubernetesResourcesDto) SetPvcs(v int64) {
	o.Pvcs = v
}

// GetSecrets returns the Secrets field value
func (o *KubernetesResourcesDto) GetSecrets() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetSecretsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secrets, true
}

// SetSecrets sets field value
func (o *KubernetesResourcesDto) SetSecrets(v int64) {
	o.Secrets = v
}

// GetServices returns the Services field value
func (o *KubernetesResourcesDto) GetServices() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetServicesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Services, true
}

// SetServices sets field value
func (o *KubernetesResourcesDto) SetServices(v int64) {
	o.Services = v
}

// GetIngresses returns the Ingresses field value
func (o *KubernetesResourcesDto) GetIngresses() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Ingresses
}

// GetIngressesOk returns a tuple with the Ingresses field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetIngressesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ingresses, true
}

// SetIngresses sets field value
func (o *KubernetesResourcesDto) SetIngresses(v int64) {
	o.Ingresses = v
}

// GetNetworkPolicies returns the NetworkPolicies field value
func (o *KubernetesResourcesDto) GetNetworkPolicies() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NetworkPolicies
}

// GetNetworkPoliciesOk returns a tuple with the NetworkPolicies field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetNetworkPoliciesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkPolicies, true
}

// SetNetworkPolicies sets field value
func (o *KubernetesResourcesDto) SetNetworkPolicies(v int64) {
	o.NetworkPolicies = v
}

// GetPdbs returns the Pdbs field value
func (o *KubernetesResourcesDto) GetPdbs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Pdbs
}

// GetPdbsOk returns a tuple with the Pdbs field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetPdbsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pdbs, true
}

// SetPdbs sets field value
func (o *KubernetesResourcesDto) SetPdbs(v int64) {
	o.Pdbs = v
}

// GetStorageClasses returns the StorageClasses field value
func (o *KubernetesResourcesDto) GetStorageClasses() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StorageClasses
}

// GetStorageClassesOk returns a tuple with the StorageClasses field value
// and a boolean to check if the value has been set.
func (o *KubernetesResourcesDto) GetStorageClassesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageClasses, true
}

// SetStorageClasses sets field value
func (o *KubernetesResourcesDto) SetStorageClasses(v int64) {
	o.StorageClasses = v
}

func (o KubernetesResourcesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesResourcesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pods"] = o.Pods
	toSerialize["crds"] = o.Crds
	toSerialize["helmreleases"] = o.Helmreleases
	toSerialize["daemonSets"] = o.DaemonSets
	toSerialize["deployments"] = o.Deployments
	toSerialize["statefulSets"] = o.StatefulSets
	toSerialize["jobs"] = o.Jobs
	toSerialize["cronJobs"] = o.CronJobs
	toSerialize["configMaps"] = o.ConfigMaps
	toSerialize["namespaces"] = o.Namespaces
	toSerialize["nodes"] = o.Nodes
	toSerialize["pvcs"] = o.Pvcs
	toSerialize["secrets"] = o.Secrets
	toSerialize["services"] = o.Services
	toSerialize["ingresses"] = o.Ingresses
	toSerialize["networkPolicies"] = o.NetworkPolicies
	toSerialize["pdbs"] = o.Pdbs
	toSerialize["storageClasses"] = o.StorageClasses
	return toSerialize, nil
}

func (o *KubernetesResourcesDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pods",
		"crds",
		"helmreleases",
		"daemonSets",
		"deployments",
		"statefulSets",
		"jobs",
		"cronJobs",
		"configMaps",
		"namespaces",
		"nodes",
		"pvcs",
		"secrets",
		"services",
		"ingresses",
		"networkPolicies",
		"pdbs",
		"storageClasses",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKubernetesResourcesDto := _KubernetesResourcesDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubernetesResourcesDto)

	if err != nil {
		return err
	}

	*o = KubernetesResourcesDto(varKubernetesResourcesDto)

	return err
}

type NullableKubernetesResourcesDto struct {
	value *KubernetesResourcesDto
	isSet bool
}

func (v NullableKubernetesResourcesDto) Get() *KubernetesResourcesDto {
	return v.value
}

func (v *NullableKubernetesResourcesDto) Set(val *KubernetesResourcesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesResourcesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesResourcesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesResourcesDto(val *KubernetesResourcesDto) *NullableKubernetesResourcesDto {
	return &NullableKubernetesResourcesDto{value: val, isSet: true}
}

func (v NullableKubernetesResourcesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesResourcesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


