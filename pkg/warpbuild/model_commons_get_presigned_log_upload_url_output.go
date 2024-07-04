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

// checks if the CommonsGetPresignedLogUploadURLOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonsGetPresignedLogUploadURLOutput{}

// CommonsGetPresignedLogUploadURLOutput struct for CommonsGetPresignedLogUploadURLOutput
type CommonsGetPresignedLogUploadURLOutput struct {
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonsGetPresignedLogUploadURLOutput CommonsGetPresignedLogUploadURLOutput

// NewCommonsGetPresignedLogUploadURLOutput instantiates a new CommonsGetPresignedLogUploadURLOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonsGetPresignedLogUploadURLOutput() *CommonsGetPresignedLogUploadURLOutput {
	this := CommonsGetPresignedLogUploadURLOutput{}
	return &this
}

// NewCommonsGetPresignedLogUploadURLOutputWithDefaults instantiates a new CommonsGetPresignedLogUploadURLOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonsGetPresignedLogUploadURLOutputWithDefaults() *CommonsGetPresignedLogUploadURLOutput {
	this := CommonsGetPresignedLogUploadURLOutput{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CommonsGetPresignedLogUploadURLOutput) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsGetPresignedLogUploadURLOutput) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CommonsGetPresignedLogUploadURLOutput) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CommonsGetPresignedLogUploadURLOutput) SetUrl(v string) {
	o.Url = &v
}

func (o CommonsGetPresignedLogUploadURLOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonsGetPresignedLogUploadURLOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonsGetPresignedLogUploadURLOutput) UnmarshalJSON(bytes []byte) (err error) {
	varCommonsGetPresignedLogUploadURLOutput := _CommonsGetPresignedLogUploadURLOutput{}

	if err = json.Unmarshal(bytes, &varCommonsGetPresignedLogUploadURLOutput); err == nil {
		*o = CommonsGetPresignedLogUploadURLOutput(varCommonsGetPresignedLogUploadURLOutput)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonsGetPresignedLogUploadURLOutput struct {
	value *CommonsGetPresignedLogUploadURLOutput
	isSet bool
}

func (v NullableCommonsGetPresignedLogUploadURLOutput) Get() *CommonsGetPresignedLogUploadURLOutput {
	return v.value
}

func (v *NullableCommonsGetPresignedLogUploadURLOutput) Set(val *CommonsGetPresignedLogUploadURLOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonsGetPresignedLogUploadURLOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonsGetPresignedLogUploadURLOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonsGetPresignedLogUploadURLOutput(val *CommonsGetPresignedLogUploadURLOutput) *NullableCommonsGetPresignedLogUploadURLOutput {
	return &NullableCommonsGetPresignedLogUploadURLOutput{value: val, isSet: true}
}

func (v NullableCommonsGetPresignedLogUploadURLOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonsGetPresignedLogUploadURLOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


