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

// checks if the FilesystemExport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesystemExport{}

// FilesystemExport Serializer for FilesystemExports.
type FilesystemExport struct {
	// A URI of the task that ran the Export.
	Task NullableString `json:"task,omitempty"`
	// A URI of the publication to be exported.
	Publication *string `json:"publication,omitempty"`
	// A URI of the repository version export.
	RepositoryVersion *string `json:"repository_version,omitempty"`
	// The URI of the last-exported-repo-version.
	StartRepositoryVersion *string `json:"start_repository_version,omitempty"`
}

// NewFilesystemExport instantiates a new FilesystemExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesystemExport() *FilesystemExport {
	this := FilesystemExport{}
	return &this
}

// NewFilesystemExportWithDefaults instantiates a new FilesystemExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesystemExportWithDefaults() *FilesystemExport {
	this := FilesystemExport{}
	return &this
}

// GetTask returns the Task field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesystemExport) GetTask() string {
	if o == nil || IsNil(o.Task.Get()) {
		var ret string
		return ret
	}
	return *o.Task.Get()
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesystemExport) GetTaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Task.Get(), o.Task.IsSet()
}

// HasTask returns a boolean if a field has been set.
func (o *FilesystemExport) HasTask() bool {
	if o != nil && o.Task.IsSet() {
		return true
	}

	return false
}

// SetTask gets a reference to the given NullableString and assigns it to the Task field.
func (o *FilesystemExport) SetTask(v string) {
	o.Task.Set(&v)
}
// SetTaskNil sets the value for Task to be an explicit nil
func (o *FilesystemExport) SetTaskNil() {
	o.Task.Set(nil)
}

// UnsetTask ensures that no value is present for Task, not even an explicit nil
func (o *FilesystemExport) UnsetTask() {
	o.Task.Unset()
}

// GetPublication returns the Publication field value if set, zero value otherwise.
func (o *FilesystemExport) GetPublication() string {
	if o == nil || IsNil(o.Publication) {
		var ret string
		return ret
	}
	return *o.Publication
}

// GetPublicationOk returns a tuple with the Publication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemExport) GetPublicationOk() (*string, bool) {
	if o == nil || IsNil(o.Publication) {
		return nil, false
	}
	return o.Publication, true
}

// HasPublication returns a boolean if a field has been set.
func (o *FilesystemExport) HasPublication() bool {
	if o != nil && !IsNil(o.Publication) {
		return true
	}

	return false
}

// SetPublication gets a reference to the given string and assigns it to the Publication field.
func (o *FilesystemExport) SetPublication(v string) {
	o.Publication = &v
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise.
func (o *FilesystemExport) GetRepositoryVersion() string {
	if o == nil || IsNil(o.RepositoryVersion) {
		var ret string
		return ret
	}
	return *o.RepositoryVersion
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemExport) GetRepositoryVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryVersion) {
		return nil, false
	}
	return o.RepositoryVersion, true
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *FilesystemExport) HasRepositoryVersion() bool {
	if o != nil && !IsNil(o.RepositoryVersion) {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given string and assigns it to the RepositoryVersion field.
func (o *FilesystemExport) SetRepositoryVersion(v string) {
	o.RepositoryVersion = &v
}

// GetStartRepositoryVersion returns the StartRepositoryVersion field value if set, zero value otherwise.
func (o *FilesystemExport) GetStartRepositoryVersion() string {
	if o == nil || IsNil(o.StartRepositoryVersion) {
		var ret string
		return ret
	}
	return *o.StartRepositoryVersion
}

// GetStartRepositoryVersionOk returns a tuple with the StartRepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemExport) GetStartRepositoryVersionOk() (*string, bool) {
	if o == nil || IsNil(o.StartRepositoryVersion) {
		return nil, false
	}
	return o.StartRepositoryVersion, true
}

// HasStartRepositoryVersion returns a boolean if a field has been set.
func (o *FilesystemExport) HasStartRepositoryVersion() bool {
	if o != nil && !IsNil(o.StartRepositoryVersion) {
		return true
	}

	return false
}

// SetStartRepositoryVersion gets a reference to the given string and assigns it to the StartRepositoryVersion field.
func (o *FilesystemExport) SetStartRepositoryVersion(v string) {
	o.StartRepositoryVersion = &v
}

func (o FilesystemExport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesystemExport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Task.IsSet() {
		toSerialize["task"] = o.Task.Get()
	}
	if !IsNil(o.Publication) {
		toSerialize["publication"] = o.Publication
	}
	if !IsNil(o.RepositoryVersion) {
		toSerialize["repository_version"] = o.RepositoryVersion
	}
	if !IsNil(o.StartRepositoryVersion) {
		toSerialize["start_repository_version"] = o.StartRepositoryVersion
	}
	return toSerialize, nil
}

type NullableFilesystemExport struct {
	value *FilesystemExport
	isSet bool
}

func (v NullableFilesystemExport) Get() *FilesystemExport {
	return v.value
}

func (v *NullableFilesystemExport) Set(val *FilesystemExport) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystemExport) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystemExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystemExport(val *FilesystemExport) *NullableFilesystemExport {
	return &NullableFilesystemExport{value: val, isSet: true}
}

func (v NullableFilesystemExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystemExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


