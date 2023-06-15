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

// checks if the CollectionVersionSearchListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionVersionSearchListResponse{}

// CollectionVersionSearchListResponse Cross-repo search results.
type CollectionVersionSearchListResponse struct {
	Repository RepositoryResponse `json:"repository"`
	CollectionVersion CollectionSummaryResponse `json:"collection_version"`
	RepositoryVersion *string `json:"repository_version,omitempty"`
	NamespaceMetadata NullableAnsibleAnsibleNamespaceMetadataResponse `json:"namespace_metadata"`
	IsHighest bool `json:"is_highest"`
	IsDeprecated bool `json:"is_deprecated"`
	IsSigned bool `json:"is_signed"`
}

// NewCollectionVersionSearchListResponse instantiates a new CollectionVersionSearchListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionVersionSearchListResponse(repository RepositoryResponse, collectionVersion CollectionSummaryResponse, namespaceMetadata NullableAnsibleAnsibleNamespaceMetadataResponse, isHighest bool, isDeprecated bool, isSigned bool) *CollectionVersionSearchListResponse {
	this := CollectionVersionSearchListResponse{}
	this.Repository = repository
	this.CollectionVersion = collectionVersion
	this.NamespaceMetadata = namespaceMetadata
	this.IsHighest = isHighest
	this.IsDeprecated = isDeprecated
	this.IsSigned = isSigned
	return &this
}

// NewCollectionVersionSearchListResponseWithDefaults instantiates a new CollectionVersionSearchListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionVersionSearchListResponseWithDefaults() *CollectionVersionSearchListResponse {
	this := CollectionVersionSearchListResponse{}
	return &this
}

// GetRepository returns the Repository field value
func (o *CollectionVersionSearchListResponse) GetRepository() RepositoryResponse {
	if o == nil {
		var ret RepositoryResponse
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *CollectionVersionSearchListResponse) GetRepositoryOk() (*RepositoryResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *CollectionVersionSearchListResponse) SetRepository(v RepositoryResponse) {
	o.Repository = v
}

// GetCollectionVersion returns the CollectionVersion field value
func (o *CollectionVersionSearchListResponse) GetCollectionVersion() CollectionSummaryResponse {
	if o == nil {
		var ret CollectionSummaryResponse
		return ret
	}

	return o.CollectionVersion
}

// GetCollectionVersionOk returns a tuple with the CollectionVersion field value
// and a boolean to check if the value has been set.
func (o *CollectionVersionSearchListResponse) GetCollectionVersionOk() (*CollectionSummaryResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionVersion, true
}

// SetCollectionVersion sets field value
func (o *CollectionVersionSearchListResponse) SetCollectionVersion(v CollectionSummaryResponse) {
	o.CollectionVersion = v
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise.
func (o *CollectionVersionSearchListResponse) GetRepositoryVersion() string {
	if o == nil || IsNil(o.RepositoryVersion) {
		var ret string
		return ret
	}
	return *o.RepositoryVersion
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionVersionSearchListResponse) GetRepositoryVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryVersion) {
		return nil, false
	}
	return o.RepositoryVersion, true
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *CollectionVersionSearchListResponse) HasRepositoryVersion() bool {
	if o != nil && !IsNil(o.RepositoryVersion) {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given string and assigns it to the RepositoryVersion field.
func (o *CollectionVersionSearchListResponse) SetRepositoryVersion(v string) {
	o.RepositoryVersion = &v
}

// GetNamespaceMetadata returns the NamespaceMetadata field value
// If the value is explicit nil, the zero value for AnsibleAnsibleNamespaceMetadataResponse will be returned
func (o *CollectionVersionSearchListResponse) GetNamespaceMetadata() AnsibleAnsibleNamespaceMetadataResponse {
	if o == nil || o.NamespaceMetadata.Get() == nil {
		var ret AnsibleAnsibleNamespaceMetadataResponse
		return ret
	}

	return *o.NamespaceMetadata.Get()
}

// GetNamespaceMetadataOk returns a tuple with the NamespaceMetadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionVersionSearchListResponse) GetNamespaceMetadataOk() (*AnsibleAnsibleNamespaceMetadataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.NamespaceMetadata.Get(), o.NamespaceMetadata.IsSet()
}

// SetNamespaceMetadata sets field value
func (o *CollectionVersionSearchListResponse) SetNamespaceMetadata(v AnsibleAnsibleNamespaceMetadataResponse) {
	o.NamespaceMetadata.Set(&v)
}

// GetIsHighest returns the IsHighest field value
func (o *CollectionVersionSearchListResponse) GetIsHighest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsHighest
}

// GetIsHighestOk returns a tuple with the IsHighest field value
// and a boolean to check if the value has been set.
func (o *CollectionVersionSearchListResponse) GetIsHighestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsHighest, true
}

// SetIsHighest sets field value
func (o *CollectionVersionSearchListResponse) SetIsHighest(v bool) {
	o.IsHighest = v
}

// GetIsDeprecated returns the IsDeprecated field value
func (o *CollectionVersionSearchListResponse) GetIsDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeprecated
}

// GetIsDeprecatedOk returns a tuple with the IsDeprecated field value
// and a boolean to check if the value has been set.
func (o *CollectionVersionSearchListResponse) GetIsDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeprecated, true
}

// SetIsDeprecated sets field value
func (o *CollectionVersionSearchListResponse) SetIsDeprecated(v bool) {
	o.IsDeprecated = v
}

// GetIsSigned returns the IsSigned field value
func (o *CollectionVersionSearchListResponse) GetIsSigned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSigned
}

// GetIsSignedOk returns a tuple with the IsSigned field value
// and a boolean to check if the value has been set.
func (o *CollectionVersionSearchListResponse) GetIsSignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSigned, true
}

// SetIsSigned sets field value
func (o *CollectionVersionSearchListResponse) SetIsSigned(v bool) {
	o.IsSigned = v
}

func (o CollectionVersionSearchListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionVersionSearchListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["repository"] = o.Repository
	toSerialize["collection_version"] = o.CollectionVersion
	if !IsNil(o.RepositoryVersion) {
		toSerialize["repository_version"] = o.RepositoryVersion
	}
	toSerialize["namespace_metadata"] = o.NamespaceMetadata.Get()
	toSerialize["is_highest"] = o.IsHighest
	toSerialize["is_deprecated"] = o.IsDeprecated
	toSerialize["is_signed"] = o.IsSigned
	return toSerialize, nil
}

type NullableCollectionVersionSearchListResponse struct {
	value *CollectionVersionSearchListResponse
	isSet bool
}

func (v NullableCollectionVersionSearchListResponse) Get() *CollectionVersionSearchListResponse {
	return v.value
}

func (v *NullableCollectionVersionSearchListResponse) Set(val *CollectionVersionSearchListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionVersionSearchListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionVersionSearchListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionVersionSearchListResponse(val *CollectionVersionSearchListResponse) *NullableCollectionVersionSearchListResponse {
	return &NullableCollectionVersionSearchListResponse{value: val, isSet: true}
}

func (v NullableCollectionVersionSearchListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionVersionSearchListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


