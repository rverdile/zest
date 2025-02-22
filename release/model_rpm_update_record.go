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
	"os"
)

// checks if the RpmUpdateRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmUpdateRecord{}

// RpmUpdateRecord A Serializer for UpdateRecord.
type RpmUpdateRecord struct {
	// A URI of a repository the new content unit should be associated with.
	Repository *string `json:"repository,omitempty"`
	// An uploaded file that may be turned into the artifact of the content unit.
	File **os.File `json:"file,omitempty"`
}

// NewRpmUpdateRecord instantiates a new RpmUpdateRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmUpdateRecord() *RpmUpdateRecord {
	this := RpmUpdateRecord{}
	return &this
}

// NewRpmUpdateRecordWithDefaults instantiates a new RpmUpdateRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmUpdateRecordWithDefaults() *RpmUpdateRecord {
	this := RpmUpdateRecord{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *RpmUpdateRecord) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecord) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *RpmUpdateRecord) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *RpmUpdateRecord) SetRepository(v string) {
	o.Repository = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *RpmUpdateRecord) GetFile() *os.File {
	if o == nil || IsNil(o.File) {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecord) GetFileOk() (**os.File, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *RpmUpdateRecord) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *RpmUpdateRecord) SetFile(v *os.File) {
	o.File = &v
}

func (o RpmUpdateRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmUpdateRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	return toSerialize, nil
}

type NullableRpmUpdateRecord struct {
	value *RpmUpdateRecord
	isSet bool
}

func (v NullableRpmUpdateRecord) Get() *RpmUpdateRecord {
	return v.value
}

func (v *NullableRpmUpdateRecord) Set(val *RpmUpdateRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmUpdateRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmUpdateRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmUpdateRecord(val *RpmUpdateRecord) *NullableRpmUpdateRecord {
	return &NullableRpmUpdateRecord{value: val, isSet: true}
}

func (v NullableRpmUpdateRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmUpdateRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


