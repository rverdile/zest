/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpGoBinding

import (
	"encoding/json"
)

// checks if the TagImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagImage{}

// TagImage A serializer for parsing and validating data associated with a manifest tagging.
type TagImage struct {
	// A tag name
	Tag string `json:"tag"`
	// sha256 of the Manifest file
	Digest string `json:"digest"`
}

// NewTagImage instantiates a new TagImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagImage(tag string, digest string) *TagImage {
	this := TagImage{}
	this.Tag = tag
	this.Digest = digest
	return &this
}

// NewTagImageWithDefaults instantiates a new TagImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagImageWithDefaults() *TagImage {
	this := TagImage{}
	return &this
}

// GetTag returns the Tag field value
func (o *TagImage) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *TagImage) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *TagImage) SetTag(v string) {
	o.Tag = v
}

// GetDigest returns the Digest field value
func (o *TagImage) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *TagImage) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *TagImage) SetDigest(v string) {
	o.Digest = v
}

func (o TagImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tag"] = o.Tag
	toSerialize["digest"] = o.Digest
	return toSerialize, nil
}

type NullableTagImage struct {
	value *TagImage
	isSet bool
}

func (v NullableTagImage) Get() *TagImage {
	return v.value
}

func (v *NullableTagImage) Set(val *TagImage) {
	v.value = val
	v.isSet = true
}

func (v NullableTagImage) IsSet() bool {
	return v.isSet
}

func (v *NullableTagImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagImage(val *TagImage) *NullableTagImage {
	return &NullableTagImage{value: val, isSet: true}
}

func (v NullableTagImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


