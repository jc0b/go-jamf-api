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

// EnrollmentCustomizationPanelLdapAuth struct for EnrollmentCustomizationPanelLdapAuth
type EnrollmentCustomizationPanelLdapAuth struct {
	DisplayName string `json:"displayName"`
	Rank int32 `json:"rank"`
	UsernameLabel string `json:"usernameLabel"`
	PasswordLabel string `json:"passwordLabel"`
	Title string `json:"title"`
	BackButtonText string `json:"backButtonText"`
	ContinueButtonText string `json:"continueButtonText"`
	LdapGroupAccess []EnrollmentCustomizationLdapGroupAccess `json:"ldapGroupAccess,omitempty"`
}

// NewEnrollmentCustomizationPanelLdapAuth instantiates a new EnrollmentCustomizationPanelLdapAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentCustomizationPanelLdapAuth(displayName string, rank int32, usernameLabel string, passwordLabel string, title string, backButtonText string, continueButtonText string) *EnrollmentCustomizationPanelLdapAuth {
	this := EnrollmentCustomizationPanelLdapAuth{}
	this.DisplayName = displayName
	this.Rank = rank
	this.UsernameLabel = usernameLabel
	this.PasswordLabel = passwordLabel
	this.Title = title
	this.BackButtonText = backButtonText
	this.ContinueButtonText = continueButtonText
	return &this
}

// NewEnrollmentCustomizationPanelLdapAuthWithDefaults instantiates a new EnrollmentCustomizationPanelLdapAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentCustomizationPanelLdapAuthWithDefaults() *EnrollmentCustomizationPanelLdapAuth {
	this := EnrollmentCustomizationPanelLdapAuth{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *EnrollmentCustomizationPanelLdapAuth) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuth) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *EnrollmentCustomizationPanelLdapAuth) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetRank returns the Rank field value
func (o *EnrollmentCustomizationPanelLdapAuth) GetRank() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuth) GetRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *EnrollmentCustomizationPanelLdapAuth) SetRank(v int32) {
	o.Rank = v
}

// GetUsernameLabel returns the UsernameLabel field value
func (o *EnrollmentCustomizationPanelLdapAuth) GetUsernameLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsernameLabel
}

// GetUsernameLabelOk returns a tuple with the UsernameLabel field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuth) GetUsernameLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsernameLabel, true
}

// SetUsernameLabel sets field value
func (o *EnrollmentCustomizationPanelLdapAuth) SetUsernameLabel(v string) {
	o.UsernameLabel = v
}

// GetPasswordLabel returns the PasswordLabel field value
func (o *EnrollmentCustomizationPanelLdapAuth) GetPasswordLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordLabel
}

// GetPasswordLabelOk returns a tuple with the PasswordLabel field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuth) GetPasswordLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordLabel, true
}

// SetPasswordLabel sets field value
func (o *EnrollmentCustomizationPanelLdapAuth) SetPasswordLabel(v string) {
	o.PasswordLabel = v
}

// GetTitle returns the Title field value
func (o *EnrollmentCustomizationPanelLdapAuth) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuth) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *EnrollmentCustomizationPanelLdapAuth) SetTitle(v string) {
	o.Title = v
}

// GetBackButtonText returns the BackButtonText field value
func (o *EnrollmentCustomizationPanelLdapAuth) GetBackButtonText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackButtonText
}

// GetBackButtonTextOk returns a tuple with the BackButtonText field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuth) GetBackButtonTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackButtonText, true
}

// SetBackButtonText sets field value
func (o *EnrollmentCustomizationPanelLdapAuth) SetBackButtonText(v string) {
	o.BackButtonText = v
}

// GetContinueButtonText returns the ContinueButtonText field value
func (o *EnrollmentCustomizationPanelLdapAuth) GetContinueButtonText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContinueButtonText
}

// GetContinueButtonTextOk returns a tuple with the ContinueButtonText field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuth) GetContinueButtonTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContinueButtonText, true
}

// SetContinueButtonText sets field value
func (o *EnrollmentCustomizationPanelLdapAuth) SetContinueButtonText(v string) {
	o.ContinueButtonText = v
}

// GetLdapGroupAccess returns the LdapGroupAccess field value if set, zero value otherwise.
func (o *EnrollmentCustomizationPanelLdapAuth) GetLdapGroupAccess() []EnrollmentCustomizationLdapGroupAccess {
	if o == nil || o.LdapGroupAccess == nil {
		var ret []EnrollmentCustomizationLdapGroupAccess
		return ret
	}
	return o.LdapGroupAccess
}

// GetLdapGroupAccessOk returns a tuple with the LdapGroupAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuth) GetLdapGroupAccessOk() ([]EnrollmentCustomizationLdapGroupAccess, bool) {
	if o == nil || o.LdapGroupAccess == nil {
		return nil, false
	}
	return o.LdapGroupAccess, true
}

// HasLdapGroupAccess returns a boolean if a field has been set.
func (o *EnrollmentCustomizationPanelLdapAuth) HasLdapGroupAccess() bool {
	if o != nil && o.LdapGroupAccess != nil {
		return true
	}

	return false
}

// SetLdapGroupAccess gets a reference to the given []EnrollmentCustomizationLdapGroupAccess and assigns it to the LdapGroupAccess field.
func (o *EnrollmentCustomizationPanelLdapAuth) SetLdapGroupAccess(v []EnrollmentCustomizationLdapGroupAccess) {
	o.LdapGroupAccess = v
}

func (o EnrollmentCustomizationPanelLdapAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["rank"] = o.Rank
	}
	if true {
		toSerialize["usernameLabel"] = o.UsernameLabel
	}
	if true {
		toSerialize["passwordLabel"] = o.PasswordLabel
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["backButtonText"] = o.BackButtonText
	}
	if true {
		toSerialize["continueButtonText"] = o.ContinueButtonText
	}
	if o.LdapGroupAccess != nil {
		toSerialize["ldapGroupAccess"] = o.LdapGroupAccess
	}
	return json.Marshal(toSerialize)
}

type NullableEnrollmentCustomizationPanelLdapAuth struct {
	value *EnrollmentCustomizationPanelLdapAuth
	isSet bool
}

func (v NullableEnrollmentCustomizationPanelLdapAuth) Get() *EnrollmentCustomizationPanelLdapAuth {
	return v.value
}

func (v *NullableEnrollmentCustomizationPanelLdapAuth) Set(val *EnrollmentCustomizationPanelLdapAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentCustomizationPanelLdapAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentCustomizationPanelLdapAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentCustomizationPanelLdapAuth(val *EnrollmentCustomizationPanelLdapAuth) *NullableEnrollmentCustomizationPanelLdapAuth {
	return &NullableEnrollmentCustomizationPanelLdapAuth{value: val, isSet: true}
}

func (v NullableEnrollmentCustomizationPanelLdapAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentCustomizationPanelLdapAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


