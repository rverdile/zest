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

// checks if the AnsibleTagResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleTagResponse{}

// AnsibleTagResponse A serializer for nesting in the CollectionVersion model.
type AnsibleTagResponse struct {
	// The name of the Tag.
	Name *string `json:"name,omitempty"`
}

// NewAnsibleTagResponse instantiates a new AnsibleTagResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleTagResponse() *AnsibleTagResponse {
	this := AnsibleTagResponse{}
	return &this
}

// NewAnsibleTagResponseWithDefaults instantiates a new AnsibleTagResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleTagResponseWithDefaults() *AnsibleTagResponse {
	this := AnsibleTagResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AnsibleTagResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleTagResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AnsibleTagResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AnsibleTagResponse) SetName(v string) {
	o.Name = &v
}

func (o AnsibleTagResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleTagResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAnsibleTagResponse struct {
	value *AnsibleTagResponse
	isSet bool
}

func (v NullableAnsibleTagResponse) Get() *AnsibleTagResponse {
	return v.value
}

func (v *NullableAnsibleTagResponse) Set(val *AnsibleTagResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleTagResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleTagResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleTagResponse(val *AnsibleTagResponse) *NullableAnsibleTagResponse {
	return &NullableAnsibleTagResponse{value: val, isSet: true}
}

func (v NullableAnsibleTagResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleTagResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


