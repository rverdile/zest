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

// checks if the ImageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageResponse{}

// ImageResponse Image serializer.
type ImageResponse struct {
	// File name.
	Name string `json:"name"`
	// File path.
	Path string `json:"path"`
	// Compatible platforms.
	Platforms string `json:"platforms"`
	Artifact NullableArtifactResponse `json:"artifact"`
}

// NewImageResponse instantiates a new ImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageResponse(name string, path string, platforms string, artifact NullableArtifactResponse) *ImageResponse {
	this := ImageResponse{}
	this.Name = name
	this.Path = path
	this.Platforms = platforms
	this.Artifact = artifact
	return &this
}

// NewImageResponseWithDefaults instantiates a new ImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageResponseWithDefaults() *ImageResponse {
	this := ImageResponse{}
	return &this
}

// GetName returns the Name field value
func (o *ImageResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImageResponse) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value
func (o *ImageResponse) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ImageResponse) SetPath(v string) {
	o.Path = v
}

// GetPlatforms returns the Platforms field value
func (o *ImageResponse) GetPlatforms() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetPlatformsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platforms, true
}

// SetPlatforms sets field value
func (o *ImageResponse) SetPlatforms(v string) {
	o.Platforms = v
}

// GetArtifact returns the Artifact field value
// If the value is explicit nil, the zero value for ArtifactResponse will be returned
func (o *ImageResponse) GetArtifact() ArtifactResponse {
	if o == nil || o.Artifact.Get() == nil {
		var ret ArtifactResponse
		return ret
	}

	return *o.Artifact.Get()
}

// GetArtifactOk returns a tuple with the Artifact field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageResponse) GetArtifactOk() (*ArtifactResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Artifact.Get(), o.Artifact.IsSet()
}

// SetArtifact sets field value
func (o *ImageResponse) SetArtifact(v ArtifactResponse) {
	o.Artifact.Set(&v)
}

func (o ImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["path"] = o.Path
	toSerialize["platforms"] = o.Platforms
	toSerialize["artifact"] = o.Artifact.Get()
	return toSerialize, nil
}

type NullableImageResponse struct {
	value *ImageResponse
	isSet bool
}

func (v NullableImageResponse) Get() *ImageResponse {
	return v.value
}

func (v *NullableImageResponse) Set(val *ImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageResponse(val *ImageResponse) *NullableImageResponse {
	return &NullableImageResponse{value: val, isSet: true}
}

func (v NullableImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


