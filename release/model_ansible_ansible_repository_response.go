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
	"time"
)

// checks if the AnsibleAnsibleRepositoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleAnsibleRepositoryResponse{}

// AnsibleAnsibleRepositoryResponse Serializer for Ansible Repositories.
type AnsibleAnsibleRepositoryResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	VersionsHref *string `json:"versions_href,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	LatestVersionHref *string `json:"latest_version_href,omitempty"`
	// A unique name for this repository.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// Retain X versions of the repository. Default is null which retains all versions.
	RetainRepoVersions NullableInt64 `json:"retain_repo_versions,omitempty"`
	// An optional remote to use by default when syncing.
	Remote NullableString `json:"remote,omitempty"`
	// Last synced metadata time.
	LastSyncedMetadataTime NullableTime `json:"last_synced_metadata_time,omitempty"`
	// Gpg public key to verify collection signatures against
	Gpgkey NullableString `json:"gpgkey,omitempty"`
}

// NewAnsibleAnsibleRepositoryResponse instantiates a new AnsibleAnsibleRepositoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleAnsibleRepositoryResponse(name string) *AnsibleAnsibleRepositoryResponse {
	this := AnsibleAnsibleRepositoryResponse{}
	this.Name = name
	return &this
}

// NewAnsibleAnsibleRepositoryResponseWithDefaults instantiates a new AnsibleAnsibleRepositoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleAnsibleRepositoryResponseWithDefaults() *AnsibleAnsibleRepositoryResponse {
	this := AnsibleAnsibleRepositoryResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *AnsibleAnsibleRepositoryResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleRepositoryResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *AnsibleAnsibleRepositoryResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *AnsibleAnsibleRepositoryResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *AnsibleAnsibleRepositoryResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleRepositoryResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *AnsibleAnsibleRepositoryResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *AnsibleAnsibleRepositoryResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetVersionsHref returns the VersionsHref field value if set, zero value otherwise.
func (o *AnsibleAnsibleRepositoryResponse) GetVersionsHref() string {
	if o == nil || IsNil(o.VersionsHref) {
		var ret string
		return ret
	}
	return *o.VersionsHref
}

// GetVersionsHrefOk returns a tuple with the VersionsHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleRepositoryResponse) GetVersionsHrefOk() (*string, bool) {
	if o == nil || IsNil(o.VersionsHref) {
		return nil, false
	}
	return o.VersionsHref, true
}

// HasVersionsHref returns a boolean if a field has been set.
func (o *AnsibleAnsibleRepositoryResponse) HasVersionsHref() bool {
	if o != nil && !IsNil(o.VersionsHref) {
		return true
	}

	return false
}

// SetVersionsHref gets a reference to the given string and assigns it to the VersionsHref field.
func (o *AnsibleAnsibleRepositoryResponse) SetVersionsHref(v string) {
	o.VersionsHref = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *AnsibleAnsibleRepositoryResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleRepositoryResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *AnsibleAnsibleRepositoryResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *AnsibleAnsibleRepositoryResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetLatestVersionHref returns the LatestVersionHref field value if set, zero value otherwise.
func (o *AnsibleAnsibleRepositoryResponse) GetLatestVersionHref() string {
	if o == nil || IsNil(o.LatestVersionHref) {
		var ret string
		return ret
	}
	return *o.LatestVersionHref
}

// GetLatestVersionHrefOk returns a tuple with the LatestVersionHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleRepositoryResponse) GetLatestVersionHrefOk() (*string, bool) {
	if o == nil || IsNil(o.LatestVersionHref) {
		return nil, false
	}
	return o.LatestVersionHref, true
}

// HasLatestVersionHref returns a boolean if a field has been set.
func (o *AnsibleAnsibleRepositoryResponse) HasLatestVersionHref() bool {
	if o != nil && !IsNil(o.LatestVersionHref) {
		return true
	}

	return false
}

// SetLatestVersionHref gets a reference to the given string and assigns it to the LatestVersionHref field.
func (o *AnsibleAnsibleRepositoryResponse) SetLatestVersionHref(v string) {
	o.LatestVersionHref = &v
}

// GetName returns the Name field value
func (o *AnsibleAnsibleRepositoryResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleRepositoryResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnsibleAnsibleRepositoryResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleAnsibleRepositoryResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleAnsibleRepositoryResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AnsibleAnsibleRepositoryResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AnsibleAnsibleRepositoryResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AnsibleAnsibleRepositoryResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AnsibleAnsibleRepositoryResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetRetainRepoVersions returns the RetainRepoVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleAnsibleRepositoryResponse) GetRetainRepoVersions() int64 {
	if o == nil || IsNil(o.RetainRepoVersions.Get()) {
		var ret int64
		return ret
	}
	return *o.RetainRepoVersions.Get()
}

// GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleAnsibleRepositoryResponse) GetRetainRepoVersionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetainRepoVersions.Get(), o.RetainRepoVersions.IsSet()
}

// HasRetainRepoVersions returns a boolean if a field has been set.
func (o *AnsibleAnsibleRepositoryResponse) HasRetainRepoVersions() bool {
	if o != nil && o.RetainRepoVersions.IsSet() {
		return true
	}

	return false
}

// SetRetainRepoVersions gets a reference to the given NullableInt64 and assigns it to the RetainRepoVersions field.
func (o *AnsibleAnsibleRepositoryResponse) SetRetainRepoVersions(v int64) {
	o.RetainRepoVersions.Set(&v)
}
// SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil
func (o *AnsibleAnsibleRepositoryResponse) SetRetainRepoVersionsNil() {
	o.RetainRepoVersions.Set(nil)
}

// UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
func (o *AnsibleAnsibleRepositoryResponse) UnsetRetainRepoVersions() {
	o.RetainRepoVersions.Unset()
}

// GetRemote returns the Remote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleAnsibleRepositoryResponse) GetRemote() string {
	if o == nil || IsNil(o.Remote.Get()) {
		var ret string
		return ret
	}
	return *o.Remote.Get()
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleAnsibleRepositoryResponse) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remote.Get(), o.Remote.IsSet()
}

// HasRemote returns a boolean if a field has been set.
func (o *AnsibleAnsibleRepositoryResponse) HasRemote() bool {
	if o != nil && o.Remote.IsSet() {
		return true
	}

	return false
}

// SetRemote gets a reference to the given NullableString and assigns it to the Remote field.
func (o *AnsibleAnsibleRepositoryResponse) SetRemote(v string) {
	o.Remote.Set(&v)
}
// SetRemoteNil sets the value for Remote to be an explicit nil
func (o *AnsibleAnsibleRepositoryResponse) SetRemoteNil() {
	o.Remote.Set(nil)
}

// UnsetRemote ensures that no value is present for Remote, not even an explicit nil
func (o *AnsibleAnsibleRepositoryResponse) UnsetRemote() {
	o.Remote.Unset()
}

// GetLastSyncedMetadataTime returns the LastSyncedMetadataTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleAnsibleRepositoryResponse) GetLastSyncedMetadataTime() time.Time {
	if o == nil || IsNil(o.LastSyncedMetadataTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastSyncedMetadataTime.Get()
}

// GetLastSyncedMetadataTimeOk returns a tuple with the LastSyncedMetadataTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleAnsibleRepositoryResponse) GetLastSyncedMetadataTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSyncedMetadataTime.Get(), o.LastSyncedMetadataTime.IsSet()
}

// HasLastSyncedMetadataTime returns a boolean if a field has been set.
func (o *AnsibleAnsibleRepositoryResponse) HasLastSyncedMetadataTime() bool {
	if o != nil && o.LastSyncedMetadataTime.IsSet() {
		return true
	}

	return false
}

// SetLastSyncedMetadataTime gets a reference to the given NullableTime and assigns it to the LastSyncedMetadataTime field.
func (o *AnsibleAnsibleRepositoryResponse) SetLastSyncedMetadataTime(v time.Time) {
	o.LastSyncedMetadataTime.Set(&v)
}
// SetLastSyncedMetadataTimeNil sets the value for LastSyncedMetadataTime to be an explicit nil
func (o *AnsibleAnsibleRepositoryResponse) SetLastSyncedMetadataTimeNil() {
	o.LastSyncedMetadataTime.Set(nil)
}

// UnsetLastSyncedMetadataTime ensures that no value is present for LastSyncedMetadataTime, not even an explicit nil
func (o *AnsibleAnsibleRepositoryResponse) UnsetLastSyncedMetadataTime() {
	o.LastSyncedMetadataTime.Unset()
}

// GetGpgkey returns the Gpgkey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleAnsibleRepositoryResponse) GetGpgkey() string {
	if o == nil || IsNil(o.Gpgkey.Get()) {
		var ret string
		return ret
	}
	return *o.Gpgkey.Get()
}

// GetGpgkeyOk returns a tuple with the Gpgkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleAnsibleRepositoryResponse) GetGpgkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gpgkey.Get(), o.Gpgkey.IsSet()
}

// HasGpgkey returns a boolean if a field has been set.
func (o *AnsibleAnsibleRepositoryResponse) HasGpgkey() bool {
	if o != nil && o.Gpgkey.IsSet() {
		return true
	}

	return false
}

// SetGpgkey gets a reference to the given NullableString and assigns it to the Gpgkey field.
func (o *AnsibleAnsibleRepositoryResponse) SetGpgkey(v string) {
	o.Gpgkey.Set(&v)
}
// SetGpgkeyNil sets the value for Gpgkey to be an explicit nil
func (o *AnsibleAnsibleRepositoryResponse) SetGpgkeyNil() {
	o.Gpgkey.Set(nil)
}

// UnsetGpgkey ensures that no value is present for Gpgkey, not even an explicit nil
func (o *AnsibleAnsibleRepositoryResponse) UnsetGpgkey() {
	o.Gpgkey.Unset()
}

func (o AnsibleAnsibleRepositoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleAnsibleRepositoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	// skip: versions_href is readOnly
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	// skip: latest_version_href is readOnly
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
	if o.LastSyncedMetadataTime.IsSet() {
		toSerialize["last_synced_metadata_time"] = o.LastSyncedMetadataTime.Get()
	}
	if o.Gpgkey.IsSet() {
		toSerialize["gpgkey"] = o.Gpgkey.Get()
	}
	return toSerialize, nil
}

type NullableAnsibleAnsibleRepositoryResponse struct {
	value *AnsibleAnsibleRepositoryResponse
	isSet bool
}

func (v NullableAnsibleAnsibleRepositoryResponse) Get() *AnsibleAnsibleRepositoryResponse {
	return v.value
}

func (v *NullableAnsibleAnsibleRepositoryResponse) Set(val *AnsibleAnsibleRepositoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleAnsibleRepositoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleAnsibleRepositoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleAnsibleRepositoryResponse(val *AnsibleAnsibleRepositoryResponse) *NullableAnsibleAnsibleRepositoryResponse {
	return &NullableAnsibleAnsibleRepositoryResponse{value: val, isSet: true}
}

func (v NullableAnsibleAnsibleRepositoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleAnsibleRepositoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


