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

// checks if the CommonsListVCSRunnerGroupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonsListVCSRunnerGroupsResponse{}

// CommonsListVCSRunnerGroupsResponse struct for CommonsListVCSRunnerGroupsResponse
type CommonsListVCSRunnerGroupsResponse struct {
	LastSyncedAt *string `json:"last_synced_at,omitempty"`
	RunnerGroups []CommonsRunnerGroup `json:"runner_groups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonsListVCSRunnerGroupsResponse CommonsListVCSRunnerGroupsResponse

// NewCommonsListVCSRunnerGroupsResponse instantiates a new CommonsListVCSRunnerGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonsListVCSRunnerGroupsResponse() *CommonsListVCSRunnerGroupsResponse {
	this := CommonsListVCSRunnerGroupsResponse{}
	return &this
}

// NewCommonsListVCSRunnerGroupsResponseWithDefaults instantiates a new CommonsListVCSRunnerGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonsListVCSRunnerGroupsResponseWithDefaults() *CommonsListVCSRunnerGroupsResponse {
	this := CommonsListVCSRunnerGroupsResponse{}
	return &this
}

// GetLastSyncedAt returns the LastSyncedAt field value if set, zero value otherwise.
func (o *CommonsListVCSRunnerGroupsResponse) GetLastSyncedAt() string {
	if o == nil || IsNil(o.LastSyncedAt) {
		var ret string
		return ret
	}
	return *o.LastSyncedAt
}

// GetLastSyncedAtOk returns a tuple with the LastSyncedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsListVCSRunnerGroupsResponse) GetLastSyncedAtOk() (*string, bool) {
	if o == nil || IsNil(o.LastSyncedAt) {
		return nil, false
	}
	return o.LastSyncedAt, true
}

// HasLastSyncedAt returns a boolean if a field has been set.
func (o *CommonsListVCSRunnerGroupsResponse) HasLastSyncedAt() bool {
	if o != nil && !IsNil(o.LastSyncedAt) {
		return true
	}

	return false
}

// SetLastSyncedAt gets a reference to the given string and assigns it to the LastSyncedAt field.
func (o *CommonsListVCSRunnerGroupsResponse) SetLastSyncedAt(v string) {
	o.LastSyncedAt = &v
}

// GetRunnerGroups returns the RunnerGroups field value if set, zero value otherwise.
func (o *CommonsListVCSRunnerGroupsResponse) GetRunnerGroups() []CommonsRunnerGroup {
	if o == nil || IsNil(o.RunnerGroups) {
		var ret []CommonsRunnerGroup
		return ret
	}
	return o.RunnerGroups
}

// GetRunnerGroupsOk returns a tuple with the RunnerGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsListVCSRunnerGroupsResponse) GetRunnerGroupsOk() ([]CommonsRunnerGroup, bool) {
	if o == nil || IsNil(o.RunnerGroups) {
		return nil, false
	}
	return o.RunnerGroups, true
}

// HasRunnerGroups returns a boolean if a field has been set.
func (o *CommonsListVCSRunnerGroupsResponse) HasRunnerGroups() bool {
	if o != nil && !IsNil(o.RunnerGroups) {
		return true
	}

	return false
}

// SetRunnerGroups gets a reference to the given []CommonsRunnerGroup and assigns it to the RunnerGroups field.
func (o *CommonsListVCSRunnerGroupsResponse) SetRunnerGroups(v []CommonsRunnerGroup) {
	o.RunnerGroups = v
}

func (o CommonsListVCSRunnerGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonsListVCSRunnerGroupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastSyncedAt) {
		toSerialize["last_synced_at"] = o.LastSyncedAt
	}
	if !IsNil(o.RunnerGroups) {
		toSerialize["runner_groups"] = o.RunnerGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonsListVCSRunnerGroupsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCommonsListVCSRunnerGroupsResponse := _CommonsListVCSRunnerGroupsResponse{}

	err = json.Unmarshal(bytes, &varCommonsListVCSRunnerGroupsResponse)

	if err != nil {
		return err
	}

	*o = CommonsListVCSRunnerGroupsResponse(varCommonsListVCSRunnerGroupsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "last_synced_at")
		delete(additionalProperties, "runner_groups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonsListVCSRunnerGroupsResponse struct {
	value *CommonsListVCSRunnerGroupsResponse
	isSet bool
}

func (v NullableCommonsListVCSRunnerGroupsResponse) Get() *CommonsListVCSRunnerGroupsResponse {
	return v.value
}

func (v *NullableCommonsListVCSRunnerGroupsResponse) Set(val *CommonsListVCSRunnerGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonsListVCSRunnerGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonsListVCSRunnerGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonsListVCSRunnerGroupsResponse(val *CommonsListVCSRunnerGroupsResponse) *NullableCommonsListVCSRunnerGroupsResponse {
	return &NullableCommonsListVCSRunnerGroupsResponse{value: val, isSet: true}
}

func (v NullableCommonsListVCSRunnerGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonsListVCSRunnerGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


