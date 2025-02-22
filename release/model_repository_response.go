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

// checks if the RepositoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryResponse{}

// RepositoryResponse Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type RepositoryResponse struct {
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
}

// NewRepositoryResponse instantiates a new RepositoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryResponse(name string) *RepositoryResponse {
	this := RepositoryResponse{}
	this.Name = name
	return &this
}

// NewRepositoryResponseWithDefaults instantiates a new RepositoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryResponseWithDefaults() *RepositoryResponse {
	this := RepositoryResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *RepositoryResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *RepositoryResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *RepositoryResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *RepositoryResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *RepositoryResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *RepositoryResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetVersionsHref returns the VersionsHref field value if set, zero value otherwise.
func (o *RepositoryResponse) GetVersionsHref() string {
	if o == nil || IsNil(o.VersionsHref) {
		var ret string
		return ret
	}
	return *o.VersionsHref
}

// GetVersionsHrefOk returns a tuple with the VersionsHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryResponse) GetVersionsHrefOk() (*string, bool) {
	if o == nil || IsNil(o.VersionsHref) {
		return nil, false
	}
	return o.VersionsHref, true
}

// HasVersionsHref returns a boolean if a field has been set.
func (o *RepositoryResponse) HasVersionsHref() bool {
	if o != nil && !IsNil(o.VersionsHref) {
		return true
	}

	return false
}

// SetVersionsHref gets a reference to the given string and assigns it to the VersionsHref field.
func (o *RepositoryResponse) SetVersionsHref(v string) {
	o.VersionsHref = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *RepositoryResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *RepositoryResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *RepositoryResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetLatestVersionHref returns the LatestVersionHref field value if set, zero value otherwise.
func (o *RepositoryResponse) GetLatestVersionHref() string {
	if o == nil || IsNil(o.LatestVersionHref) {
		var ret string
		return ret
	}
	return *o.LatestVersionHref
}

// GetLatestVersionHrefOk returns a tuple with the LatestVersionHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryResponse) GetLatestVersionHrefOk() (*string, bool) {
	if o == nil || IsNil(o.LatestVersionHref) {
		return nil, false
	}
	return o.LatestVersionHref, true
}

// HasLatestVersionHref returns a boolean if a field has been set.
func (o *RepositoryResponse) HasLatestVersionHref() bool {
	if o != nil && !IsNil(o.LatestVersionHref) {
		return true
	}

	return false
}

// SetLatestVersionHref gets a reference to the given string and assigns it to the LatestVersionHref field.
func (o *RepositoryResponse) SetLatestVersionHref(v string) {
	o.LatestVersionHref = &v
}

// GetName returns the Name field value
func (o *RepositoryResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RepositoryResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RepositoryResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RepositoryResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *RepositoryResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *RepositoryResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *RepositoryResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetRetainRepoVersions returns the RetainRepoVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryResponse) GetRetainRepoVersions() int64 {
	if o == nil || IsNil(o.RetainRepoVersions.Get()) {
		var ret int64
		return ret
	}
	return *o.RetainRepoVersions.Get()
}

// GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryResponse) GetRetainRepoVersionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetainRepoVersions.Get(), o.RetainRepoVersions.IsSet()
}

// HasRetainRepoVersions returns a boolean if a field has been set.
func (o *RepositoryResponse) HasRetainRepoVersions() bool {
	if o != nil && o.RetainRepoVersions.IsSet() {
		return true
	}

	return false
}

// SetRetainRepoVersions gets a reference to the given NullableInt64 and assigns it to the RetainRepoVersions field.
func (o *RepositoryResponse) SetRetainRepoVersions(v int64) {
	o.RetainRepoVersions.Set(&v)
}
// SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil
func (o *RepositoryResponse) SetRetainRepoVersionsNil() {
	o.RetainRepoVersions.Set(nil)
}

// UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
func (o *RepositoryResponse) UnsetRetainRepoVersions() {
	o.RetainRepoVersions.Unset()
}

// GetRemote returns the Remote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryResponse) GetRemote() string {
	if o == nil || IsNil(o.Remote.Get()) {
		var ret string
		return ret
	}
	return *o.Remote.Get()
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryResponse) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remote.Get(), o.Remote.IsSet()
}

// HasRemote returns a boolean if a field has been set.
func (o *RepositoryResponse) HasRemote() bool {
	if o != nil && o.Remote.IsSet() {
		return true
	}

	return false
}

// SetRemote gets a reference to the given NullableString and assigns it to the Remote field.
func (o *RepositoryResponse) SetRemote(v string) {
	o.Remote.Set(&v)
}
// SetRemoteNil sets the value for Remote to be an explicit nil
func (o *RepositoryResponse) SetRemoteNil() {
	o.Remote.Set(nil)
}

// UnsetRemote ensures that no value is present for Remote, not even an explicit nil
func (o *RepositoryResponse) UnsetRemote() {
	o.Remote.Unset()
}

func (o RepositoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if !IsNil(o.VersionsHref) {
		toSerialize["versions_href"] = o.VersionsHref
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if !IsNil(o.LatestVersionHref) {
		toSerialize["latest_version_href"] = o.LatestVersionHref
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
	return toSerialize, nil
}

type NullableRepositoryResponse struct {
	value *RepositoryResponse
	isSet bool
}

func (v NullableRepositoryResponse) Get() *RepositoryResponse {
	return v.value
}

func (v *NullableRepositoryResponse) Set(val *RepositoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryResponse(val *RepositoryResponse) *NullableRepositoryResponse {
	return &NullableRepositoryResponse{value: val, isSet: true}
}

func (v NullableRepositoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


