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

// MobileDevicePrestageNamesV2 struct for MobileDevicePrestageNamesV2
type MobileDevicePrestageNamesV2 struct {
	AssignNamesUsing *string `json:"assignNamesUsing,omitempty"`
	PrestageDeviceNames []MobileDevicePrestageNameV2 `json:"prestageDeviceNames,omitempty"`
	DeviceNamePrefix *string `json:"deviceNamePrefix,omitempty"`
	DeviceNameSuffix *string `json:"deviceNameSuffix,omitempty"`
	SingleDeviceName *string `json:"singleDeviceName,omitempty"`
	ManageNames *bool `json:"manageNames,omitempty"`
	DeviceNamingConfigured *bool `json:"deviceNamingConfigured,omitempty"`
}

// NewMobileDevicePrestageNamesV2 instantiates a new MobileDevicePrestageNamesV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDevicePrestageNamesV2() *MobileDevicePrestageNamesV2 {
	this := MobileDevicePrestageNamesV2{}
	return &this
}

// NewMobileDevicePrestageNamesV2WithDefaults instantiates a new MobileDevicePrestageNamesV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDevicePrestageNamesV2WithDefaults() *MobileDevicePrestageNamesV2 {
	this := MobileDevicePrestageNamesV2{}
	return &this
}

// GetAssignNamesUsing returns the AssignNamesUsing field value if set, zero value otherwise.
func (o *MobileDevicePrestageNamesV2) GetAssignNamesUsing() string {
	if o == nil || o.AssignNamesUsing == nil {
		var ret string
		return ret
	}
	return *o.AssignNamesUsing
}

// GetAssignNamesUsingOk returns a tuple with the AssignNamesUsing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNamesV2) GetAssignNamesUsingOk() (*string, bool) {
	if o == nil || o.AssignNamesUsing == nil {
		return nil, false
	}
	return o.AssignNamesUsing, true
}

// HasAssignNamesUsing returns a boolean if a field has been set.
func (o *MobileDevicePrestageNamesV2) HasAssignNamesUsing() bool {
	if o != nil && o.AssignNamesUsing != nil {
		return true
	}

	return false
}

// SetAssignNamesUsing gets a reference to the given string and assigns it to the AssignNamesUsing field.
func (o *MobileDevicePrestageNamesV2) SetAssignNamesUsing(v string) {
	o.AssignNamesUsing = &v
}

// GetPrestageDeviceNames returns the PrestageDeviceNames field value if set, zero value otherwise.
func (o *MobileDevicePrestageNamesV2) GetPrestageDeviceNames() []MobileDevicePrestageNameV2 {
	if o == nil || o.PrestageDeviceNames == nil {
		var ret []MobileDevicePrestageNameV2
		return ret
	}
	return o.PrestageDeviceNames
}

// GetPrestageDeviceNamesOk returns a tuple with the PrestageDeviceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNamesV2) GetPrestageDeviceNamesOk() ([]MobileDevicePrestageNameV2, bool) {
	if o == nil || o.PrestageDeviceNames == nil {
		return nil, false
	}
	return o.PrestageDeviceNames, true
}

// HasPrestageDeviceNames returns a boolean if a field has been set.
func (o *MobileDevicePrestageNamesV2) HasPrestageDeviceNames() bool {
	if o != nil && o.PrestageDeviceNames != nil {
		return true
	}

	return false
}

// SetPrestageDeviceNames gets a reference to the given []MobileDevicePrestageNameV2 and assigns it to the PrestageDeviceNames field.
func (o *MobileDevicePrestageNamesV2) SetPrestageDeviceNames(v []MobileDevicePrestageNameV2) {
	o.PrestageDeviceNames = v
}

// GetDeviceNamePrefix returns the DeviceNamePrefix field value if set, zero value otherwise.
func (o *MobileDevicePrestageNamesV2) GetDeviceNamePrefix() string {
	if o == nil || o.DeviceNamePrefix == nil {
		var ret string
		return ret
	}
	return *o.DeviceNamePrefix
}

// GetDeviceNamePrefixOk returns a tuple with the DeviceNamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNamesV2) GetDeviceNamePrefixOk() (*string, bool) {
	if o == nil || o.DeviceNamePrefix == nil {
		return nil, false
	}
	return o.DeviceNamePrefix, true
}

// HasDeviceNamePrefix returns a boolean if a field has been set.
func (o *MobileDevicePrestageNamesV2) HasDeviceNamePrefix() bool {
	if o != nil && o.DeviceNamePrefix != nil {
		return true
	}

	return false
}

// SetDeviceNamePrefix gets a reference to the given string and assigns it to the DeviceNamePrefix field.
func (o *MobileDevicePrestageNamesV2) SetDeviceNamePrefix(v string) {
	o.DeviceNamePrefix = &v
}

// GetDeviceNameSuffix returns the DeviceNameSuffix field value if set, zero value otherwise.
func (o *MobileDevicePrestageNamesV2) GetDeviceNameSuffix() string {
	if o == nil || o.DeviceNameSuffix == nil {
		var ret string
		return ret
	}
	return *o.DeviceNameSuffix
}

// GetDeviceNameSuffixOk returns a tuple with the DeviceNameSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNamesV2) GetDeviceNameSuffixOk() (*string, bool) {
	if o == nil || o.DeviceNameSuffix == nil {
		return nil, false
	}
	return o.DeviceNameSuffix, true
}

// HasDeviceNameSuffix returns a boolean if a field has been set.
func (o *MobileDevicePrestageNamesV2) HasDeviceNameSuffix() bool {
	if o != nil && o.DeviceNameSuffix != nil {
		return true
	}

	return false
}

// SetDeviceNameSuffix gets a reference to the given string and assigns it to the DeviceNameSuffix field.
func (o *MobileDevicePrestageNamesV2) SetDeviceNameSuffix(v string) {
	o.DeviceNameSuffix = &v
}

// GetSingleDeviceName returns the SingleDeviceName field value if set, zero value otherwise.
func (o *MobileDevicePrestageNamesV2) GetSingleDeviceName() string {
	if o == nil || o.SingleDeviceName == nil {
		var ret string
		return ret
	}
	return *o.SingleDeviceName
}

// GetSingleDeviceNameOk returns a tuple with the SingleDeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNamesV2) GetSingleDeviceNameOk() (*string, bool) {
	if o == nil || o.SingleDeviceName == nil {
		return nil, false
	}
	return o.SingleDeviceName, true
}

// HasSingleDeviceName returns a boolean if a field has been set.
func (o *MobileDevicePrestageNamesV2) HasSingleDeviceName() bool {
	if o != nil && o.SingleDeviceName != nil {
		return true
	}

	return false
}

// SetSingleDeviceName gets a reference to the given string and assigns it to the SingleDeviceName field.
func (o *MobileDevicePrestageNamesV2) SetSingleDeviceName(v string) {
	o.SingleDeviceName = &v
}

// GetManageNames returns the ManageNames field value if set, zero value otherwise.
func (o *MobileDevicePrestageNamesV2) GetManageNames() bool {
	if o == nil || o.ManageNames == nil {
		var ret bool
		return ret
	}
	return *o.ManageNames
}

// GetManageNamesOk returns a tuple with the ManageNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNamesV2) GetManageNamesOk() (*bool, bool) {
	if o == nil || o.ManageNames == nil {
		return nil, false
	}
	return o.ManageNames, true
}

// HasManageNames returns a boolean if a field has been set.
func (o *MobileDevicePrestageNamesV2) HasManageNames() bool {
	if o != nil && o.ManageNames != nil {
		return true
	}

	return false
}

// SetManageNames gets a reference to the given bool and assigns it to the ManageNames field.
func (o *MobileDevicePrestageNamesV2) SetManageNames(v bool) {
	o.ManageNames = &v
}

// GetDeviceNamingConfigured returns the DeviceNamingConfigured field value if set, zero value otherwise.
func (o *MobileDevicePrestageNamesV2) GetDeviceNamingConfigured() bool {
	if o == nil || o.DeviceNamingConfigured == nil {
		var ret bool
		return ret
	}
	return *o.DeviceNamingConfigured
}

// GetDeviceNamingConfiguredOk returns a tuple with the DeviceNamingConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNamesV2) GetDeviceNamingConfiguredOk() (*bool, bool) {
	if o == nil || o.DeviceNamingConfigured == nil {
		return nil, false
	}
	return o.DeviceNamingConfigured, true
}

// HasDeviceNamingConfigured returns a boolean if a field has been set.
func (o *MobileDevicePrestageNamesV2) HasDeviceNamingConfigured() bool {
	if o != nil && o.DeviceNamingConfigured != nil {
		return true
	}

	return false
}

// SetDeviceNamingConfigured gets a reference to the given bool and assigns it to the DeviceNamingConfigured field.
func (o *MobileDevicePrestageNamesV2) SetDeviceNamingConfigured(v bool) {
	o.DeviceNamingConfigured = &v
}

func (o MobileDevicePrestageNamesV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignNamesUsing != nil {
		toSerialize["assignNamesUsing"] = o.AssignNamesUsing
	}
	if o.PrestageDeviceNames != nil {
		toSerialize["prestageDeviceNames"] = o.PrestageDeviceNames
	}
	if o.DeviceNamePrefix != nil {
		toSerialize["deviceNamePrefix"] = o.DeviceNamePrefix
	}
	if o.DeviceNameSuffix != nil {
		toSerialize["deviceNameSuffix"] = o.DeviceNameSuffix
	}
	if o.SingleDeviceName != nil {
		toSerialize["singleDeviceName"] = o.SingleDeviceName
	}
	if o.ManageNames != nil {
		toSerialize["manageNames"] = o.ManageNames
	}
	if o.DeviceNamingConfigured != nil {
		toSerialize["deviceNamingConfigured"] = o.DeviceNamingConfigured
	}
	return json.Marshal(toSerialize)
}

type NullableMobileDevicePrestageNamesV2 struct {
	value *MobileDevicePrestageNamesV2
	isSet bool
}

func (v NullableMobileDevicePrestageNamesV2) Get() *MobileDevicePrestageNamesV2 {
	return v.value
}

func (v *NullableMobileDevicePrestageNamesV2) Set(val *MobileDevicePrestageNamesV2) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDevicePrestageNamesV2) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDevicePrestageNamesV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDevicePrestageNamesV2(val *MobileDevicePrestageNamesV2) *NullableMobileDevicePrestageNamesV2 {
	return &NullableMobileDevicePrestageNamesV2{value: val, isSet: true}
}

func (v NullableMobileDevicePrestageNamesV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDevicePrestageNamesV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

