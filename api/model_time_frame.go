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

// TimeFrame struct for TimeFrame
type TimeFrame struct {
	BeginTime *string `json:"beginTime,omitempty"`
	EndTime *string `json:"endTime,omitempty"`
}

// NewTimeFrame instantiates a new TimeFrame object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeFrame() *TimeFrame {
	this := TimeFrame{}
	return &this
}

// NewTimeFrameWithDefaults instantiates a new TimeFrame object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeFrameWithDefaults() *TimeFrame {
	this := TimeFrame{}
	return &this
}

// GetBeginTime returns the BeginTime field value if set, zero value otherwise.
func (o *TimeFrame) GetBeginTime() string {
	if o == nil || o.BeginTime == nil {
		var ret string
		return ret
	}
	return *o.BeginTime
}

// GetBeginTimeOk returns a tuple with the BeginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeFrame) GetBeginTimeOk() (*string, bool) {
	if o == nil || o.BeginTime == nil {
		return nil, false
	}
	return o.BeginTime, true
}

// HasBeginTime returns a boolean if a field has been set.
func (o *TimeFrame) HasBeginTime() bool {
	if o != nil && o.BeginTime != nil {
		return true
	}

	return false
}

// SetBeginTime gets a reference to the given string and assigns it to the BeginTime field.
func (o *TimeFrame) SetBeginTime(v string) {
	o.BeginTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TimeFrame) GetEndTime() string {
	if o == nil || o.EndTime == nil {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeFrame) GetEndTimeOk() (*string, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TimeFrame) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *TimeFrame) SetEndTime(v string) {
	o.EndTime = &v
}

func (o TimeFrame) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BeginTime != nil {
		toSerialize["beginTime"] = o.BeginTime
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}
	return json.Marshal(toSerialize)
}

type NullableTimeFrame struct {
	value *TimeFrame
	isSet bool
}

func (v NullableTimeFrame) Get() *TimeFrame {
	return v.value
}

func (v *NullableTimeFrame) Set(val *TimeFrame) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeFrame) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeFrame) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeFrame(val *TimeFrame) *NullableTimeFrame {
	return &NullableTimeFrame{value: val, isSet: true}
}

func (v NullableTimeFrame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeFrame) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


