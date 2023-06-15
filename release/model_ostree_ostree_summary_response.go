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

// checks if the OstreeOstreeSummaryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OstreeOstreeSummaryResponse{}

// OstreeOstreeSummaryResponse A Serializer class for an OSTree summary file.
type OstreeOstreeSummaryResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Artifact file representing the physical content
	Artifact string `json:"artifact"`
	// Path where the artifact is located relative to distributions base_path
	RelativePath string `json:"relative_path"`
}

// NewOstreeOstreeSummaryResponse instantiates a new OstreeOstreeSummaryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOstreeOstreeSummaryResponse(artifact string, relativePath string) *OstreeOstreeSummaryResponse {
	this := OstreeOstreeSummaryResponse{}
	this.Artifact = artifact
	this.RelativePath = relativePath
	return &this
}

// NewOstreeOstreeSummaryResponseWithDefaults instantiates a new OstreeOstreeSummaryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOstreeOstreeSummaryResponseWithDefaults() *OstreeOstreeSummaryResponse {
	this := OstreeOstreeSummaryResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *OstreeOstreeSummaryResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeSummaryResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *OstreeOstreeSummaryResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *OstreeOstreeSummaryResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *OstreeOstreeSummaryResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeSummaryResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *OstreeOstreeSummaryResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *OstreeOstreeSummaryResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetArtifact returns the Artifact field value
func (o *OstreeOstreeSummaryResponse) GetArtifact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeSummaryResponse) GetArtifactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Artifact, true
}

// SetArtifact sets field value
func (o *OstreeOstreeSummaryResponse) SetArtifact(v string) {
	o.Artifact = v
}

// GetRelativePath returns the RelativePath field value
func (o *OstreeOstreeSummaryResponse) GetRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeSummaryResponse) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativePath, true
}

// SetRelativePath sets field value
func (o *OstreeOstreeSummaryResponse) SetRelativePath(v string) {
	o.RelativePath = v
}

func (o OstreeOstreeSummaryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OstreeOstreeSummaryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	toSerialize["artifact"] = o.Artifact
	toSerialize["relative_path"] = o.RelativePath
	return toSerialize, nil
}

type NullableOstreeOstreeSummaryResponse struct {
	value *OstreeOstreeSummaryResponse
	isSet bool
}

func (v NullableOstreeOstreeSummaryResponse) Get() *OstreeOstreeSummaryResponse {
	return v.value
}

func (v *NullableOstreeOstreeSummaryResponse) Set(val *OstreeOstreeSummaryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOstreeOstreeSummaryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOstreeOstreeSummaryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOstreeOstreeSummaryResponse(val *OstreeOstreeSummaryResponse) *NullableOstreeOstreeSummaryResponse {
	return &NullableOstreeOstreeSummaryResponse{value: val, isSet: true}
}

func (v NullableOstreeOstreeSummaryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOstreeOstreeSummaryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


