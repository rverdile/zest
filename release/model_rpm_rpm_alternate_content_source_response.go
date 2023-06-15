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

// checks if the RpmRpmAlternateContentSourceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmRpmAlternateContentSourceResponse{}

// RpmRpmAlternateContentSourceResponse Serializer for RPM alternate content source.
type RpmRpmAlternateContentSourceResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Name of Alternate Content Source.
	Name string `json:"name"`
	// Date of last refresh of AlternateContentSource.
	LastRefreshed NullableTime `json:"last_refreshed,omitempty"`
	// List of paths that will be appended to the Remote url when searching for content.
	Paths []string `json:"paths,omitempty"`
	// The remote to provide alternate content source.
	Remote string `json:"remote"`
}

// NewRpmRpmAlternateContentSourceResponse instantiates a new RpmRpmAlternateContentSourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmRpmAlternateContentSourceResponse(name string, remote string) *RpmRpmAlternateContentSourceResponse {
	this := RpmRpmAlternateContentSourceResponse{}
	this.Name = name
	this.Remote = remote
	return &this
}

// NewRpmRpmAlternateContentSourceResponseWithDefaults instantiates a new RpmRpmAlternateContentSourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmRpmAlternateContentSourceResponseWithDefaults() *RpmRpmAlternateContentSourceResponse {
	this := RpmRpmAlternateContentSourceResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *RpmRpmAlternateContentSourceResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmAlternateContentSourceResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *RpmRpmAlternateContentSourceResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *RpmRpmAlternateContentSourceResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *RpmRpmAlternateContentSourceResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmAlternateContentSourceResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *RpmRpmAlternateContentSourceResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *RpmRpmAlternateContentSourceResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value
func (o *RpmRpmAlternateContentSourceResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RpmRpmAlternateContentSourceResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RpmRpmAlternateContentSourceResponse) SetName(v string) {
	o.Name = v
}

// GetLastRefreshed returns the LastRefreshed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmAlternateContentSourceResponse) GetLastRefreshed() time.Time {
	if o == nil || IsNil(o.LastRefreshed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastRefreshed.Get()
}

// GetLastRefreshedOk returns a tuple with the LastRefreshed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmAlternateContentSourceResponse) GetLastRefreshedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastRefreshed.Get(), o.LastRefreshed.IsSet()
}

// HasLastRefreshed returns a boolean if a field has been set.
func (o *RpmRpmAlternateContentSourceResponse) HasLastRefreshed() bool {
	if o != nil && o.LastRefreshed.IsSet() {
		return true
	}

	return false
}

// SetLastRefreshed gets a reference to the given NullableTime and assigns it to the LastRefreshed field.
func (o *RpmRpmAlternateContentSourceResponse) SetLastRefreshed(v time.Time) {
	o.LastRefreshed.Set(&v)
}
// SetLastRefreshedNil sets the value for LastRefreshed to be an explicit nil
func (o *RpmRpmAlternateContentSourceResponse) SetLastRefreshedNil() {
	o.LastRefreshed.Set(nil)
}

// UnsetLastRefreshed ensures that no value is present for LastRefreshed, not even an explicit nil
func (o *RpmRpmAlternateContentSourceResponse) UnsetLastRefreshed() {
	o.LastRefreshed.Unset()
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *RpmRpmAlternateContentSourceResponse) GetPaths() []string {
	if o == nil || IsNil(o.Paths) {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmAlternateContentSourceResponse) GetPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.Paths) {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *RpmRpmAlternateContentSourceResponse) HasPaths() bool {
	if o != nil && !IsNil(o.Paths) {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *RpmRpmAlternateContentSourceResponse) SetPaths(v []string) {
	o.Paths = v
}

// GetRemote returns the Remote field value
func (o *RpmRpmAlternateContentSourceResponse) GetRemote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value
// and a boolean to check if the value has been set.
func (o *RpmRpmAlternateContentSourceResponse) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remote, true
}

// SetRemote sets field value
func (o *RpmRpmAlternateContentSourceResponse) SetRemote(v string) {
	o.Remote = v
}

func (o RpmRpmAlternateContentSourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmRpmAlternateContentSourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	toSerialize["name"] = o.Name
	if o.LastRefreshed.IsSet() {
		toSerialize["last_refreshed"] = o.LastRefreshed.Get()
	}
	if !IsNil(o.Paths) {
		toSerialize["paths"] = o.Paths
	}
	toSerialize["remote"] = o.Remote
	return toSerialize, nil
}

type NullableRpmRpmAlternateContentSourceResponse struct {
	value *RpmRpmAlternateContentSourceResponse
	isSet bool
}

func (v NullableRpmRpmAlternateContentSourceResponse) Get() *RpmRpmAlternateContentSourceResponse {
	return v.value
}

func (v *NullableRpmRpmAlternateContentSourceResponse) Set(val *RpmRpmAlternateContentSourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmRpmAlternateContentSourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmRpmAlternateContentSourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmRpmAlternateContentSourceResponse(val *RpmRpmAlternateContentSourceResponse) *NullableRpmRpmAlternateContentSourceResponse {
	return &NullableRpmRpmAlternateContentSourceResponse{value: val, isSet: true}
}

func (v NullableRpmRpmAlternateContentSourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmRpmAlternateContentSourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


