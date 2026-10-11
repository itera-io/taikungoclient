# ExportLokiLogsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to [**LokiLogsQuery**](LokiLogsQuery.md) |  | [optional] 

## Methods

### NewExportLokiLogsCommand

`func NewExportLokiLogsCommand() *ExportLokiLogsCommand`

NewExportLokiLogsCommand instantiates a new ExportLokiLogsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportLokiLogsCommandWithDefaults

`func NewExportLokiLogsCommandWithDefaults() *ExportLokiLogsCommand`

NewExportLokiLogsCommandWithDefaults instantiates a new ExportLokiLogsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ExportLokiLogsCommand) GetQuery() LokiLogsQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ExportLokiLogsCommand) GetQueryOk() (*LokiLogsQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ExportLokiLogsCommand) SetQuery(v LokiLogsQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ExportLokiLogsCommand) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


