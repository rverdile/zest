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

// checks if the CollectionVersionListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionVersionListResponse{}

// CollectionVersionListResponse A serializer for a CollectionVersion list item.
type CollectionVersionListResponse struct {
	Version *string `json:"version,omitempty"`
	Href *string `json:"href,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RequiresAnsible NullableString `json:"requires_ansible,omitempty"`
}

// NewCollectionVersionListResponse instantiates a new CollectionVersionListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionVersionListResponse(createdAt time.Time, updatedAt time.Time) *CollectionVersionListResponse {
	this := CollectionVersionListResponse{}
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewCollectionVersionListResponseWithDefaults instantiates a new CollectionVersionListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionVersionListResponseWithDefaults() *CollectionVersionListResponse {
	this := CollectionVersionListResponse{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CollectionVersionListResponse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionVersionListResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CollectionVersionListResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CollectionVersionListResponse) SetVersion(v string) {
	o.Version = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *CollectionVersionListResponse) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionVersionListResponse) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *CollectionVersionListResponse) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *CollectionVersionListResponse) SetHref(v string) {
	o.Href = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CollectionVersionListResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CollectionVersionListResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CollectionVersionListResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CollectionVersionListResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CollectionVersionListResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CollectionVersionListResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetRequiresAnsible returns the RequiresAnsible field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionVersionListResponse) GetRequiresAnsible() string {
	if o == nil || IsNil(o.RequiresAnsible.Get()) {
		var ret string
		return ret
	}
	return *o.RequiresAnsible.Get()
}

// GetRequiresAnsibleOk returns a tuple with the RequiresAnsible field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionVersionListResponse) GetRequiresAnsibleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiresAnsible.Get(), o.RequiresAnsible.IsSet()
}

// HasRequiresAnsible returns a boolean if a field has been set.
func (o *CollectionVersionListResponse) HasRequiresAnsible() bool {
	if o != nil && o.RequiresAnsible.IsSet() {
		return true
	}

	return false
}

// SetRequiresAnsible gets a reference to the given NullableString and assigns it to the RequiresAnsible field.
func (o *CollectionVersionListResponse) SetRequiresAnsible(v string) {
	o.RequiresAnsible.Set(&v)
}
// SetRequiresAnsibleNil sets the value for RequiresAnsible to be an explicit nil
func (o *CollectionVersionListResponse) SetRequiresAnsibleNil() {
	o.RequiresAnsible.Set(nil)
}

// UnsetRequiresAnsible ensures that no value is present for RequiresAnsible, not even an explicit nil
func (o *CollectionVersionListResponse) UnsetRequiresAnsible() {
	o.RequiresAnsible.Unset()
}

func (o CollectionVersionListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionVersionListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: version is readOnly
	// skip: href is readOnly
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if o.RequiresAnsible.IsSet() {
		toSerialize["requires_ansible"] = o.RequiresAnsible.Get()
	}
	return toSerialize, nil
}

type NullableCollectionVersionListResponse struct {
	value *CollectionVersionListResponse
	isSet bool
}

func (v NullableCollectionVersionListResponse) Get() *CollectionVersionListResponse {
	return v.value
}

func (v *NullableCollectionVersionListResponse) Set(val *CollectionVersionListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionVersionListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionVersionListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionVersionListResponse(val *CollectionVersionListResponse) *NullableCollectionVersionListResponse {
	return &NullableCollectionVersionListResponse{value: val, isSet: true}
}

func (v NullableCollectionVersionListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionVersionListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


