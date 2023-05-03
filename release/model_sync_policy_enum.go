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
	"fmt"
)

// SyncPolicyEnum * `additive` - additive * `mirror_complete` - mirror_complete * `mirror_content_only` - mirror_content_only
type SyncPolicyEnum string

// List of SyncPolicyEnum
const (
	SYNCPOLICYENUM_ADDITIVE SyncPolicyEnum = "additive"
	SYNCPOLICYENUM_MIRROR_COMPLETE SyncPolicyEnum = "mirror_complete"
	SYNCPOLICYENUM_MIRROR_CONTENT_ONLY SyncPolicyEnum = "mirror_content_only"
)

// All allowed values of SyncPolicyEnum enum
var AllowedSyncPolicyEnumEnumValues = []SyncPolicyEnum{
	"additive",
	"mirror_complete",
	"mirror_content_only",
}

func (v *SyncPolicyEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SyncPolicyEnum(value)
	for _, existing := range AllowedSyncPolicyEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SyncPolicyEnum", value)
}

// NewSyncPolicyEnumFromValue returns a pointer to a valid SyncPolicyEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyncPolicyEnumFromValue(v string) (*SyncPolicyEnum, error) {
	ev := SyncPolicyEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyncPolicyEnum: valid values are %v", v, AllowedSyncPolicyEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyncPolicyEnum) IsValid() bool {
	for _, existing := range AllowedSyncPolicyEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyncPolicyEnum value
func (v SyncPolicyEnum) Ptr() *SyncPolicyEnum {
	return &v
}

type NullableSyncPolicyEnum struct {
	value *SyncPolicyEnum
	isSet bool
}

func (v NullableSyncPolicyEnum) Get() *SyncPolicyEnum {
	return v.value
}

func (v *NullableSyncPolicyEnum) Set(val *SyncPolicyEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncPolicyEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncPolicyEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncPolicyEnum(val *SyncPolicyEnum) *NullableSyncPolicyEnum {
	return &NullableSyncPolicyEnum{value: val, isSet: true}
}

func (v NullableSyncPolicyEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncPolicyEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

