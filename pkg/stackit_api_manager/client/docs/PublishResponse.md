# PublishResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiUrl** | Pointer to **string** | URL under which published API is available | [optional] 
**LinterWarnings** | Pointer to **[]string** | List of specific warnings returned by the API specification linter | [optional] 
**LinterWarningsCount** | Pointer to **string** | Number of warnings returned by the API specification linter | [optional] 
**PullRequestCreated** | Pointer to **bool** | Indicates whether the PR was created | [optional] 
**PullRequestUrl** | Pointer to **string** | URL of the PR which is created with the published specification | [optional] 

## Methods

### NewPublishResponse

`func NewPublishResponse() *PublishResponse`

NewPublishResponse instantiates a new PublishResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishResponseWithDefaults

`func NewPublishResponseWithDefaults() *PublishResponse`

NewPublishResponseWithDefaults instantiates a new PublishResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiUrl

`func (o *PublishResponse) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *PublishResponse) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *PublishResponse) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *PublishResponse) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetLinterWarnings

`func (o *PublishResponse) GetLinterWarnings() []string`

GetLinterWarnings returns the LinterWarnings field if non-nil, zero value otherwise.

### GetLinterWarningsOk

`func (o *PublishResponse) GetLinterWarningsOk() (*[]string, bool)`

GetLinterWarningsOk returns a tuple with the LinterWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinterWarnings

`func (o *PublishResponse) SetLinterWarnings(v []string)`

SetLinterWarnings sets LinterWarnings field to given value.

### HasLinterWarnings

`func (o *PublishResponse) HasLinterWarnings() bool`

HasLinterWarnings returns a boolean if a field has been set.

### GetLinterWarningsCount

`func (o *PublishResponse) GetLinterWarningsCount() string`

GetLinterWarningsCount returns the LinterWarningsCount field if non-nil, zero value otherwise.

### GetLinterWarningsCountOk

`func (o *PublishResponse) GetLinterWarningsCountOk() (*string, bool)`

GetLinterWarningsCountOk returns a tuple with the LinterWarningsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinterWarningsCount

`func (o *PublishResponse) SetLinterWarningsCount(v string)`

SetLinterWarningsCount sets LinterWarningsCount field to given value.

### HasLinterWarningsCount

`func (o *PublishResponse) HasLinterWarningsCount() bool`

HasLinterWarningsCount returns a boolean if a field has been set.

### GetPullRequestCreated

`func (o *PublishResponse) GetPullRequestCreated() bool`

GetPullRequestCreated returns the PullRequestCreated field if non-nil, zero value otherwise.

### GetPullRequestCreatedOk

`func (o *PublishResponse) GetPullRequestCreatedOk() (*bool, bool)`

GetPullRequestCreatedOk returns a tuple with the PullRequestCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestCreated

`func (o *PublishResponse) SetPullRequestCreated(v bool)`

SetPullRequestCreated sets PullRequestCreated field to given value.

### HasPullRequestCreated

`func (o *PublishResponse) HasPullRequestCreated() bool`

HasPullRequestCreated returns a boolean if a field has been set.

### GetPullRequestUrl

`func (o *PublishResponse) GetPullRequestUrl() string`

GetPullRequestUrl returns the PullRequestUrl field if non-nil, zero value otherwise.

### GetPullRequestUrlOk

`func (o *PublishResponse) GetPullRequestUrlOk() (*string, bool)`

GetPullRequestUrlOk returns a tuple with the PullRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestUrl

`func (o *PublishResponse) SetPullRequestUrl(v string)`

SetPullRequestUrl sets PullRequestUrl field to given value.

### HasPullRequestUrl

`func (o *PublishResponse) HasPullRequestUrl() bool`

HasPullRequestUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


