# KubernetesResourcesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pods** | **int64** |  | 
**Crds** | **int64** |  | 
**Helmreleases** | **int64** |  | 
**DaemonSets** | **int64** |  | 
**Deployments** | **int64** |  | 
**StatefulSets** | **int64** |  | 
**Jobs** | **int64** |  | 
**CronJobs** | **int64** |  | 
**ConfigMaps** | **int64** |  | 
**Namespaces** | **int64** |  | 
**Nodes** | **int64** |  | 
**Pvcs** | **int64** |  | 
**Secrets** | **int64** |  | 
**Services** | **int64** |  | 
**Ingresses** | **int64** |  | 
**NetworkPolicies** | **int64** |  | 
**Pdbs** | **int64** |  | 
**StorageClasses** | **int64** |  | 

## Methods

### NewKubernetesResourcesDto

`func NewKubernetesResourcesDto(pods int64, crds int64, helmreleases int64, daemonSets int64, deployments int64, statefulSets int64, jobs int64, cronJobs int64, configMaps int64, namespaces int64, nodes int64, pvcs int64, secrets int64, services int64, ingresses int64, networkPolicies int64, pdbs int64, storageClasses int64, ) *KubernetesResourcesDto`

NewKubernetesResourcesDto instantiates a new KubernetesResourcesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesResourcesDtoWithDefaults

`func NewKubernetesResourcesDtoWithDefaults() *KubernetesResourcesDto`

NewKubernetesResourcesDtoWithDefaults instantiates a new KubernetesResourcesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPods

`func (o *KubernetesResourcesDto) GetPods() int64`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *KubernetesResourcesDto) GetPodsOk() (*int64, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *KubernetesResourcesDto) SetPods(v int64)`

SetPods sets Pods field to given value.


### GetCrds

`func (o *KubernetesResourcesDto) GetCrds() int64`

GetCrds returns the Crds field if non-nil, zero value otherwise.

### GetCrdsOk

`func (o *KubernetesResourcesDto) GetCrdsOk() (*int64, bool)`

GetCrdsOk returns a tuple with the Crds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrds

`func (o *KubernetesResourcesDto) SetCrds(v int64)`

SetCrds sets Crds field to given value.


### GetHelmreleases

`func (o *KubernetesResourcesDto) GetHelmreleases() int64`

GetHelmreleases returns the Helmreleases field if non-nil, zero value otherwise.

### GetHelmreleasesOk

`func (o *KubernetesResourcesDto) GetHelmreleasesOk() (*int64, bool)`

GetHelmreleasesOk returns a tuple with the Helmreleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmreleases

`func (o *KubernetesResourcesDto) SetHelmreleases(v int64)`

SetHelmreleases sets Helmreleases field to given value.


### GetDaemonSets

`func (o *KubernetesResourcesDto) GetDaemonSets() int64`

GetDaemonSets returns the DaemonSets field if non-nil, zero value otherwise.

### GetDaemonSetsOk

`func (o *KubernetesResourcesDto) GetDaemonSetsOk() (*int64, bool)`

GetDaemonSetsOk returns a tuple with the DaemonSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonSets

`func (o *KubernetesResourcesDto) SetDaemonSets(v int64)`

SetDaemonSets sets DaemonSets field to given value.


### GetDeployments

`func (o *KubernetesResourcesDto) GetDeployments() int64`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *KubernetesResourcesDto) GetDeploymentsOk() (*int64, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *KubernetesResourcesDto) SetDeployments(v int64)`

SetDeployments sets Deployments field to given value.


### GetStatefulSets

`func (o *KubernetesResourcesDto) GetStatefulSets() int64`

GetStatefulSets returns the StatefulSets field if non-nil, zero value otherwise.

### GetStatefulSetsOk

`func (o *KubernetesResourcesDto) GetStatefulSetsOk() (*int64, bool)`

GetStatefulSetsOk returns a tuple with the StatefulSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulSets

`func (o *KubernetesResourcesDto) SetStatefulSets(v int64)`

SetStatefulSets sets StatefulSets field to given value.


### GetJobs

`func (o *KubernetesResourcesDto) GetJobs() int64`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *KubernetesResourcesDto) GetJobsOk() (*int64, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *KubernetesResourcesDto) SetJobs(v int64)`

SetJobs sets Jobs field to given value.


### GetCronJobs

`func (o *KubernetesResourcesDto) GetCronJobs() int64`

GetCronJobs returns the CronJobs field if non-nil, zero value otherwise.

### GetCronJobsOk

`func (o *KubernetesResourcesDto) GetCronJobsOk() (*int64, bool)`

GetCronJobsOk returns a tuple with the CronJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronJobs

`func (o *KubernetesResourcesDto) SetCronJobs(v int64)`

SetCronJobs sets CronJobs field to given value.


### GetConfigMaps

`func (o *KubernetesResourcesDto) GetConfigMaps() int64`

GetConfigMaps returns the ConfigMaps field if non-nil, zero value otherwise.

### GetConfigMapsOk

`func (o *KubernetesResourcesDto) GetConfigMapsOk() (*int64, bool)`

GetConfigMapsOk returns a tuple with the ConfigMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMaps

`func (o *KubernetesResourcesDto) SetConfigMaps(v int64)`

SetConfigMaps sets ConfigMaps field to given value.


### GetNamespaces

`func (o *KubernetesResourcesDto) GetNamespaces() int64`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *KubernetesResourcesDto) GetNamespacesOk() (*int64, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *KubernetesResourcesDto) SetNamespaces(v int64)`

SetNamespaces sets Namespaces field to given value.


### GetNodes

`func (o *KubernetesResourcesDto) GetNodes() int64`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *KubernetesResourcesDto) GetNodesOk() (*int64, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *KubernetesResourcesDto) SetNodes(v int64)`

SetNodes sets Nodes field to given value.


### GetPvcs

`func (o *KubernetesResourcesDto) GetPvcs() int64`

GetPvcs returns the Pvcs field if non-nil, zero value otherwise.

### GetPvcsOk

`func (o *KubernetesResourcesDto) GetPvcsOk() (*int64, bool)`

GetPvcsOk returns a tuple with the Pvcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcs

`func (o *KubernetesResourcesDto) SetPvcs(v int64)`

SetPvcs sets Pvcs field to given value.


### GetSecrets

`func (o *KubernetesResourcesDto) GetSecrets() int64`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *KubernetesResourcesDto) GetSecretsOk() (*int64, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *KubernetesResourcesDto) SetSecrets(v int64)`

SetSecrets sets Secrets field to given value.


### GetServices

`func (o *KubernetesResourcesDto) GetServices() int64`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *KubernetesResourcesDto) GetServicesOk() (*int64, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *KubernetesResourcesDto) SetServices(v int64)`

SetServices sets Services field to given value.


### GetIngresses

`func (o *KubernetesResourcesDto) GetIngresses() int64`

GetIngresses returns the Ingresses field if non-nil, zero value otherwise.

### GetIngressesOk

`func (o *KubernetesResourcesDto) GetIngressesOk() (*int64, bool)`

GetIngressesOk returns a tuple with the Ingresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngresses

`func (o *KubernetesResourcesDto) SetIngresses(v int64)`

SetIngresses sets Ingresses field to given value.


### GetNetworkPolicies

`func (o *KubernetesResourcesDto) GetNetworkPolicies() int64`

GetNetworkPolicies returns the NetworkPolicies field if non-nil, zero value otherwise.

### GetNetworkPoliciesOk

`func (o *KubernetesResourcesDto) GetNetworkPoliciesOk() (*int64, bool)`

GetNetworkPoliciesOk returns a tuple with the NetworkPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicies

`func (o *KubernetesResourcesDto) SetNetworkPolicies(v int64)`

SetNetworkPolicies sets NetworkPolicies field to given value.


### GetPdbs

`func (o *KubernetesResourcesDto) GetPdbs() int64`

GetPdbs returns the Pdbs field if non-nil, zero value otherwise.

### GetPdbsOk

`func (o *KubernetesResourcesDto) GetPdbsOk() (*int64, bool)`

GetPdbsOk returns a tuple with the Pdbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdbs

`func (o *KubernetesResourcesDto) SetPdbs(v int64)`

SetPdbs sets Pdbs field to given value.


### GetStorageClasses

`func (o *KubernetesResourcesDto) GetStorageClasses() int64`

GetStorageClasses returns the StorageClasses field if non-nil, zero value otherwise.

### GetStorageClassesOk

`func (o *KubernetesResourcesDto) GetStorageClassesOk() (*int64, bool)`

GetStorageClassesOk returns a tuple with the StorageClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClasses

`func (o *KubernetesResourcesDto) SetStorageClasses(v int64)`

SetStorageClasses sets StorageClasses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


