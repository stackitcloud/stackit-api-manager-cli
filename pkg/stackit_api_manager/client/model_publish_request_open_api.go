/*
api-manager-api

STACKIT API Manager API

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PublishRequestOpenApi struct for PublishRequestOpenApi
type PublishRequestOpenApi struct {
	Base64Encoded *string `json:"base64Encoded,omitempty"`
}

// NewPublishRequestOpenApi instantiates a new PublishRequestOpenApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishRequestOpenApi() *PublishRequestOpenApi {
	this := PublishRequestOpenApi{}
	return &this
}

// NewPublishRequestOpenApiWithDefaults instantiates a new PublishRequestOpenApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishRequestOpenApiWithDefaults() *PublishRequestOpenApi {
	this := PublishRequestOpenApi{}
	return &this
}

// GetBase64Encoded returns the Base64Encoded field value if set, zero value otherwise.
func (o *PublishRequestOpenApi) GetBase64Encoded() string {
	if o == nil || o.Base64Encoded == nil {
		var ret string
		return ret
	}
	return *o.Base64Encoded
}

// GetBase64EncodedOk returns a tuple with the Base64Encoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishRequestOpenApi) GetBase64EncodedOk() (*string, bool) {
	if o == nil || o.Base64Encoded == nil {
		return nil, false
	}
	return o.Base64Encoded, true
}

// HasBase64Encoded returns a boolean if a field has been set.
func (o *PublishRequestOpenApi) HasBase64Encoded() bool {
	if o != nil && o.Base64Encoded != nil {
		return true
	}

	return false
}

// SetBase64Encoded gets a reference to the given string and assigns it to the Base64Encoded field.
func (o *PublishRequestOpenApi) SetBase64Encoded(v string) {
	o.Base64Encoded = &v
}

func (o PublishRequestOpenApi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Base64Encoded != nil {
		toSerialize["base64Encoded"] = o.Base64Encoded
	}
	return json.Marshal(toSerialize)
}

type NullablePublishRequestOpenApi struct {
	value *PublishRequestOpenApi
	isSet bool
}

func (v NullablePublishRequestOpenApi) Get() *PublishRequestOpenApi {
	return v.value
}

func (v *NullablePublishRequestOpenApi) Set(val *PublishRequestOpenApi) {
	v.value = val
	v.isSet = true
}

func (v NullablePublishRequestOpenApi) IsSet() bool {
	return v.isSet
}

func (v *NullablePublishRequestOpenApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublishRequestOpenApi(val *PublishRequestOpenApi) *NullablePublishRequestOpenApi {
	return &NullablePublishRequestOpenApi{value: val, isSet: true}
}

func (v NullablePublishRequestOpenApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublishRequestOpenApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


