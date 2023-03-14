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

// checks if the UploadDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadDetailResponse{}

// UploadDetailResponse Serializer for chunked uploads.
type UploadDetailResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// The size of the upload in bytes.
	Size int64 `json:"size"`
	// Timestamp when upload is committed.
	Completed *time.Time `json:"completed,omitempty"`
	Chunks []UploadChunkResponse `json:"chunks,omitempty"`
}

// NewUploadDetailResponse instantiates a new UploadDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadDetailResponse(size int64) *UploadDetailResponse {
	this := UploadDetailResponse{}
	this.Size = size
	return &this
}

// NewUploadDetailResponseWithDefaults instantiates a new UploadDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadDetailResponseWithDefaults() *UploadDetailResponse {
	this := UploadDetailResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *UploadDetailResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadDetailResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *UploadDetailResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *UploadDetailResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *UploadDetailResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadDetailResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *UploadDetailResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *UploadDetailResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetSize returns the Size field value
func (o *UploadDetailResponse) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *UploadDetailResponse) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *UploadDetailResponse) SetSize(v int64) {
	o.Size = v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *UploadDetailResponse) GetCompleted() time.Time {
	if o == nil || IsNil(o.Completed) {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadDetailResponse) GetCompletedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *UploadDetailResponse) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given time.Time and assigns it to the Completed field.
func (o *UploadDetailResponse) SetCompleted(v time.Time) {
	o.Completed = &v
}

// GetChunks returns the Chunks field value if set, zero value otherwise.
func (o *UploadDetailResponse) GetChunks() []UploadChunkResponse {
	if o == nil || IsNil(o.Chunks) {
		var ret []UploadChunkResponse
		return ret
	}
	return o.Chunks
}

// GetChunksOk returns a tuple with the Chunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadDetailResponse) GetChunksOk() ([]UploadChunkResponse, bool) {
	if o == nil || IsNil(o.Chunks) {
		return nil, false
	}
	return o.Chunks, true
}

// HasChunks returns a boolean if a field has been set.
func (o *UploadDetailResponse) HasChunks() bool {
	if o != nil && !IsNil(o.Chunks) {
		return true
	}

	return false
}

// SetChunks gets a reference to the given []UploadChunkResponse and assigns it to the Chunks field.
func (o *UploadDetailResponse) SetChunks(v []UploadChunkResponse) {
	o.Chunks = v
}

func (o UploadDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	toSerialize["size"] = o.Size
	// skip: completed is readOnly
	// skip: chunks is readOnly
	return toSerialize, nil
}

type NullableUploadDetailResponse struct {
	value *UploadDetailResponse
	isSet bool
}

func (v NullableUploadDetailResponse) Get() *UploadDetailResponse {
	return v.value
}

func (v *NullableUploadDetailResponse) Set(val *UploadDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadDetailResponse(val *UploadDetailResponse) *NullableUploadDetailResponse {
	return &NullableUploadDetailResponse{value: val, isSet: true}
}

func (v NullableUploadDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


