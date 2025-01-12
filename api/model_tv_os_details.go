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

// TvOsDetails will be populated if the type is appleTv.
type TvOsDetails struct {
	Model *string `json:"model,omitempty"`
	ModelIdentifier *string `json:"modelIdentifier,omitempty"`
	ModelNumber *string `json:"modelNumber,omitempty"`
	Supervised *bool `json:"supervised,omitempty"`
	AirplayPassword *string `json:"airplayPassword,omitempty"`
	DeviceId *string `json:"deviceId,omitempty"`
	Locales *string `json:"locales,omitempty"`
	Purchasing *PurchasingV2 `json:"purchasing,omitempty"`
	ConfigurationProfiles []ConfigurationProfile `json:"configurationProfiles,omitempty"`
}

// NewTvOsDetails instantiates a new TvOsDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTvOsDetails() *TvOsDetails {
	this := TvOsDetails{}
	return &this
}

// NewTvOsDetailsWithDefaults instantiates a new TvOsDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTvOsDetailsWithDefaults() *TvOsDetails {
	this := TvOsDetails{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *TvOsDetails) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvOsDetails) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *TvOsDetails) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *TvOsDetails) SetModel(v string) {
	o.Model = &v
}

// GetModelIdentifier returns the ModelIdentifier field value if set, zero value otherwise.
func (o *TvOsDetails) GetModelIdentifier() string {
	if o == nil || o.ModelIdentifier == nil {
		var ret string
		return ret
	}
	return *o.ModelIdentifier
}

// GetModelIdentifierOk returns a tuple with the ModelIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvOsDetails) GetModelIdentifierOk() (*string, bool) {
	if o == nil || o.ModelIdentifier == nil {
		return nil, false
	}
	return o.ModelIdentifier, true
}

// HasModelIdentifier returns a boolean if a field has been set.
func (o *TvOsDetails) HasModelIdentifier() bool {
	if o != nil && o.ModelIdentifier != nil {
		return true
	}

	return false
}

// SetModelIdentifier gets a reference to the given string and assigns it to the ModelIdentifier field.
func (o *TvOsDetails) SetModelIdentifier(v string) {
	o.ModelIdentifier = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *TvOsDetails) GetModelNumber() string {
	if o == nil || o.ModelNumber == nil {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvOsDetails) GetModelNumberOk() (*string, bool) {
	if o == nil || o.ModelNumber == nil {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *TvOsDetails) HasModelNumber() bool {
	if o != nil && o.ModelNumber != nil {
		return true
	}

	return false
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *TvOsDetails) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetSupervised returns the Supervised field value if set, zero value otherwise.
func (o *TvOsDetails) GetSupervised() bool {
	if o == nil || o.Supervised == nil {
		var ret bool
		return ret
	}
	return *o.Supervised
}

// GetSupervisedOk returns a tuple with the Supervised field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvOsDetails) GetSupervisedOk() (*bool, bool) {
	if o == nil || o.Supervised == nil {
		return nil, false
	}
	return o.Supervised, true
}

// HasSupervised returns a boolean if a field has been set.
func (o *TvOsDetails) HasSupervised() bool {
	if o != nil && o.Supervised != nil {
		return true
	}

	return false
}

// SetSupervised gets a reference to the given bool and assigns it to the Supervised field.
func (o *TvOsDetails) SetSupervised(v bool) {
	o.Supervised = &v
}

// GetAirplayPassword returns the AirplayPassword field value if set, zero value otherwise.
func (o *TvOsDetails) GetAirplayPassword() string {
	if o == nil || o.AirplayPassword == nil {
		var ret string
		return ret
	}
	return *o.AirplayPassword
}

// GetAirplayPasswordOk returns a tuple with the AirplayPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvOsDetails) GetAirplayPasswordOk() (*string, bool) {
	if o == nil || o.AirplayPassword == nil {
		return nil, false
	}
	return o.AirplayPassword, true
}

// HasAirplayPassword returns a boolean if a field has been set.
func (o *TvOsDetails) HasAirplayPassword() bool {
	if o != nil && o.AirplayPassword != nil {
		return true
	}

	return false
}

// SetAirplayPassword gets a reference to the given string and assigns it to the AirplayPassword field.
func (o *TvOsDetails) SetAirplayPassword(v string) {
	o.AirplayPassword = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *TvOsDetails) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvOsDetails) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *TvOsDetails) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *TvOsDetails) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetLocales returns the Locales field value if set, zero value otherwise.
func (o *TvOsDetails) GetLocales() string {
	if o == nil || o.Locales == nil {
		var ret string
		return ret
	}
	return *o.Locales
}

// GetLocalesOk returns a tuple with the Locales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvOsDetails) GetLocalesOk() (*string, bool) {
	if o == nil || o.Locales == nil {
		return nil, false
	}
	return o.Locales, true
}

// HasLocales returns a boolean if a field has been set.
func (o *TvOsDetails) HasLocales() bool {
	if o != nil && o.Locales != nil {
		return true
	}

	return false
}

// SetLocales gets a reference to the given string and assigns it to the Locales field.
func (o *TvOsDetails) SetLocales(v string) {
	o.Locales = &v
}

// GetPurchasing returns the Purchasing field value if set, zero value otherwise.
func (o *TvOsDetails) GetPurchasing() PurchasingV2 {
	if o == nil || o.Purchasing == nil {
		var ret PurchasingV2
		return ret
	}
	return *o.Purchasing
}

// GetPurchasingOk returns a tuple with the Purchasing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvOsDetails) GetPurchasingOk() (*PurchasingV2, bool) {
	if o == nil || o.Purchasing == nil {
		return nil, false
	}
	return o.Purchasing, true
}

// HasPurchasing returns a boolean if a field has been set.
func (o *TvOsDetails) HasPurchasing() bool {
	if o != nil && o.Purchasing != nil {
		return true
	}

	return false
}

// SetPurchasing gets a reference to the given PurchasingV2 and assigns it to the Purchasing field.
func (o *TvOsDetails) SetPurchasing(v PurchasingV2) {
	o.Purchasing = &v
}

// GetConfigurationProfiles returns the ConfigurationProfiles field value if set, zero value otherwise.
func (o *TvOsDetails) GetConfigurationProfiles() []ConfigurationProfile {
	if o == nil || o.ConfigurationProfiles == nil {
		var ret []ConfigurationProfile
		return ret
	}
	return o.ConfigurationProfiles
}

// GetConfigurationProfilesOk returns a tuple with the ConfigurationProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvOsDetails) GetConfigurationProfilesOk() ([]ConfigurationProfile, bool) {
	if o == nil || o.ConfigurationProfiles == nil {
		return nil, false
	}
	return o.ConfigurationProfiles, true
}

// HasConfigurationProfiles returns a boolean if a field has been set.
func (o *TvOsDetails) HasConfigurationProfiles() bool {
	if o != nil && o.ConfigurationProfiles != nil {
		return true
	}

	return false
}

// SetConfigurationProfiles gets a reference to the given []ConfigurationProfile and assigns it to the ConfigurationProfiles field.
func (o *TvOsDetails) SetConfigurationProfiles(v []ConfigurationProfile) {
	o.ConfigurationProfiles = v
}

func (o TvOsDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.ModelIdentifier != nil {
		toSerialize["modelIdentifier"] = o.ModelIdentifier
	}
	if o.ModelNumber != nil {
		toSerialize["modelNumber"] = o.ModelNumber
	}
	if o.Supervised != nil {
		toSerialize["supervised"] = o.Supervised
	}
	if o.AirplayPassword != nil {
		toSerialize["airplayPassword"] = o.AirplayPassword
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.Locales != nil {
		toSerialize["locales"] = o.Locales
	}
	if o.Purchasing != nil {
		toSerialize["purchasing"] = o.Purchasing
	}
	if o.ConfigurationProfiles != nil {
		toSerialize["configurationProfiles"] = o.ConfigurationProfiles
	}
	return json.Marshal(toSerialize)
}

type NullableTvOsDetails struct {
	value *TvOsDetails
	isSet bool
}

func (v NullableTvOsDetails) Get() *TvOsDetails {
	return v.value
}

func (v *NullableTvOsDetails) Set(val *TvOsDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTvOsDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTvOsDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTvOsDetails(val *TvOsDetails) *NullableTvOsDetails {
	return &NullableTvOsDetails{value: val, isSet: true}
}

func (v NullableTvOsDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTvOsDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


