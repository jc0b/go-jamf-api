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

// ComputerLocation struct for ComputerLocation
type ComputerLocation struct {
	Username *string `json:"username,omitempty"`
	Position *string `json:"position,omitempty"`
	Room *string `json:"room,omitempty"`
}

// NewComputerLocation instantiates a new ComputerLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerLocation() *ComputerLocation {
	this := ComputerLocation{}
	return &this
}

// NewComputerLocationWithDefaults instantiates a new ComputerLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerLocationWithDefaults() *ComputerLocation {
	this := ComputerLocation{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ComputerLocation) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocation) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ComputerLocation) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ComputerLocation) SetUsername(v string) {
	o.Username = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ComputerLocation) GetPosition() string {
	if o == nil || o.Position == nil {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocation) GetPositionOk() (*string, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ComputerLocation) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *ComputerLocation) SetPosition(v string) {
	o.Position = &v
}

// GetRoom returns the Room field value if set, zero value otherwise.
func (o *ComputerLocation) GetRoom() string {
	if o == nil || o.Room == nil {
		var ret string
		return ret
	}
	return *o.Room
}

// GetRoomOk returns a tuple with the Room field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocation) GetRoomOk() (*string, bool) {
	if o == nil || o.Room == nil {
		return nil, false
	}
	return o.Room, true
}

// HasRoom returns a boolean if a field has been set.
func (o *ComputerLocation) HasRoom() bool {
	if o != nil && o.Room != nil {
		return true
	}

	return false
}

// SetRoom gets a reference to the given string and assigns it to the Room field.
func (o *ComputerLocation) SetRoom(v string) {
	o.Room = &v
}

func (o ComputerLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Room != nil {
		toSerialize["room"] = o.Room
	}
	return json.Marshal(toSerialize)
}

type NullableComputerLocation struct {
	value *ComputerLocation
	isSet bool
}

func (v NullableComputerLocation) Get() *ComputerLocation {
	return v.value
}

func (v *NullableComputerLocation) Set(val *ComputerLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerLocation(val *ComputerLocation) *NullableComputerLocation {
	return &NullableComputerLocation{value: val, isSet: true}
}

func (v NullableComputerLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


