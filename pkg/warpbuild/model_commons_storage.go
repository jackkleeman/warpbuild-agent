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

// checks if the CommonsStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonsStorage{}

// CommonsStorage struct for CommonsStorage
type CommonsStorage struct {
	Iops *int32 `json:"iops,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Throughput *int32 `json:"throughput,omitempty"`
	Tier *string `json:"tier,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonsStorage CommonsStorage

// NewCommonsStorage instantiates a new CommonsStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonsStorage() *CommonsStorage {
	this := CommonsStorage{}
	return &this
}

// NewCommonsStorageWithDefaults instantiates a new CommonsStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonsStorageWithDefaults() *CommonsStorage {
	this := CommonsStorage{}
	return &this
}

// GetIops returns the Iops field value if set, zero value otherwise.
func (o *CommonsStorage) GetIops() int32 {
	if o == nil || IsNil(o.Iops) {
		var ret int32
		return ret
	}
	return *o.Iops
}

// GetIopsOk returns a tuple with the Iops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsStorage) GetIopsOk() (*int32, bool) {
	if o == nil || IsNil(o.Iops) {
		return nil, false
	}
	return o.Iops, true
}

// HasIops returns a boolean if a field has been set.
func (o *CommonsStorage) HasIops() bool {
	if o != nil && !IsNil(o.Iops) {
		return true
	}

	return false
}

// SetIops gets a reference to the given int32 and assigns it to the Iops field.
func (o *CommonsStorage) SetIops(v int32) {
	o.Iops = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CommonsStorage) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsStorage) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CommonsStorage) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *CommonsStorage) SetSize(v int32) {
	o.Size = &v
}

// GetThroughput returns the Throughput field value if set, zero value otherwise.
func (o *CommonsStorage) GetThroughput() int32 {
	if o == nil || IsNil(o.Throughput) {
		var ret int32
		return ret
	}
	return *o.Throughput
}

// GetThroughputOk returns a tuple with the Throughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsStorage) GetThroughputOk() (*int32, bool) {
	if o == nil || IsNil(o.Throughput) {
		return nil, false
	}
	return o.Throughput, true
}

// HasThroughput returns a boolean if a field has been set.
func (o *CommonsStorage) HasThroughput() bool {
	if o != nil && !IsNil(o.Throughput) {
		return true
	}

	return false
}

// SetThroughput gets a reference to the given int32 and assigns it to the Throughput field.
func (o *CommonsStorage) SetThroughput(v int32) {
	o.Throughput = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *CommonsStorage) GetTier() string {
	if o == nil || IsNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsStorage) GetTierOk() (*string, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *CommonsStorage) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *CommonsStorage) SetTier(v string) {
	o.Tier = &v
}

func (o CommonsStorage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonsStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Iops) {
		toSerialize["iops"] = o.Iops
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Throughput) {
		toSerialize["throughput"] = o.Throughput
	}
	if !IsNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonsStorage) UnmarshalJSON(bytes []byte) (err error) {
	varCommonsStorage := _CommonsStorage{}

	if err = json.Unmarshal(bytes, &varCommonsStorage); err == nil {
		*o = CommonsStorage(varCommonsStorage)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "iops")
		delete(additionalProperties, "size")
		delete(additionalProperties, "throughput")
		delete(additionalProperties, "tier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonsStorage struct {
	value *CommonsStorage
	isSet bool
}

func (v NullableCommonsStorage) Get() *CommonsStorage {
	return v.value
}

func (v *NullableCommonsStorage) Set(val *CommonsStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonsStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonsStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonsStorage(val *CommonsStorage) *NullableCommonsStorage {
	return &NullableCommonsStorage{value: val, isSet: true}
}

func (v NullableCommonsStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonsStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


