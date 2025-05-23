/*
Hyperledger Cactus Plugin - Connector Iroha V2

Can perform basic tasks on a Iroha V2 ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-iroha2

import (
	"encoding/json"
	"fmt"
)

// IrohaTransactionDefinitionV1Instruction - struct for IrohaTransactionDefinitionV1Instruction
type IrohaTransactionDefinitionV1Instruction struct {
	IrohaInstructionRequestV1 *IrohaInstructionRequestV1
	ArrayOfIrohaInstructionRequestV1 *[]IrohaInstructionRequestV1
}

// IrohaInstructionRequestV1AsIrohaTransactionDefinitionV1Instruction is a convenience function that returns IrohaInstructionRequestV1 wrapped in IrohaTransactionDefinitionV1Instruction
func IrohaInstructionRequestV1AsIrohaTransactionDefinitionV1Instruction(v *IrohaInstructionRequestV1) IrohaTransactionDefinitionV1Instruction {
	return IrohaTransactionDefinitionV1Instruction{
		IrohaInstructionRequestV1: v,
	}
}

// []IrohaInstructionRequestV1AsIrohaTransactionDefinitionV1Instruction is a convenience function that returns []IrohaInstructionRequestV1 wrapped in IrohaTransactionDefinitionV1Instruction
func ArrayOfIrohaInstructionRequestV1AsIrohaTransactionDefinitionV1Instruction(v *[]IrohaInstructionRequestV1) IrohaTransactionDefinitionV1Instruction {
	return IrohaTransactionDefinitionV1Instruction{
		ArrayOfIrohaInstructionRequestV1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IrohaTransactionDefinitionV1Instruction) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IrohaInstructionRequestV1
	err = newStrictDecoder(data).Decode(&dst.IrohaInstructionRequestV1)
	if err == nil {
		jsonIrohaInstructionRequestV1, _ := json.Marshal(dst.IrohaInstructionRequestV1)
		if string(jsonIrohaInstructionRequestV1) == "{}" { // empty struct
			dst.IrohaInstructionRequestV1 = nil
		} else {
			match++
		}
	} else {
		dst.IrohaInstructionRequestV1 = nil
	}

	// try to unmarshal data into ArrayOfIrohaInstructionRequestV1
	err = newStrictDecoder(data).Decode(&dst.ArrayOfIrohaInstructionRequestV1)
	if err == nil {
		jsonArrayOfIrohaInstructionRequestV1, _ := json.Marshal(dst.ArrayOfIrohaInstructionRequestV1)
		if string(jsonArrayOfIrohaInstructionRequestV1) == "{}" { // empty struct
			dst.ArrayOfIrohaInstructionRequestV1 = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfIrohaInstructionRequestV1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IrohaInstructionRequestV1 = nil
		dst.ArrayOfIrohaInstructionRequestV1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IrohaTransactionDefinitionV1Instruction)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IrohaTransactionDefinitionV1Instruction)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IrohaTransactionDefinitionV1Instruction) MarshalJSON() ([]byte, error) {
	if src.IrohaInstructionRequestV1 != nil {
		return json.Marshal(&src.IrohaInstructionRequestV1)
	}

	if src.ArrayOfIrohaInstructionRequestV1 != nil {
		return json.Marshal(&src.ArrayOfIrohaInstructionRequestV1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IrohaTransactionDefinitionV1Instruction) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.IrohaInstructionRequestV1 != nil {
		return obj.IrohaInstructionRequestV1
	}

	if obj.ArrayOfIrohaInstructionRequestV1 != nil {
		return obj.ArrayOfIrohaInstructionRequestV1
	}

	// all schemas are nil
	return nil
}

type NullableIrohaTransactionDefinitionV1Instruction struct {
	value *IrohaTransactionDefinitionV1Instruction
	isSet bool
}

func (v NullableIrohaTransactionDefinitionV1Instruction) Get() *IrohaTransactionDefinitionV1Instruction {
	return v.value
}

func (v *NullableIrohaTransactionDefinitionV1Instruction) Set(val *IrohaTransactionDefinitionV1Instruction) {
	v.value = val
	v.isSet = true
}

func (v NullableIrohaTransactionDefinitionV1Instruction) IsSet() bool {
	return v.isSet
}

func (v *NullableIrohaTransactionDefinitionV1Instruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIrohaTransactionDefinitionV1Instruction(val *IrohaTransactionDefinitionV1Instruction) *NullableIrohaTransactionDefinitionV1Instruction {
	return &NullableIrohaTransactionDefinitionV1Instruction{value: val, isSet: true}
}

func (v NullableIrohaTransactionDefinitionV1Instruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIrohaTransactionDefinitionV1Instruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


