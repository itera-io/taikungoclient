# Stream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 
**Pod** | **NullableString** |  | 
**PodTemplateHash** | **NullableString** |  | 
**StreamStream** | **NullableString** |  | 
**App** | **NullableString** |  | 
**Container** | **NullableString** |  | 
**Filename** | **NullableString** |  | 
**Tier** | **NullableString** |  | 
**Component** | **NullableString** |  | 

## Methods

### NewStream

`func NewStream(job NullableString, namespace NullableString, pod NullableString, podTemplateHash NullableString, streamStream NullableString, app NullableString, container NullableString, filename NullableString, tier NullableString, component NullableString, ) *Stream`

NewStream instantiates a new Stream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamWithDefaults

`func NewStreamWithDefaults() *Stream`

NewStreamWithDefaults instantiates a new Stream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *Stream) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *Stream) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *Stream) SetJob(v string)`

SetJob sets Job field to given value.


### SetJobNil

`func (o *Stream) SetJobNil(b bool)`

 SetJobNil sets the value for Job to be an explicit nil

### UnsetJob
`func (o *Stream) UnsetJob()`

UnsetJob ensures that no value is present for Job, not even an explicit nil
### GetNamespace

`func (o *Stream) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Stream) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Stream) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *Stream) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *Stream) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetPod

`func (o *Stream) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *Stream) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *Stream) SetPod(v string)`

SetPod sets Pod field to given value.


### SetPodNil

`func (o *Stream) SetPodNil(b bool)`

 SetPodNil sets the value for Pod to be an explicit nil

### UnsetPod
`func (o *Stream) UnsetPod()`

UnsetPod ensures that no value is present for Pod, not even an explicit nil
### GetPodTemplateHash

`func (o *Stream) GetPodTemplateHash() string`

GetPodTemplateHash returns the PodTemplateHash field if non-nil, zero value otherwise.

### GetPodTemplateHashOk

`func (o *Stream) GetPodTemplateHashOk() (*string, bool)`

GetPodTemplateHashOk returns a tuple with the PodTemplateHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodTemplateHash

`func (o *Stream) SetPodTemplateHash(v string)`

SetPodTemplateHash sets PodTemplateHash field to given value.


### SetPodTemplateHashNil

`func (o *Stream) SetPodTemplateHashNil(b bool)`

 SetPodTemplateHashNil sets the value for PodTemplateHash to be an explicit nil

### UnsetPodTemplateHash
`func (o *Stream) UnsetPodTemplateHash()`

UnsetPodTemplateHash ensures that no value is present for PodTemplateHash, not even an explicit nil
### GetStreamStream

`func (o *Stream) GetStreamStream() string`

GetStreamStream returns the StreamStream field if non-nil, zero value otherwise.

### GetStreamStreamOk

`func (o *Stream) GetStreamStreamOk() (*string, bool)`

GetStreamStreamOk returns a tuple with the StreamStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamStream

`func (o *Stream) SetStreamStream(v string)`

SetStreamStream sets StreamStream field to given value.


### SetStreamStreamNil

`func (o *Stream) SetStreamStreamNil(b bool)`

 SetStreamStreamNil sets the value for StreamStream to be an explicit nil

### UnsetStreamStream
`func (o *Stream) UnsetStreamStream()`

UnsetStreamStream ensures that no value is present for StreamStream, not even an explicit nil
### GetApp

`func (o *Stream) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *Stream) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *Stream) SetApp(v string)`

SetApp sets App field to given value.


### SetAppNil

`func (o *Stream) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *Stream) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil
### GetContainer

`func (o *Stream) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *Stream) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *Stream) SetContainer(v string)`

SetContainer sets Container field to given value.


### SetContainerNil

`func (o *Stream) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *Stream) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetFilename

`func (o *Stream) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Stream) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Stream) SetFilename(v string)`

SetFilename sets Filename field to given value.


### SetFilenameNil

`func (o *Stream) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *Stream) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetTier

`func (o *Stream) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Stream) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Stream) SetTier(v string)`

SetTier sets Tier field to given value.


### SetTierNil

`func (o *Stream) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *Stream) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetComponent

`func (o *Stream) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *Stream) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *Stream) SetComponent(v string)`

SetComponent sets Component field to given value.


### SetComponentNil

`func (o *Stream) SetComponentNil(b bool)`

 SetComponentNil sets the value for Component to be an explicit nil

### UnsetComponent
`func (o *Stream) UnsetComponent()`

UnsetComponent ensures that no value is present for Component, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


