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

// checks if the AuthUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthUserResponse{}

// AuthUserResponse struct for AuthUserResponse
type AuthUserResponse struct {
	AccessToken string `json:"access_token"`
<<<<<<< HEAD
	Organization V1Organization `json:"organization"`
	RefreshToken string `json:"refresh_token"`
	ShouldShowVcsIntegration bool `json:"should_show_vcs_integration"`
=======
	IsDifferentOrg bool `json:"is_different_org"`
	Organization V1Organization `json:"organization"`
	RefreshToken string `json:"refresh_token"`
>>>>>>> prajjwal-warp-323
	User V1User `json:"user"`
	AdditionalProperties map[string]interface{}
}

type _AuthUserResponse AuthUserResponse

// NewAuthUserResponse instantiates a new AuthUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
<<<<<<< HEAD
func NewAuthUserResponse(accessToken string, organization V1Organization, refreshToken string, shouldShowVcsIntegration bool, user V1User) *AuthUserResponse {
	this := AuthUserResponse{}
	this.AccessToken = accessToken
	this.Organization = organization
	this.RefreshToken = refreshToken
	this.ShouldShowVcsIntegration = shouldShowVcsIntegration
=======
func NewAuthUserResponse(accessToken string, isDifferentOrg bool, organization V1Organization, refreshToken string, user V1User) *AuthUserResponse {
	this := AuthUserResponse{}
	this.AccessToken = accessToken
	this.IsDifferentOrg = isDifferentOrg
	this.Organization = organization
	this.RefreshToken = refreshToken
>>>>>>> prajjwal-warp-323
	this.User = user
	return &this
}

// NewAuthUserResponseWithDefaults instantiates a new AuthUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthUserResponseWithDefaults() *AuthUserResponse {
	this := AuthUserResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *AuthUserResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *AuthUserResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *AuthUserResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

<<<<<<< HEAD
=======
// GetIsDifferentOrg returns the IsDifferentOrg field value
func (o *AuthUserResponse) GetIsDifferentOrg() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDifferentOrg
}

// GetIsDifferentOrgOk returns a tuple with the IsDifferentOrg field value
// and a boolean to check if the value has been set.
func (o *AuthUserResponse) GetIsDifferentOrgOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDifferentOrg, true
}

// SetIsDifferentOrg sets field value
func (o *AuthUserResponse) SetIsDifferentOrg(v bool) {
	o.IsDifferentOrg = v
}

>>>>>>> prajjwal-warp-323
// GetOrganization returns the Organization field value
func (o *AuthUserResponse) GetOrganization() V1Organization {
	if o == nil {
		var ret V1Organization
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *AuthUserResponse) GetOrganizationOk() (*V1Organization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *AuthUserResponse) SetOrganization(v V1Organization) {
	o.Organization = v
}

// GetRefreshToken returns the RefreshToken field value
func (o *AuthUserResponse) GetRefreshToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *AuthUserResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *AuthUserResponse) SetRefreshToken(v string) {
	o.RefreshToken = v
}

<<<<<<< HEAD
// GetShouldShowVcsIntegration returns the ShouldShowVcsIntegration field value
func (o *AuthUserResponse) GetShouldShowVcsIntegration() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShouldShowVcsIntegration
}

// GetShouldShowVcsIntegrationOk returns a tuple with the ShouldShowVcsIntegration field value
// and a boolean to check if the value has been set.
func (o *AuthUserResponse) GetShouldShowVcsIntegrationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShouldShowVcsIntegration, true
}

// SetShouldShowVcsIntegration sets field value
func (o *AuthUserResponse) SetShouldShowVcsIntegration(v bool) {
	o.ShouldShowVcsIntegration = v
}

=======
>>>>>>> prajjwal-warp-323
// GetUser returns the User field value
func (o *AuthUserResponse) GetUser() V1User {
	if o == nil {
		var ret V1User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *AuthUserResponse) GetUserOk() (*V1User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *AuthUserResponse) SetUser(v V1User) {
	o.User = v
}

func (o AuthUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
<<<<<<< HEAD
	toSerialize["organization"] = o.Organization
	toSerialize["refresh_token"] = o.RefreshToken
	toSerialize["should_show_vcs_integration"] = o.ShouldShowVcsIntegration
=======
	toSerialize["is_different_org"] = o.IsDifferentOrg
	toSerialize["organization"] = o.Organization
	toSerialize["refresh_token"] = o.RefreshToken
>>>>>>> prajjwal-warp-323
	toSerialize["user"] = o.User

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthUserResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAuthUserResponse := _AuthUserResponse{}

<<<<<<< HEAD
	err = json.Unmarshal(bytes, &varAuthUserResponse)

	if err != nil {
		return err
	}

	*o = AuthUserResponse(varAuthUserResponse)

=======
	if err = json.Unmarshal(bytes, &varAuthUserResponse); err == nil {
		*o = AuthUserResponse(varAuthUserResponse)
	}

>>>>>>> prajjwal-warp-323
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access_token")
<<<<<<< HEAD
		delete(additionalProperties, "organization")
		delete(additionalProperties, "refresh_token")
		delete(additionalProperties, "should_show_vcs_integration")
=======
		delete(additionalProperties, "is_different_org")
		delete(additionalProperties, "organization")
		delete(additionalProperties, "refresh_token")
>>>>>>> prajjwal-warp-323
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthUserResponse struct {
	value *AuthUserResponse
	isSet bool
}

func (v NullableAuthUserResponse) Get() *AuthUserResponse {
	return v.value
}

func (v *NullableAuthUserResponse) Set(val *AuthUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthUserResponse(val *AuthUserResponse) *NullableAuthUserResponse {
	return &NullableAuthUserResponse{value: val, isSet: true}
}

func (v NullableAuthUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


