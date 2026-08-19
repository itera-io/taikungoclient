# BackupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludedNamespace** | **[]string** |  | 
**ExcludedNamespace** | **[]string** |  | 

## Methods

### NewBackupDto

`func NewBackupDto(includedNamespace []string, excludedNamespace []string, ) *BackupDto`

NewBackupDto instantiates a new BackupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupDtoWithDefaults

`func NewBackupDtoWithDefaults() *BackupDto`

NewBackupDtoWithDefaults instantiates a new BackupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludedNamespace

`func (o *BackupDto) GetIncludedNamespace() []string`

GetIncludedNamespace returns the IncludedNamespace field if non-nil, zero value otherwise.

### GetIncludedNamespaceOk

`func (o *BackupDto) GetIncludedNamespaceOk() (*[]string, bool)`

GetIncludedNamespaceOk returns a tuple with the IncludedNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedNamespace

`func (o *BackupDto) SetIncludedNamespace(v []string)`

SetIncludedNamespace sets IncludedNamespace field to given value.


### SetIncludedNamespaceNil

`func (o *BackupDto) SetIncludedNamespaceNil(b bool)`

 SetIncludedNamespaceNil sets the value for IncludedNamespace to be an explicit nil

### UnsetIncludedNamespace
`func (o *BackupDto) UnsetIncludedNamespace()`

UnsetIncludedNamespace ensures that no value is present for IncludedNamespace, not even an explicit nil
### GetExcludedNamespace

`func (o *BackupDto) GetExcludedNamespace() []string`

GetExcludedNamespace returns the ExcludedNamespace field if non-nil, zero value otherwise.

### GetExcludedNamespaceOk

`func (o *BackupDto) GetExcludedNamespaceOk() (*[]string, bool)`

GetExcludedNamespaceOk returns a tuple with the ExcludedNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNamespace

`func (o *BackupDto) SetExcludedNamespace(v []string)`

SetExcludedNamespace sets ExcludedNamespace field to given value.


### SetExcludedNamespaceNil

`func (o *BackupDto) SetExcludedNamespaceNil(b bool)`

 SetExcludedNamespaceNil sets the value for ExcludedNamespace to be an explicit nil

### UnsetExcludedNamespace
`func (o *BackupDto) UnsetExcludedNamespace()`

UnsetExcludedNamespace ensures that no value is present for ExcludedNamespace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


