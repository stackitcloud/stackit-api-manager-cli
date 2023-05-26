/*
STACKIT API Management Service

STACKIT API Manager

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PublishResponse OpenAPI specification was published successfully
type PublishResponse struct {
	ApiUrl *string `json:"apiUrl,omitempty"`
}

// NewPublishResponse instantiates a new PublishResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishResponse() *PublishResponse {
	this := PublishResponse{}
	return &this
}

// NewPublishResponseWithDefaults instantiates a new PublishResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishResponseWithDefaults() *PublishResponse {
	this := PublishResponse{}
	return &this
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise.
func (o *PublishResponse) GetApiUrl() string {
	if o == nil || o.ApiUrl == nil {
		var ret string
		return ret
	}
	return *o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishResponse) GetApiUrlOk() (*string, bool) {
	if o == nil || o.ApiUrl == nil {
		return nil, false
	}
	return o.ApiUrl, true
}

// HasApiUrl returns a boolean if a field has been set.
func (o *PublishResponse) HasApiUrl() bool {
	if o != nil && o.ApiUrl != nil {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given string and assigns it to the ApiUrl field.
func (o *PublishResponse) SetApiUrl(v string) {
	o.ApiUrl = &v
}

func (o PublishResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiUrl != nil {
		toSerialize["apiUrl"] = o.ApiUrl
	}
	return json.Marshal(toSerialize)
}

type NullablePublishResponse struct {
	value *PublishResponse
	isSet bool
}

func (v NullablePublishResponse) Get() *PublishResponse {
	return v.value
}

func (v *NullablePublishResponse) Set(val *PublishResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePublishResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePublishResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublishResponse(val *PublishResponse) *NullablePublishResponse {
	return &NullablePublishResponse{value: val, isSet: true}
}

func (v NullablePublishResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublishResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
