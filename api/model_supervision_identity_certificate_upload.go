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

// SupervisionIdentityCertificateUpload struct for SupervisionIdentityCertificateUpload
type SupervisionIdentityCertificateUpload struct {
	DisplayName string `json:"displayName"`
	Password string `json:"password"`
	// The base 64 encoded supervision identity certificate data
	CertificateData *string `json:"certificateData,omitempty"`
}

// NewSupervisionIdentityCertificateUpload instantiates a new SupervisionIdentityCertificateUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupervisionIdentityCertificateUpload(displayName string, password string) *SupervisionIdentityCertificateUpload {
	this := SupervisionIdentityCertificateUpload{}
	this.DisplayName = displayName
	this.Password = password
	return &this
}

// NewSupervisionIdentityCertificateUploadWithDefaults instantiates a new SupervisionIdentityCertificateUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupervisionIdentityCertificateUploadWithDefaults() *SupervisionIdentityCertificateUpload {
	this := SupervisionIdentityCertificateUpload{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *SupervisionIdentityCertificateUpload) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *SupervisionIdentityCertificateUpload) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *SupervisionIdentityCertificateUpload) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetPassword returns the Password field value
func (o *SupervisionIdentityCertificateUpload) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SupervisionIdentityCertificateUpload) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SupervisionIdentityCertificateUpload) SetPassword(v string) {
	o.Password = v
}

// GetCertificateData returns the CertificateData field value if set, zero value otherwise.
func (o *SupervisionIdentityCertificateUpload) GetCertificateData() string {
	if o == nil || o.CertificateData == nil {
		var ret string
		return ret
	}
	return *o.CertificateData
}

// GetCertificateDataOk returns a tuple with the CertificateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupervisionIdentityCertificateUpload) GetCertificateDataOk() (*string, bool) {
	if o == nil || o.CertificateData == nil {
		return nil, false
	}
	return o.CertificateData, true
}

// HasCertificateData returns a boolean if a field has been set.
func (o *SupervisionIdentityCertificateUpload) HasCertificateData() bool {
	if o != nil && o.CertificateData != nil {
		return true
	}

	return false
}

// SetCertificateData gets a reference to the given string and assigns it to the CertificateData field.
func (o *SupervisionIdentityCertificateUpload) SetCertificateData(v string) {
	o.CertificateData = &v
}

func (o SupervisionIdentityCertificateUpload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.CertificateData != nil {
		toSerialize["certificateData"] = o.CertificateData
	}
	return json.Marshal(toSerialize)
}

type NullableSupervisionIdentityCertificateUpload struct {
	value *SupervisionIdentityCertificateUpload
	isSet bool
}

func (v NullableSupervisionIdentityCertificateUpload) Get() *SupervisionIdentityCertificateUpload {
	return v.value
}

func (v *NullableSupervisionIdentityCertificateUpload) Set(val *SupervisionIdentityCertificateUpload) {
	v.value = val
	v.isSet = true
}

func (v NullableSupervisionIdentityCertificateUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableSupervisionIdentityCertificateUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupervisionIdentityCertificateUpload(val *SupervisionIdentityCertificateUpload) *NullableSupervisionIdentityCertificateUpload {
	return &NullableSupervisionIdentityCertificateUpload{value: val, isSet: true}
}

func (v NullableSupervisionIdentityCertificateUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupervisionIdentityCertificateUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


