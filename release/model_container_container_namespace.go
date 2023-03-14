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

// checks if the ContainerContainerNamespace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerContainerNamespace{}

// ContainerContainerNamespace Serializer for ContainerNamespaces.
type ContainerContainerNamespace struct {
	Name string `json:"name"`
}

// NewContainerContainerNamespace instantiates a new ContainerContainerNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerContainerNamespace(name string) *ContainerContainerNamespace {
	this := ContainerContainerNamespace{}
	this.Name = name
	return &this
}

// NewContainerContainerNamespaceWithDefaults instantiates a new ContainerContainerNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerContainerNamespaceWithDefaults() *ContainerContainerNamespace {
	this := ContainerContainerNamespace{}
	return &this
}

// GetName returns the Name field value
func (o *ContainerContainerNamespace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerNamespace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerContainerNamespace) SetName(v string) {
	o.Name = v
}

func (o ContainerContainerNamespace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerContainerNamespace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableContainerContainerNamespace struct {
	value *ContainerContainerNamespace
	isSet bool
}

func (v NullableContainerContainerNamespace) Get() *ContainerContainerNamespace {
	return v.value
}

func (v *NullableContainerContainerNamespace) Set(val *ContainerContainerNamespace) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerContainerNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerContainerNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerContainerNamespace(val *ContainerContainerNamespace) *NullableContainerContainerNamespace {
	return &NullableContainerContainerNamespace{value: val, isSet: true}
}

func (v NullableContainerContainerNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerContainerNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


