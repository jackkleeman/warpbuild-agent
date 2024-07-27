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

// checks if the CommonsRunnerImagePullSecretAWS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonsRunnerImagePullSecretAWS{}

// CommonsRunnerImagePullSecretAWS struct for CommonsRunnerImagePullSecretAWS
type CommonsRunnerImagePullSecretAWS struct {
	AccessKeyId *string `json:"access_key_id,omitempty"`
	// AWSECRRepository is the short name of the ecr repository For example, if the complete uri for an image is <account_id>.dkr.ecr.<region>.amazonaws.com/acme/customrunners:v1.5.0 The AWS ECR Repo is `acme/customrunners`
	AwsEcrRepository *string `json:"aws_ecr_repository,omitempty"`
	Region *string `json:"region,omitempty"`
	SecretAccessKey *string `json:"secret_access_key,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonsRunnerImagePullSecretAWS CommonsRunnerImagePullSecretAWS

// NewCommonsRunnerImagePullSecretAWS instantiates a new CommonsRunnerImagePullSecretAWS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonsRunnerImagePullSecretAWS() *CommonsRunnerImagePullSecretAWS {
	this := CommonsRunnerImagePullSecretAWS{}
	return &this
}

// NewCommonsRunnerImagePullSecretAWSWithDefaults instantiates a new CommonsRunnerImagePullSecretAWS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonsRunnerImagePullSecretAWSWithDefaults() *CommonsRunnerImagePullSecretAWS {
	this := CommonsRunnerImagePullSecretAWS{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *CommonsRunnerImagePullSecretAWS) GetAccessKeyId() string {
	if o == nil || IsNil(o.AccessKeyId) {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRunnerImagePullSecretAWS) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKeyId) {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *CommonsRunnerImagePullSecretAWS) HasAccessKeyId() bool {
	if o != nil && !IsNil(o.AccessKeyId) {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *CommonsRunnerImagePullSecretAWS) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetAwsEcrRepository returns the AwsEcrRepository field value if set, zero value otherwise.
func (o *CommonsRunnerImagePullSecretAWS) GetAwsEcrRepository() string {
	if o == nil || IsNil(o.AwsEcrRepository) {
		var ret string
		return ret
	}
	return *o.AwsEcrRepository
}

// GetAwsEcrRepositoryOk returns a tuple with the AwsEcrRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRunnerImagePullSecretAWS) GetAwsEcrRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.AwsEcrRepository) {
		return nil, false
	}
	return o.AwsEcrRepository, true
}

// HasAwsEcrRepository returns a boolean if a field has been set.
func (o *CommonsRunnerImagePullSecretAWS) HasAwsEcrRepository() bool {
	if o != nil && !IsNil(o.AwsEcrRepository) {
		return true
	}

	return false
}

// SetAwsEcrRepository gets a reference to the given string and assigns it to the AwsEcrRepository field.
func (o *CommonsRunnerImagePullSecretAWS) SetAwsEcrRepository(v string) {
	o.AwsEcrRepository = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CommonsRunnerImagePullSecretAWS) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRunnerImagePullSecretAWS) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CommonsRunnerImagePullSecretAWS) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CommonsRunnerImagePullSecretAWS) SetRegion(v string) {
	o.Region = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *CommonsRunnerImagePullSecretAWS) GetSecretAccessKey() string {
	if o == nil || IsNil(o.SecretAccessKey) {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRunnerImagePullSecretAWS) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretAccessKey) {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *CommonsRunnerImagePullSecretAWS) HasSecretAccessKey() bool {
	if o != nil && !IsNil(o.SecretAccessKey) {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *CommonsRunnerImagePullSecretAWS) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

func (o CommonsRunnerImagePullSecretAWS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonsRunnerImagePullSecretAWS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessKeyId) {
		toSerialize["access_key_id"] = o.AccessKeyId
	}
	if !IsNil(o.AwsEcrRepository) {
		toSerialize["aws_ecr_repository"] = o.AwsEcrRepository
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.SecretAccessKey) {
		toSerialize["secret_access_key"] = o.SecretAccessKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonsRunnerImagePullSecretAWS) UnmarshalJSON(bytes []byte) (err error) {
	varCommonsRunnerImagePullSecretAWS := _CommonsRunnerImagePullSecretAWS{}

	err = json.Unmarshal(bytes, &varCommonsRunnerImagePullSecretAWS)

	if err != nil {
		return err
	}

	*o = CommonsRunnerImagePullSecretAWS(varCommonsRunnerImagePullSecretAWS)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access_key_id")
		delete(additionalProperties, "aws_ecr_repository")
		delete(additionalProperties, "region")
		delete(additionalProperties, "secret_access_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonsRunnerImagePullSecretAWS struct {
	value *CommonsRunnerImagePullSecretAWS
	isSet bool
}

func (v NullableCommonsRunnerImagePullSecretAWS) Get() *CommonsRunnerImagePullSecretAWS {
	return v.value
}

func (v *NullableCommonsRunnerImagePullSecretAWS) Set(val *CommonsRunnerImagePullSecretAWS) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonsRunnerImagePullSecretAWS) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonsRunnerImagePullSecretAWS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonsRunnerImagePullSecretAWS(val *CommonsRunnerImagePullSecretAWS) *NullableCommonsRunnerImagePullSecretAWS {
	return &NullableCommonsRunnerImagePullSecretAWS{value: val, isSet: true}
}

func (v NullableCommonsRunnerImagePullSecretAWS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonsRunnerImagePullSecretAWS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


