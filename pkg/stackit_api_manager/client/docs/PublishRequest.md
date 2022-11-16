# PublishRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Spec** | Pointer to [**PublishRequestSpec**](PublishRequestSpec.md) |  | [optional] 

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

### GetMetadata

`func (o *PublishRequest) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PublishRequest) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PublishRequest) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PublishRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *PublishRequest) GetSpec() PublishRequestSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PublishRequest) GetSpecOk() (*PublishRequestSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PublishRequest) SetSpec(v PublishRequestSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *PublishRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


