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

// MobileDeviceSearchParams struct for MobileDeviceSearchParams
type MobileDeviceSearchParams struct {
	PageNumber *int32 `json:"pageNumber,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	IsLoadToEnd *bool `json:"isLoadToEnd,omitempty"`
	OrderBy []OrderBy `json:"orderBy,omitempty"`
	Udid *string `json:"udid,omitempty"`
	MacAddress *string `json:"macAddress,omitempty"`
	Name *string `json:"name,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	OsType *string `json:"osType,omitempty"`
	IsManaged *bool `json:"isManaged,omitempty"`
	ExcludedIds []int32 `json:"excludedIds,omitempty"`
}

// NewMobileDeviceSearchParams instantiates a new MobileDeviceSearchParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDeviceSearchParams() *MobileDeviceSearchParams {
	this := MobileDeviceSearchParams{}
	return &this
}

// NewMobileDeviceSearchParamsWithDefaults instantiates a new MobileDeviceSearchParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDeviceSearchParamsWithDefaults() *MobileDeviceSearchParams {
	this := MobileDeviceSearchParams{}
	return &this
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *MobileDeviceSearchParams) GetPageNumber() int32 {
	if o == nil || o.PageNumber == nil {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSearchParams) GetPageNumberOk() (*int32, bool) {
	if o == nil || o.PageNumber == nil {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *MobileDeviceSearchParams) HasPageNumber() bool {
	if o != nil && o.PageNumber != nil {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *MobileDeviceSearchParams) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *MobileDeviceSearchParams) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSearchParams) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *MobileDeviceSearchParams) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *MobileDeviceSearchParams) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetIsLoadToEnd returns the IsLoadToEnd field value if set, zero value otherwise.
func (o *MobileDeviceSearchParams) GetIsLoadToEnd() bool {
	if o == nil || o.IsLoadToEnd == nil {
		var ret bool
		return ret
	}
	return *o.IsLoadToEnd
}

// GetIsLoadToEndOk returns a tuple with the IsLoadToEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSearchParams) GetIsLoadToEndOk() (*bool, bool) {
	if o == nil || o.IsLoadToEnd == nil {
		return nil, false
	}
	return o.IsLoadToEnd, true
}

// HasIsLoadToEnd returns a boolean if a field has been set.
func (o *MobileDeviceSearchParams) HasIsLoadToEnd() bool {
	if o != nil && o.IsLoadToEnd != nil {
		return true
	}

	return false
}

// SetIsLoadToEnd gets a reference to the given bool and assigns it to the IsLoadToEnd field.
func (o *MobileDeviceSearchParams) SetIsLoadToEnd(v bool) {
	o.IsLoadToEnd = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *MobileDeviceSearchParams) GetOrderBy() []OrderBy {
	if o == nil || o.OrderBy == nil {
		var ret []OrderBy
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSearchParams) GetOrderByOk() ([]OrderBy, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *MobileDeviceSearchParams) HasOrderBy() bool {
	if o != nil && o.OrderBy != nil {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given []OrderBy and assigns it to the OrderBy field.
func (o *MobileDeviceSearchParams) SetOrderBy(v []OrderBy) {
	o.OrderBy = v
}

// GetUdid returns the Udid field value if set, zero value otherwise.
func (o *MobileDeviceSearchParams) GetUdid() string {
	if o == nil || o.Udid == nil {
		var ret string
		return ret
	}
	return *o.Udid
}

// GetUdidOk returns a tuple with the Udid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSearchParams) GetUdidOk() (*string, bool) {
	if o == nil || o.Udid == nil {
		return nil, false
	}
	return o.Udid, true
}

// HasUdid returns a boolean if a field has been set.
func (o *MobileDeviceSearchParams) HasUdid() bool {
	if o != nil && o.Udid != nil {
		return true
	}

	return false
}

// SetUdid gets a reference to the given string and assigns it to the Udid field.
func (o *MobileDeviceSearchParams) SetUdid(v string) {
	o.Udid = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *MobileDeviceSearchParams) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSearchParams) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *MobileDeviceSearchParams) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *MobileDeviceSearchParams) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MobileDeviceSearchParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSearchParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MobileDeviceSearchParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MobileDeviceSearchParams) SetName(v string) {
	o.Name = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *MobileDeviceSearchParams) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSearchParams) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *MobileDeviceSearchParams) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *MobileDeviceSearchParams) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetOsType returns the OsType field value if set, zero value otherwise.
func (o *MobileDeviceSearchParams) GetOsType() string {
	if o == nil || o.OsType == nil {
		var ret string
		return ret
	}
	return *o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSearchParams) GetOsTypeOk() (*string, bool) {
	if o == nil || o.OsType == nil {
		return nil, false
	}
	return o.OsType, true
}

// HasOsType returns a boolean if a field has been set.
func (o *MobileDeviceSearchParams) HasOsType() bool {
	if o != nil && o.OsType != nil {
		return true
	}

	return false
}

// SetOsType gets a reference to the given string and assigns it to the OsType field.
func (o *MobileDeviceSearchParams) SetOsType(v string) {
	o.OsType = &v
}

// GetIsManaged returns the IsManaged field value if set, zero value otherwise.
func (o *MobileDeviceSearchParams) GetIsManaged() bool {
	if o == nil || o.IsManaged == nil {
		var ret bool
		return ret
	}
	return *o.IsManaged
}

// GetIsManagedOk returns a tuple with the IsManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSearchParams) GetIsManagedOk() (*bool, bool) {
	if o == nil || o.IsManaged == nil {
		return nil, false
	}
	return o.IsManaged, true
}

// HasIsManaged returns a boolean if a field has been set.
func (o *MobileDeviceSearchParams) HasIsManaged() bool {
	if o != nil && o.IsManaged != nil {
		return true
	}

	return false
}

// SetIsManaged gets a reference to the given bool and assigns it to the IsManaged field.
func (o *MobileDeviceSearchParams) SetIsManaged(v bool) {
	o.IsManaged = &v
}

// GetExcludedIds returns the ExcludedIds field value if set, zero value otherwise.
func (o *MobileDeviceSearchParams) GetExcludedIds() []int32 {
	if o == nil || o.ExcludedIds == nil {
		var ret []int32
		return ret
	}
	return o.ExcludedIds
}

// GetExcludedIdsOk returns a tuple with the ExcludedIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceSearchParams) GetExcludedIdsOk() ([]int32, bool) {
	if o == nil || o.ExcludedIds == nil {
		return nil, false
	}
	return o.ExcludedIds, true
}

// HasExcludedIds returns a boolean if a field has been set.
func (o *MobileDeviceSearchParams) HasExcludedIds() bool {
	if o != nil && o.ExcludedIds != nil {
		return true
	}

	return false
}

// SetExcludedIds gets a reference to the given []int32 and assigns it to the ExcludedIds field.
func (o *MobileDeviceSearchParams) SetExcludedIds(v []int32) {
	o.ExcludedIds = v
}

func (o MobileDeviceSearchParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PageNumber != nil {
		toSerialize["pageNumber"] = o.PageNumber
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.IsLoadToEnd != nil {
		toSerialize["isLoadToEnd"] = o.IsLoadToEnd
	}
	if o.OrderBy != nil {
		toSerialize["orderBy"] = o.OrderBy
	}
	if o.Udid != nil {
		toSerialize["udid"] = o.Udid
	}
	if o.MacAddress != nil {
		toSerialize["macAddress"] = o.MacAddress
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if o.OsType != nil {
		toSerialize["osType"] = o.OsType
	}
	if o.IsManaged != nil {
		toSerialize["isManaged"] = o.IsManaged
	}
	if o.ExcludedIds != nil {
		toSerialize["excludedIds"] = o.ExcludedIds
	}
	return json.Marshal(toSerialize)
}

type NullableMobileDeviceSearchParams struct {
	value *MobileDeviceSearchParams
	isSet bool
}

func (v NullableMobileDeviceSearchParams) Get() *MobileDeviceSearchParams {
	return v.value
}

func (v *NullableMobileDeviceSearchParams) Set(val *MobileDeviceSearchParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDeviceSearchParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDeviceSearchParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDeviceSearchParams(val *MobileDeviceSearchParams) *NullableMobileDeviceSearchParams {
	return &NullableMobileDeviceSearchParams{value: val, isSet: true}
}

func (v NullableMobileDeviceSearchParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDeviceSearchParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


