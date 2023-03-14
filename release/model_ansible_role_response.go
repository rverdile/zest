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

// checks if the AnsibleRoleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleRoleResponse{}

// AnsibleRoleResponse A serializer for Role versions.
type AnsibleRoleResponse struct {
	// Artifact file representing the physical content
	Artifact string `json:"artifact"`
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	Version string `json:"version"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewAnsibleRoleResponse instantiates a new AnsibleRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleRoleResponse(artifact string, version string, name string, namespace string) *AnsibleRoleResponse {
	this := AnsibleRoleResponse{}
	this.Artifact = artifact
	this.Version = version
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewAnsibleRoleResponseWithDefaults instantiates a new AnsibleRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleRoleResponseWithDefaults() *AnsibleRoleResponse {
	this := AnsibleRoleResponse{}
	return &this
}

// GetArtifact returns the Artifact field value
func (o *AnsibleRoleResponse) GetArtifact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value
// and a boolean to check if the value has been set.
func (o *AnsibleRoleResponse) GetArtifactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Artifact, true
}

// SetArtifact sets field value
func (o *AnsibleRoleResponse) SetArtifact(v string) {
	o.Artifact = v
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *AnsibleRoleResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleRoleResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *AnsibleRoleResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *AnsibleRoleResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *AnsibleRoleResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleRoleResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *AnsibleRoleResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *AnsibleRoleResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetVersion returns the Version field value
func (o *AnsibleRoleResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AnsibleRoleResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AnsibleRoleResponse) SetVersion(v string) {
	o.Version = v
}

// GetName returns the Name field value
func (o *AnsibleRoleResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnsibleRoleResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnsibleRoleResponse) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *AnsibleRoleResponse) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *AnsibleRoleResponse) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *AnsibleRoleResponse) SetNamespace(v string) {
	o.Namespace = v
}

func (o AnsibleRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleRoleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["artifact"] = o.Artifact
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	toSerialize["version"] = o.Version
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableAnsibleRoleResponse struct {
	value *AnsibleRoleResponse
	isSet bool
}

func (v NullableAnsibleRoleResponse) Get() *AnsibleRoleResponse {
	return v.value
}

func (v *NullableAnsibleRoleResponse) Set(val *AnsibleRoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleRoleResponse(val *AnsibleRoleResponse) *NullableAnsibleRoleResponse {
	return &NullableAnsibleRoleResponse{value: val, isSet: true}
}

func (v NullableAnsibleRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


