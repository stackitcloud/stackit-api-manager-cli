# RetireVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | Identifier of API to be retired | [optional] [readonly] 
**ProjectId** | Pointer to **string** | Project ID for API to be retired | [optional] [readonly] 
**Version** | Pointer to **string** | API version to be retired | [optional] [readonly] 

## Methods

### NewRetireVersionRequest

`func NewRetireVersionRequest() *RetireVersionRequest`

NewRetireVersionRequest instantiates a new RetireVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetireVersionRequestWithDefaults

`func NewRetireVersionRequestWithDefaults() *RetireVersionRequest`

NewRetireVersionRequestWithDefaults instantiates a new RetireVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *RetireVersionRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RetireVersionRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RetireVersionRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *RetireVersionRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetProjectId

`func (o *RetireVersionRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RetireVersionRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RetireVersionRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RetireVersionRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetVersion

`func (o *RetireVersionRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RetireVersionRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RetireVersionRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RetireVersionRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


