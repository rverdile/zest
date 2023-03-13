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

// checks if the UnpaginatedCollectionVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnpaginatedCollectionVersionResponse{}

// UnpaginatedCollectionVersionResponse A serializer for unpaginated CollectionVersion.
type UnpaginatedCollectionVersionResponse struct {
	Version *string `json:"version,omitempty"`
	Href *string `json:"href,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RequiresAnsible NullableString `json:"requires_ansible,omitempty"`
	Artifact *ArtifactRefResponse `json:"artifact,omitempty"`
	Collection *CollectionRefResponse `json:"collection,omitempty"`
	DownloadUrl *string `json:"download_url,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *CollectionNamespaceResponse `json:"namespace,omitempty"`
	Signatures *string `json:"signatures,omitempty"`
	Metadata *CollectionMetadataResponse `json:"metadata,omitempty"`
	GitUrl *string `json:"git_url,omitempty"`
	GitCommitSha *string `json:"git_commit_sha,omitempty"`
}

// NewUnpaginatedCollectionVersionResponse instantiates a new UnpaginatedCollectionVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnpaginatedCollectionVersionResponse(createdAt time.Time, updatedAt time.Time) *UnpaginatedCollectionVersionResponse {
	this := UnpaginatedCollectionVersionResponse{}
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewUnpaginatedCollectionVersionResponseWithDefaults instantiates a new UnpaginatedCollectionVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnpaginatedCollectionVersionResponseWithDefaults() *UnpaginatedCollectionVersionResponse {
	this := UnpaginatedCollectionVersionResponse{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UnpaginatedCollectionVersionResponse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnpaginatedCollectionVersionResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UnpaginatedCollectionVersionResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *UnpaginatedCollectionVersionResponse) SetVersion(v string) {
	o.Version = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *UnpaginatedCollectionVersionResponse) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnpaginatedCollectionVersionResponse) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *UnpaginatedCollectionVersionResponse) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *UnpaginatedCollectionVersionResponse) SetHref(v string) {
	o.Href = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *UnpaginatedCollectionVersionResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UnpaginatedCollectionVersionResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UnpaginatedCollectionVersionResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *UnpaginatedCollectionVersionResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UnpaginatedCollectionVersionResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *UnpaginatedCollectionVersionResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetRequiresAnsible returns the RequiresAnsible field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnpaginatedCollectionVersionResponse) GetRequiresAnsible() string {
	if o == nil || IsNil(o.RequiresAnsible.Get()) {
		var ret string
		return ret
	}
	return *o.RequiresAnsible.Get()
}

// GetRequiresAnsibleOk returns a tuple with the RequiresAnsible field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnpaginatedCollectionVersionResponse) GetRequiresAnsibleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiresAnsible.Get(), o.RequiresAnsible.IsSet()
}

// HasRequiresAnsible returns a boolean if a field has been set.
func (o *UnpaginatedCollectionVersionResponse) HasRequiresAnsible() bool {
	if o != nil && o.RequiresAnsible.IsSet() {
		return true
	}

	return false
}

// SetRequiresAnsible gets a reference to the given NullableString and assigns it to the RequiresAnsible field.
func (o *UnpaginatedCollectionVersionResponse) SetRequiresAnsible(v string) {
	o.RequiresAnsible.Set(&v)
}
// SetRequiresAnsibleNil sets the value for RequiresAnsible to be an explicit nil
func (o *UnpaginatedCollectionVersionResponse) SetRequiresAnsibleNil() {
	o.RequiresAnsible.Set(nil)
}

// UnsetRequiresAnsible ensures that no value is present for RequiresAnsible, not even an explicit nil
func (o *UnpaginatedCollectionVersionResponse) UnsetRequiresAnsible() {
	o.RequiresAnsible.Unset()
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *UnpaginatedCollectionVersionResponse) GetArtifact() ArtifactRefResponse {
	if o == nil || IsNil(o.Artifact) {
		var ret ArtifactRefResponse
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnpaginatedCollectionVersionResponse) GetArtifactOk() (*ArtifactRefResponse, bool) {
	if o == nil || IsNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *UnpaginatedCollectionVersionResponse) HasArtifact() bool {
	if o != nil && !IsNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given ArtifactRefResponse and assigns it to the Artifact field.
func (o *UnpaginatedCollectionVersionResponse) SetArtifact(v ArtifactRefResponse) {
	o.Artifact = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *UnpaginatedCollectionVersionResponse) GetCollection() CollectionRefResponse {
	if o == nil || IsNil(o.Collection) {
		var ret CollectionRefResponse
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnpaginatedCollectionVersionResponse) GetCollectionOk() (*CollectionRefResponse, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *UnpaginatedCollectionVersionResponse) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given CollectionRefResponse and assigns it to the Collection field.
func (o *UnpaginatedCollectionVersionResponse) SetCollection(v CollectionRefResponse) {
	o.Collection = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *UnpaginatedCollectionVersionResponse) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnpaginatedCollectionVersionResponse) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *UnpaginatedCollectionVersionResponse) HasDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *UnpaginatedCollectionVersionResponse) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnpaginatedCollectionVersionResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnpaginatedCollectionVersionResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnpaginatedCollectionVersionResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnpaginatedCollectionVersionResponse) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *UnpaginatedCollectionVersionResponse) GetNamespace() CollectionNamespaceResponse {
	if o == nil || IsNil(o.Namespace) {
		var ret CollectionNamespaceResponse
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnpaginatedCollectionVersionResponse) GetNamespaceOk() (*CollectionNamespaceResponse, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *UnpaginatedCollectionVersionResponse) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given CollectionNamespaceResponse and assigns it to the Namespace field.
func (o *UnpaginatedCollectionVersionResponse) SetNamespace(v CollectionNamespaceResponse) {
	o.Namespace = &v
}

// GetSignatures returns the Signatures field value if set, zero value otherwise.
func (o *UnpaginatedCollectionVersionResponse) GetSignatures() string {
	if o == nil || IsNil(o.Signatures) {
		var ret string
		return ret
	}
	return *o.Signatures
}

// GetSignaturesOk returns a tuple with the Signatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnpaginatedCollectionVersionResponse) GetSignaturesOk() (*string, bool) {
	if o == nil || IsNil(o.Signatures) {
		return nil, false
	}
	return o.Signatures, true
}

// HasSignatures returns a boolean if a field has been set.
func (o *UnpaginatedCollectionVersionResponse) HasSignatures() bool {
	if o != nil && !IsNil(o.Signatures) {
		return true
	}

	return false
}

// SetSignatures gets a reference to the given string and assigns it to the Signatures field.
func (o *UnpaginatedCollectionVersionResponse) SetSignatures(v string) {
	o.Signatures = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnpaginatedCollectionVersionResponse) GetMetadata() CollectionMetadataResponse {
	if o == nil || IsNil(o.Metadata) {
		var ret CollectionMetadataResponse
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnpaginatedCollectionVersionResponse) GetMetadataOk() (*CollectionMetadataResponse, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnpaginatedCollectionVersionResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given CollectionMetadataResponse and assigns it to the Metadata field.
func (o *UnpaginatedCollectionVersionResponse) SetMetadata(v CollectionMetadataResponse) {
	o.Metadata = &v
}

// GetGitUrl returns the GitUrl field value if set, zero value otherwise.
func (o *UnpaginatedCollectionVersionResponse) GetGitUrl() string {
	if o == nil || IsNil(o.GitUrl) {
		var ret string
		return ret
	}
	return *o.GitUrl
}

// GetGitUrlOk returns a tuple with the GitUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnpaginatedCollectionVersionResponse) GetGitUrlOk() (*string, bool) {
	if o == nil || IsNil(o.GitUrl) {
		return nil, false
	}
	return o.GitUrl, true
}

// HasGitUrl returns a boolean if a field has been set.
func (o *UnpaginatedCollectionVersionResponse) HasGitUrl() bool {
	if o != nil && !IsNil(o.GitUrl) {
		return true
	}

	return false
}

// SetGitUrl gets a reference to the given string and assigns it to the GitUrl field.
func (o *UnpaginatedCollectionVersionResponse) SetGitUrl(v string) {
	o.GitUrl = &v
}

// GetGitCommitSha returns the GitCommitSha field value if set, zero value otherwise.
func (o *UnpaginatedCollectionVersionResponse) GetGitCommitSha() string {
	if o == nil || IsNil(o.GitCommitSha) {
		var ret string
		return ret
	}
	return *o.GitCommitSha
}

// GetGitCommitShaOk returns a tuple with the GitCommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnpaginatedCollectionVersionResponse) GetGitCommitShaOk() (*string, bool) {
	if o == nil || IsNil(o.GitCommitSha) {
		return nil, false
	}
	return o.GitCommitSha, true
}

// HasGitCommitSha returns a boolean if a field has been set.
func (o *UnpaginatedCollectionVersionResponse) HasGitCommitSha() bool {
	if o != nil && !IsNil(o.GitCommitSha) {
		return true
	}

	return false
}

// SetGitCommitSha gets a reference to the given string and assigns it to the GitCommitSha field.
func (o *UnpaginatedCollectionVersionResponse) SetGitCommitSha(v string) {
	o.GitCommitSha = &v
}

func (o UnpaginatedCollectionVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnpaginatedCollectionVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: version is readOnly
	// skip: href is readOnly
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if o.RequiresAnsible.IsSet() {
		toSerialize["requires_ansible"] = o.RequiresAnsible.Get()
	}
	// skip: artifact is readOnly
	// skip: collection is readOnly
	// skip: download_url is readOnly
	// skip: name is readOnly
	// skip: namespace is readOnly
	// skip: signatures is readOnly
	// skip: metadata is readOnly
	// skip: git_url is readOnly
	// skip: git_commit_sha is readOnly
	return toSerialize, nil
}

type NullableUnpaginatedCollectionVersionResponse struct {
	value *UnpaginatedCollectionVersionResponse
	isSet bool
}

func (v NullableUnpaginatedCollectionVersionResponse) Get() *UnpaginatedCollectionVersionResponse {
	return v.value
}

func (v *NullableUnpaginatedCollectionVersionResponse) Set(val *UnpaginatedCollectionVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUnpaginatedCollectionVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUnpaginatedCollectionVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnpaginatedCollectionVersionResponse(val *UnpaginatedCollectionVersionResponse) *NullableUnpaginatedCollectionVersionResponse {
	return &NullableUnpaginatedCollectionVersionResponse{value: val, isSet: true}
}

func (v NullableUnpaginatedCollectionVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnpaginatedCollectionVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


