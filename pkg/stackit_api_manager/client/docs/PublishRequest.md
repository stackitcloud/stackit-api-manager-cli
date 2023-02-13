# PublishRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** |  | [optional] [readonly] 
**Metadata** | Pointer to [**PublishMetadata**](PublishMetadata.md) |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] [readonly] 
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


