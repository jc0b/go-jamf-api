/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// InitializeV1 Initial Jamf Pro setup data
type InitializeV1 struct {
	ActivationCode string `json:"activationCode"`
	InstitutionName string `json:"institutionName"`
	EulaAccepted bool `json:"eulaAccepted"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email *string `json:"email,omitempty"`
	JssUrl string `json:"jssUrl"`
}

// NewInitializeV1 instantiates a new InitializeV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitializeV1(activationCode string, institutionName string, eulaAccepted bool, username string, password string, jssUrl string) *InitializeV1 {
	this := InitializeV1{}
	this.ActivationCode = activationCode
	this.InstitutionName = institutionName
	this.EulaAccepted = eulaAccepted
	this.Username = username
	this.Password = password
	this.JssUrl = jssUrl
	return &this
}

// NewInitializeV1WithDefaults instantiates a new InitializeV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitializeV1WithDefaults() *InitializeV1 {
	this := InitializeV1{}
	return &this
}

// GetActivationCode returns the ActivationCode field value
func (o *InitializeV1) GetActivationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActivationCode
}

// GetActivationCodeOk returns a tuple with the ActivationCode field value
// and a boolean to check if the value has been set.
func (o *InitializeV1) GetActivationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivationCode, true
}

// SetActivationCode sets field value
func (o *InitializeV1) SetActivationCode(v string) {
	o.ActivationCode = v
}

// GetInstitutionName returns the InstitutionName field value
func (o *InitializeV1) GetInstitutionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionName
}

// GetInstitutionNameOk returns a tuple with the InstitutionName field value
// and a boolean to check if the value has been set.
func (o *InitializeV1) GetInstitutionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstitutionName, true
}

// SetInstitutionName sets field value
func (o *InitializeV1) SetInstitutionName(v string) {
	o.InstitutionName = v
}

// GetEulaAccepted returns the EulaAccepted field value
func (o *InitializeV1) GetEulaAccepted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EulaAccepted
}

// GetEulaAcceptedOk returns a tuple with the EulaAccepted field value
// and a boolean to check if the value has been set.
func (o *InitializeV1) GetEulaAcceptedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EulaAccepted, true
}

// SetEulaAccepted sets field value
func (o *InitializeV1) SetEulaAccepted(v bool) {
	o.EulaAccepted = v
}

// GetUsername returns the Username field value
func (o *InitializeV1) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *InitializeV1) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *InitializeV1) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *InitializeV1) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *InitializeV1) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *InitializeV1) SetPassword(v string) {
	o.Password = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InitializeV1) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitializeV1) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InitializeV1) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InitializeV1) SetEmail(v string) {
	o.Email = &v
}

// GetJssUrl returns the JssUrl field value
func (o *InitializeV1) GetJssUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JssUrl
}

// GetJssUrlOk returns a tuple with the JssUrl field value
// and a boolean to check if the value has been set.
func (o *InitializeV1) GetJssUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JssUrl, true
}

// SetJssUrl sets field value
func (o *InitializeV1) SetJssUrl(v string) {
	o.JssUrl = v
}

func (o InitializeV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["activationCode"] = o.ActivationCode
	}
	if true {
		toSerialize["institutionName"] = o.InstitutionName
	}
	if true {
		toSerialize["eulaAccepted"] = o.EulaAccepted
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["jssUrl"] = o.JssUrl
	}
	return json.Marshal(toSerialize)
}

type NullableInitializeV1 struct {
	value *InitializeV1
	isSet bool
}

func (v NullableInitializeV1) Get() *InitializeV1 {
	return v.value
}

func (v *NullableInitializeV1) Set(val *InitializeV1) {
	v.value = val
	v.isSet = true
}

func (v NullableInitializeV1) IsSet() bool {
	return v.isSet
}

func (v *NullableInitializeV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitializeV1(val *InitializeV1) *NullableInitializeV1 {
	return &NullableInitializeV1{value: val, isSet: true}
}

func (v NullableInitializeV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitializeV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


