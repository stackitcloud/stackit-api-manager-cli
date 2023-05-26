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

// PublishRequest struct for PublishRequest
type PublishRequest struct {
	Identifier *string `json:"identifier,omitempty"`
	IgnoreLintingErrors *bool `json:"ignoreLintingErrors,omitempty"`
	IgnoreBreakingChanges *bool `json:"ignoreBreakingChanges,omitempty"`
	Metadata *PublishMetadata `json:"metadata,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
	Spec *Spec `json:"spec,omitempty"`
}

// NewPublishRequest instantiates a new PublishRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishRequest() *PublishRequest {
	this := PublishRequest{}
	return &this
}

// NewPublishRequestWithDefaults instantiates a new PublishRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishRequestWithDefaults() *PublishRequest {
	this := PublishRequest{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *PublishRequest) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishRequest) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PublishRequest) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *PublishRequest) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetIgnoreLintingErrors returns the IgnoreLintingErrors field value if set, zero value otherwise.
func (o *PublishRequest) GetIgnoreLintingErrors() bool {
	if o == nil || o.IgnoreLintingErrors == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreLintingErrors
}

// GetIgnoreLintingErrorsOk returns a tuple with the IgnoreLintingErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishRequest) GetIgnoreLintingErrorsOk() (*bool, bool) {
	if o == nil || o.IgnoreLintingErrors == nil {
		return nil, false
	}
	return o.IgnoreLintingErrors, true
}

// HasIgnoreLintingErrors returns a boolean if a field has been set.
func (o *PublishRequest) HasIgnoreLintingErrors() bool {
	if o != nil && o.IgnoreLintingErrors != nil {
		return true
	}

	return false
}

// SetIgnoreLintingErrors gets a reference to the given bool and assigns it to the IgnoreLintingErrors field.
func (o *PublishRequest) SetIgnoreLintingErrors(v bool) {
	o.IgnoreLintingErrors = &v
}

// GetIgnoreBreakingChanges returns the IgnoreBreakingChanges field value if set, zero value otherwise.
func (o *PublishRequest) GetIgnoreBreakingChanges() bool {
	if o == nil || o.IgnoreBreakingChanges == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreBreakingChanges
}

// GetIgnoreBreakingChangesOk returns a tuple with the IgnoreBreakingChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishRequest) GetIgnoreBreakingChangesOk() (*bool, bool) {
	if o == nil || o.IgnoreBreakingChanges == nil {
		return nil, false
	}
	return o.IgnoreBreakingChanges, true
}

// HasIgnoreBreakingChanges returns a boolean if a field has been set.
func (o *PublishRequest) HasIgnoreBreakingChanges() bool {
	if o != nil && o.IgnoreBreakingChanges != nil {
		return true
	}

	return false
}

// SetIgnoreBreakingChanges gets a reference to the given bool and assigns it to the IgnoreBreakingChanges field.
func (o *PublishRequest) SetIgnoreBreakingChanges(v bool) {
	o.IgnoreBreakingChanges = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PublishRequest) GetMetadata() PublishMetadata {
	if o == nil || o.Metadata == nil {
		var ret PublishMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishRequest) GetMetadataOk() (*PublishMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PublishRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given PublishMetadata and assigns it to the Metadata field.
func (o *PublishRequest) SetMetadata(v PublishMetadata) {
	o.Metadata = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *PublishRequest) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishRequest) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *PublishRequest) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *PublishRequest) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *PublishRequest) GetSpec() Spec {
	if o == nil || o.Spec == nil {
		var ret Spec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishRequest) GetSpecOk() (*Spec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *PublishRequest) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given Spec and assigns it to the Spec field.
func (o *PublishRequest) SetSpec(v Spec) {
	o.Spec = &v
}

func (o PublishRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.IgnoreLintingErrors != nil {
		toSerialize["ignoreLintingErrors"] = o.IgnoreLintingErrors
	}
	if o.IgnoreBreakingChanges != nil {
		toSerialize["ignoreBreakingChanges"] = o.IgnoreBreakingChanges
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullablePublishRequest struct {
	value *PublishRequest
	isSet bool
}

func (v NullablePublishRequest) Get() *PublishRequest {
	return v.value
}

func (v *NullablePublishRequest) Set(val *PublishRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePublishRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePublishRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublishRequest(val *PublishRequest) *NullablePublishRequest {
	return &NullablePublishRequest{value: val, isSet: true}
}

func (v NullablePublishRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublishRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


