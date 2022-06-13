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

// ComputerGeneral struct for ComputerGeneral
type ComputerGeneral struct {
	Name *string `json:"name,omitempty"`
	LastIpAddress *string `json:"lastIpAddress,omitempty"`
	LastReportedIp *string `json:"lastReportedIp,omitempty"`
	JamfBinaryVersion *string `json:"jamfBinaryVersion,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Barcode1 *string `json:"barcode1,omitempty"`
	Barcode2 *string `json:"barcode2,omitempty"`
	AssetTag *string `json:"assetTag,omitempty"`
	RemoteManagement *ComputerRemoteManagement `json:"remoteManagement,omitempty"`
	Supervised *bool `json:"supervised,omitempty"`
	MdmCapable *ComputerMdmCapability `json:"mdmCapable,omitempty"`
	ReportDate *time.Time `json:"reportDate,omitempty"`
	LastContactTime *time.Time `json:"lastContactTime,omitempty"`
	LastCloudBackupDate *time.Time `json:"lastCloudBackupDate,omitempty"`
	LastEnrolledDate *time.Time `json:"lastEnrolledDate,omitempty"`
	MdmProfileExpiration *time.Time `json:"mdmProfileExpiration,omitempty"`
	InitialEntryDate *string `json:"initialEntryDate,omitempty"`
	DistributionPoint *string `json:"distributionPoint,omitempty"`
	EnrollmentMethod *EnrollmentMethod `json:"enrollmentMethod,omitempty"`
	Site *V1Site `json:"site,omitempty"`
	ItunesStoreAccountActive *bool `json:"itunesStoreAccountActive,omitempty"`
	EnrolledViaAutomatedDeviceEnrollment *bool `json:"enrolledViaAutomatedDeviceEnrollment,omitempty"`
	UserApprovedMdm *bool `json:"userApprovedMdm,omitempty"`
	ExtensionAttributes []ComputerExtensionAttribute `json:"extensionAttributes,omitempty"`
}

// NewComputerGeneral instantiates a new ComputerGeneral object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerGeneral() *ComputerGeneral {
	this := ComputerGeneral{}
	return &this
}

// NewComputerGeneralWithDefaults instantiates a new ComputerGeneral object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerGeneralWithDefaults() *ComputerGeneral {
	this := ComputerGeneral{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputerGeneral) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputerGeneral) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputerGeneral) SetName(v string) {
	o.Name = &v
}

// GetLastIpAddress returns the LastIpAddress field value if set, zero value otherwise.
func (o *ComputerGeneral) GetLastIpAddress() string {
	if o == nil || o.LastIpAddress == nil {
		var ret string
		return ret
	}
	return *o.LastIpAddress
}

// GetLastIpAddressOk returns a tuple with the LastIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetLastIpAddressOk() (*string, bool) {
	if o == nil || o.LastIpAddress == nil {
		return nil, false
	}
	return o.LastIpAddress, true
}

// HasLastIpAddress returns a boolean if a field has been set.
func (o *ComputerGeneral) HasLastIpAddress() bool {
	if o != nil && o.LastIpAddress != nil {
		return true
	}

	return false
}

// SetLastIpAddress gets a reference to the given string and assigns it to the LastIpAddress field.
func (o *ComputerGeneral) SetLastIpAddress(v string) {
	o.LastIpAddress = &v
}

// GetLastReportedIp returns the LastReportedIp field value if set, zero value otherwise.
func (o *ComputerGeneral) GetLastReportedIp() string {
	if o == nil || o.LastReportedIp == nil {
		var ret string
		return ret
	}
	return *o.LastReportedIp
}

// GetLastReportedIpOk returns a tuple with the LastReportedIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetLastReportedIpOk() (*string, bool) {
	if o == nil || o.LastReportedIp == nil {
		return nil, false
	}
	return o.LastReportedIp, true
}

// HasLastReportedIp returns a boolean if a field has been set.
func (o *ComputerGeneral) HasLastReportedIp() bool {
	if o != nil && o.LastReportedIp != nil {
		return true
	}

	return false
}

// SetLastReportedIp gets a reference to the given string and assigns it to the LastReportedIp field.
func (o *ComputerGeneral) SetLastReportedIp(v string) {
	o.LastReportedIp = &v
}

// GetJamfBinaryVersion returns the JamfBinaryVersion field value if set, zero value otherwise.
func (o *ComputerGeneral) GetJamfBinaryVersion() string {
	if o == nil || o.JamfBinaryVersion == nil {
		var ret string
		return ret
	}
	return *o.JamfBinaryVersion
}

// GetJamfBinaryVersionOk returns a tuple with the JamfBinaryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetJamfBinaryVersionOk() (*string, bool) {
	if o == nil || o.JamfBinaryVersion == nil {
		return nil, false
	}
	return o.JamfBinaryVersion, true
}

// HasJamfBinaryVersion returns a boolean if a field has been set.
func (o *ComputerGeneral) HasJamfBinaryVersion() bool {
	if o != nil && o.JamfBinaryVersion != nil {
		return true
	}

	return false
}

// SetJamfBinaryVersion gets a reference to the given string and assigns it to the JamfBinaryVersion field.
func (o *ComputerGeneral) SetJamfBinaryVersion(v string) {
	o.JamfBinaryVersion = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *ComputerGeneral) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *ComputerGeneral) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *ComputerGeneral) SetPlatform(v string) {
	o.Platform = &v
}

// GetBarcode1 returns the Barcode1 field value if set, zero value otherwise.
func (o *ComputerGeneral) GetBarcode1() string {
	if o == nil || o.Barcode1 == nil {
		var ret string
		return ret
	}
	return *o.Barcode1
}

// GetBarcode1Ok returns a tuple with the Barcode1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetBarcode1Ok() (*string, bool) {
	if o == nil || o.Barcode1 == nil {
		return nil, false
	}
	return o.Barcode1, true
}

// HasBarcode1 returns a boolean if a field has been set.
func (o *ComputerGeneral) HasBarcode1() bool {
	if o != nil && o.Barcode1 != nil {
		return true
	}

	return false
}

// SetBarcode1 gets a reference to the given string and assigns it to the Barcode1 field.
func (o *ComputerGeneral) SetBarcode1(v string) {
	o.Barcode1 = &v
}

// GetBarcode2 returns the Barcode2 field value if set, zero value otherwise.
func (o *ComputerGeneral) GetBarcode2() string {
	if o == nil || o.Barcode2 == nil {
		var ret string
		return ret
	}
	return *o.Barcode2
}

// GetBarcode2Ok returns a tuple with the Barcode2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetBarcode2Ok() (*string, bool) {
	if o == nil || o.Barcode2 == nil {
		return nil, false
	}
	return o.Barcode2, true
}

// HasBarcode2 returns a boolean if a field has been set.
func (o *ComputerGeneral) HasBarcode2() bool {
	if o != nil && o.Barcode2 != nil {
		return true
	}

	return false
}

// SetBarcode2 gets a reference to the given string and assigns it to the Barcode2 field.
func (o *ComputerGeneral) SetBarcode2(v string) {
	o.Barcode2 = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise.
func (o *ComputerGeneral) GetAssetTag() string {
	if o == nil || o.AssetTag == nil {
		var ret string
		return ret
	}
	return *o.AssetTag
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetAssetTagOk() (*string, bool) {
	if o == nil || o.AssetTag == nil {
		return nil, false
	}
	return o.AssetTag, true
}

// HasAssetTag returns a boolean if a field has been set.
func (o *ComputerGeneral) HasAssetTag() bool {
	if o != nil && o.AssetTag != nil {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given string and assigns it to the AssetTag field.
func (o *ComputerGeneral) SetAssetTag(v string) {
	o.AssetTag = &v
}

// GetRemoteManagement returns the RemoteManagement field value if set, zero value otherwise.
func (o *ComputerGeneral) GetRemoteManagement() ComputerRemoteManagement {
	if o == nil || o.RemoteManagement == nil {
		var ret ComputerRemoteManagement
		return ret
	}
	return *o.RemoteManagement
}

// GetRemoteManagementOk returns a tuple with the RemoteManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetRemoteManagementOk() (*ComputerRemoteManagement, bool) {
	if o == nil || o.RemoteManagement == nil {
		return nil, false
	}
	return o.RemoteManagement, true
}

// HasRemoteManagement returns a boolean if a field has been set.
func (o *ComputerGeneral) HasRemoteManagement() bool {
	if o != nil && o.RemoteManagement != nil {
		return true
	}

	return false
}

// SetRemoteManagement gets a reference to the given ComputerRemoteManagement and assigns it to the RemoteManagement field.
func (o *ComputerGeneral) SetRemoteManagement(v ComputerRemoteManagement) {
	o.RemoteManagement = &v
}

// GetSupervised returns the Supervised field value if set, zero value otherwise.
func (o *ComputerGeneral) GetSupervised() bool {
	if o == nil || o.Supervised == nil {
		var ret bool
		return ret
	}
	return *o.Supervised
}

// GetSupervisedOk returns a tuple with the Supervised field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetSupervisedOk() (*bool, bool) {
	if o == nil || o.Supervised == nil {
		return nil, false
	}
	return o.Supervised, true
}

// HasSupervised returns a boolean if a field has been set.
func (o *ComputerGeneral) HasSupervised() bool {
	if o != nil && o.Supervised != nil {
		return true
	}

	return false
}

// SetSupervised gets a reference to the given bool and assigns it to the Supervised field.
func (o *ComputerGeneral) SetSupervised(v bool) {
	o.Supervised = &v
}

// GetMdmCapable returns the MdmCapable field value if set, zero value otherwise.
func (o *ComputerGeneral) GetMdmCapable() ComputerMdmCapability {
	if o == nil || o.MdmCapable == nil {
		var ret ComputerMdmCapability
		return ret
	}
	return *o.MdmCapable
}

// GetMdmCapableOk returns a tuple with the MdmCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetMdmCapableOk() (*ComputerMdmCapability, bool) {
	if o == nil || o.MdmCapable == nil {
		return nil, false
	}
	return o.MdmCapable, true
}

// HasMdmCapable returns a boolean if a field has been set.
func (o *ComputerGeneral) HasMdmCapable() bool {
	if o != nil && o.MdmCapable != nil {
		return true
	}

	return false
}

// SetMdmCapable gets a reference to the given ComputerMdmCapability and assigns it to the MdmCapable field.
func (o *ComputerGeneral) SetMdmCapable(v ComputerMdmCapability) {
	o.MdmCapable = &v
}

// GetReportDate returns the ReportDate field value if set, zero value otherwise.
func (o *ComputerGeneral) GetReportDate() time.Time {
	if o == nil || o.ReportDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ReportDate
}

// GetReportDateOk returns a tuple with the ReportDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetReportDateOk() (*time.Time, bool) {
	if o == nil || o.ReportDate == nil {
		return nil, false
	}
	return o.ReportDate, true
}

// HasReportDate returns a boolean if a field has been set.
func (o *ComputerGeneral) HasReportDate() bool {
	if o != nil && o.ReportDate != nil {
		return true
	}

	return false
}

// SetReportDate gets a reference to the given time.Time and assigns it to the ReportDate field.
func (o *ComputerGeneral) SetReportDate(v time.Time) {
	o.ReportDate = &v
}

// GetLastContactTime returns the LastContactTime field value if set, zero value otherwise.
func (o *ComputerGeneral) GetLastContactTime() time.Time {
	if o == nil || o.LastContactTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastContactTime
}

// GetLastContactTimeOk returns a tuple with the LastContactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetLastContactTimeOk() (*time.Time, bool) {
	if o == nil || o.LastContactTime == nil {
		return nil, false
	}
	return o.LastContactTime, true
}

// HasLastContactTime returns a boolean if a field has been set.
func (o *ComputerGeneral) HasLastContactTime() bool {
	if o != nil && o.LastContactTime != nil {
		return true
	}

	return false
}

// SetLastContactTime gets a reference to the given time.Time and assigns it to the LastContactTime field.
func (o *ComputerGeneral) SetLastContactTime(v time.Time) {
	o.LastContactTime = &v
}

// GetLastCloudBackupDate returns the LastCloudBackupDate field value if set, zero value otherwise.
func (o *ComputerGeneral) GetLastCloudBackupDate() time.Time {
	if o == nil || o.LastCloudBackupDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastCloudBackupDate
}

// GetLastCloudBackupDateOk returns a tuple with the LastCloudBackupDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetLastCloudBackupDateOk() (*time.Time, bool) {
	if o == nil || o.LastCloudBackupDate == nil {
		return nil, false
	}
	return o.LastCloudBackupDate, true
}

// HasLastCloudBackupDate returns a boolean if a field has been set.
func (o *ComputerGeneral) HasLastCloudBackupDate() bool {
	if o != nil && o.LastCloudBackupDate != nil {
		return true
	}

	return false
}

// SetLastCloudBackupDate gets a reference to the given time.Time and assigns it to the LastCloudBackupDate field.
func (o *ComputerGeneral) SetLastCloudBackupDate(v time.Time) {
	o.LastCloudBackupDate = &v
}

// GetLastEnrolledDate returns the LastEnrolledDate field value if set, zero value otherwise.
func (o *ComputerGeneral) GetLastEnrolledDate() time.Time {
	if o == nil || o.LastEnrolledDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastEnrolledDate
}

// GetLastEnrolledDateOk returns a tuple with the LastEnrolledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetLastEnrolledDateOk() (*time.Time, bool) {
	if o == nil || o.LastEnrolledDate == nil {
		return nil, false
	}
	return o.LastEnrolledDate, true
}

// HasLastEnrolledDate returns a boolean if a field has been set.
func (o *ComputerGeneral) HasLastEnrolledDate() bool {
	if o != nil && o.LastEnrolledDate != nil {
		return true
	}

	return false
}

// SetLastEnrolledDate gets a reference to the given time.Time and assigns it to the LastEnrolledDate field.
func (o *ComputerGeneral) SetLastEnrolledDate(v time.Time) {
	o.LastEnrolledDate = &v
}

// GetMdmProfileExpiration returns the MdmProfileExpiration field value if set, zero value otherwise.
func (o *ComputerGeneral) GetMdmProfileExpiration() time.Time {
	if o == nil || o.MdmProfileExpiration == nil {
		var ret time.Time
		return ret
	}
	return *o.MdmProfileExpiration
}

// GetMdmProfileExpirationOk returns a tuple with the MdmProfileExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetMdmProfileExpirationOk() (*time.Time, bool) {
	if o == nil || o.MdmProfileExpiration == nil {
		return nil, false
	}
	return o.MdmProfileExpiration, true
}

// HasMdmProfileExpiration returns a boolean if a field has been set.
func (o *ComputerGeneral) HasMdmProfileExpiration() bool {
	if o != nil && o.MdmProfileExpiration != nil {
		return true
	}

	return false
}

// SetMdmProfileExpiration gets a reference to the given time.Time and assigns it to the MdmProfileExpiration field.
func (o *ComputerGeneral) SetMdmProfileExpiration(v time.Time) {
	o.MdmProfileExpiration = &v
}

// GetInitialEntryDate returns the InitialEntryDate field value if set, zero value otherwise.
func (o *ComputerGeneral) GetInitialEntryDate() string {
	if o == nil || o.InitialEntryDate == nil {
		var ret string
		return ret
	}
	return *o.InitialEntryDate
}

// GetInitialEntryDateOk returns a tuple with the InitialEntryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetInitialEntryDateOk() (*string, bool) {
	if o == nil || o.InitialEntryDate == nil {
		return nil, false
	}
	return o.InitialEntryDate, true
}

// HasInitialEntryDate returns a boolean if a field has been set.
func (o *ComputerGeneral) HasInitialEntryDate() bool {
	if o != nil && o.InitialEntryDate != nil {
		return true
	}

	return false
}

// SetInitialEntryDate gets a reference to the given string and assigns it to the InitialEntryDate field.
func (o *ComputerGeneral) SetInitialEntryDate(v string) {
	o.InitialEntryDate = &v
}

// GetDistributionPoint returns the DistributionPoint field value if set, zero value otherwise.
func (o *ComputerGeneral) GetDistributionPoint() string {
	if o == nil || o.DistributionPoint == nil {
		var ret string
		return ret
	}
	return *o.DistributionPoint
}

// GetDistributionPointOk returns a tuple with the DistributionPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetDistributionPointOk() (*string, bool) {
	if o == nil || o.DistributionPoint == nil {
		return nil, false
	}
	return o.DistributionPoint, true
}

// HasDistributionPoint returns a boolean if a field has been set.
func (o *ComputerGeneral) HasDistributionPoint() bool {
	if o != nil && o.DistributionPoint != nil {
		return true
	}

	return false
}

// SetDistributionPoint gets a reference to the given string and assigns it to the DistributionPoint field.
func (o *ComputerGeneral) SetDistributionPoint(v string) {
	o.DistributionPoint = &v
}

// GetEnrollmentMethod returns the EnrollmentMethod field value if set, zero value otherwise.
func (o *ComputerGeneral) GetEnrollmentMethod() EnrollmentMethod {
	if o == nil || o.EnrollmentMethod == nil {
		var ret EnrollmentMethod
		return ret
	}
	return *o.EnrollmentMethod
}

// GetEnrollmentMethodOk returns a tuple with the EnrollmentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetEnrollmentMethodOk() (*EnrollmentMethod, bool) {
	if o == nil || o.EnrollmentMethod == nil {
		return nil, false
	}
	return o.EnrollmentMethod, true
}

// HasEnrollmentMethod returns a boolean if a field has been set.
func (o *ComputerGeneral) HasEnrollmentMethod() bool {
	if o != nil && o.EnrollmentMethod != nil {
		return true
	}

	return false
}

// SetEnrollmentMethod gets a reference to the given EnrollmentMethod and assigns it to the EnrollmentMethod field.
func (o *ComputerGeneral) SetEnrollmentMethod(v EnrollmentMethod) {
	o.EnrollmentMethod = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *ComputerGeneral) GetSite() V1Site {
	if o == nil || o.Site == nil {
		var ret V1Site
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetSiteOk() (*V1Site, bool) {
	if o == nil || o.Site == nil {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *ComputerGeneral) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given V1Site and assigns it to the Site field.
func (o *ComputerGeneral) SetSite(v V1Site) {
	o.Site = &v
}

// GetItunesStoreAccountActive returns the ItunesStoreAccountActive field value if set, zero value otherwise.
func (o *ComputerGeneral) GetItunesStoreAccountActive() bool {
	if o == nil || o.ItunesStoreAccountActive == nil {
		var ret bool
		return ret
	}
	return *o.ItunesStoreAccountActive
}

// GetItunesStoreAccountActiveOk returns a tuple with the ItunesStoreAccountActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetItunesStoreAccountActiveOk() (*bool, bool) {
	if o == nil || o.ItunesStoreAccountActive == nil {
		return nil, false
	}
	return o.ItunesStoreAccountActive, true
}

// HasItunesStoreAccountActive returns a boolean if a field has been set.
func (o *ComputerGeneral) HasItunesStoreAccountActive() bool {
	if o != nil && o.ItunesStoreAccountActive != nil {
		return true
	}

	return false
}

// SetItunesStoreAccountActive gets a reference to the given bool and assigns it to the ItunesStoreAccountActive field.
func (o *ComputerGeneral) SetItunesStoreAccountActive(v bool) {
	o.ItunesStoreAccountActive = &v
}

// GetEnrolledViaAutomatedDeviceEnrollment returns the EnrolledViaAutomatedDeviceEnrollment field value if set, zero value otherwise.
func (o *ComputerGeneral) GetEnrolledViaAutomatedDeviceEnrollment() bool {
	if o == nil || o.EnrolledViaAutomatedDeviceEnrollment == nil {
		var ret bool
		return ret
	}
	return *o.EnrolledViaAutomatedDeviceEnrollment
}

// GetEnrolledViaAutomatedDeviceEnrollmentOk returns a tuple with the EnrolledViaAutomatedDeviceEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetEnrolledViaAutomatedDeviceEnrollmentOk() (*bool, bool) {
	if o == nil || o.EnrolledViaAutomatedDeviceEnrollment == nil {
		return nil, false
	}
	return o.EnrolledViaAutomatedDeviceEnrollment, true
}

// HasEnrolledViaAutomatedDeviceEnrollment returns a boolean if a field has been set.
func (o *ComputerGeneral) HasEnrolledViaAutomatedDeviceEnrollment() bool {
	if o != nil && o.EnrolledViaAutomatedDeviceEnrollment != nil {
		return true
	}

	return false
}

// SetEnrolledViaAutomatedDeviceEnrollment gets a reference to the given bool and assigns it to the EnrolledViaAutomatedDeviceEnrollment field.
func (o *ComputerGeneral) SetEnrolledViaAutomatedDeviceEnrollment(v bool) {
	o.EnrolledViaAutomatedDeviceEnrollment = &v
}

// GetUserApprovedMdm returns the UserApprovedMdm field value if set, zero value otherwise.
func (o *ComputerGeneral) GetUserApprovedMdm() bool {
	if o == nil || o.UserApprovedMdm == nil {
		var ret bool
		return ret
	}
	return *o.UserApprovedMdm
}

// GetUserApprovedMdmOk returns a tuple with the UserApprovedMdm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetUserApprovedMdmOk() (*bool, bool) {
	if o == nil || o.UserApprovedMdm == nil {
		return nil, false
	}
	return o.UserApprovedMdm, true
}

// HasUserApprovedMdm returns a boolean if a field has been set.
func (o *ComputerGeneral) HasUserApprovedMdm() bool {
	if o != nil && o.UserApprovedMdm != nil {
		return true
	}

	return false
}

// SetUserApprovedMdm gets a reference to the given bool and assigns it to the UserApprovedMdm field.
func (o *ComputerGeneral) SetUserApprovedMdm(v bool) {
	o.UserApprovedMdm = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *ComputerGeneral) GetExtensionAttributes() []ComputerExtensionAttribute {
	if o == nil || o.ExtensionAttributes == nil {
		var ret []ComputerExtensionAttribute
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneral) GetExtensionAttributesOk() ([]ComputerExtensionAttribute, bool) {
	if o == nil || o.ExtensionAttributes == nil {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *ComputerGeneral) HasExtensionAttributes() bool {
	if o != nil && o.ExtensionAttributes != nil {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given []ComputerExtensionAttribute and assigns it to the ExtensionAttributes field.
func (o *ComputerGeneral) SetExtensionAttributes(v []ComputerExtensionAttribute) {
	o.ExtensionAttributes = v
}

func (o ComputerGeneral) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.LastIpAddress != nil {
		toSerialize["lastIpAddress"] = o.LastIpAddress
	}
	if o.LastReportedIp != nil {
		toSerialize["lastReportedIp"] = o.LastReportedIp
	}
	if o.JamfBinaryVersion != nil {
		toSerialize["jamfBinaryVersion"] = o.JamfBinaryVersion
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.Barcode1 != nil {
		toSerialize["barcode1"] = o.Barcode1
	}
	if o.Barcode2 != nil {
		toSerialize["barcode2"] = o.Barcode2
	}
	if o.AssetTag != nil {
		toSerialize["assetTag"] = o.AssetTag
	}
	if o.RemoteManagement != nil {
		toSerialize["remoteManagement"] = o.RemoteManagement
	}
	if o.Supervised != nil {
		toSerialize["supervised"] = o.Supervised
	}
	if o.MdmCapable != nil {
		toSerialize["mdmCapable"] = o.MdmCapable
	}
	if o.ReportDate != nil {
		toSerialize["reportDate"] = o.ReportDate
	}
	if o.LastContactTime != nil {
		toSerialize["lastContactTime"] = o.LastContactTime
	}
	if o.LastCloudBackupDate != nil {
		toSerialize["lastCloudBackupDate"] = o.LastCloudBackupDate
	}
	if o.LastEnrolledDate != nil {
		toSerialize["lastEnrolledDate"] = o.LastEnrolledDate
	}
	if o.MdmProfileExpiration != nil {
		toSerialize["mdmProfileExpiration"] = o.MdmProfileExpiration
	}
	if o.InitialEntryDate != nil {
		toSerialize["initialEntryDate"] = o.InitialEntryDate
	}
	if o.DistributionPoint != nil {
		toSerialize["distributionPoint"] = o.DistributionPoint
	}
	if o.EnrollmentMethod != nil {
		toSerialize["enrollmentMethod"] = o.EnrollmentMethod
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	if o.ItunesStoreAccountActive != nil {
		toSerialize["itunesStoreAccountActive"] = o.ItunesStoreAccountActive
	}
	if o.EnrolledViaAutomatedDeviceEnrollment != nil {
		toSerialize["enrolledViaAutomatedDeviceEnrollment"] = o.EnrolledViaAutomatedDeviceEnrollment
	}
	if o.UserApprovedMdm != nil {
		toSerialize["userApprovedMdm"] = o.UserApprovedMdm
	}
	if o.ExtensionAttributes != nil {
		toSerialize["extensionAttributes"] = o.ExtensionAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableComputerGeneral struct {
	value *ComputerGeneral
	isSet bool
}

func (v NullableComputerGeneral) Get() *ComputerGeneral {
	return v.value
}

func (v *NullableComputerGeneral) Set(val *ComputerGeneral) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerGeneral) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerGeneral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerGeneral(val *ComputerGeneral) *NullableComputerGeneral {
	return &NullableComputerGeneral{value: val, isSet: true}
}

func (v NullableComputerGeneral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerGeneral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

