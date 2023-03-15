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

// checks if the ContainerContainerDistribution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerContainerDistribution{}

// ContainerContainerDistribution A serializer for ContainerDistribution.
type ContainerContainerDistribution struct {
	// A unique name. Ex, `rawhide` and `stable`.
	Name string `json:"name"`
	// The latest RepositoryVersion for this Repository will be served.
	Repository NullableString `json:"repository,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// An optional content-guard. If none is specified, a default one will be used.
	ContentGuard *string `json:"content_guard,omitempty"`
	// The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \"foo\" and \"foo/bar\")
	BasePath string `json:"base_path"`
	// RepositoryVersion to be served
	RepositoryVersion NullableString `json:"repository_version,omitempty"`
	// Restrict pull access to explicitly authorized users. Defaults to unrestricted pull access.
	Private *bool `json:"private,omitempty"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
}

// NewContainerContainerDistribution instantiates a new ContainerContainerDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerContainerDistribution(name string, basePath string) *ContainerContainerDistribution {
	this := ContainerContainerDistribution{}
	this.Name = name
	this.BasePath = basePath
	return &this
}

// NewContainerContainerDistributionWithDefaults instantiates a new ContainerContainerDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerContainerDistributionWithDefaults() *ContainerContainerDistribution {
	this := ContainerContainerDistribution{}
	return &this
}

// GetName returns the Name field value
func (o *ContainerContainerDistribution) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistribution) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerContainerDistribution) SetName(v string) {
	o.Name = v
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerDistribution) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerDistribution) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *ContainerContainerDistribution) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *ContainerContainerDistribution) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *ContainerContainerDistribution) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *ContainerContainerDistribution) UnsetRepository() {
	o.Repository.Unset()
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *ContainerContainerDistribution) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistribution) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *ContainerContainerDistribution) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *ContainerContainerDistribution) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetContentGuard returns the ContentGuard field value if set, zero value otherwise.
func (o *ContainerContainerDistribution) GetContentGuard() string {
	if o == nil || IsNil(o.ContentGuard) {
		var ret string
		return ret
	}
	return *o.ContentGuard
}

// GetContentGuardOk returns a tuple with the ContentGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistribution) GetContentGuardOk() (*string, bool) {
	if o == nil || IsNil(o.ContentGuard) {
		return nil, false
	}
	return o.ContentGuard, true
}

// HasContentGuard returns a boolean if a field has been set.
func (o *ContainerContainerDistribution) HasContentGuard() bool {
	if o != nil && !IsNil(o.ContentGuard) {
		return true
	}

	return false
}

// SetContentGuard gets a reference to the given string and assigns it to the ContentGuard field.
func (o *ContainerContainerDistribution) SetContentGuard(v string) {
	o.ContentGuard = &v
}

// GetBasePath returns the BasePath field value
func (o *ContainerContainerDistribution) GetBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistribution) GetBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasePath, true
}

// SetBasePath sets field value
func (o *ContainerContainerDistribution) SetBasePath(v string) {
	o.BasePath = v
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerDistribution) GetRepositoryVersion() string {
	if o == nil || IsNil(o.RepositoryVersion.Get()) {
		var ret string
		return ret
	}
	return *o.RepositoryVersion.Get()
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerDistribution) GetRepositoryVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepositoryVersion.Get(), o.RepositoryVersion.IsSet()
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *ContainerContainerDistribution) HasRepositoryVersion() bool {
	if o != nil && o.RepositoryVersion.IsSet() {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given NullableString and assigns it to the RepositoryVersion field.
func (o *ContainerContainerDistribution) SetRepositoryVersion(v string) {
	o.RepositoryVersion.Set(&v)
}
// SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil
func (o *ContainerContainerDistribution) SetRepositoryVersionNil() {
	o.RepositoryVersion.Set(nil)
}

// UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
func (o *ContainerContainerDistribution) UnsetRepositoryVersion() {
	o.RepositoryVersion.Unset()
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *ContainerContainerDistribution) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistribution) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ContainerContainerDistribution) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *ContainerContainerDistribution) SetPrivate(v bool) {
	o.Private = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerDistribution) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerDistribution) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ContainerContainerDistribution) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ContainerContainerDistribution) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ContainerContainerDistribution) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ContainerContainerDistribution) UnsetDescription() {
	o.Description.Unset()
}

func (o ContainerContainerDistribution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerContainerDistribution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if !IsNil(o.ContentGuard) {
		toSerialize["content_guard"] = o.ContentGuard
	}
	toSerialize["base_path"] = o.BasePath
	if o.RepositoryVersion.IsSet() {
		toSerialize["repository_version"] = o.RepositoryVersion.Get()
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableContainerContainerDistribution struct {
	value *ContainerContainerDistribution
	isSet bool
}

func (v NullableContainerContainerDistribution) Get() *ContainerContainerDistribution {
	return v.value
}

func (v *NullableContainerContainerDistribution) Set(val *ContainerContainerDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerContainerDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerContainerDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerContainerDistribution(val *ContainerContainerDistribution) *NullableContainerContainerDistribution {
	return &NullableContainerContainerDistribution{value: val, isSet: true}
}

func (v NullableContainerContainerDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerContainerDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


