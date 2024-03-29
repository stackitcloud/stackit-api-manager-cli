# PublishRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | Identifier of API to be published | [optional] [readonly] 
**IgnoreBreakingChanges** | Pointer to **bool** |  | [optional] 
**IgnoreLintingErrors** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to [**PublishMetadata**](PublishMetadata.md) |  | [optional] 
**ProjectId** | Pointer to **string** | Project ID under which the API is to be published | [optional] [readonly] 
**Spec** | Pointer to [**Spec**](Spec.md) |  | [optional] 

## Methods

### NewPublishRequest

`func NewPublishRequest() *PublishRequest`

NewPublishRequest instantiates a new PublishRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishRequestWithDefaults

`func NewPublishRequestWithDefaults() *PublishRequest`

NewPublishRequestWithDefaults instantiates a new PublishRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *PublishRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PublishRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PublishRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PublishRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIgnoreBreakingChanges

`func (o *PublishRequest) GetIgnoreBreakingChanges() bool`

GetIgnoreBreakingChanges returns the IgnoreBreakingChanges field if non-nil, zero value otherwise.

### GetIgnoreBreakingChangesOk

`func (o *PublishRequest) GetIgnoreBreakingChangesOk() (*bool, bool)`

GetIgnoreBreakingChangesOk returns a tuple with the IgnoreBreakingChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreBreakingChanges

`func (o *PublishRequest) SetIgnoreBreakingChanges(v bool)`

SetIgnoreBreakingChanges sets IgnoreBreakingChanges field to given value.

### HasIgnoreBreakingChanges

`func (o *PublishRequest) HasIgnoreBreakingChanges() bool`

HasIgnoreBreakingChanges returns a boolean if a field has been set.

### GetIgnoreLintingErrors

`func (o *PublishRequest) GetIgnoreLintingErrors() bool`

GetIgnoreLintingErrors returns the IgnoreLintingErrors field if non-nil, zero value otherwise.

### GetIgnoreLintingErrorsOk

`func (o *PublishRequest) GetIgnoreLintingErrorsOk() (*bool, bool)`

GetIgnoreLintingErrorsOk returns a tuple with the IgnoreLintingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreLintingErrors

`func (o *PublishRequest) SetIgnoreLintingErrors(v bool)`

SetIgnoreLintingErrors sets IgnoreLintingErrors field to given value.

### HasIgnoreLintingErrors

`func (o *PublishRequest) HasIgnoreLintingErrors() bool`

HasIgnoreLintingErrors returns a boolean if a field has been set.

### GetMetadata

`func (o *PublishRequest) GetMetadata() PublishMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PublishRequest) GetMetadataOk() (*PublishMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PublishRequest) SetMetadata(v PublishMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PublishRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProjectId

`func (o *PublishRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PublishRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PublishRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PublishRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSpec

`func (o *PublishRequest) GetSpec() Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PublishRequest) GetSpecOk() (*Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PublishRequest) SetSpec(v Spec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *PublishRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


