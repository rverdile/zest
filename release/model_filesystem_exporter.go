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

// checks if the FilesystemExporter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesystemExporter{}

// FilesystemExporter Serializer for FilesystemExporters.
type FilesystemExporter struct {
	// Unique name of the file system exporter.
	Name string `json:"name"`
	// File system location to export to.
	Path string `json:"path"`
	Method *MethodEnum `json:"method,omitempty"`
}

// NewFilesystemExporter instantiates a new FilesystemExporter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesystemExporter(name string, path string) *FilesystemExporter {
	this := FilesystemExporter{}
	this.Name = name
	this.Path = path
	var method MethodEnum = METHODENUM_WRITE
	this.Method = &method
	return &this
}

// NewFilesystemExporterWithDefaults instantiates a new FilesystemExporter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesystemExporterWithDefaults() *FilesystemExporter {
	this := FilesystemExporter{}
	var method MethodEnum = METHODENUM_WRITE
	this.Method = &method
	return &this
}

// GetName returns the Name field value
func (o *FilesystemExporter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FilesystemExporter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FilesystemExporter) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value
func (o *FilesystemExporter) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *FilesystemExporter) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *FilesystemExporter) SetPath(v string) {
	o.Path = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *FilesystemExporter) GetMethod() MethodEnum {
	if o == nil || IsNil(o.Method) {
		var ret MethodEnum
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemExporter) GetMethodOk() (*MethodEnum, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *FilesystemExporter) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given MethodEnum and assigns it to the Method field.
func (o *FilesystemExporter) SetMethod(v MethodEnum) {
	o.Method = &v
}

func (o FilesystemExporter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesystemExporter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["path"] = o.Path
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	return toSerialize, nil
}

type NullableFilesystemExporter struct {
	value *FilesystemExporter
	isSet bool
}

func (v NullableFilesystemExporter) Get() *FilesystemExporter {
	return v.value
}

func (v *NullableFilesystemExporter) Set(val *FilesystemExporter) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystemExporter) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystemExporter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystemExporter(val *FilesystemExporter) *NullableFilesystemExporter {
	return &NullableFilesystemExporter{value: val, isSet: true}
}

func (v NullableFilesystemExporter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystemExporter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


