# EnumList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudTypes** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**ProjectStatuses** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**ServerRoles** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**ServerStatuses** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**SecurityGroupRules** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**PrometheusTypes** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**AuditLogs** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**RebootOptions** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**Availability** | [**[]CommonAvailabilityDto**](CommonAvailabilityDto.md) |  | 
**SlackTypes** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**RequestLogs** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**AzureQuotas** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**ShowbackKinds** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**AlertTypes** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**ReminderTypes** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**AwsPlatforms** | [**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md) |  | 
**CronPeriods** | [**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md) |  | 
**ValidityPeriods** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**AlertingIntegrationTypes** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**GoogleImageTypes** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**StandaloneVmStatuses** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**OpenstackContinents** | [**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md) |  | 
**RetentionPeriods** | [**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md) |  | 
**ProjectResourceTypes** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**ProjectResourceUnits** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**NonGlobalRoles** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**LbProviders** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**FilteredUserRoles** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 

## Methods

### NewEnumList

`func NewEnumList(cloudTypes []CommonDropdownDto, projectStatuses []CommonDropdownDto, serverRoles []CommonDropdownDto, serverStatuses []CommonDropdownDto, securityGroupRules []CommonDropdownDto, prometheusTypes []CommonDropdownDto, auditLogs []CommonDropdownDto, rebootOptions []CommonDropdownDto, availability []CommonAvailabilityDto, slackTypes []CommonDropdownDto, requestLogs []CommonDropdownDto, azureQuotas []CommonDropdownDto, showbackKinds []CommonDropdownDto, alertTypes []CommonDropdownDto, reminderTypes []CommonDropdownDto, awsPlatforms []CommonStringBasedDropdownDto, cronPeriods []CommonStringBasedDropdownDto, validityPeriods []CommonDropdownDto, alertingIntegrationTypes []CommonDropdownDto, googleImageTypes []CommonDropdownDto, standaloneVmStatuses []CommonDropdownDto, openstackContinents []CommonStringBasedDropdownDto, retentionPeriods []CommonStringBasedDropdownDto, projectResourceTypes []CommonDropdownDto, projectResourceUnits []CommonDropdownDto, nonGlobalRoles []CommonDropdownDto, lbProviders []CommonDropdownDto, filteredUserRoles []CommonDropdownDto, ) *EnumList`

NewEnumList instantiates a new EnumList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumListWithDefaults

`func NewEnumListWithDefaults() *EnumList`

NewEnumListWithDefaults instantiates a new EnumList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudTypes

`func (o *EnumList) GetCloudTypes() []CommonDropdownDto`

GetCloudTypes returns the CloudTypes field if non-nil, zero value otherwise.

### GetCloudTypesOk

`func (o *EnumList) GetCloudTypesOk() (*[]CommonDropdownDto, bool)`

GetCloudTypesOk returns a tuple with the CloudTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudTypes

`func (o *EnumList) SetCloudTypes(v []CommonDropdownDto)`

SetCloudTypes sets CloudTypes field to given value.


### SetCloudTypesNil

`func (o *EnumList) SetCloudTypesNil(b bool)`

 SetCloudTypesNil sets the value for CloudTypes to be an explicit nil

### UnsetCloudTypes
`func (o *EnumList) UnsetCloudTypes()`

UnsetCloudTypes ensures that no value is present for CloudTypes, not even an explicit nil
### GetProjectStatuses

`func (o *EnumList) GetProjectStatuses() []CommonDropdownDto`

GetProjectStatuses returns the ProjectStatuses field if non-nil, zero value otherwise.

### GetProjectStatusesOk

`func (o *EnumList) GetProjectStatusesOk() (*[]CommonDropdownDto, bool)`

GetProjectStatusesOk returns a tuple with the ProjectStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStatuses

`func (o *EnumList) SetProjectStatuses(v []CommonDropdownDto)`

SetProjectStatuses sets ProjectStatuses field to given value.


### SetProjectStatusesNil

`func (o *EnumList) SetProjectStatusesNil(b bool)`

 SetProjectStatusesNil sets the value for ProjectStatuses to be an explicit nil

### UnsetProjectStatuses
`func (o *EnumList) UnsetProjectStatuses()`

UnsetProjectStatuses ensures that no value is present for ProjectStatuses, not even an explicit nil
### GetServerRoles

`func (o *EnumList) GetServerRoles() []CommonDropdownDto`

GetServerRoles returns the ServerRoles field if non-nil, zero value otherwise.

### GetServerRolesOk

`func (o *EnumList) GetServerRolesOk() (*[]CommonDropdownDto, bool)`

GetServerRolesOk returns a tuple with the ServerRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRoles

`func (o *EnumList) SetServerRoles(v []CommonDropdownDto)`

SetServerRoles sets ServerRoles field to given value.


### SetServerRolesNil

`func (o *EnumList) SetServerRolesNil(b bool)`

 SetServerRolesNil sets the value for ServerRoles to be an explicit nil

### UnsetServerRoles
`func (o *EnumList) UnsetServerRoles()`

UnsetServerRoles ensures that no value is present for ServerRoles, not even an explicit nil
### GetServerStatuses

`func (o *EnumList) GetServerStatuses() []CommonDropdownDto`

GetServerStatuses returns the ServerStatuses field if non-nil, zero value otherwise.

### GetServerStatusesOk

`func (o *EnumList) GetServerStatusesOk() (*[]CommonDropdownDto, bool)`

GetServerStatusesOk returns a tuple with the ServerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatuses

`func (o *EnumList) SetServerStatuses(v []CommonDropdownDto)`

SetServerStatuses sets ServerStatuses field to given value.


### SetServerStatusesNil

`func (o *EnumList) SetServerStatusesNil(b bool)`

 SetServerStatusesNil sets the value for ServerStatuses to be an explicit nil

### UnsetServerStatuses
`func (o *EnumList) UnsetServerStatuses()`

UnsetServerStatuses ensures that no value is present for ServerStatuses, not even an explicit nil
### GetSecurityGroupRules

`func (o *EnumList) GetSecurityGroupRules() []CommonDropdownDto`

GetSecurityGroupRules returns the SecurityGroupRules field if non-nil, zero value otherwise.

### GetSecurityGroupRulesOk

`func (o *EnumList) GetSecurityGroupRulesOk() (*[]CommonDropdownDto, bool)`

GetSecurityGroupRulesOk returns a tuple with the SecurityGroupRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupRules

`func (o *EnumList) SetSecurityGroupRules(v []CommonDropdownDto)`

SetSecurityGroupRules sets SecurityGroupRules field to given value.


### SetSecurityGroupRulesNil

`func (o *EnumList) SetSecurityGroupRulesNil(b bool)`

 SetSecurityGroupRulesNil sets the value for SecurityGroupRules to be an explicit nil

### UnsetSecurityGroupRules
`func (o *EnumList) UnsetSecurityGroupRules()`

UnsetSecurityGroupRules ensures that no value is present for SecurityGroupRules, not even an explicit nil
### GetPrometheusTypes

`func (o *EnumList) GetPrometheusTypes() []CommonDropdownDto`

GetPrometheusTypes returns the PrometheusTypes field if non-nil, zero value otherwise.

### GetPrometheusTypesOk

`func (o *EnumList) GetPrometheusTypesOk() (*[]CommonDropdownDto, bool)`

GetPrometheusTypesOk returns a tuple with the PrometheusTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusTypes

`func (o *EnumList) SetPrometheusTypes(v []CommonDropdownDto)`

SetPrometheusTypes sets PrometheusTypes field to given value.


### SetPrometheusTypesNil

`func (o *EnumList) SetPrometheusTypesNil(b bool)`

 SetPrometheusTypesNil sets the value for PrometheusTypes to be an explicit nil

### UnsetPrometheusTypes
`func (o *EnumList) UnsetPrometheusTypes()`

UnsetPrometheusTypes ensures that no value is present for PrometheusTypes, not even an explicit nil
### GetAuditLogs

`func (o *EnumList) GetAuditLogs() []CommonDropdownDto`

GetAuditLogs returns the AuditLogs field if non-nil, zero value otherwise.

### GetAuditLogsOk

`func (o *EnumList) GetAuditLogsOk() (*[]CommonDropdownDto, bool)`

GetAuditLogsOk returns a tuple with the AuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogs

`func (o *EnumList) SetAuditLogs(v []CommonDropdownDto)`

SetAuditLogs sets AuditLogs field to given value.


### SetAuditLogsNil

`func (o *EnumList) SetAuditLogsNil(b bool)`

 SetAuditLogsNil sets the value for AuditLogs to be an explicit nil

### UnsetAuditLogs
`func (o *EnumList) UnsetAuditLogs()`

UnsetAuditLogs ensures that no value is present for AuditLogs, not even an explicit nil
### GetRebootOptions

`func (o *EnumList) GetRebootOptions() []CommonDropdownDto`

GetRebootOptions returns the RebootOptions field if non-nil, zero value otherwise.

### GetRebootOptionsOk

`func (o *EnumList) GetRebootOptionsOk() (*[]CommonDropdownDto, bool)`

GetRebootOptionsOk returns a tuple with the RebootOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootOptions

`func (o *EnumList) SetRebootOptions(v []CommonDropdownDto)`

SetRebootOptions sets RebootOptions field to given value.


### SetRebootOptionsNil

`func (o *EnumList) SetRebootOptionsNil(b bool)`

 SetRebootOptionsNil sets the value for RebootOptions to be an explicit nil

### UnsetRebootOptions
`func (o *EnumList) UnsetRebootOptions()`

UnsetRebootOptions ensures that no value is present for RebootOptions, not even an explicit nil
### GetAvailability

`func (o *EnumList) GetAvailability() []CommonAvailabilityDto`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *EnumList) GetAvailabilityOk() (*[]CommonAvailabilityDto, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *EnumList) SetAvailability(v []CommonAvailabilityDto)`

SetAvailability sets Availability field to given value.


### SetAvailabilityNil

`func (o *EnumList) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *EnumList) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetSlackTypes

`func (o *EnumList) GetSlackTypes() []CommonDropdownDto`

GetSlackTypes returns the SlackTypes field if non-nil, zero value otherwise.

### GetSlackTypesOk

`func (o *EnumList) GetSlackTypesOk() (*[]CommonDropdownDto, bool)`

GetSlackTypesOk returns a tuple with the SlackTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackTypes

`func (o *EnumList) SetSlackTypes(v []CommonDropdownDto)`

SetSlackTypes sets SlackTypes field to given value.


### SetSlackTypesNil

`func (o *EnumList) SetSlackTypesNil(b bool)`

 SetSlackTypesNil sets the value for SlackTypes to be an explicit nil

### UnsetSlackTypes
`func (o *EnumList) UnsetSlackTypes()`

UnsetSlackTypes ensures that no value is present for SlackTypes, not even an explicit nil
### GetRequestLogs

`func (o *EnumList) GetRequestLogs() []CommonDropdownDto`

GetRequestLogs returns the RequestLogs field if non-nil, zero value otherwise.

### GetRequestLogsOk

`func (o *EnumList) GetRequestLogsOk() (*[]CommonDropdownDto, bool)`

GetRequestLogsOk returns a tuple with the RequestLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLogs

`func (o *EnumList) SetRequestLogs(v []CommonDropdownDto)`

SetRequestLogs sets RequestLogs field to given value.


### SetRequestLogsNil

`func (o *EnumList) SetRequestLogsNil(b bool)`

 SetRequestLogsNil sets the value for RequestLogs to be an explicit nil

### UnsetRequestLogs
`func (o *EnumList) UnsetRequestLogs()`

UnsetRequestLogs ensures that no value is present for RequestLogs, not even an explicit nil
### GetAzureQuotas

`func (o *EnumList) GetAzureQuotas() []CommonDropdownDto`

GetAzureQuotas returns the AzureQuotas field if non-nil, zero value otherwise.

### GetAzureQuotasOk

`func (o *EnumList) GetAzureQuotasOk() (*[]CommonDropdownDto, bool)`

GetAzureQuotasOk returns a tuple with the AzureQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureQuotas

`func (o *EnumList) SetAzureQuotas(v []CommonDropdownDto)`

SetAzureQuotas sets AzureQuotas field to given value.


### SetAzureQuotasNil

`func (o *EnumList) SetAzureQuotasNil(b bool)`

 SetAzureQuotasNil sets the value for AzureQuotas to be an explicit nil

### UnsetAzureQuotas
`func (o *EnumList) UnsetAzureQuotas()`

UnsetAzureQuotas ensures that no value is present for AzureQuotas, not even an explicit nil
### GetShowbackKinds

`func (o *EnumList) GetShowbackKinds() []CommonDropdownDto`

GetShowbackKinds returns the ShowbackKinds field if non-nil, zero value otherwise.

### GetShowbackKindsOk

`func (o *EnumList) GetShowbackKindsOk() (*[]CommonDropdownDto, bool)`

GetShowbackKindsOk returns a tuple with the ShowbackKinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowbackKinds

`func (o *EnumList) SetShowbackKinds(v []CommonDropdownDto)`

SetShowbackKinds sets ShowbackKinds field to given value.


### SetShowbackKindsNil

`func (o *EnumList) SetShowbackKindsNil(b bool)`

 SetShowbackKindsNil sets the value for ShowbackKinds to be an explicit nil

### UnsetShowbackKinds
`func (o *EnumList) UnsetShowbackKinds()`

UnsetShowbackKinds ensures that no value is present for ShowbackKinds, not even an explicit nil
### GetAlertTypes

`func (o *EnumList) GetAlertTypes() []CommonDropdownDto`

GetAlertTypes returns the AlertTypes field if non-nil, zero value otherwise.

### GetAlertTypesOk

`func (o *EnumList) GetAlertTypesOk() (*[]CommonDropdownDto, bool)`

GetAlertTypesOk returns a tuple with the AlertTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypes

`func (o *EnumList) SetAlertTypes(v []CommonDropdownDto)`

SetAlertTypes sets AlertTypes field to given value.


### SetAlertTypesNil

`func (o *EnumList) SetAlertTypesNil(b bool)`

 SetAlertTypesNil sets the value for AlertTypes to be an explicit nil

### UnsetAlertTypes
`func (o *EnumList) UnsetAlertTypes()`

UnsetAlertTypes ensures that no value is present for AlertTypes, not even an explicit nil
### GetReminderTypes

`func (o *EnumList) GetReminderTypes() []CommonDropdownDto`

GetReminderTypes returns the ReminderTypes field if non-nil, zero value otherwise.

### GetReminderTypesOk

`func (o *EnumList) GetReminderTypesOk() (*[]CommonDropdownDto, bool)`

GetReminderTypesOk returns a tuple with the ReminderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderTypes

`func (o *EnumList) SetReminderTypes(v []CommonDropdownDto)`

SetReminderTypes sets ReminderTypes field to given value.


### SetReminderTypesNil

`func (o *EnumList) SetReminderTypesNil(b bool)`

 SetReminderTypesNil sets the value for ReminderTypes to be an explicit nil

### UnsetReminderTypes
`func (o *EnumList) UnsetReminderTypes()`

UnsetReminderTypes ensures that no value is present for ReminderTypes, not even an explicit nil
### GetAwsPlatforms

`func (o *EnumList) GetAwsPlatforms() []CommonStringBasedDropdownDto`

GetAwsPlatforms returns the AwsPlatforms field if non-nil, zero value otherwise.

### GetAwsPlatformsOk

`func (o *EnumList) GetAwsPlatformsOk() (*[]CommonStringBasedDropdownDto, bool)`

GetAwsPlatformsOk returns a tuple with the AwsPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsPlatforms

`func (o *EnumList) SetAwsPlatforms(v []CommonStringBasedDropdownDto)`

SetAwsPlatforms sets AwsPlatforms field to given value.


### SetAwsPlatformsNil

`func (o *EnumList) SetAwsPlatformsNil(b bool)`

 SetAwsPlatformsNil sets the value for AwsPlatforms to be an explicit nil

### UnsetAwsPlatforms
`func (o *EnumList) UnsetAwsPlatforms()`

UnsetAwsPlatforms ensures that no value is present for AwsPlatforms, not even an explicit nil
### GetCronPeriods

`func (o *EnumList) GetCronPeriods() []CommonStringBasedDropdownDto`

GetCronPeriods returns the CronPeriods field if non-nil, zero value otherwise.

### GetCronPeriodsOk

`func (o *EnumList) GetCronPeriodsOk() (*[]CommonStringBasedDropdownDto, bool)`

GetCronPeriodsOk returns a tuple with the CronPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronPeriods

`func (o *EnumList) SetCronPeriods(v []CommonStringBasedDropdownDto)`

SetCronPeriods sets CronPeriods field to given value.


### SetCronPeriodsNil

`func (o *EnumList) SetCronPeriodsNil(b bool)`

 SetCronPeriodsNil sets the value for CronPeriods to be an explicit nil

### UnsetCronPeriods
`func (o *EnumList) UnsetCronPeriods()`

UnsetCronPeriods ensures that no value is present for CronPeriods, not even an explicit nil
### GetValidityPeriods

`func (o *EnumList) GetValidityPeriods() []CommonDropdownDto`

GetValidityPeriods returns the ValidityPeriods field if non-nil, zero value otherwise.

### GetValidityPeriodsOk

`func (o *EnumList) GetValidityPeriodsOk() (*[]CommonDropdownDto, bool)`

GetValidityPeriodsOk returns a tuple with the ValidityPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriods

`func (o *EnumList) SetValidityPeriods(v []CommonDropdownDto)`

SetValidityPeriods sets ValidityPeriods field to given value.


### SetValidityPeriodsNil

`func (o *EnumList) SetValidityPeriodsNil(b bool)`

 SetValidityPeriodsNil sets the value for ValidityPeriods to be an explicit nil

### UnsetValidityPeriods
`func (o *EnumList) UnsetValidityPeriods()`

UnsetValidityPeriods ensures that no value is present for ValidityPeriods, not even an explicit nil
### GetAlertingIntegrationTypes

`func (o *EnumList) GetAlertingIntegrationTypes() []CommonDropdownDto`

GetAlertingIntegrationTypes returns the AlertingIntegrationTypes field if non-nil, zero value otherwise.

### GetAlertingIntegrationTypesOk

`func (o *EnumList) GetAlertingIntegrationTypesOk() (*[]CommonDropdownDto, bool)`

GetAlertingIntegrationTypesOk returns a tuple with the AlertingIntegrationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingIntegrationTypes

`func (o *EnumList) SetAlertingIntegrationTypes(v []CommonDropdownDto)`

SetAlertingIntegrationTypes sets AlertingIntegrationTypes field to given value.


### SetAlertingIntegrationTypesNil

`func (o *EnumList) SetAlertingIntegrationTypesNil(b bool)`

 SetAlertingIntegrationTypesNil sets the value for AlertingIntegrationTypes to be an explicit nil

### UnsetAlertingIntegrationTypes
`func (o *EnumList) UnsetAlertingIntegrationTypes()`

UnsetAlertingIntegrationTypes ensures that no value is present for AlertingIntegrationTypes, not even an explicit nil
### GetGoogleImageTypes

`func (o *EnumList) GetGoogleImageTypes() []CommonDropdownDto`

GetGoogleImageTypes returns the GoogleImageTypes field if non-nil, zero value otherwise.

### GetGoogleImageTypesOk

`func (o *EnumList) GetGoogleImageTypesOk() (*[]CommonDropdownDto, bool)`

GetGoogleImageTypesOk returns a tuple with the GoogleImageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleImageTypes

`func (o *EnumList) SetGoogleImageTypes(v []CommonDropdownDto)`

SetGoogleImageTypes sets GoogleImageTypes field to given value.


### SetGoogleImageTypesNil

`func (o *EnumList) SetGoogleImageTypesNil(b bool)`

 SetGoogleImageTypesNil sets the value for GoogleImageTypes to be an explicit nil

### UnsetGoogleImageTypes
`func (o *EnumList) UnsetGoogleImageTypes()`

UnsetGoogleImageTypes ensures that no value is present for GoogleImageTypes, not even an explicit nil
### GetStandaloneVmStatuses

`func (o *EnumList) GetStandaloneVmStatuses() []CommonDropdownDto`

GetStandaloneVmStatuses returns the StandaloneVmStatuses field if non-nil, zero value otherwise.

### GetStandaloneVmStatusesOk

`func (o *EnumList) GetStandaloneVmStatusesOk() (*[]CommonDropdownDto, bool)`

GetStandaloneVmStatusesOk returns a tuple with the StandaloneVmStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneVmStatuses

`func (o *EnumList) SetStandaloneVmStatuses(v []CommonDropdownDto)`

SetStandaloneVmStatuses sets StandaloneVmStatuses field to given value.


### SetStandaloneVmStatusesNil

`func (o *EnumList) SetStandaloneVmStatusesNil(b bool)`

 SetStandaloneVmStatusesNil sets the value for StandaloneVmStatuses to be an explicit nil

### UnsetStandaloneVmStatuses
`func (o *EnumList) UnsetStandaloneVmStatuses()`

UnsetStandaloneVmStatuses ensures that no value is present for StandaloneVmStatuses, not even an explicit nil
### GetOpenstackContinents

`func (o *EnumList) GetOpenstackContinents() []CommonStringBasedDropdownDto`

GetOpenstackContinents returns the OpenstackContinents field if non-nil, zero value otherwise.

### GetOpenstackContinentsOk

`func (o *EnumList) GetOpenstackContinentsOk() (*[]CommonStringBasedDropdownDto, bool)`

GetOpenstackContinentsOk returns a tuple with the OpenstackContinents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackContinents

`func (o *EnumList) SetOpenstackContinents(v []CommonStringBasedDropdownDto)`

SetOpenstackContinents sets OpenstackContinents field to given value.


### SetOpenstackContinentsNil

`func (o *EnumList) SetOpenstackContinentsNil(b bool)`

 SetOpenstackContinentsNil sets the value for OpenstackContinents to be an explicit nil

### UnsetOpenstackContinents
`func (o *EnumList) UnsetOpenstackContinents()`

UnsetOpenstackContinents ensures that no value is present for OpenstackContinents, not even an explicit nil
### GetRetentionPeriods

`func (o *EnumList) GetRetentionPeriods() []CommonStringBasedDropdownDto`

GetRetentionPeriods returns the RetentionPeriods field if non-nil, zero value otherwise.

### GetRetentionPeriodsOk

`func (o *EnumList) GetRetentionPeriodsOk() (*[]CommonStringBasedDropdownDto, bool)`

GetRetentionPeriodsOk returns a tuple with the RetentionPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriods

`func (o *EnumList) SetRetentionPeriods(v []CommonStringBasedDropdownDto)`

SetRetentionPeriods sets RetentionPeriods field to given value.


### SetRetentionPeriodsNil

`func (o *EnumList) SetRetentionPeriodsNil(b bool)`

 SetRetentionPeriodsNil sets the value for RetentionPeriods to be an explicit nil

### UnsetRetentionPeriods
`func (o *EnumList) UnsetRetentionPeriods()`

UnsetRetentionPeriods ensures that no value is present for RetentionPeriods, not even an explicit nil
### GetProjectResourceTypes

`func (o *EnumList) GetProjectResourceTypes() []CommonDropdownDto`

GetProjectResourceTypes returns the ProjectResourceTypes field if non-nil, zero value otherwise.

### GetProjectResourceTypesOk

`func (o *EnumList) GetProjectResourceTypesOk() (*[]CommonDropdownDto, bool)`

GetProjectResourceTypesOk returns a tuple with the ProjectResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectResourceTypes

`func (o *EnumList) SetProjectResourceTypes(v []CommonDropdownDto)`

SetProjectResourceTypes sets ProjectResourceTypes field to given value.


### SetProjectResourceTypesNil

`func (o *EnumList) SetProjectResourceTypesNil(b bool)`

 SetProjectResourceTypesNil sets the value for ProjectResourceTypes to be an explicit nil

### UnsetProjectResourceTypes
`func (o *EnumList) UnsetProjectResourceTypes()`

UnsetProjectResourceTypes ensures that no value is present for ProjectResourceTypes, not even an explicit nil
### GetProjectResourceUnits

`func (o *EnumList) GetProjectResourceUnits() []CommonDropdownDto`

GetProjectResourceUnits returns the ProjectResourceUnits field if non-nil, zero value otherwise.

### GetProjectResourceUnitsOk

`func (o *EnumList) GetProjectResourceUnitsOk() (*[]CommonDropdownDto, bool)`

GetProjectResourceUnitsOk returns a tuple with the ProjectResourceUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectResourceUnits

`func (o *EnumList) SetProjectResourceUnits(v []CommonDropdownDto)`

SetProjectResourceUnits sets ProjectResourceUnits field to given value.


### SetProjectResourceUnitsNil

`func (o *EnumList) SetProjectResourceUnitsNil(b bool)`

 SetProjectResourceUnitsNil sets the value for ProjectResourceUnits to be an explicit nil

### UnsetProjectResourceUnits
`func (o *EnumList) UnsetProjectResourceUnits()`

UnsetProjectResourceUnits ensures that no value is present for ProjectResourceUnits, not even an explicit nil
### GetNonGlobalRoles

`func (o *EnumList) GetNonGlobalRoles() []CommonDropdownDto`

GetNonGlobalRoles returns the NonGlobalRoles field if non-nil, zero value otherwise.

### GetNonGlobalRolesOk

`func (o *EnumList) GetNonGlobalRolesOk() (*[]CommonDropdownDto, bool)`

GetNonGlobalRolesOk returns a tuple with the NonGlobalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonGlobalRoles

`func (o *EnumList) SetNonGlobalRoles(v []CommonDropdownDto)`

SetNonGlobalRoles sets NonGlobalRoles field to given value.


### SetNonGlobalRolesNil

`func (o *EnumList) SetNonGlobalRolesNil(b bool)`

 SetNonGlobalRolesNil sets the value for NonGlobalRoles to be an explicit nil

### UnsetNonGlobalRoles
`func (o *EnumList) UnsetNonGlobalRoles()`

UnsetNonGlobalRoles ensures that no value is present for NonGlobalRoles, not even an explicit nil
### GetLbProviders

`func (o *EnumList) GetLbProviders() []CommonDropdownDto`

GetLbProviders returns the LbProviders field if non-nil, zero value otherwise.

### GetLbProvidersOk

`func (o *EnumList) GetLbProvidersOk() (*[]CommonDropdownDto, bool)`

GetLbProvidersOk returns a tuple with the LbProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbProviders

`func (o *EnumList) SetLbProviders(v []CommonDropdownDto)`

SetLbProviders sets LbProviders field to given value.


### SetLbProvidersNil

`func (o *EnumList) SetLbProvidersNil(b bool)`

 SetLbProvidersNil sets the value for LbProviders to be an explicit nil

### UnsetLbProviders
`func (o *EnumList) UnsetLbProviders()`

UnsetLbProviders ensures that no value is present for LbProviders, not even an explicit nil
### GetFilteredUserRoles

`func (o *EnumList) GetFilteredUserRoles() []CommonDropdownDto`

GetFilteredUserRoles returns the FilteredUserRoles field if non-nil, zero value otherwise.

### GetFilteredUserRolesOk

`func (o *EnumList) GetFilteredUserRolesOk() (*[]CommonDropdownDto, bool)`

GetFilteredUserRolesOk returns a tuple with the FilteredUserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredUserRoles

`func (o *EnumList) SetFilteredUserRoles(v []CommonDropdownDto)`

SetFilteredUserRoles sets FilteredUserRoles field to given value.


### SetFilteredUserRolesNil

`func (o *EnumList) SetFilteredUserRolesNil(b bool)`

 SetFilteredUserRolesNil sets the value for FilteredUserRoles to be an explicit nil

### UnsetFilteredUserRoles
`func (o *EnumList) UnsetFilteredUserRoles()`

UnsetFilteredUserRoles ensures that no value is present for FilteredUserRoles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


