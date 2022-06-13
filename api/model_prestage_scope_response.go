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

// PrestageScopeResponse struct for PrestageScopeResponse
type PrestageScopeResponse struct {
	PrestageId *int32 `json:"prestageId,omitempty"`
	Assignments []PrestageScopeAssignment `json:"assignments,omitempty"`
	VersionLock *int32 `json:"versionLock,omitempty"`
}

// NewPrestageScopeResponse instantiates a new PrestageScopeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrestageScopeResponse() *PrestageScopeResponse {
	this := PrestageScopeResponse{}
	return &this
}

// NewPrestageScopeResponseWithDefaults instantiates a new PrestageScopeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrestageScopeResponseWithDefaults() *PrestageScopeResponse {
	this := PrestageScopeResponse{}
	return &this
}

// GetPrestageId returns the PrestageId field value if set, zero value otherwise.
func (o *PrestageScopeResponse) GetPrestageId() int32 {
	if o == nil || o.PrestageId == nil {
		var ret int32
		return ret
	}
	return *o.PrestageId
}

// GetPrestageIdOk returns a tuple with the PrestageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageScopeResponse) GetPrestageIdOk() (*int32, bool) {
	if o == nil || o.PrestageId == nil {
		return nil, false
	}
	return o.PrestageId, true
}

// HasPrestageId returns a boolean if a field has been set.
func (o *PrestageScopeResponse) HasPrestageId() bool {
	if o != nil && o.PrestageId != nil {
		return true
	}

	return false
}

// SetPrestageId gets a reference to the given int32 and assigns it to the PrestageId field.
func (o *PrestageScopeResponse) SetPrestageId(v int32) {
	o.PrestageId = &v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *PrestageScopeResponse) GetAssignments() []PrestageScopeAssignment {
	if o == nil || o.Assignments == nil {
		var ret []PrestageScopeAssignment
		return ret
	}
	return o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageScopeResponse) GetAssignmentsOk() ([]PrestageScopeAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *PrestageScopeResponse) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []PrestageScopeAssignment and assigns it to the Assignments field.
func (o *PrestageScopeResponse) SetAssignments(v []PrestageScopeAssignment) {
	o.Assignments = v
}

// GetVersionLock returns the VersionLock field value if set, zero value otherwise.
func (o *PrestageScopeResponse) GetVersionLock() int32 {
	if o == nil || o.VersionLock == nil {
		var ret int32
		return ret
	}
	return *o.VersionLock
}

// GetVersionLockOk returns a tuple with the VersionLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageScopeResponse) GetVersionLockOk() (*int32, bool) {
	if o == nil || o.VersionLock == nil {
		return nil, false
	}
	return o.VersionLock, true
}

// HasVersionLock returns a boolean if a field has been set.
func (o *PrestageScopeResponse) HasVersionLock() bool {
	if o != nil && o.VersionLock != nil {
		return true
	}

	return false
}

// SetVersionLock gets a reference to the given int32 and assigns it to the VersionLock field.
func (o *PrestageScopeResponse) SetVersionLock(v int32) {
	o.VersionLock = &v
}

func (o PrestageScopeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrestageId != nil {
		toSerialize["prestageId"] = o.PrestageId
	}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	if o.VersionLock != nil {
		toSerialize["versionLock"] = o.VersionLock
	}
	return json.Marshal(toSerialize)
}

type NullablePrestageScopeResponse struct {
	value *PrestageScopeResponse
	isSet bool
}

func (v NullablePrestageScopeResponse) Get() *PrestageScopeResponse {
	return v.value
}

func (v *NullablePrestageScopeResponse) Set(val *PrestageScopeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrestageScopeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrestageScopeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrestageScopeResponse(val *PrestageScopeResponse) *NullablePrestageScopeResponse {
	return &NullablePrestageScopeResponse{value: val, isSet: true}
}

func (v NullablePrestageScopeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrestageScopeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

