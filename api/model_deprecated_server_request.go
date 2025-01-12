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

// DeprecatedServerRequest An old Cloud Identity Provider LDAP server configuration for requests
type DeprecatedServerRequest struct {
	Enabled bool `json:"enabled"`
	ProviderName string `json:"providerName"`
	DisplayName string `json:"displayName"`
	ServerUrl string `json:"serverUrl"`
	DomainName string `json:"domainName"`
	Port int32 `json:"port"`
	Keystore CloudLdapKeystoreFile `json:"keystore"`
	ConnectionTimeout int32 `json:"connectionTimeout"`
	SearchTimeout int32 `json:"searchTimeout"`
	UseWildcards bool `json:"useWildcards"`
	ConnectionType string `json:"connectionType"`
}

// NewDeprecatedServerRequest instantiates a new DeprecatedServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedServerRequest(enabled bool, providerName string, displayName string, serverUrl string, domainName string, port int32, keystore CloudLdapKeystoreFile, connectionTimeout int32, searchTimeout int32, useWildcards bool, connectionType string) *DeprecatedServerRequest {
	this := DeprecatedServerRequest{}
	this.Enabled = enabled
	this.ProviderName = providerName
	this.DisplayName = displayName
	this.ServerUrl = serverUrl
	this.DomainName = domainName
	this.Port = port
	this.Keystore = keystore
	this.ConnectionTimeout = connectionTimeout
	this.SearchTimeout = searchTimeout
	this.UseWildcards = useWildcards
	this.ConnectionType = connectionType
	return &this
}

// NewDeprecatedServerRequestWithDefaults instantiates a new DeprecatedServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedServerRequestWithDefaults() *DeprecatedServerRequest {
	this := DeprecatedServerRequest{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DeprecatedServerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DeprecatedServerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProviderName returns the ProviderName field value
func (o *DeprecatedServerRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *DeprecatedServerRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetDisplayName returns the DisplayName field value
func (o *DeprecatedServerRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *DeprecatedServerRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetServerUrl returns the ServerUrl field value
func (o *DeprecatedServerRequest) GetServerUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerUrl
}

// GetServerUrlOk returns a tuple with the ServerUrl field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerRequest) GetServerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerUrl, true
}

// SetServerUrl sets field value
func (o *DeprecatedServerRequest) SetServerUrl(v string) {
	o.ServerUrl = v
}

// GetDomainName returns the DomainName field value
func (o *DeprecatedServerRequest) GetDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerRequest) GetDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainName, true
}

// SetDomainName sets field value
func (o *DeprecatedServerRequest) SetDomainName(v string) {
	o.DomainName = v
}

// GetPort returns the Port field value
func (o *DeprecatedServerRequest) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerRequest) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *DeprecatedServerRequest) SetPort(v int32) {
	o.Port = v
}

// GetKeystore returns the Keystore field value
func (o *DeprecatedServerRequest) GetKeystore() CloudLdapKeystoreFile {
	if o == nil {
		var ret CloudLdapKeystoreFile
		return ret
	}

	return o.Keystore
}

// GetKeystoreOk returns a tuple with the Keystore field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerRequest) GetKeystoreOk() (*CloudLdapKeystoreFile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keystore, true
}

// SetKeystore sets field value
func (o *DeprecatedServerRequest) SetKeystore(v CloudLdapKeystoreFile) {
	o.Keystore = v
}

// GetConnectionTimeout returns the ConnectionTimeout field value
func (o *DeprecatedServerRequest) GetConnectionTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConnectionTimeout
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerRequest) GetConnectionTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionTimeout, true
}

// SetConnectionTimeout sets field value
func (o *DeprecatedServerRequest) SetConnectionTimeout(v int32) {
	o.ConnectionTimeout = v
}

// GetSearchTimeout returns the SearchTimeout field value
func (o *DeprecatedServerRequest) GetSearchTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SearchTimeout
}

// GetSearchTimeoutOk returns a tuple with the SearchTimeout field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerRequest) GetSearchTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchTimeout, true
}

// SetSearchTimeout sets field value
func (o *DeprecatedServerRequest) SetSearchTimeout(v int32) {
	o.SearchTimeout = v
}

// GetUseWildcards returns the UseWildcards field value
func (o *DeprecatedServerRequest) GetUseWildcards() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseWildcards
}

// GetUseWildcardsOk returns a tuple with the UseWildcards field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerRequest) GetUseWildcardsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseWildcards, true
}

// SetUseWildcards sets field value
func (o *DeprecatedServerRequest) SetUseWildcards(v bool) {
	o.UseWildcards = v
}

// GetConnectionType returns the ConnectionType field value
func (o *DeprecatedServerRequest) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServerRequest) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *DeprecatedServerRequest) SetConnectionType(v string) {
	o.ConnectionType = v
}

func (o DeprecatedServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["providerName"] = o.ProviderName
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["serverUrl"] = o.ServerUrl
	}
	if true {
		toSerialize["domainName"] = o.DomainName
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["keystore"] = o.Keystore
	}
	if true {
		toSerialize["connectionTimeout"] = o.ConnectionTimeout
	}
	if true {
		toSerialize["searchTimeout"] = o.SearchTimeout
	}
	if true {
		toSerialize["useWildcards"] = o.UseWildcards
	}
	if true {
		toSerialize["connectionType"] = o.ConnectionType
	}
	return json.Marshal(toSerialize)
}

type NullableDeprecatedServerRequest struct {
	value *DeprecatedServerRequest
	isSet bool
}

func (v NullableDeprecatedServerRequest) Get() *DeprecatedServerRequest {
	return v.value
}

func (v *NullableDeprecatedServerRequest) Set(val *DeprecatedServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedServerRequest(val *DeprecatedServerRequest) *NullableDeprecatedServerRequest {
	return &NullableDeprecatedServerRequest{value: val, isSet: true}
}

func (v NullableDeprecatedServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


