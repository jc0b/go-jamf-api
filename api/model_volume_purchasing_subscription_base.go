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

// VolumePurchasingSubscriptionBase struct for VolumePurchasingSubscriptionBase
type VolumePurchasingSubscriptionBase struct {
	Name string `json:"name"`
	Enabled *bool `json:"enabled,omitempty"`
	Triggers []string `json:"triggers,omitempty"`
	LocationIds []string `json:"locationIds,omitempty"`
	InternalRecipients []InternalRecipient `json:"internalRecipients,omitempty"`
	ExternalRecipients []ExternalRecipient `json:"externalRecipients,omitempty"`
	SiteId *string `json:"siteId,omitempty"`
}

// NewVolumePurchasingSubscriptionBase instantiates a new VolumePurchasingSubscriptionBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumePurchasingSubscriptionBase(name string) *VolumePurchasingSubscriptionBase {
	this := VolumePurchasingSubscriptionBase{}
	this.Name = name
	var enabled bool = true
	this.Enabled = &enabled
	var siteId string = "-1"
	this.SiteId = &siteId
	return &this
}

// NewVolumePurchasingSubscriptionBaseWithDefaults instantiates a new VolumePurchasingSubscriptionBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumePurchasingSubscriptionBaseWithDefaults() *VolumePurchasingSubscriptionBase {
	this := VolumePurchasingSubscriptionBase{}
	var enabled bool = true
	this.Enabled = &enabled
	var siteId string = "-1"
	this.SiteId = &siteId
	return &this
}

// GetName returns the Name field value
func (o *VolumePurchasingSubscriptionBase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscriptionBase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VolumePurchasingSubscriptionBase) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VolumePurchasingSubscriptionBase) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscriptionBase) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VolumePurchasingSubscriptionBase) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VolumePurchasingSubscriptionBase) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *VolumePurchasingSubscriptionBase) GetTriggers() []string {
	if o == nil || o.Triggers == nil {
		var ret []string
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscriptionBase) GetTriggersOk() ([]string, bool) {
	if o == nil || o.Triggers == nil {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *VolumePurchasingSubscriptionBase) HasTriggers() bool {
	if o != nil && o.Triggers != nil {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []string and assigns it to the Triggers field.
func (o *VolumePurchasingSubscriptionBase) SetTriggers(v []string) {
	o.Triggers = v
}

// GetLocationIds returns the LocationIds field value if set, zero value otherwise.
func (o *VolumePurchasingSubscriptionBase) GetLocationIds() []string {
	if o == nil || o.LocationIds == nil {
		var ret []string
		return ret
	}
	return o.LocationIds
}

// GetLocationIdsOk returns a tuple with the LocationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscriptionBase) GetLocationIdsOk() ([]string, bool) {
	if o == nil || o.LocationIds == nil {
		return nil, false
	}
	return o.LocationIds, true
}

// HasLocationIds returns a boolean if a field has been set.
func (o *VolumePurchasingSubscriptionBase) HasLocationIds() bool {
	if o != nil && o.LocationIds != nil {
		return true
	}

	return false
}

// SetLocationIds gets a reference to the given []string and assigns it to the LocationIds field.
func (o *VolumePurchasingSubscriptionBase) SetLocationIds(v []string) {
	o.LocationIds = v
}

// GetInternalRecipients returns the InternalRecipients field value if set, zero value otherwise.
func (o *VolumePurchasingSubscriptionBase) GetInternalRecipients() []InternalRecipient {
	if o == nil || o.InternalRecipients == nil {
		var ret []InternalRecipient
		return ret
	}
	return o.InternalRecipients
}

// GetInternalRecipientsOk returns a tuple with the InternalRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscriptionBase) GetInternalRecipientsOk() ([]InternalRecipient, bool) {
	if o == nil || o.InternalRecipients == nil {
		return nil, false
	}
	return o.InternalRecipients, true
}

// HasInternalRecipients returns a boolean if a field has been set.
func (o *VolumePurchasingSubscriptionBase) HasInternalRecipients() bool {
	if o != nil && o.InternalRecipients != nil {
		return true
	}

	return false
}

// SetInternalRecipients gets a reference to the given []InternalRecipient and assigns it to the InternalRecipients field.
func (o *VolumePurchasingSubscriptionBase) SetInternalRecipients(v []InternalRecipient) {
	o.InternalRecipients = v
}

// GetExternalRecipients returns the ExternalRecipients field value if set, zero value otherwise.
func (o *VolumePurchasingSubscriptionBase) GetExternalRecipients() []ExternalRecipient {
	if o == nil || o.ExternalRecipients == nil {
		var ret []ExternalRecipient
		return ret
	}
	return o.ExternalRecipients
}

// GetExternalRecipientsOk returns a tuple with the ExternalRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscriptionBase) GetExternalRecipientsOk() ([]ExternalRecipient, bool) {
	if o == nil || o.ExternalRecipients == nil {
		return nil, false
	}
	return o.ExternalRecipients, true
}

// HasExternalRecipients returns a boolean if a field has been set.
func (o *VolumePurchasingSubscriptionBase) HasExternalRecipients() bool {
	if o != nil && o.ExternalRecipients != nil {
		return true
	}

	return false
}

// SetExternalRecipients gets a reference to the given []ExternalRecipient and assigns it to the ExternalRecipients field.
func (o *VolumePurchasingSubscriptionBase) SetExternalRecipients(v []ExternalRecipient) {
	o.ExternalRecipients = v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *VolumePurchasingSubscriptionBase) GetSiteId() string {
	if o == nil || o.SiteId == nil {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscriptionBase) GetSiteIdOk() (*string, bool) {
	if o == nil || o.SiteId == nil {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *VolumePurchasingSubscriptionBase) HasSiteId() bool {
	if o != nil && o.SiteId != nil {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *VolumePurchasingSubscriptionBase) SetSiteId(v string) {
	o.SiteId = &v
}

func (o VolumePurchasingSubscriptionBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Triggers != nil {
		toSerialize["triggers"] = o.Triggers
	}
	if o.LocationIds != nil {
		toSerialize["locationIds"] = o.LocationIds
	}
	if o.InternalRecipients != nil {
		toSerialize["internalRecipients"] = o.InternalRecipients
	}
	if o.ExternalRecipients != nil {
		toSerialize["externalRecipients"] = o.ExternalRecipients
	}
	if o.SiteId != nil {
		toSerialize["siteId"] = o.SiteId
	}
	return json.Marshal(toSerialize)
}

type NullableVolumePurchasingSubscriptionBase struct {
	value *VolumePurchasingSubscriptionBase
	isSet bool
}

func (v NullableVolumePurchasingSubscriptionBase) Get() *VolumePurchasingSubscriptionBase {
	return v.value
}

func (v *NullableVolumePurchasingSubscriptionBase) Set(val *VolumePurchasingSubscriptionBase) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumePurchasingSubscriptionBase) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumePurchasingSubscriptionBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumePurchasingSubscriptionBase(val *VolumePurchasingSubscriptionBase) *NullableVolumePurchasingSubscriptionBase {
	return &NullableVolumePurchasingSubscriptionBase{value: val, isSet: true}
}

func (v NullableVolumePurchasingSubscriptionBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumePurchasingSubscriptionBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


