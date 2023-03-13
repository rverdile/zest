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

// checks if the UploadCommit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadCommit{}

// UploadCommit A mixin for validating unknown serializers' fields.
type UploadCommit struct {
	// The expected sha256 checksum for the file.
	Sha256 string `json:"sha256"`
}

// NewUploadCommit instantiates a new UploadCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadCommit(sha256 string) *UploadCommit {
	this := UploadCommit{}
	this.Sha256 = sha256
	return &this
}

// NewUploadCommitWithDefaults instantiates a new UploadCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadCommitWithDefaults() *UploadCommit {
	this := UploadCommit{}
	return &this
}

// GetSha256 returns the Sha256 field value
func (o *UploadCommit) GetSha256() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value
// and a boolean to check if the value has been set.
func (o *UploadCommit) GetSha256Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha256, true
}

// SetSha256 sets field value
func (o *UploadCommit) SetSha256(v string) {
	o.Sha256 = v
}

func (o UploadCommit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadCommit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sha256"] = o.Sha256
	return toSerialize, nil
}

type NullableUploadCommit struct {
	value *UploadCommit
	isSet bool
}

func (v NullableUploadCommit) Get() *UploadCommit {
	return v.value
}

func (v *NullableUploadCommit) Set(val *UploadCommit) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadCommit(val *UploadCommit) *NullableUploadCommit {
	return &NullableUploadCommit{value: val, isSet: true}
}

func (v NullableUploadCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


