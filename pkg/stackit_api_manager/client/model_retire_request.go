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

// RetireRequest struct for RetireRequest
type RetireRequest struct {
	Metadata *RetireMetadata `json:"metadata,omitempty"`
}

// NewRetireRequest instantiates a new RetireRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetireRequest() *RetireRequest {
	this := RetireRequest{}
	return &this
}

// NewRetireRequestWithDefaults instantiates a new RetireRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetireRequestWithDefaults() *RetireRequest {
	this := RetireRequest{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RetireRequest) GetMetadata() RetireMetadata {
	if o == nil || o.Metadata == nil {
		var ret RetireMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetireRequest) GetMetadataOk() (*RetireMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RetireRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given RetireMetadata and assigns it to the Metadata field.
func (o *RetireRequest) SetMetadata(v RetireMetadata) {
	o.Metadata = &v
}

func (o RetireRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableRetireRequest struct {
	value *RetireRequest
	isSet bool
}

func (v NullableRetireRequest) Get() *RetireRequest {
	return v.value
}

func (v *NullableRetireRequest) Set(val *RetireRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRetireRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRetireRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetireRequest(val *RetireRequest) *NullableRetireRequest {
	return &NullableRetireRequest{value: val, isSet: true}
}

func (v NullableRetireRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetireRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
