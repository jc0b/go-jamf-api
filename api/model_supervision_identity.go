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

// SupervisionIdentity struct for SupervisionIdentity
type SupervisionIdentity struct {
	Id *int32 `json:"id,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	CommonName *string `json:"commonName,omitempty"`
	ExpirationDate *string `json:"expirationDate,omitempty"`
}

// NewSupervisionIdentity instantiates a new SupervisionIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupervisionIdentity() *SupervisionIdentity {
	this := SupervisionIdentity{}
	return &this
}

// NewSupervisionIdentityWithDefaults instantiates a new SupervisionIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupervisionIdentityWithDefaults() *SupervisionIdentity {
	this := SupervisionIdentity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SupervisionIdentity) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupervisionIdentity) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupervisionIdentity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SupervisionIdentity) SetId(v int32) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SupervisionIdentity) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupervisionIdentity) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SupervisionIdentity) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SupervisionIdentity) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *SupervisionIdentity) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupervisionIdentity) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *SupervisionIdentity) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *SupervisionIdentity) SetCommonName(v string) {
	o.CommonName = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *SupervisionIdentity) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupervisionIdentity) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *SupervisionIdentity) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *SupervisionIdentity) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

func (o SupervisionIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.CommonName != nil {
		toSerialize["commonName"] = o.CommonName
	}
	if o.ExpirationDate != nil {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	return json.Marshal(toSerialize)
}

type NullableSupervisionIdentity struct {
	value *SupervisionIdentity
	isSet bool
}

func (v NullableSupervisionIdentity) Get() *SupervisionIdentity {
	return v.value
}

func (v *NullableSupervisionIdentity) Set(val *SupervisionIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableSupervisionIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableSupervisionIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupervisionIdentity(val *SupervisionIdentity) *NullableSupervisionIdentity {
	return &NullableSupervisionIdentity{value: val, isSet: true}
}

func (v NullableSupervisionIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupervisionIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

