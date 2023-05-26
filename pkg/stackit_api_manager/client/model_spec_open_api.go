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

// SpecOpenApi struct for SpecOpenApi
type SpecOpenApi struct {
	// now the spec will be base64 string, later we might add a url to a spec file or something else
	Base64Encoded *string `json:"base64Encoded,omitempty"`
}

// NewSpecOpenApi instantiates a new SpecOpenApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecOpenApi() *SpecOpenApi {
	this := SpecOpenApi{}
	return &this
}

// NewSpecOpenApiWithDefaults instantiates a new SpecOpenApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecOpenApiWithDefaults() *SpecOpenApi {
	this := SpecOpenApi{}
	return &this
}

// GetBase64Encoded returns the Base64Encoded field value if set, zero value otherwise.
func (o *SpecOpenApi) GetBase64Encoded() string {
	if o == nil || o.Base64Encoded == nil {
		var ret string
		return ret
	}
	return *o.Base64Encoded
}

// GetBase64EncodedOk returns a tuple with the Base64Encoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecOpenApi) GetBase64EncodedOk() (*string, bool) {
	if o == nil || o.Base64Encoded == nil {
		return nil, false
	}
	return o.Base64Encoded, true
}

// HasBase64Encoded returns a boolean if a field has been set.
func (o *SpecOpenApi) HasBase64Encoded() bool {
	if o != nil && o.Base64Encoded != nil {
		return true
	}

	return false
}

// SetBase64Encoded gets a reference to the given string and assigns it to the Base64Encoded field.
func (o *SpecOpenApi) SetBase64Encoded(v string) {
	o.Base64Encoded = &v
}

func (o SpecOpenApi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Base64Encoded != nil {
		toSerialize["base64Encoded"] = o.Base64Encoded
	}
	return json.Marshal(toSerialize)
}

type NullableSpecOpenApi struct {
	value *SpecOpenApi
	isSet bool
}

func (v NullableSpecOpenApi) Get() *SpecOpenApi {
	return v.value
}

func (v *NullableSpecOpenApi) Set(val *SpecOpenApi) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecOpenApi) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecOpenApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecOpenApi(val *SpecOpenApi) *NullableSpecOpenApi {
	return &NullableSpecOpenApi{value: val, isSet: true}
}

func (v NullableSpecOpenApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecOpenApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


