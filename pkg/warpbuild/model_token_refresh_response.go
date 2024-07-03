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

// checks if the TokenRefreshResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenRefreshResponse{}

// TokenRefreshResponse struct for TokenRefreshResponse
type TokenRefreshResponse struct {
	AccessToken string `json:"access_token"`
	AdditionalProperties map[string]interface{}
}

type _TokenRefreshResponse TokenRefreshResponse

// NewTokenRefreshResponse instantiates a new TokenRefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRefreshResponse(accessToken string) *TokenRefreshResponse {
	this := TokenRefreshResponse{}
	this.AccessToken = accessToken
	return &this
}

// NewTokenRefreshResponseWithDefaults instantiates a new TokenRefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRefreshResponseWithDefaults() *TokenRefreshResponse {
	this := TokenRefreshResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *TokenRefreshResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TokenRefreshResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TokenRefreshResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

func (o TokenRefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenRefreshResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenRefreshResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTokenRefreshResponse := _TokenRefreshResponse{}

	if err = json.Unmarshal(bytes, &varTokenRefreshResponse); err == nil {
		*o = TokenRefreshResponse(varTokenRefreshResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenRefreshResponse struct {
	value *TokenRefreshResponse
	isSet bool
}

func (v NullableTokenRefreshResponse) Get() *TokenRefreshResponse {
	return v.value
}

func (v *NullableTokenRefreshResponse) Set(val *TokenRefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRefreshResponse(val *TokenRefreshResponse) *NullableTokenRefreshResponse {
	return &NullableTokenRefreshResponse{value: val, isSet: true}
}

func (v NullableTokenRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


