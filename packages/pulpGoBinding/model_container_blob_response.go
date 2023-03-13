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
	"time"
)

// checks if the ContainerBlobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerBlobResponse{}

// ContainerBlobResponse Serializer for Blobs.
type ContainerBlobResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Artifact file representing the physical content
	Artifact string `json:"artifact"`
	// sha256 of the Blob file
	Digest string `json:"digest"`
}

// NewContainerBlobResponse instantiates a new ContainerBlobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerBlobResponse(artifact string, digest string) *ContainerBlobResponse {
	this := ContainerBlobResponse{}
	this.Artifact = artifact
	this.Digest = digest
	return &this
}

// NewContainerBlobResponseWithDefaults instantiates a new ContainerBlobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerBlobResponseWithDefaults() *ContainerBlobResponse {
	this := ContainerBlobResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *ContainerBlobResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerBlobResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *ContainerBlobResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *ContainerBlobResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *ContainerBlobResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerBlobResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *ContainerBlobResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *ContainerBlobResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetArtifact returns the Artifact field value
func (o *ContainerBlobResponse) GetArtifact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value
// and a boolean to check if the value has been set.
func (o *ContainerBlobResponse) GetArtifactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Artifact, true
}

// SetArtifact sets field value
func (o *ContainerBlobResponse) SetArtifact(v string) {
	o.Artifact = v
}

// GetDigest returns the Digest field value
func (o *ContainerBlobResponse) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *ContainerBlobResponse) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *ContainerBlobResponse) SetDigest(v string) {
	o.Digest = v
}

func (o ContainerBlobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerBlobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	toSerialize["artifact"] = o.Artifact
	toSerialize["digest"] = o.Digest
	return toSerialize, nil
}

type NullableContainerBlobResponse struct {
	value *ContainerBlobResponse
	isSet bool
}

func (v NullableContainerBlobResponse) Get() *ContainerBlobResponse {
	return v.value
}

func (v *NullableContainerBlobResponse) Set(val *ContainerBlobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerBlobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerBlobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerBlobResponse(val *ContainerBlobResponse) *NullableContainerBlobResponse {
	return &NullableContainerBlobResponse{value: val, isSet: true}
}

func (v NullableContainerBlobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerBlobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


