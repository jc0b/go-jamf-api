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

// PutComputerPrestageV2AllOf struct for PutComputerPrestageV2AllOf
type PutComputerPrestageV2AllOf struct {
	RecoveryLockPassword *string `json:"recoveryLockPassword,omitempty"`
	VersionLock *int32 `json:"versionLock,omitempty"`
}

// NewPutComputerPrestageV2AllOf instantiates a new PutComputerPrestageV2AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutComputerPrestageV2AllOf() *PutComputerPrestageV2AllOf {
	this := PutComputerPrestageV2AllOf{}
	return &this
}

// NewPutComputerPrestageV2AllOfWithDefaults instantiates a new PutComputerPrestageV2AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutComputerPrestageV2AllOfWithDefaults() *PutComputerPrestageV2AllOf {
	this := PutComputerPrestageV2AllOf{}
	return &this
}

// GetRecoveryLockPassword returns the RecoveryLockPassword field value if set, zero value otherwise.
func (o *PutComputerPrestageV2AllOf) GetRecoveryLockPassword() string {
	if o == nil || o.RecoveryLockPassword == nil {
		var ret string
		return ret
	}
	return *o.RecoveryLockPassword
}

// GetRecoveryLockPasswordOk returns a tuple with the RecoveryLockPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2AllOf) GetRecoveryLockPasswordOk() (*string, bool) {
	if o == nil || o.RecoveryLockPassword == nil {
		return nil, false
	}
	return o.RecoveryLockPassword, true
}

// HasRecoveryLockPassword returns a boolean if a field has been set.
func (o *PutComputerPrestageV2AllOf) HasRecoveryLockPassword() bool {
	if o != nil && o.RecoveryLockPassword != nil {
		return true
	}

	return false
}

// SetRecoveryLockPassword gets a reference to the given string and assigns it to the RecoveryLockPassword field.
func (o *PutComputerPrestageV2AllOf) SetRecoveryLockPassword(v string) {
	o.RecoveryLockPassword = &v
}

// GetVersionLock returns the VersionLock field value if set, zero value otherwise.
func (o *PutComputerPrestageV2AllOf) GetVersionLock() int32 {
	if o == nil || o.VersionLock == nil {
		var ret int32
		return ret
	}
	return *o.VersionLock
}

// GetVersionLockOk returns a tuple with the VersionLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2AllOf) GetVersionLockOk() (*int32, bool) {
	if o == nil || o.VersionLock == nil {
		return nil, false
	}
	return o.VersionLock, true
}

// HasVersionLock returns a boolean if a field has been set.
func (o *PutComputerPrestageV2AllOf) HasVersionLock() bool {
	if o != nil && o.VersionLock != nil {
		return true
	}

	return false
}

// SetVersionLock gets a reference to the given int32 and assigns it to the VersionLock field.
func (o *PutComputerPrestageV2AllOf) SetVersionLock(v int32) {
	o.VersionLock = &v
}

func (o PutComputerPrestageV2AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecoveryLockPassword != nil {
		toSerialize["recoveryLockPassword"] = o.RecoveryLockPassword
	}
	if o.VersionLock != nil {
		toSerialize["versionLock"] = o.VersionLock
	}
	return json.Marshal(toSerialize)
}

type NullablePutComputerPrestageV2AllOf struct {
	value *PutComputerPrestageV2AllOf
	isSet bool
}

func (v NullablePutComputerPrestageV2AllOf) Get() *PutComputerPrestageV2AllOf {
	return v.value
}

func (v *NullablePutComputerPrestageV2AllOf) Set(val *PutComputerPrestageV2AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePutComputerPrestageV2AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePutComputerPrestageV2AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutComputerPrestageV2AllOf(val *PutComputerPrestageV2AllOf) *NullablePutComputerPrestageV2AllOf {
	return &NullablePutComputerPrestageV2AllOf{value: val, isSet: true}
}

func (v NullablePutComputerPrestageV2AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutComputerPrestageV2AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


