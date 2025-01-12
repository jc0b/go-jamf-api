/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// AndroidDetails will be populated if the type is android.
type AndroidDetails struct {
	OsName *string `json:"osName,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty"`
	Model *string `json:"model,omitempty"`
	InternalCapacityMb *int32 `json:"internalCapacityMb,omitempty"`
	InternalAvailableMb *int32 `json:"internalAvailableMb,omitempty"`
	InternalPercentUsed *int32 `json:"internalPercentUsed,omitempty"`
	ExternalCapacityMb *int32 `json:"externalCapacityMb,omitempty"`
	ExternalAvailableMb *int32 `json:"externalAvailableMb,omitempty"`
	ExternalPercentUsed *int32 `json:"externalPercentUsed,omitempty"`
	BatteryLevel *int32 `json:"batteryLevel,omitempty"`
	LastBackupTimestamp *time.Time `json:"lastBackupTimestamp,omitempty"`
	ApiVersion *int32 `json:"apiVersion,omitempty"`
	Computer *IdAndName `json:"computer,omitempty"`
	Security *Security `json:"security,omitempty"`
}

// NewAndroidDetails instantiates a new AndroidDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAndroidDetails() *AndroidDetails {
	this := AndroidDetails{}
	return &this
}

// NewAndroidDetailsWithDefaults instantiates a new AndroidDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAndroidDetailsWithDefaults() *AndroidDetails {
	this := AndroidDetails{}
	return &this
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *AndroidDetails) GetOsName() string {
	if o == nil || o.OsName == nil {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetOsNameOk() (*string, bool) {
	if o == nil || o.OsName == nil {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *AndroidDetails) HasOsName() bool {
	if o != nil && o.OsName != nil {
		return true
	}

	return false
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *AndroidDetails) SetOsName(v string) {
	o.OsName = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *AndroidDetails) GetManufacturer() string {
	if o == nil || o.Manufacturer == nil {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetManufacturerOk() (*string, bool) {
	if o == nil || o.Manufacturer == nil {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *AndroidDetails) HasManufacturer() bool {
	if o != nil && o.Manufacturer != nil {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *AndroidDetails) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AndroidDetails) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AndroidDetails) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AndroidDetails) SetModel(v string) {
	o.Model = &v
}

// GetInternalCapacityMb returns the InternalCapacityMb field value if set, zero value otherwise.
func (o *AndroidDetails) GetInternalCapacityMb() int32 {
	if o == nil || o.InternalCapacityMb == nil {
		var ret int32
		return ret
	}
	return *o.InternalCapacityMb
}

// GetInternalCapacityMbOk returns a tuple with the InternalCapacityMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetInternalCapacityMbOk() (*int32, bool) {
	if o == nil || o.InternalCapacityMb == nil {
		return nil, false
	}
	return o.InternalCapacityMb, true
}

// HasInternalCapacityMb returns a boolean if a field has been set.
func (o *AndroidDetails) HasInternalCapacityMb() bool {
	if o != nil && o.InternalCapacityMb != nil {
		return true
	}

	return false
}

// SetInternalCapacityMb gets a reference to the given int32 and assigns it to the InternalCapacityMb field.
func (o *AndroidDetails) SetInternalCapacityMb(v int32) {
	o.InternalCapacityMb = &v
}

// GetInternalAvailableMb returns the InternalAvailableMb field value if set, zero value otherwise.
func (o *AndroidDetails) GetInternalAvailableMb() int32 {
	if o == nil || o.InternalAvailableMb == nil {
		var ret int32
		return ret
	}
	return *o.InternalAvailableMb
}

// GetInternalAvailableMbOk returns a tuple with the InternalAvailableMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetInternalAvailableMbOk() (*int32, bool) {
	if o == nil || o.InternalAvailableMb == nil {
		return nil, false
	}
	return o.InternalAvailableMb, true
}

// HasInternalAvailableMb returns a boolean if a field has been set.
func (o *AndroidDetails) HasInternalAvailableMb() bool {
	if o != nil && o.InternalAvailableMb != nil {
		return true
	}

	return false
}

// SetInternalAvailableMb gets a reference to the given int32 and assigns it to the InternalAvailableMb field.
func (o *AndroidDetails) SetInternalAvailableMb(v int32) {
	o.InternalAvailableMb = &v
}

// GetInternalPercentUsed returns the InternalPercentUsed field value if set, zero value otherwise.
func (o *AndroidDetails) GetInternalPercentUsed() int32 {
	if o == nil || o.InternalPercentUsed == nil {
		var ret int32
		return ret
	}
	return *o.InternalPercentUsed
}

// GetInternalPercentUsedOk returns a tuple with the InternalPercentUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetInternalPercentUsedOk() (*int32, bool) {
	if o == nil || o.InternalPercentUsed == nil {
		return nil, false
	}
	return o.InternalPercentUsed, true
}

// HasInternalPercentUsed returns a boolean if a field has been set.
func (o *AndroidDetails) HasInternalPercentUsed() bool {
	if o != nil && o.InternalPercentUsed != nil {
		return true
	}

	return false
}

// SetInternalPercentUsed gets a reference to the given int32 and assigns it to the InternalPercentUsed field.
func (o *AndroidDetails) SetInternalPercentUsed(v int32) {
	o.InternalPercentUsed = &v
}

// GetExternalCapacityMb returns the ExternalCapacityMb field value if set, zero value otherwise.
func (o *AndroidDetails) GetExternalCapacityMb() int32 {
	if o == nil || o.ExternalCapacityMb == nil {
		var ret int32
		return ret
	}
	return *o.ExternalCapacityMb
}

// GetExternalCapacityMbOk returns a tuple with the ExternalCapacityMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetExternalCapacityMbOk() (*int32, bool) {
	if o == nil || o.ExternalCapacityMb == nil {
		return nil, false
	}
	return o.ExternalCapacityMb, true
}

// HasExternalCapacityMb returns a boolean if a field has been set.
func (o *AndroidDetails) HasExternalCapacityMb() bool {
	if o != nil && o.ExternalCapacityMb != nil {
		return true
	}

	return false
}

// SetExternalCapacityMb gets a reference to the given int32 and assigns it to the ExternalCapacityMb field.
func (o *AndroidDetails) SetExternalCapacityMb(v int32) {
	o.ExternalCapacityMb = &v
}

// GetExternalAvailableMb returns the ExternalAvailableMb field value if set, zero value otherwise.
func (o *AndroidDetails) GetExternalAvailableMb() int32 {
	if o == nil || o.ExternalAvailableMb == nil {
		var ret int32
		return ret
	}
	return *o.ExternalAvailableMb
}

// GetExternalAvailableMbOk returns a tuple with the ExternalAvailableMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetExternalAvailableMbOk() (*int32, bool) {
	if o == nil || o.ExternalAvailableMb == nil {
		return nil, false
	}
	return o.ExternalAvailableMb, true
}

// HasExternalAvailableMb returns a boolean if a field has been set.
func (o *AndroidDetails) HasExternalAvailableMb() bool {
	if o != nil && o.ExternalAvailableMb != nil {
		return true
	}

	return false
}

// SetExternalAvailableMb gets a reference to the given int32 and assigns it to the ExternalAvailableMb field.
func (o *AndroidDetails) SetExternalAvailableMb(v int32) {
	o.ExternalAvailableMb = &v
}

// GetExternalPercentUsed returns the ExternalPercentUsed field value if set, zero value otherwise.
func (o *AndroidDetails) GetExternalPercentUsed() int32 {
	if o == nil || o.ExternalPercentUsed == nil {
		var ret int32
		return ret
	}
	return *o.ExternalPercentUsed
}

// GetExternalPercentUsedOk returns a tuple with the ExternalPercentUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetExternalPercentUsedOk() (*int32, bool) {
	if o == nil || o.ExternalPercentUsed == nil {
		return nil, false
	}
	return o.ExternalPercentUsed, true
}

// HasExternalPercentUsed returns a boolean if a field has been set.
func (o *AndroidDetails) HasExternalPercentUsed() bool {
	if o != nil && o.ExternalPercentUsed != nil {
		return true
	}

	return false
}

// SetExternalPercentUsed gets a reference to the given int32 and assigns it to the ExternalPercentUsed field.
func (o *AndroidDetails) SetExternalPercentUsed(v int32) {
	o.ExternalPercentUsed = &v
}

// GetBatteryLevel returns the BatteryLevel field value if set, zero value otherwise.
func (o *AndroidDetails) GetBatteryLevel() int32 {
	if o == nil || o.BatteryLevel == nil {
		var ret int32
		return ret
	}
	return *o.BatteryLevel
}

// GetBatteryLevelOk returns a tuple with the BatteryLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetBatteryLevelOk() (*int32, bool) {
	if o == nil || o.BatteryLevel == nil {
		return nil, false
	}
	return o.BatteryLevel, true
}

// HasBatteryLevel returns a boolean if a field has been set.
func (o *AndroidDetails) HasBatteryLevel() bool {
	if o != nil && o.BatteryLevel != nil {
		return true
	}

	return false
}

// SetBatteryLevel gets a reference to the given int32 and assigns it to the BatteryLevel field.
func (o *AndroidDetails) SetBatteryLevel(v int32) {
	o.BatteryLevel = &v
}

// GetLastBackupTimestamp returns the LastBackupTimestamp field value if set, zero value otherwise.
func (o *AndroidDetails) GetLastBackupTimestamp() time.Time {
	if o == nil || o.LastBackupTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.LastBackupTimestamp
}

// GetLastBackupTimestampOk returns a tuple with the LastBackupTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetLastBackupTimestampOk() (*time.Time, bool) {
	if o == nil || o.LastBackupTimestamp == nil {
		return nil, false
	}
	return o.LastBackupTimestamp, true
}

// HasLastBackupTimestamp returns a boolean if a field has been set.
func (o *AndroidDetails) HasLastBackupTimestamp() bool {
	if o != nil && o.LastBackupTimestamp != nil {
		return true
	}

	return false
}

// SetLastBackupTimestamp gets a reference to the given time.Time and assigns it to the LastBackupTimestamp field.
func (o *AndroidDetails) SetLastBackupTimestamp(v time.Time) {
	o.LastBackupTimestamp = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *AndroidDetails) GetApiVersion() int32 {
	if o == nil || o.ApiVersion == nil {
		var ret int32
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetApiVersionOk() (*int32, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *AndroidDetails) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given int32 and assigns it to the ApiVersion field.
func (o *AndroidDetails) SetApiVersion(v int32) {
	o.ApiVersion = &v
}

// GetComputer returns the Computer field value if set, zero value otherwise.
func (o *AndroidDetails) GetComputer() IdAndName {
	if o == nil || o.Computer == nil {
		var ret IdAndName
		return ret
	}
	return *o.Computer
}

// GetComputerOk returns a tuple with the Computer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetComputerOk() (*IdAndName, bool) {
	if o == nil || o.Computer == nil {
		return nil, false
	}
	return o.Computer, true
}

// HasComputer returns a boolean if a field has been set.
func (o *AndroidDetails) HasComputer() bool {
	if o != nil && o.Computer != nil {
		return true
	}

	return false
}

// SetComputer gets a reference to the given IdAndName and assigns it to the Computer field.
func (o *AndroidDetails) SetComputer(v IdAndName) {
	o.Computer = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *AndroidDetails) GetSecurity() Security {
	if o == nil || o.Security == nil {
		var ret Security
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDetails) GetSecurityOk() (*Security, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *AndroidDetails) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given Security and assigns it to the Security field.
func (o *AndroidDetails) SetSecurity(v Security) {
	o.Security = &v
}

func (o AndroidDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OsName != nil {
		toSerialize["osName"] = o.OsName
	}
	if o.Manufacturer != nil {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.InternalCapacityMb != nil {
		toSerialize["internalCapacityMb"] = o.InternalCapacityMb
	}
	if o.InternalAvailableMb != nil {
		toSerialize["internalAvailableMb"] = o.InternalAvailableMb
	}
	if o.InternalPercentUsed != nil {
		toSerialize["internalPercentUsed"] = o.InternalPercentUsed
	}
	if o.ExternalCapacityMb != nil {
		toSerialize["externalCapacityMb"] = o.ExternalCapacityMb
	}
	if o.ExternalAvailableMb != nil {
		toSerialize["externalAvailableMb"] = o.ExternalAvailableMb
	}
	if o.ExternalPercentUsed != nil {
		toSerialize["externalPercentUsed"] = o.ExternalPercentUsed
	}
	if o.BatteryLevel != nil {
		toSerialize["batteryLevel"] = o.BatteryLevel
	}
	if o.LastBackupTimestamp != nil {
		toSerialize["lastBackupTimestamp"] = o.LastBackupTimestamp
	}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.Computer != nil {
		toSerialize["computer"] = o.Computer
	}
	if o.Security != nil {
		toSerialize["security"] = o.Security
	}
	return json.Marshal(toSerialize)
}

type NullableAndroidDetails struct {
	value *AndroidDetails
	isSet bool
}

func (v NullableAndroidDetails) Get() *AndroidDetails {
	return v.value
}

func (v *NullableAndroidDetails) Set(val *AndroidDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAndroidDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAndroidDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAndroidDetails(val *AndroidDetails) *NullableAndroidDetails {
	return &NullableAndroidDetails{value: val, isSet: true}
}

func (v NullableAndroidDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAndroidDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


