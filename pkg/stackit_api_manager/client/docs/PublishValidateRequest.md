# PublishValidateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**PublishMetadata**](PublishMetadata.md) |  | [optional] 
**Spec** | Pointer to [**Spec**](Spec.md) |  | [optional] 

## Methods

### NewPublishValidateRequest

`func NewPublishValidateRequest() *PublishValidateRequest`

NewPublishValidateRequest instantiates a new PublishValidateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishValidateRequestWithDefaults

`func NewPublishValidateRequestWithDefaults() *PublishValidateRequest`

NewPublishValidateRequestWithDefaults instantiates a new PublishValidateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *PublishValidateRequest) GetMetadata() PublishMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PublishValidateRequest) GetMetadataOk() (*PublishMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PublishValidateRequest) SetMetadata(v PublishMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PublishValidateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *PublishValidateRequest) GetSpec() Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PublishValidateRequest) GetSpecOk() (*Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PublishValidateRequest) SetSpec(v Spec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *PublishValidateRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


