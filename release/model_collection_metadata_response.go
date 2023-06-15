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

// checks if the CollectionMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionMetadataResponse{}

// CollectionMetadataResponse A serializer for a CollectionVersion metadata.
type CollectionMetadataResponse struct {
	Authors []string `json:"authors,omitempty"`
	Contents map[string]interface{} `json:"contents,omitempty"`
	Dependencies map[string]interface{} `json:"dependencies,omitempty"`
	Description *string `json:"description,omitempty"`
	Documentation *string `json:"documentation,omitempty"`
	Homepage *string `json:"homepage,omitempty"`
	Issues *string `json:"issues,omitempty"`
	License []string `json:"license,omitempty"`
	Repository *string `json:"repository,omitempty"`
	Tags []string `json:"tags"`
}

// NewCollectionMetadataResponse instantiates a new CollectionMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionMetadataResponse(tags []string) *CollectionMetadataResponse {
	this := CollectionMetadataResponse{}
	this.Tags = tags
	return &this
}

// NewCollectionMetadataResponseWithDefaults instantiates a new CollectionMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionMetadataResponseWithDefaults() *CollectionMetadataResponse {
	this := CollectionMetadataResponse{}
	return &this
}

// GetAuthors returns the Authors field value if set, zero value otherwise.
func (o *CollectionMetadataResponse) GetAuthors() []string {
	if o == nil || IsNil(o.Authors) {
		var ret []string
		return ret
	}
	return o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMetadataResponse) GetAuthorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Authors) {
		return nil, false
	}
	return o.Authors, true
}

// HasAuthors returns a boolean if a field has been set.
func (o *CollectionMetadataResponse) HasAuthors() bool {
	if o != nil && !IsNil(o.Authors) {
		return true
	}

	return false
}

// SetAuthors gets a reference to the given []string and assigns it to the Authors field.
func (o *CollectionMetadataResponse) SetAuthors(v []string) {
	o.Authors = v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *CollectionMetadataResponse) GetContents() map[string]interface{} {
	if o == nil || IsNil(o.Contents) {
		var ret map[string]interface{}
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMetadataResponse) GetContentsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Contents) {
		return map[string]interface{}{}, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *CollectionMetadataResponse) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given map[string]interface{} and assigns it to the Contents field.
func (o *CollectionMetadataResponse) SetContents(v map[string]interface{}) {
	o.Contents = v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *CollectionMetadataResponse) GetDependencies() map[string]interface{} {
	if o == nil || IsNil(o.Dependencies) {
		var ret map[string]interface{}
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMetadataResponse) GetDependenciesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return map[string]interface{}{}, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *CollectionMetadataResponse) HasDependencies() bool {
	if o != nil && !IsNil(o.Dependencies) {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given map[string]interface{} and assigns it to the Dependencies field.
func (o *CollectionMetadataResponse) SetDependencies(v map[string]interface{}) {
	o.Dependencies = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CollectionMetadataResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMetadataResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CollectionMetadataResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CollectionMetadataResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *CollectionMetadataResponse) GetDocumentation() string {
	if o == nil || IsNil(o.Documentation) {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMetadataResponse) GetDocumentationOk() (*string, bool) {
	if o == nil || IsNil(o.Documentation) {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *CollectionMetadataResponse) HasDocumentation() bool {
	if o != nil && !IsNil(o.Documentation) {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *CollectionMetadataResponse) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetHomepage returns the Homepage field value if set, zero value otherwise.
func (o *CollectionMetadataResponse) GetHomepage() string {
	if o == nil || IsNil(o.Homepage) {
		var ret string
		return ret
	}
	return *o.Homepage
}

// GetHomepageOk returns a tuple with the Homepage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMetadataResponse) GetHomepageOk() (*string, bool) {
	if o == nil || IsNil(o.Homepage) {
		return nil, false
	}
	return o.Homepage, true
}

// HasHomepage returns a boolean if a field has been set.
func (o *CollectionMetadataResponse) HasHomepage() bool {
	if o != nil && !IsNil(o.Homepage) {
		return true
	}

	return false
}

// SetHomepage gets a reference to the given string and assigns it to the Homepage field.
func (o *CollectionMetadataResponse) SetHomepage(v string) {
	o.Homepage = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *CollectionMetadataResponse) GetIssues() string {
	if o == nil || IsNil(o.Issues) {
		var ret string
		return ret
	}
	return *o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMetadataResponse) GetIssuesOk() (*string, bool) {
	if o == nil || IsNil(o.Issues) {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *CollectionMetadataResponse) HasIssues() bool {
	if o != nil && !IsNil(o.Issues) {
		return true
	}

	return false
}

// SetIssues gets a reference to the given string and assigns it to the Issues field.
func (o *CollectionMetadataResponse) SetIssues(v string) {
	o.Issues = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *CollectionMetadataResponse) GetLicense() []string {
	if o == nil || IsNil(o.License) {
		var ret []string
		return ret
	}
	return o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMetadataResponse) GetLicenseOk() ([]string, bool) {
	if o == nil || IsNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *CollectionMetadataResponse) HasLicense() bool {
	if o != nil && !IsNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given []string and assigns it to the License field.
func (o *CollectionMetadataResponse) SetLicense(v []string) {
	o.License = v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *CollectionMetadataResponse) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMetadataResponse) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *CollectionMetadataResponse) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *CollectionMetadataResponse) SetRepository(v string) {
	o.Repository = &v
}

// GetTags returns the Tags field value
func (o *CollectionMetadataResponse) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CollectionMetadataResponse) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *CollectionMetadataResponse) SetTags(v []string) {
	o.Tags = v
}

func (o CollectionMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authors) {
		toSerialize["authors"] = o.Authors
	}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	if !IsNil(o.Dependencies) {
		toSerialize["dependencies"] = o.Dependencies
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Documentation) {
		toSerialize["documentation"] = o.Documentation
	}
	if !IsNil(o.Homepage) {
		toSerialize["homepage"] = o.Homepage
	}
	if !IsNil(o.Issues) {
		toSerialize["issues"] = o.Issues
	}
	if !IsNil(o.License) {
		toSerialize["license"] = o.License
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

type NullableCollectionMetadataResponse struct {
	value *CollectionMetadataResponse
	isSet bool
}

func (v NullableCollectionMetadataResponse) Get() *CollectionMetadataResponse {
	return v.value
}

func (v *NullableCollectionMetadataResponse) Set(val *CollectionMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionMetadataResponse(val *CollectionMetadataResponse) *NullableCollectionMetadataResponse {
	return &NullableCollectionMetadataResponse{value: val, isSet: true}
}

func (v NullableCollectionMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


