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

// RetireVersionRequest struct for RetireVersionRequest
type RetireVersionRequest struct {
	// Identifier of API to be retired
	Identifier *string `json:"identifier,omitempty"`
	// Project ID for API to be retired
	ProjectId *string `json:"projectId,omitempty"`
	// API version to be retired
	Version *string `json:"version,omitempty"`
}

// NewRetireVersionRequest instantiates a new RetireVersionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetireVersionRequest() *RetireVersionRequest {
	this := RetireVersionRequest{}
	return &this
}

// NewRetireVersionRequestWithDefaults instantiates a new RetireVersionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetireVersionRequestWithDefaults() *RetireVersionRequest {
	this := RetireVersionRequest{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *RetireVersionRequest) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetireVersionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *RetireVersionRequest) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *RetireVersionRequest) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *RetireVersionRequest) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetireVersionRequest) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *RetireVersionRequest) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *RetireVersionRequest) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RetireVersionRequest) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetireVersionRequest) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RetireVersionRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *RetireVersionRequest) SetVersion(v string) {
	o.Version = &v
}

func (o RetireVersionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableRetireVersionRequest struct {
	value *RetireVersionRequest
	isSet bool
}

func (v NullableRetireVersionRequest) Get() *RetireVersionRequest {
	return v.value
}

func (v *NullableRetireVersionRequest) Set(val *RetireVersionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRetireVersionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRetireVersionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetireVersionRequest(val *RetireVersionRequest) *NullableRetireVersionRequest {
	return &NullableRetireVersionRequest{value: val, isSet: true}
}

func (v NullableRetireVersionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetireVersionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
