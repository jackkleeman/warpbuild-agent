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

// checks if the DebuggerPublishEventInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebuggerPublishEventInput{}

// DebuggerPublishEventInput struct for DebuggerPublishEventInput
type DebuggerPublishEventInput struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Id *string `json:"id,omitempty"`
	Publisher *string `json:"publisher,omitempty"`
	Topic *string `json:"topic,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DebuggerPublishEventInput DebuggerPublishEventInput

// NewDebuggerPublishEventInput instantiates a new DebuggerPublishEventInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebuggerPublishEventInput() *DebuggerPublishEventInput {
	this := DebuggerPublishEventInput{}
	return &this
}

// NewDebuggerPublishEventInputWithDefaults instantiates a new DebuggerPublishEventInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebuggerPublishEventInputWithDefaults() *DebuggerPublishEventInput {
	this := DebuggerPublishEventInput{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DebuggerPublishEventInput) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebuggerPublishEventInput) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DebuggerPublishEventInput) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *DebuggerPublishEventInput) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DebuggerPublishEventInput) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebuggerPublishEventInput) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DebuggerPublishEventInput) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DebuggerPublishEventInput) SetId(v string) {
	o.Id = &v
}

// GetPublisher returns the Publisher field value if set, zero value otherwise.
func (o *DebuggerPublishEventInput) GetPublisher() string {
	if o == nil || IsNil(o.Publisher) {
		var ret string
		return ret
	}
	return *o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebuggerPublishEventInput) GetPublisherOk() (*string, bool) {
	if o == nil || IsNil(o.Publisher) {
		return nil, false
	}
	return o.Publisher, true
}

// HasPublisher returns a boolean if a field has been set.
func (o *DebuggerPublishEventInput) HasPublisher() bool {
	if o != nil && !IsNil(o.Publisher) {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given string and assigns it to the Publisher field.
func (o *DebuggerPublishEventInput) SetPublisher(v string) {
	o.Publisher = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *DebuggerPublishEventInput) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebuggerPublishEventInput) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *DebuggerPublishEventInput) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *DebuggerPublishEventInput) SetTopic(v string) {
	o.Topic = &v
}

func (o DebuggerPublishEventInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebuggerPublishEventInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Publisher) {
		toSerialize["publisher"] = o.Publisher
	}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DebuggerPublishEventInput) UnmarshalJSON(bytes []byte) (err error) {
	varDebuggerPublishEventInput := _DebuggerPublishEventInput{}

	err = json.Unmarshal(bytes, &varDebuggerPublishEventInput)

	if err != nil {
		return err
	}

	*o = DebuggerPublishEventInput(varDebuggerPublishEventInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "id")
		delete(additionalProperties, "publisher")
		delete(additionalProperties, "topic")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDebuggerPublishEventInput struct {
	value *DebuggerPublishEventInput
	isSet bool
}

func (v NullableDebuggerPublishEventInput) Get() *DebuggerPublishEventInput {
	return v.value
}

func (v *NullableDebuggerPublishEventInput) Set(val *DebuggerPublishEventInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDebuggerPublishEventInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDebuggerPublishEventInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebuggerPublishEventInput(val *DebuggerPublishEventInput) *NullableDebuggerPublishEventInput {
	return &NullableDebuggerPublishEventInput{value: val, isSet: true}
}

func (v NullableDebuggerPublishEventInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebuggerPublishEventInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


