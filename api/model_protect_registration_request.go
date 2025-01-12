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

// ProtectRegistrationRequest Create an API Client in the Jamf Protect web console to obtain these values.
type ProtectRegistrationRequest struct {
	ProtectUrl string `json:"protectUrl"`
	ClientId string `json:"clientId"`
	Password string `json:"password"`
}

// NewProtectRegistrationRequest instantiates a new ProtectRegistrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectRegistrationRequest(protectUrl string, clientId string, password string) *ProtectRegistrationRequest {
	this := ProtectRegistrationRequest{}
	this.ProtectUrl = protectUrl
	this.ClientId = clientId
	this.Password = password
	return &this
}

// NewProtectRegistrationRequestWithDefaults instantiates a new ProtectRegistrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectRegistrationRequestWithDefaults() *ProtectRegistrationRequest {
	this := ProtectRegistrationRequest{}
	return &this
}

// GetProtectUrl returns the ProtectUrl field value
func (o *ProtectRegistrationRequest) GetProtectUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProtectUrl
}

// GetProtectUrlOk returns a tuple with the ProtectUrl field value
// and a boolean to check if the value has been set.
func (o *ProtectRegistrationRequest) GetProtectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtectUrl, true
}

// SetProtectUrl sets field value
func (o *ProtectRegistrationRequest) SetProtectUrl(v string) {
	o.ProtectUrl = v
}

// GetClientId returns the ClientId field value
func (o *ProtectRegistrationRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ProtectRegistrationRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ProtectRegistrationRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetPassword returns the Password field value
func (o *ProtectRegistrationRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ProtectRegistrationRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ProtectRegistrationRequest) SetPassword(v string) {
	o.Password = v
}

func (o ProtectRegistrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["protectUrl"] = o.ProtectUrl
	}
	if true {
		toSerialize["clientId"] = o.ClientId
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableProtectRegistrationRequest struct {
	value *ProtectRegistrationRequest
	isSet bool
}

func (v NullableProtectRegistrationRequest) Get() *ProtectRegistrationRequest {
	return v.value
}

func (v *NullableProtectRegistrationRequest) Set(val *ProtectRegistrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectRegistrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectRegistrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectRegistrationRequest(val *ProtectRegistrationRequest) *NullableProtectRegistrationRequest {
	return &NullableProtectRegistrationRequest{value: val, isSet: true}
}

func (v NullableProtectRegistrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectRegistrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


