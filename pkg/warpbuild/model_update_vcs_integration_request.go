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

// checks if the UpdateVCSIntegrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVCSIntegrationRequest{}

// UpdateVCSIntegrationRequest struct for UpdateVCSIntegrationRequest
type UpdateVCSIntegrationRequest struct {
	// Code is  - 'code' param in the callback from gitlab
	Code *string `json:"code,omitempty"`
	Id string `json:"id"`
	// InstallationID  - 'installation_id' param from github installation
	InstallationId *string `json:"installation_id,omitempty"`
	// SetupAction  - 'setup_action' param from github installation
	SetupAction *string `json:"setup_action,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateVCSIntegrationRequest UpdateVCSIntegrationRequest

// NewUpdateVCSIntegrationRequest instantiates a new UpdateVCSIntegrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVCSIntegrationRequest(id string) *UpdateVCSIntegrationRequest {
	this := UpdateVCSIntegrationRequest{}
	this.Id = id
	return &this
}

// NewUpdateVCSIntegrationRequestWithDefaults instantiates a new UpdateVCSIntegrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVCSIntegrationRequestWithDefaults() *UpdateVCSIntegrationRequest {
	this := UpdateVCSIntegrationRequest{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UpdateVCSIntegrationRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVCSIntegrationRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UpdateVCSIntegrationRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UpdateVCSIntegrationRequest) SetCode(v string) {
	o.Code = &v
}

// GetId returns the Id field value
func (o *UpdateVCSIntegrationRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateVCSIntegrationRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateVCSIntegrationRequest) SetId(v string) {
	o.Id = v
}

// GetInstallationId returns the InstallationId field value if set, zero value otherwise.
func (o *UpdateVCSIntegrationRequest) GetInstallationId() string {
	if o == nil || IsNil(o.InstallationId) {
		var ret string
		return ret
	}
	return *o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVCSIntegrationRequest) GetInstallationIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstallationId) {
		return nil, false
	}
	return o.InstallationId, true
}

// HasInstallationId returns a boolean if a field has been set.
func (o *UpdateVCSIntegrationRequest) HasInstallationId() bool {
	if o != nil && !IsNil(o.InstallationId) {
		return true
	}

	return false
}

// SetInstallationId gets a reference to the given string and assigns it to the InstallationId field.
func (o *UpdateVCSIntegrationRequest) SetInstallationId(v string) {
	o.InstallationId = &v
}

// GetSetupAction returns the SetupAction field value if set, zero value otherwise.
func (o *UpdateVCSIntegrationRequest) GetSetupAction() string {
	if o == nil || IsNil(o.SetupAction) {
		var ret string
		return ret
	}
	return *o.SetupAction
}

// GetSetupActionOk returns a tuple with the SetupAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVCSIntegrationRequest) GetSetupActionOk() (*string, bool) {
	if o == nil || IsNil(o.SetupAction) {
		return nil, false
	}
	return o.SetupAction, true
}

// HasSetupAction returns a boolean if a field has been set.
func (o *UpdateVCSIntegrationRequest) HasSetupAction() bool {
	if o != nil && !IsNil(o.SetupAction) {
		return true
	}

	return false
}

// SetSetupAction gets a reference to the given string and assigns it to the SetupAction field.
func (o *UpdateVCSIntegrationRequest) SetSetupAction(v string) {
	o.SetupAction = &v
}

func (o UpdateVCSIntegrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVCSIntegrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.InstallationId) {
		toSerialize["installation_id"] = o.InstallationId
	}
	if !IsNil(o.SetupAction) {
		toSerialize["setup_action"] = o.SetupAction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateVCSIntegrationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateVCSIntegrationRequest := _UpdateVCSIntegrationRequest{}

	err = json.Unmarshal(bytes, &varUpdateVCSIntegrationRequest)

	if err != nil {
		return err
	}

	*o = UpdateVCSIntegrationRequest(varUpdateVCSIntegrationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "id")
		delete(additionalProperties, "installation_id")
		delete(additionalProperties, "setup_action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateVCSIntegrationRequest struct {
	value *UpdateVCSIntegrationRequest
	isSet bool
}

func (v NullableUpdateVCSIntegrationRequest) Get() *UpdateVCSIntegrationRequest {
	return v.value
}

func (v *NullableUpdateVCSIntegrationRequest) Set(val *UpdateVCSIntegrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVCSIntegrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVCSIntegrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVCSIntegrationRequest(val *UpdateVCSIntegrationRequest) *NullableUpdateVCSIntegrationRequest {
	return &NullableUpdateVCSIntegrationRequest{value: val, isSet: true}
}

func (v NullableUpdateVCSIntegrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVCSIntegrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


