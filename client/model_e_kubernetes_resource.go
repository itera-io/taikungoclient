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
	"fmt"
)

// EKubernetesResource the model 'EKubernetesResource'
type EKubernetesResource string

// List of EKubernetesResource
const (
	EKUBERNETESRESOURCE_NONE EKubernetesResource = "None"
	EKUBERNETESRESOURCE_DAEMON_SET EKubernetesResource = "DaemonSet"
	EKUBERNETESRESOURCE_PVC EKubernetesResource = "Pvc"
	EKUBERNETESRESOURCE_DEPLOYMENT EKubernetesResource = "Deployment"
	EKUBERNETESRESOURCE_CONFIG_MAP EKubernetesResource = "ConfigMap"
	EKUBERNETESRESOURCE_SECRET EKubernetesResource = "Secret"
	EKUBERNETESRESOURCE_STS EKubernetesResource = "Sts"
	EKUBERNETESRESOURCE_SERVICE EKubernetesResource = "Service"
	EKUBERNETESRESOURCE_NODE EKubernetesResource = "Node"
	EKUBERNETESRESOURCE_CRD EKubernetesResource = "Crd"
	EKUBERNETESRESOURCE_STORAGE_CLASS EKubernetesResource = "StorageClass"
	EKUBERNETESRESOURCE_POD EKubernetesResource = "Pod"
	EKUBERNETESRESOURCE_INGRESS EKubernetesResource = "Ingress"
	EKUBERNETESRESOURCE_PDB EKubernetesResource = "Pdb"
	EKUBERNETESRESOURCE_CRON_JOB EKubernetesResource = "CronJob"
	EKUBERNETESRESOURCE_JOB EKubernetesResource = "Job"
	EKUBERNETESRESOURCE_NETWORK_POLICY EKubernetesResource = "NetworkPolicy"
	EKUBERNETESRESOURCE_NAMESPACE EKubernetesResource = "Namespace"
	EKUBERNETESRESOURCE_PERSISTENT_VOLUME EKubernetesResource = "PersistentVolume"
	EKUBERNETESRESOURCE_LIMIT_RANGE EKubernetesResource = "LimitRange"
	EKUBERNETESRESOURCE_RESOURCE_QUOTA EKubernetesResource = "ResourceQuota"
)

// All allowed values of EKubernetesResource enum
var AllowedEKubernetesResourceEnumValues = []EKubernetesResource{
	"None",
	"DaemonSet",
	"Pvc",
	"Deployment",
	"ConfigMap",
	"Secret",
	"Sts",
	"Service",
	"Node",
	"Crd",
	"StorageClass",
	"Pod",
	"Ingress",
	"Pdb",
	"CronJob",
	"Job",
	"NetworkPolicy",
	"Namespace",
	"PersistentVolume",
	"LimitRange",
	"ResourceQuota",
}

func (v *EKubernetesResource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EKubernetesResource(value)
	for _, existing := range AllowedEKubernetesResourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EKubernetesResource", value)
}

// NewEKubernetesResourceFromValue returns a pointer to a valid EKubernetesResource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEKubernetesResourceFromValue(v string) (*EKubernetesResource, error) {
	ev := EKubernetesResource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EKubernetesResource: valid values are %v", v, AllowedEKubernetesResourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EKubernetesResource) IsValid() bool {
	for _, existing := range AllowedEKubernetesResourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EKubernetesResource value
func (v EKubernetesResource) Ptr() *EKubernetesResource {
	return &v
}

type NullableEKubernetesResource struct {
	value *EKubernetesResource
	isSet bool
}

func (v NullableEKubernetesResource) Get() *EKubernetesResource {
	return v.value
}

func (v *NullableEKubernetesResource) Set(val *EKubernetesResource) {
	v.value = val
	v.isSet = true
}

func (v NullableEKubernetesResource) IsSet() bool {
	return v.isSet
}

func (v *NullableEKubernetesResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEKubernetesResource(val *EKubernetesResource) *NullableEKubernetesResource {
	return &NullableEKubernetesResource{value: val, isSet: true}
}

func (v NullableEKubernetesResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEKubernetesResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

