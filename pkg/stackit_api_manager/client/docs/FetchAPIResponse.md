# FetchAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiUrl** | Pointer to **string** | URL under which fetched API is available | [optional] 
**Spec** | Pointer to [**Spec**](Spec.md) |  | [optional] 
**Stage** | Pointer to **string** | Server stage under which fetched API is published | [optional] 
**UpstreamUrl** | Pointer to **string** | URL for the upstream server targeted by the fetched API | [optional] 

## Methods

### NewFetchAPIResponse

`func NewFetchAPIResponse() *FetchAPIResponse`

NewFetchAPIResponse instantiates a new FetchAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchAPIResponseWithDefaults

`func NewFetchAPIResponseWithDefaults() *FetchAPIResponse`

NewFetchAPIResponseWithDefaults instantiates a new FetchAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiUrl

`func (o *FetchAPIResponse) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *FetchAPIResponse) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *FetchAPIResponse) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *FetchAPIResponse) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetSpec

`func (o *FetchAPIResponse) GetSpec() Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *FetchAPIResponse) GetSpecOk() (*Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *FetchAPIResponse) SetSpec(v Spec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *FetchAPIResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStage

`func (o *FetchAPIResponse) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *FetchAPIResponse) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *FetchAPIResponse) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *FetchAPIResponse) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetUpstreamUrl

`func (o *FetchAPIResponse) GetUpstreamUrl() string`

GetUpstreamUrl returns the UpstreamUrl field if non-nil, zero value otherwise.

### GetUpstreamUrlOk

`func (o *FetchAPIResponse) GetUpstreamUrlOk() (*string, bool)`

GetUpstreamUrlOk returns a tuple with the UpstreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamUrl

`func (o *FetchAPIResponse) SetUpstreamUrl(v string)`

SetUpstreamUrl sets UpstreamUrl field to given value.

### HasUpstreamUrl

`func (o *FetchAPIResponse) HasUpstreamUrl() bool`

HasUpstreamUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


