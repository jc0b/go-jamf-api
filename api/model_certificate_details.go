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

// CertificateDetails struct for CertificateDetails
type CertificateDetails struct {
	Subject *string `json:"subject,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
}

// NewCertificateDetails instantiates a new CertificateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateDetails() *CertificateDetails {
	this := CertificateDetails{}
	var subject string = ""
	this.Subject = &subject
	var serialNumber string = ""
	this.SerialNumber = &serialNumber
	return &this
}

// NewCertificateDetailsWithDefaults instantiates a new CertificateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateDetailsWithDefaults() *CertificateDetails {
	this := CertificateDetails{}
	var subject string = ""
	this.Subject = &subject
	var serialNumber string = ""
	this.SerialNumber = &serialNumber
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CertificateDetails) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetails) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CertificateDetails) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *CertificateDetails) SetSubject(v string) {
	o.Subject = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *CertificateDetails) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetails) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *CertificateDetails) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *CertificateDetails) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

func (o CertificateDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateDetails struct {
	value *CertificateDetails
	isSet bool
}

func (v NullableCertificateDetails) Get() *CertificateDetails {
	return v.value
}

func (v *NullableCertificateDetails) Set(val *CertificateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateDetails(val *CertificateDetails) *NullableCertificateDetails {
	return &NullableCertificateDetails{value: val, isSet: true}
}

func (v NullableCertificateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


