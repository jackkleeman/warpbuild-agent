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

// checks if the AuthUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthUserRequest{}

// AuthUserRequest struct for AuthUserRequest
type AuthUserRequest struct {
	Code string `json:"code"`
	State string `json:"state"`
	AdditionalProperties map[string]interface{}
}

type _AuthUserRequest AuthUserRequest

// NewAuthUserRequest instantiates a new AuthUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthUserRequest(code string, state string) *AuthUserRequest {
	this := AuthUserRequest{}
	this.Code = code
	this.State = state
	return &this
}

// NewAuthUserRequestWithDefaults instantiates a new AuthUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthUserRequestWithDefaults() *AuthUserRequest {
	this := AuthUserRequest{}
	return &this
}

// GetCode returns the Code field value
func (o *AuthUserRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AuthUserRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AuthUserRequest) SetCode(v string) {
	o.Code = v
}

// GetState returns the State field value
func (o *AuthUserRequest) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *AuthUserRequest) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *AuthUserRequest) SetState(v string) {
	o.State = v
}

func (o AuthUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["state"] = o.State

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthUserRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAuthUserRequest := _AuthUserRequest{}

<<<<<<< HEAD
	err = json.Unmarshal(bytes, &varAuthUserRequest)

	if err != nil {
		return err
	}

	*o = AuthUserRequest(varAuthUserRequest)

=======
	if err = json.Unmarshal(bytes, &varAuthUserRequest); err == nil {
		*o = AuthUserRequest(varAuthUserRequest)
	}

>>>>>>> prajjwal-warp-323
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthUserRequest struct {
	value *AuthUserRequest
	isSet bool
}

func (v NullableAuthUserRequest) Get() *AuthUserRequest {
	return v.value
}

func (v *NullableAuthUserRequest) Set(val *AuthUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthUserRequest(val *AuthUserRequest) *NullableAuthUserRequest {
	return &NullableAuthUserRequest{value: val, isSet: true}
}

func (v NullableAuthUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


