/*
Hyperledger Cacti Plugin - Connector CDL

Can perform basic tasks on Fujitsu CDL service.

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-cdl

import (
	"encoding/json"
	"fmt"
)

// GetLineageOptionDirectionV1 the model 'GetLineageOptionDirectionV1'
type GetLineageOptionDirectionV1 string

// List of GetLineageOptionDirectionV1
const (
	Backward GetLineageOptionDirectionV1 = "backward"
	Forward GetLineageOptionDirectionV1 = "forward"
	Both GetLineageOptionDirectionV1 = "both"
)

// All allowed values of GetLineageOptionDirectionV1 enum
var AllowedGetLineageOptionDirectionV1EnumValues = []GetLineageOptionDirectionV1{
	"backward",
	"forward",
	"both",
}

func (v *GetLineageOptionDirectionV1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetLineageOptionDirectionV1(value)
	for _, existing := range AllowedGetLineageOptionDirectionV1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetLineageOptionDirectionV1", value)
}

// NewGetLineageOptionDirectionV1FromValue returns a pointer to a valid GetLineageOptionDirectionV1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetLineageOptionDirectionV1FromValue(v string) (*GetLineageOptionDirectionV1, error) {
	ev := GetLineageOptionDirectionV1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetLineageOptionDirectionV1: valid values are %v", v, AllowedGetLineageOptionDirectionV1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetLineageOptionDirectionV1) IsValid() bool {
	for _, existing := range AllowedGetLineageOptionDirectionV1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetLineageOptionDirectionV1 value
func (v GetLineageOptionDirectionV1) Ptr() *GetLineageOptionDirectionV1 {
	return &v
}

type NullableGetLineageOptionDirectionV1 struct {
	value *GetLineageOptionDirectionV1
	isSet bool
}

func (v NullableGetLineageOptionDirectionV1) Get() *GetLineageOptionDirectionV1 {
	return v.value
}

func (v *NullableGetLineageOptionDirectionV1) Set(val *GetLineageOptionDirectionV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLineageOptionDirectionV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLineageOptionDirectionV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLineageOptionDirectionV1(val *GetLineageOptionDirectionV1) *NullableGetLineageOptionDirectionV1 {
	return &NullableGetLineageOptionDirectionV1{value: val, isSet: true}
}

func (v NullableGetLineageOptionDirectionV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLineageOptionDirectionV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

