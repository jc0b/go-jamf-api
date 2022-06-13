/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// PersonalHotspotSetting the model 'PersonalHotspotSetting'
type PersonalHotspotSetting string

// List of PersonalHotspotSetting
const (
	PERSONALHOTSPOTSETTING_ENABLE_PERSONAL_HOTSPOT PersonalHotspotSetting = "ENABLE_PERSONAL_HOTSPOT"
	PERSONALHOTSPOTSETTING_DISABLE_PERSONAL_HOTSPOT PersonalHotspotSetting = "DISABLE_PERSONAL_HOTSPOT"
)

// All allowed values of PersonalHotspotSetting enum
var AllowedPersonalHotspotSettingEnumValues = []PersonalHotspotSetting{
	"ENABLE_PERSONAL_HOTSPOT",
	"DISABLE_PERSONAL_HOTSPOT",
}

func (v *PersonalHotspotSetting) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PersonalHotspotSetting(value)
	for _, existing := range AllowedPersonalHotspotSettingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PersonalHotspotSetting", value)
}

// NewPersonalHotspotSettingFromValue returns a pointer to a valid PersonalHotspotSetting
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPersonalHotspotSettingFromValue(v string) (*PersonalHotspotSetting, error) {
	ev := PersonalHotspotSetting(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PersonalHotspotSetting: valid values are %v", v, AllowedPersonalHotspotSettingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PersonalHotspotSetting) IsValid() bool {
	for _, existing := range AllowedPersonalHotspotSettingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PersonalHotspotSetting value
func (v PersonalHotspotSetting) Ptr() *PersonalHotspotSetting {
	return &v
}

type NullablePersonalHotspotSetting struct {
	value *PersonalHotspotSetting
	isSet bool
}

func (v NullablePersonalHotspotSetting) Get() *PersonalHotspotSetting {
	return v.value
}

func (v *NullablePersonalHotspotSetting) Set(val *PersonalHotspotSetting) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonalHotspotSetting) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonalHotspotSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonalHotspotSetting(val *PersonalHotspotSetting) *NullablePersonalHotspotSetting {
	return &NullablePersonalHotspotSetting{value: val, isSet: true}
}

func (v NullablePersonalHotspotSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonalHotspotSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
