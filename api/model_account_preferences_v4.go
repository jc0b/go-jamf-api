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

// AccountPreferencesV4 struct for AccountPreferencesV4
type AccountPreferencesV4 struct {
	Language string `json:"language"`
	DateFormat string `json:"dateFormat"`
	Timezone string `json:"timezone"`
	DisableRelativeDates bool `json:"disableRelativeDates"`
	DisablePageLeaveCheck bool `json:"disablePageLeaveCheck"`
	DisableTablePagination bool `json:"disableTablePagination"`
}

// NewAccountPreferencesV4 instantiates a new AccountPreferencesV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountPreferencesV4(language string, dateFormat string, timezone string, disableRelativeDates bool, disablePageLeaveCheck bool, disableTablePagination bool) *AccountPreferencesV4 {
	this := AccountPreferencesV4{}
	this.Language = language
	this.DateFormat = dateFormat
	this.Timezone = timezone
	this.DisableRelativeDates = disableRelativeDates
	this.DisablePageLeaveCheck = disablePageLeaveCheck
	this.DisableTablePagination = disableTablePagination
	return &this
}

// NewAccountPreferencesV4WithDefaults instantiates a new AccountPreferencesV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountPreferencesV4WithDefaults() *AccountPreferencesV4 {
	this := AccountPreferencesV4{}
	return &this
}

// GetLanguage returns the Language field value
func (o *AccountPreferencesV4) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *AccountPreferencesV4) SetLanguage(v string) {
	o.Language = v
}

// GetDateFormat returns the DateFormat field value
func (o *AccountPreferencesV4) GetDateFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetDateFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateFormat, true
}

// SetDateFormat sets field value
func (o *AccountPreferencesV4) SetDateFormat(v string) {
	o.DateFormat = v
}

// GetTimezone returns the Timezone field value
func (o *AccountPreferencesV4) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *AccountPreferencesV4) SetTimezone(v string) {
	o.Timezone = v
}

// GetDisableRelativeDates returns the DisableRelativeDates field value
func (o *AccountPreferencesV4) GetDisableRelativeDates() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableRelativeDates
}

// GetDisableRelativeDatesOk returns a tuple with the DisableRelativeDates field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetDisableRelativeDatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableRelativeDates, true
}

// SetDisableRelativeDates sets field value
func (o *AccountPreferencesV4) SetDisableRelativeDates(v bool) {
	o.DisableRelativeDates = v
}

// GetDisablePageLeaveCheck returns the DisablePageLeaveCheck field value
func (o *AccountPreferencesV4) GetDisablePageLeaveCheck() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisablePageLeaveCheck
}

// GetDisablePageLeaveCheckOk returns a tuple with the DisablePageLeaveCheck field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetDisablePageLeaveCheckOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisablePageLeaveCheck, true
}

// SetDisablePageLeaveCheck sets field value
func (o *AccountPreferencesV4) SetDisablePageLeaveCheck(v bool) {
	o.DisablePageLeaveCheck = v
}

// GetDisableTablePagination returns the DisableTablePagination field value
func (o *AccountPreferencesV4) GetDisableTablePagination() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableTablePagination
}

// GetDisableTablePaginationOk returns a tuple with the DisableTablePagination field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetDisableTablePaginationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableTablePagination, true
}

// SetDisableTablePagination sets field value
func (o *AccountPreferencesV4) SetDisableTablePagination(v bool) {
	o.DisableTablePagination = v
}

func (o AccountPreferencesV4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["language"] = o.Language
	}
	if true {
		toSerialize["dateFormat"] = o.DateFormat
	}
	if true {
		toSerialize["timezone"] = o.Timezone
	}
	if true {
		toSerialize["disableRelativeDates"] = o.DisableRelativeDates
	}
	if true {
		toSerialize["disablePageLeaveCheck"] = o.DisablePageLeaveCheck
	}
	if true {
		toSerialize["disableTablePagination"] = o.DisableTablePagination
	}
	return json.Marshal(toSerialize)
}

type NullableAccountPreferencesV4 struct {
	value *AccountPreferencesV4
	isSet bool
}

func (v NullableAccountPreferencesV4) Get() *AccountPreferencesV4 {
	return v.value
}

func (v *NullableAccountPreferencesV4) Set(val *AccountPreferencesV4) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountPreferencesV4) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountPreferencesV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountPreferencesV4(val *AccountPreferencesV4) *NullableAccountPreferencesV4 {
	return &NullableAccountPreferencesV4{value: val, isSet: true}
}

func (v NullableAccountPreferencesV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountPreferencesV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


