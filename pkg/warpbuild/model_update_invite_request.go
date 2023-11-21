/*
Warp Builds API Docs

This is the docs for warp builds api for argonaut

API version: 0.4.0
Contact: support@swagger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package warpbuild

import (
	"encoding/json"
)

// checks if the UpdateInviteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInviteRequest{}

// UpdateInviteRequest struct for UpdateInviteRequest
type UpdateInviteRequest struct {
	Status *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateInviteRequest UpdateInviteRequest

// NewUpdateInviteRequest instantiates a new UpdateInviteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInviteRequest() *UpdateInviteRequest {
	this := UpdateInviteRequest{}
	return &this
}

// NewUpdateInviteRequestWithDefaults instantiates a new UpdateInviteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInviteRequestWithDefaults() *UpdateInviteRequest {
	this := UpdateInviteRequest{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateInviteRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInviteRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateInviteRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateInviteRequest) SetStatus(v string) {
	o.Status = &v
}

func (o UpdateInviteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInviteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateInviteRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateInviteRequest := _UpdateInviteRequest{}

	err = json.Unmarshal(bytes, &varUpdateInviteRequest)

	if err != nil {
		return err
	}

	*o = UpdateInviteRequest(varUpdateInviteRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateInviteRequest struct {
	value *UpdateInviteRequest
	isSet bool
}

func (v NullableUpdateInviteRequest) Get() *UpdateInviteRequest {
	return v.value
}

func (v *NullableUpdateInviteRequest) Set(val *UpdateInviteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInviteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInviteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInviteRequest(val *UpdateInviteRequest) *NullableUpdateInviteRequest {
	return &NullableUpdateInviteRequest{value: val, isSet: true}
}

func (v NullableUpdateInviteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInviteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


