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

// checks if the ApproveVCSIntegrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApproveVCSIntegrationRequest{}

// ApproveVCSIntegrationRequest struct for ApproveVCSIntegrationRequest
type ApproveVCSIntegrationRequest struct {
	Code *string `json:"code,omitempty"`
	InstallationId *string `json:"installation_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApproveVCSIntegrationRequest ApproveVCSIntegrationRequest

// NewApproveVCSIntegrationRequest instantiates a new ApproveVCSIntegrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproveVCSIntegrationRequest() *ApproveVCSIntegrationRequest {
	this := ApproveVCSIntegrationRequest{}
	return &this
}

// NewApproveVCSIntegrationRequestWithDefaults instantiates a new ApproveVCSIntegrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApproveVCSIntegrationRequestWithDefaults() *ApproveVCSIntegrationRequest {
	this := ApproveVCSIntegrationRequest{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApproveVCSIntegrationRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApproveVCSIntegrationRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApproveVCSIntegrationRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApproveVCSIntegrationRequest) SetCode(v string) {
	o.Code = &v
}

// GetInstallationId returns the InstallationId field value if set, zero value otherwise.
func (o *ApproveVCSIntegrationRequest) GetInstallationId() string {
	if o == nil || IsNil(o.InstallationId) {
		var ret string
		return ret
	}
	return *o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApproveVCSIntegrationRequest) GetInstallationIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstallationId) {
		return nil, false
	}
	return o.InstallationId, true
}

// HasInstallationId returns a boolean if a field has been set.
func (o *ApproveVCSIntegrationRequest) HasInstallationId() bool {
	if o != nil && !IsNil(o.InstallationId) {
		return true
	}

	return false
}

// SetInstallationId gets a reference to the given string and assigns it to the InstallationId field.
func (o *ApproveVCSIntegrationRequest) SetInstallationId(v string) {
	o.InstallationId = &v
}

func (o ApproveVCSIntegrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApproveVCSIntegrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.InstallationId) {
		toSerialize["installation_id"] = o.InstallationId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApproveVCSIntegrationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varApproveVCSIntegrationRequest := _ApproveVCSIntegrationRequest{}

	if err = json.Unmarshal(bytes, &varApproveVCSIntegrationRequest); err == nil {
		*o = ApproveVCSIntegrationRequest(varApproveVCSIntegrationRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "installation_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApproveVCSIntegrationRequest struct {
	value *ApproveVCSIntegrationRequest
	isSet bool
}

func (v NullableApproveVCSIntegrationRequest) Get() *ApproveVCSIntegrationRequest {
	return v.value
}

func (v *NullableApproveVCSIntegrationRequest) Set(val *ApproveVCSIntegrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApproveVCSIntegrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApproveVCSIntegrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproveVCSIntegrationRequest(val *ApproveVCSIntegrationRequest) *NullableApproveVCSIntegrationRequest {
	return &NullableApproveVCSIntegrationRequest{value: val, isSet: true}
}

func (v NullableApproveVCSIntegrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproveVCSIntegrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


