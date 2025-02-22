/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest

import (
	"encoding/json"
)

// checks if the RpmRpmRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmRpmRepository{}

// RpmRpmRepository Serializer for Rpm Repositories.
type RpmRpmRepository struct {
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// A unique name for this repository.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// Retain X versions of the repository. Default is null which retains all versions.
	RetainRepoVersions NullableInt64 `json:"retain_repo_versions,omitempty"`
	// An optional remote to use by default when syncing.
	Remote NullableString `json:"remote,omitempty"`
	// Whether to automatically create publications for new repository versions, and update any distributions pointing to this repository.
	Autopublish *bool `json:"autopublish,omitempty"`
	// A reference to an associated signing service.
	MetadataSigningService NullableString `json:"metadata_signing_service,omitempty"`
	// The number of versions of each package to keep in the repository; older versions will be purged. The default is '0', which will disable this feature and keep all versions of each package.
	RetainPackageVersions *int64 `json:"retain_package_versions,omitempty"`
	MetadataChecksumType NullableMetadataChecksumTypeEnum `json:"metadata_checksum_type,omitempty"`
	PackageChecksumType NullablePackageChecksumTypeEnum `json:"package_checksum_type,omitempty"`
	// An option specifying whether a client should perform a GPG signature check on packages.
	Gpgcheck *int64 `json:"gpgcheck,omitempty"`
	// An option specifying whether a client should perform a GPG signature check on the repodata.
	RepoGpgcheck *int64 `json:"repo_gpgcheck,omitempty"`
	// DEPRECATED: An option specifying whether Pulp should generate SQLite metadata.
	SqliteMetadata *bool `json:"sqlite_metadata,omitempty"`
}

// NewRpmRpmRepository instantiates a new RpmRpmRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmRpmRepository(name string) *RpmRpmRepository {
	this := RpmRpmRepository{}
	this.Name = name
	var autopublish bool = false
	this.Autopublish = &autopublish
	var gpgcheck int64 = 0
	this.Gpgcheck = &gpgcheck
	var repoGpgcheck int64 = 0
	this.RepoGpgcheck = &repoGpgcheck
	var sqliteMetadata bool = false
	this.SqliteMetadata = &sqliteMetadata
	return &this
}

// NewRpmRpmRepositoryWithDefaults instantiates a new RpmRpmRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmRpmRepositoryWithDefaults() *RpmRpmRepository {
	this := RpmRpmRepository{}
	var autopublish bool = false
	this.Autopublish = &autopublish
	var gpgcheck int64 = 0
	this.Gpgcheck = &gpgcheck
	var repoGpgcheck int64 = 0
	this.RepoGpgcheck = &repoGpgcheck
	var sqliteMetadata bool = false
	this.SqliteMetadata = &sqliteMetadata
	return &this
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *RpmRpmRepository) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmRepository) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *RpmRpmRepository) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *RpmRpmRepository) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetName returns the Name field value
func (o *RpmRpmRepository) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RpmRpmRepository) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RpmRpmRepository) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRepository) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRepository) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RpmRpmRepository) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *RpmRpmRepository) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *RpmRpmRepository) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *RpmRpmRepository) UnsetDescription() {
	o.Description.Unset()
}

// GetRetainRepoVersions returns the RetainRepoVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRepository) GetRetainRepoVersions() int64 {
	if o == nil || IsNil(o.RetainRepoVersions.Get()) {
		var ret int64
		return ret
	}
	return *o.RetainRepoVersions.Get()
}

// GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRepository) GetRetainRepoVersionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetainRepoVersions.Get(), o.RetainRepoVersions.IsSet()
}

// HasRetainRepoVersions returns a boolean if a field has been set.
func (o *RpmRpmRepository) HasRetainRepoVersions() bool {
	if o != nil && o.RetainRepoVersions.IsSet() {
		return true
	}

	return false
}

// SetRetainRepoVersions gets a reference to the given NullableInt64 and assigns it to the RetainRepoVersions field.
func (o *RpmRpmRepository) SetRetainRepoVersions(v int64) {
	o.RetainRepoVersions.Set(&v)
}
// SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil
func (o *RpmRpmRepository) SetRetainRepoVersionsNil() {
	o.RetainRepoVersions.Set(nil)
}

// UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
func (o *RpmRpmRepository) UnsetRetainRepoVersions() {
	o.RetainRepoVersions.Unset()
}

// GetRemote returns the Remote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRepository) GetRemote() string {
	if o == nil || IsNil(o.Remote.Get()) {
		var ret string
		return ret
	}
	return *o.Remote.Get()
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRepository) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remote.Get(), o.Remote.IsSet()
}

// HasRemote returns a boolean if a field has been set.
func (o *RpmRpmRepository) HasRemote() bool {
	if o != nil && o.Remote.IsSet() {
		return true
	}

	return false
}

// SetRemote gets a reference to the given NullableString and assigns it to the Remote field.
func (o *RpmRpmRepository) SetRemote(v string) {
	o.Remote.Set(&v)
}
// SetRemoteNil sets the value for Remote to be an explicit nil
func (o *RpmRpmRepository) SetRemoteNil() {
	o.Remote.Set(nil)
}

// UnsetRemote ensures that no value is present for Remote, not even an explicit nil
func (o *RpmRpmRepository) UnsetRemote() {
	o.Remote.Unset()
}

// GetAutopublish returns the Autopublish field value if set, zero value otherwise.
func (o *RpmRpmRepository) GetAutopublish() bool {
	if o == nil || IsNil(o.Autopublish) {
		var ret bool
		return ret
	}
	return *o.Autopublish
}

// GetAutopublishOk returns a tuple with the Autopublish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmRepository) GetAutopublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Autopublish) {
		return nil, false
	}
	return o.Autopublish, true
}

// HasAutopublish returns a boolean if a field has been set.
func (o *RpmRpmRepository) HasAutopublish() bool {
	if o != nil && !IsNil(o.Autopublish) {
		return true
	}

	return false
}

// SetAutopublish gets a reference to the given bool and assigns it to the Autopublish field.
func (o *RpmRpmRepository) SetAutopublish(v bool) {
	o.Autopublish = &v
}

// GetMetadataSigningService returns the MetadataSigningService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRepository) GetMetadataSigningService() string {
	if o == nil || IsNil(o.MetadataSigningService.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataSigningService.Get()
}

// GetMetadataSigningServiceOk returns a tuple with the MetadataSigningService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRepository) GetMetadataSigningServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataSigningService.Get(), o.MetadataSigningService.IsSet()
}

// HasMetadataSigningService returns a boolean if a field has been set.
func (o *RpmRpmRepository) HasMetadataSigningService() bool {
	if o != nil && o.MetadataSigningService.IsSet() {
		return true
	}

	return false
}

// SetMetadataSigningService gets a reference to the given NullableString and assigns it to the MetadataSigningService field.
func (o *RpmRpmRepository) SetMetadataSigningService(v string) {
	o.MetadataSigningService.Set(&v)
}
// SetMetadataSigningServiceNil sets the value for MetadataSigningService to be an explicit nil
func (o *RpmRpmRepository) SetMetadataSigningServiceNil() {
	o.MetadataSigningService.Set(nil)
}

// UnsetMetadataSigningService ensures that no value is present for MetadataSigningService, not even an explicit nil
func (o *RpmRpmRepository) UnsetMetadataSigningService() {
	o.MetadataSigningService.Unset()
}

// GetRetainPackageVersions returns the RetainPackageVersions field value if set, zero value otherwise.
func (o *RpmRpmRepository) GetRetainPackageVersions() int64 {
	if o == nil || IsNil(o.RetainPackageVersions) {
		var ret int64
		return ret
	}
	return *o.RetainPackageVersions
}

// GetRetainPackageVersionsOk returns a tuple with the RetainPackageVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmRepository) GetRetainPackageVersionsOk() (*int64, bool) {
	if o == nil || IsNil(o.RetainPackageVersions) {
		return nil, false
	}
	return o.RetainPackageVersions, true
}

// HasRetainPackageVersions returns a boolean if a field has been set.
func (o *RpmRpmRepository) HasRetainPackageVersions() bool {
	if o != nil && !IsNil(o.RetainPackageVersions) {
		return true
	}

	return false
}

// SetRetainPackageVersions gets a reference to the given int64 and assigns it to the RetainPackageVersions field.
func (o *RpmRpmRepository) SetRetainPackageVersions(v int64) {
	o.RetainPackageVersions = &v
}

// GetMetadataChecksumType returns the MetadataChecksumType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRepository) GetMetadataChecksumType() MetadataChecksumTypeEnum {
	if o == nil || IsNil(o.MetadataChecksumType.Get()) {
		var ret MetadataChecksumTypeEnum
		return ret
	}
	return *o.MetadataChecksumType.Get()
}

// GetMetadataChecksumTypeOk returns a tuple with the MetadataChecksumType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRepository) GetMetadataChecksumTypeOk() (*MetadataChecksumTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataChecksumType.Get(), o.MetadataChecksumType.IsSet()
}

// HasMetadataChecksumType returns a boolean if a field has been set.
func (o *RpmRpmRepository) HasMetadataChecksumType() bool {
	if o != nil && o.MetadataChecksumType.IsSet() {
		return true
	}

	return false
}

// SetMetadataChecksumType gets a reference to the given NullableMetadataChecksumTypeEnum and assigns it to the MetadataChecksumType field.
func (o *RpmRpmRepository) SetMetadataChecksumType(v MetadataChecksumTypeEnum) {
	o.MetadataChecksumType.Set(&v)
}
// SetMetadataChecksumTypeNil sets the value for MetadataChecksumType to be an explicit nil
func (o *RpmRpmRepository) SetMetadataChecksumTypeNil() {
	o.MetadataChecksumType.Set(nil)
}

// UnsetMetadataChecksumType ensures that no value is present for MetadataChecksumType, not even an explicit nil
func (o *RpmRpmRepository) UnsetMetadataChecksumType() {
	o.MetadataChecksumType.Unset()
}

// GetPackageChecksumType returns the PackageChecksumType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRepository) GetPackageChecksumType() PackageChecksumTypeEnum {
	if o == nil || IsNil(o.PackageChecksumType.Get()) {
		var ret PackageChecksumTypeEnum
		return ret
	}
	return *o.PackageChecksumType.Get()
}

// GetPackageChecksumTypeOk returns a tuple with the PackageChecksumType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRepository) GetPackageChecksumTypeOk() (*PackageChecksumTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageChecksumType.Get(), o.PackageChecksumType.IsSet()
}

// HasPackageChecksumType returns a boolean if a field has been set.
func (o *RpmRpmRepository) HasPackageChecksumType() bool {
	if o != nil && o.PackageChecksumType.IsSet() {
		return true
	}

	return false
}

// SetPackageChecksumType gets a reference to the given NullablePackageChecksumTypeEnum and assigns it to the PackageChecksumType field.
func (o *RpmRpmRepository) SetPackageChecksumType(v PackageChecksumTypeEnum) {
	o.PackageChecksumType.Set(&v)
}
// SetPackageChecksumTypeNil sets the value for PackageChecksumType to be an explicit nil
func (o *RpmRpmRepository) SetPackageChecksumTypeNil() {
	o.PackageChecksumType.Set(nil)
}

// UnsetPackageChecksumType ensures that no value is present for PackageChecksumType, not even an explicit nil
func (o *RpmRpmRepository) UnsetPackageChecksumType() {
	o.PackageChecksumType.Unset()
}

// GetGpgcheck returns the Gpgcheck field value if set, zero value otherwise.
func (o *RpmRpmRepository) GetGpgcheck() int64 {
	if o == nil || IsNil(o.Gpgcheck) {
		var ret int64
		return ret
	}
	return *o.Gpgcheck
}

// GetGpgcheckOk returns a tuple with the Gpgcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmRepository) GetGpgcheckOk() (*int64, bool) {
	if o == nil || IsNil(o.Gpgcheck) {
		return nil, false
	}
	return o.Gpgcheck, true
}

// HasGpgcheck returns a boolean if a field has been set.
func (o *RpmRpmRepository) HasGpgcheck() bool {
	if o != nil && !IsNil(o.Gpgcheck) {
		return true
	}

	return false
}

// SetGpgcheck gets a reference to the given int64 and assigns it to the Gpgcheck field.
func (o *RpmRpmRepository) SetGpgcheck(v int64) {
	o.Gpgcheck = &v
}

// GetRepoGpgcheck returns the RepoGpgcheck field value if set, zero value otherwise.
func (o *RpmRpmRepository) GetRepoGpgcheck() int64 {
	if o == nil || IsNil(o.RepoGpgcheck) {
		var ret int64
		return ret
	}
	return *o.RepoGpgcheck
}

// GetRepoGpgcheckOk returns a tuple with the RepoGpgcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmRepository) GetRepoGpgcheckOk() (*int64, bool) {
	if o == nil || IsNil(o.RepoGpgcheck) {
		return nil, false
	}
	return o.RepoGpgcheck, true
}

// HasRepoGpgcheck returns a boolean if a field has been set.
func (o *RpmRpmRepository) HasRepoGpgcheck() bool {
	if o != nil && !IsNil(o.RepoGpgcheck) {
		return true
	}

	return false
}

// SetRepoGpgcheck gets a reference to the given int64 and assigns it to the RepoGpgcheck field.
func (o *RpmRpmRepository) SetRepoGpgcheck(v int64) {
	o.RepoGpgcheck = &v
}

// GetSqliteMetadata returns the SqliteMetadata field value if set, zero value otherwise.
func (o *RpmRpmRepository) GetSqliteMetadata() bool {
	if o == nil || IsNil(o.SqliteMetadata) {
		var ret bool
		return ret
	}
	return *o.SqliteMetadata
}

// GetSqliteMetadataOk returns a tuple with the SqliteMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmRepository) GetSqliteMetadataOk() (*bool, bool) {
	if o == nil || IsNil(o.SqliteMetadata) {
		return nil, false
	}
	return o.SqliteMetadata, true
}

// HasSqliteMetadata returns a boolean if a field has been set.
func (o *RpmRpmRepository) HasSqliteMetadata() bool {
	if o != nil && !IsNil(o.SqliteMetadata) {
		return true
	}

	return false
}

// SetSqliteMetadata gets a reference to the given bool and assigns it to the SqliteMetadata field.
func (o *RpmRpmRepository) SetSqliteMetadata(v bool) {
	o.SqliteMetadata = &v
}

func (o RpmRpmRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmRpmRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.RetainRepoVersions.IsSet() {
		toSerialize["retain_repo_versions"] = o.RetainRepoVersions.Get()
	}
	if o.Remote.IsSet() {
		toSerialize["remote"] = o.Remote.Get()
	}
	if !IsNil(o.Autopublish) {
		toSerialize["autopublish"] = o.Autopublish
	}
	if o.MetadataSigningService.IsSet() {
		toSerialize["metadata_signing_service"] = o.MetadataSigningService.Get()
	}
	if !IsNil(o.RetainPackageVersions) {
		toSerialize["retain_package_versions"] = o.RetainPackageVersions
	}
	if o.MetadataChecksumType.IsSet() {
		toSerialize["metadata_checksum_type"] = o.MetadataChecksumType.Get()
	}
	if o.PackageChecksumType.IsSet() {
		toSerialize["package_checksum_type"] = o.PackageChecksumType.Get()
	}
	if !IsNil(o.Gpgcheck) {
		toSerialize["gpgcheck"] = o.Gpgcheck
	}
	if !IsNil(o.RepoGpgcheck) {
		toSerialize["repo_gpgcheck"] = o.RepoGpgcheck
	}
	if !IsNil(o.SqliteMetadata) {
		toSerialize["sqlite_metadata"] = o.SqliteMetadata
	}
	return toSerialize, nil
}

type NullableRpmRpmRepository struct {
	value *RpmRpmRepository
	isSet bool
}

func (v NullableRpmRpmRepository) Get() *RpmRpmRepository {
	return v.value
}

func (v *NullableRpmRpmRepository) Set(val *RpmRpmRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmRpmRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmRpmRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmRpmRepository(val *RpmRpmRepository) *NullableRpmRpmRepository {
	return &NullableRpmRpmRepository{value: val, isSet: true}
}

func (v NullableRpmRpmRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmRpmRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


