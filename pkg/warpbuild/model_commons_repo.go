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

// checks if the CommonsRepo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonsRepo{}

// CommonsRepo struct for CommonsRepo
type CommonsRepo struct {
	AccountId *string `json:"account_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	DefaultBranch *string `json:"default_branch,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Id *string `json:"id,omitempty"`
	IntegrationId *string `json:"integration_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Owner *string `json:"owner,omitempty"`
	Provider *string `json:"provider,omitempty"`
	Sshurl *string `json:"sshurl,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	VcsId *string `json:"vcs_id,omitempty"`
	WebUrl *string `json:"web_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonsRepo CommonsRepo

// NewCommonsRepo instantiates a new CommonsRepo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonsRepo() *CommonsRepo {
	this := CommonsRepo{}
	return &this
}

// NewCommonsRepoWithDefaults instantiates a new CommonsRepo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonsRepoWithDefaults() *CommonsRepo {
	this := CommonsRepo{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *CommonsRepo) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRepo) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *CommonsRepo) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *CommonsRepo) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CommonsRepo) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRepo) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CommonsRepo) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CommonsRepo) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise.
func (o *CommonsRepo) GetDefaultBranch() string {
	if o == nil || IsNil(o.DefaultBranch) {
		var ret string
		return ret
	}
	return *o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRepo) GetDefaultBranchOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultBranch) {
		return nil, false
	}
	return o.DefaultBranch, true
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *CommonsRepo) HasDefaultBranch() bool {
	if o != nil && !IsNil(o.DefaultBranch) {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given string and assigns it to the DefaultBranch field.
func (o *CommonsRepo) SetDefaultBranch(v string) {
	o.DefaultBranch = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CommonsRepo) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRepo) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CommonsRepo) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CommonsRepo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CommonsRepo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRepo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CommonsRepo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CommonsRepo) SetId(v string) {
	o.Id = &v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *CommonsRepo) GetIntegrationId() string {
	if o == nil || IsNil(o.IntegrationId) {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRepo) GetIntegrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationId) {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *CommonsRepo) HasIntegrationId() bool {
	if o != nil && !IsNil(o.IntegrationId) {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *CommonsRepo) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CommonsRepo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRepo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CommonsRepo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CommonsRepo) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *CommonsRepo) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRepo) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *CommonsRepo) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *CommonsRepo) SetOwner(v string) {
	o.Owner = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CommonsRepo) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRepo) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CommonsRepo) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *CommonsRepo) SetProvider(v string) {
	o.Provider = &v
}

// GetSshurl returns the Sshurl field value if set, zero value otherwise.
func (o *CommonsRepo) GetSshurl() string {
	if o == nil || IsNil(o.Sshurl) {
		var ret string
		return ret
	}
	return *o.Sshurl
}

// GetSshurlOk returns a tuple with the Sshurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRepo) GetSshurlOk() (*string, bool) {
	if o == nil || IsNil(o.Sshurl) {
		return nil, false
	}
	return o.Sshurl, true
}

// HasSshurl returns a boolean if a field has been set.
func (o *CommonsRepo) HasSshurl() bool {
	if o != nil && !IsNil(o.Sshurl) {
		return true
	}

	return false
}

// SetSshurl gets a reference to the given string and assigns it to the Sshurl field.
func (o *CommonsRepo) SetSshurl(v string) {
	o.Sshurl = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CommonsRepo) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRepo) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CommonsRepo) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *CommonsRepo) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetVcsId returns the VcsId field value if set, zero value otherwise.
func (o *CommonsRepo) GetVcsId() string {
	if o == nil || IsNil(o.VcsId) {
		var ret string
		return ret
	}
	return *o.VcsId
}

// GetVcsIdOk returns a tuple with the VcsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRepo) GetVcsIdOk() (*string, bool) {
	if o == nil || IsNil(o.VcsId) {
		return nil, false
	}
	return o.VcsId, true
}

// HasVcsId returns a boolean if a field has been set.
func (o *CommonsRepo) HasVcsId() bool {
	if o != nil && !IsNil(o.VcsId) {
		return true
	}

	return false
}

// SetVcsId gets a reference to the given string and assigns it to the VcsId field.
func (o *CommonsRepo) SetVcsId(v string) {
	o.VcsId = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise.
func (o *CommonsRepo) GetWebUrl() string {
	if o == nil || IsNil(o.WebUrl) {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsRepo) GetWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebUrl) {
		return nil, false
	}
	return o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *CommonsRepo) HasWebUrl() bool {
	if o != nil && !IsNil(o.WebUrl) {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *CommonsRepo) SetWebUrl(v string) {
	o.WebUrl = &v
}

func (o CommonsRepo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonsRepo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DefaultBranch) {
		toSerialize["default_branch"] = o.DefaultBranch
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IntegrationId) {
		toSerialize["integration_id"] = o.IntegrationId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Sshurl) {
		toSerialize["sshurl"] = o.Sshurl
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.VcsId) {
		toSerialize["vcs_id"] = o.VcsId
	}
	if !IsNil(o.WebUrl) {
		toSerialize["web_url"] = o.WebUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonsRepo) UnmarshalJSON(bytes []byte) (err error) {
	varCommonsRepo := _CommonsRepo{}

	if err = json.Unmarshal(bytes, &varCommonsRepo); err == nil {
		*o = CommonsRepo(varCommonsRepo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "default_branch")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "id")
		delete(additionalProperties, "integration_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "sshurl")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "vcs_id")
		delete(additionalProperties, "web_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonsRepo struct {
	value *CommonsRepo
	isSet bool
}

func (v NullableCommonsRepo) Get() *CommonsRepo {
	return v.value
}

func (v *NullableCommonsRepo) Set(val *CommonsRepo) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonsRepo) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonsRepo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonsRepo(val *CommonsRepo) *NullableCommonsRepo {
	return &NullableCommonsRepo{value: val, isSet: true}
}

func (v NullableCommonsRepo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonsRepo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


