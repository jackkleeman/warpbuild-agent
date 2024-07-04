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

// checks if the CommonsListTokensOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonsListTokensOptions{}

// CommonsListTokensOptions struct for CommonsListTokensOptions
type CommonsListTokensOptions struct {
	OrganizationId *string `json:"organization_id,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonsListTokensOptions CommonsListTokensOptions

// NewCommonsListTokensOptions instantiates a new CommonsListTokensOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonsListTokensOptions() *CommonsListTokensOptions {
	this := CommonsListTokensOptions{}
	return &this
}

// NewCommonsListTokensOptionsWithDefaults instantiates a new CommonsListTokensOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonsListTokensOptionsWithDefaults() *CommonsListTokensOptions {
	this := CommonsListTokensOptions{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *CommonsListTokensOptions) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsListTokensOptions) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CommonsListTokensOptions) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *CommonsListTokensOptions) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CommonsListTokensOptions) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsListTokensOptions) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CommonsListTokensOptions) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *CommonsListTokensOptions) SetUserId(v string) {
	o.UserId = &v
}

func (o CommonsListTokensOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonsListTokensOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonsListTokensOptions) UnmarshalJSON(bytes []byte) (err error) {
	varCommonsListTokensOptions := _CommonsListTokensOptions{}

<<<<<<< HEAD
	err = json.Unmarshal(bytes, &varCommonsListTokensOptions)

	if err != nil {
		return err
	}

	*o = CommonsListTokensOptions(varCommonsListTokensOptions)

=======
	if err = json.Unmarshal(bytes, &varCommonsListTokensOptions); err == nil {
		*o = CommonsListTokensOptions(varCommonsListTokensOptions)
	}

>>>>>>> prajjwal-warp-323
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "organization_id")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonsListTokensOptions struct {
	value *CommonsListTokensOptions
	isSet bool
}

func (v NullableCommonsListTokensOptions) Get() *CommonsListTokensOptions {
	return v.value
}

func (v *NullableCommonsListTokensOptions) Set(val *CommonsListTokensOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonsListTokensOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonsListTokensOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonsListTokensOptions(val *CommonsListTokensOptions) *NullableCommonsListTokensOptions {
	return &NullableCommonsListTokensOptions{value: val, isSet: true}
}

func (v NullableCommonsListTokensOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonsListTokensOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


