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

// checks if the GroupUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupUserResponse{}

// GroupUserResponse Serializer for Users that belong to a Group.
type GroupUserResponse struct {
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
	PulpHref *string `json:"pulp_href,omitempty"`
}

// NewGroupUserResponse instantiates a new GroupUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUserResponse(username string) *GroupUserResponse {
	this := GroupUserResponse{}
	this.Username = username
	return &this
}

// NewGroupUserResponseWithDefaults instantiates a new GroupUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUserResponseWithDefaults() *GroupUserResponse {
	this := GroupUserResponse{}
	return &this
}

// GetUsername returns the Username field value
func (o *GroupUserResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *GroupUserResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *GroupUserResponse) SetUsername(v string) {
	o.Username = v
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *GroupUserResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUserResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *GroupUserResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *GroupUserResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

func (o GroupUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	// skip: pulp_href is readOnly
	return toSerialize, nil
}

type NullableGroupUserResponse struct {
	value *GroupUserResponse
	isSet bool
}

func (v NullableGroupUserResponse) Get() *GroupUserResponse {
	return v.value
}

func (v *NullableGroupUserResponse) Set(val *GroupUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUserResponse(val *GroupUserResponse) *NullableGroupUserResponse {
	return &NullableGroupUserResponse{value: val, isSet: true}
}

func (v NullableGroupUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


