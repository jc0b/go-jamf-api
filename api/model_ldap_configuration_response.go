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

// LdapConfigurationResponse A Cloud Identity Provider LDAP configuration for responses
type LdapConfigurationResponse struct {
	CloudIdPCommon CloudIdPCommon `json:"cloudIdPCommon"`
	Server CloudLdapServerResponse `json:"server"`
	Mappings *CloudLdapMappingsResponse `json:"mappings,omitempty"`
}

// NewLdapConfigurationResponse instantiates a new LdapConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapConfigurationResponse(cloudIdPCommon CloudIdPCommon, server CloudLdapServerResponse) *LdapConfigurationResponse {
	this := LdapConfigurationResponse{}
	this.CloudIdPCommon = cloudIdPCommon
	this.Server = server
	return &this
}

// NewLdapConfigurationResponseWithDefaults instantiates a new LdapConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapConfigurationResponseWithDefaults() *LdapConfigurationResponse {
	this := LdapConfigurationResponse{}
	return &this
}

// GetCloudIdPCommon returns the CloudIdPCommon field value
func (o *LdapConfigurationResponse) GetCloudIdPCommon() CloudIdPCommon {
	if o == nil {
		var ret CloudIdPCommon
		return ret
	}

	return o.CloudIdPCommon
}

// GetCloudIdPCommonOk returns a tuple with the CloudIdPCommon field value
// and a boolean to check if the value has been set.
func (o *LdapConfigurationResponse) GetCloudIdPCommonOk() (*CloudIdPCommon, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudIdPCommon, true
}

// SetCloudIdPCommon sets field value
func (o *LdapConfigurationResponse) SetCloudIdPCommon(v CloudIdPCommon) {
	o.CloudIdPCommon = v
}

// GetServer returns the Server field value
func (o *LdapConfigurationResponse) GetServer() CloudLdapServerResponse {
	if o == nil {
		var ret CloudLdapServerResponse
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *LdapConfigurationResponse) GetServerOk() (*CloudLdapServerResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *LdapConfigurationResponse) SetServer(v CloudLdapServerResponse) {
	o.Server = v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *LdapConfigurationResponse) GetMappings() CloudLdapMappingsResponse {
	if o == nil || o.Mappings == nil {
		var ret CloudLdapMappingsResponse
		return ret
	}
	return *o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigurationResponse) GetMappingsOk() (*CloudLdapMappingsResponse, bool) {
	if o == nil || o.Mappings == nil {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *LdapConfigurationResponse) HasMappings() bool {
	if o != nil && o.Mappings != nil {
		return true
	}

	return false
}

// SetMappings gets a reference to the given CloudLdapMappingsResponse and assigns it to the Mappings field.
func (o *LdapConfigurationResponse) SetMappings(v CloudLdapMappingsResponse) {
	o.Mappings = &v
}

func (o LdapConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cloudIdPCommon"] = o.CloudIdPCommon
	}
	if true {
		toSerialize["server"] = o.Server
	}
	if o.Mappings != nil {
		toSerialize["mappings"] = o.Mappings
	}
	return json.Marshal(toSerialize)
}

type NullableLdapConfigurationResponse struct {
	value *LdapConfigurationResponse
	isSet bool
}

func (v NullableLdapConfigurationResponse) Get() *LdapConfigurationResponse {
	return v.value
}

func (v *NullableLdapConfigurationResponse) Set(val *LdapConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapConfigurationResponse(val *LdapConfigurationResponse) *NullableLdapConfigurationResponse {
	return &NullableLdapConfigurationResponse{value: val, isSet: true}
}

func (v NullableLdapConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


