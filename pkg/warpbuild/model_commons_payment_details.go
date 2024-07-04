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

// checks if the CommonsPaymentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonsPaymentDetails{}

// CommonsPaymentDetails struct for CommonsPaymentDetails
type CommonsPaymentDetails struct {
	Amount *int32 `json:"amount,omitempty"`
	BillingEmail *string `json:"billing_email,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Method *string `json:"method,omitempty"`
	MethodMeta *string `json:"method_meta,omitempty"`
	Ref *string `json:"ref,omitempty"`
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonsPaymentDetails CommonsPaymentDetails

// NewCommonsPaymentDetails instantiates a new CommonsPaymentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonsPaymentDetails() *CommonsPaymentDetails {
	this := CommonsPaymentDetails{}
	return &this
}

// NewCommonsPaymentDetailsWithDefaults instantiates a new CommonsPaymentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonsPaymentDetailsWithDefaults() *CommonsPaymentDetails {
	this := CommonsPaymentDetails{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CommonsPaymentDetails) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsPaymentDetails) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CommonsPaymentDetails) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *CommonsPaymentDetails) SetAmount(v int32) {
	o.Amount = &v
}

// GetBillingEmail returns the BillingEmail field value if set, zero value otherwise.
func (o *CommonsPaymentDetails) GetBillingEmail() string {
	if o == nil || IsNil(o.BillingEmail) {
		var ret string
		return ret
	}
	return *o.BillingEmail
}

// GetBillingEmailOk returns a tuple with the BillingEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsPaymentDetails) GetBillingEmailOk() (*string, bool) {
	if o == nil || IsNil(o.BillingEmail) {
		return nil, false
	}
	return o.BillingEmail, true
}

// HasBillingEmail returns a boolean if a field has been set.
func (o *CommonsPaymentDetails) HasBillingEmail() bool {
	if o != nil && !IsNil(o.BillingEmail) {
		return true
	}

	return false
}

// SetBillingEmail gets a reference to the given string and assigns it to the BillingEmail field.
func (o *CommonsPaymentDetails) SetBillingEmail(v string) {
	o.BillingEmail = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CommonsPaymentDetails) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsPaymentDetails) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CommonsPaymentDetails) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CommonsPaymentDetails) SetCurrency(v string) {
	o.Currency = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *CommonsPaymentDetails) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsPaymentDetails) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *CommonsPaymentDetails) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *CommonsPaymentDetails) SetMethod(v string) {
	o.Method = &v
}

// GetMethodMeta returns the MethodMeta field value if set, zero value otherwise.
func (o *CommonsPaymentDetails) GetMethodMeta() string {
	if o == nil || IsNil(o.MethodMeta) {
		var ret string
		return ret
	}
	return *o.MethodMeta
}

// GetMethodMetaOk returns a tuple with the MethodMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsPaymentDetails) GetMethodMetaOk() (*string, bool) {
	if o == nil || IsNil(o.MethodMeta) {
		return nil, false
	}
	return o.MethodMeta, true
}

// HasMethodMeta returns a boolean if a field has been set.
func (o *CommonsPaymentDetails) HasMethodMeta() bool {
	if o != nil && !IsNil(o.MethodMeta) {
		return true
	}

	return false
}

// SetMethodMeta gets a reference to the given string and assigns it to the MethodMeta field.
func (o *CommonsPaymentDetails) SetMethodMeta(v string) {
	o.MethodMeta = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *CommonsPaymentDetails) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsPaymentDetails) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *CommonsPaymentDetails) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *CommonsPaymentDetails) SetRef(v string) {
	o.Ref = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CommonsPaymentDetails) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonsPaymentDetails) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CommonsPaymentDetails) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CommonsPaymentDetails) SetStatus(v string) {
	o.Status = &v
}

func (o CommonsPaymentDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonsPaymentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BillingEmail) {
		toSerialize["billing_email"] = o.BillingEmail
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.MethodMeta) {
		toSerialize["method_meta"] = o.MethodMeta
	}
	if !IsNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonsPaymentDetails) UnmarshalJSON(bytes []byte) (err error) {
	varCommonsPaymentDetails := _CommonsPaymentDetails{}

<<<<<<< HEAD
	err = json.Unmarshal(bytes, &varCommonsPaymentDetails)

	if err != nil {
		return err
	}

	*o = CommonsPaymentDetails(varCommonsPaymentDetails)

=======
	if err = json.Unmarshal(bytes, &varCommonsPaymentDetails); err == nil {
		*o = CommonsPaymentDetails(varCommonsPaymentDetails)
	}

>>>>>>> prajjwal-warp-323
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "billing_email")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "method")
		delete(additionalProperties, "method_meta")
		delete(additionalProperties, "ref")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonsPaymentDetails struct {
	value *CommonsPaymentDetails
	isSet bool
}

func (v NullableCommonsPaymentDetails) Get() *CommonsPaymentDetails {
	return v.value
}

func (v *NullableCommonsPaymentDetails) Set(val *CommonsPaymentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonsPaymentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonsPaymentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonsPaymentDetails(val *CommonsPaymentDetails) *NullableCommonsPaymentDetails {
	return &NullableCommonsPaymentDetails{value: val, isSet: true}
}

func (v NullableCommonsPaymentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonsPaymentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


