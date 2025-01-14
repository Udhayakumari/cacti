/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
)

// checks if the WatchBlocksV1Progress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchBlocksV1Progress{}

// WatchBlocksV1Progress struct for WatchBlocksV1Progress
type WatchBlocksV1Progress struct {
	BlockHeader *Web3BlockHeader `json:"blockHeader,omitempty"`
	BlockData *WatchBlocksV1BlockData `json:"blockData,omitempty"`
}

// NewWatchBlocksV1Progress instantiates a new WatchBlocksV1Progress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchBlocksV1Progress() *WatchBlocksV1Progress {
	this := WatchBlocksV1Progress{}
	return &this
}

// NewWatchBlocksV1ProgressWithDefaults instantiates a new WatchBlocksV1Progress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchBlocksV1ProgressWithDefaults() *WatchBlocksV1Progress {
	this := WatchBlocksV1Progress{}
	return &this
}

// GetBlockHeader returns the BlockHeader field value if set, zero value otherwise.
func (o *WatchBlocksV1Progress) GetBlockHeader() Web3BlockHeader {
	if o == nil || IsNil(o.BlockHeader) {
		var ret Web3BlockHeader
		return ret
	}
	return *o.BlockHeader
}

// GetBlockHeaderOk returns a tuple with the BlockHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchBlocksV1Progress) GetBlockHeaderOk() (*Web3BlockHeader, bool) {
	if o == nil || IsNil(o.BlockHeader) {
		return nil, false
	}
	return o.BlockHeader, true
}

// HasBlockHeader returns a boolean if a field has been set.
func (o *WatchBlocksV1Progress) HasBlockHeader() bool {
	if o != nil && !IsNil(o.BlockHeader) {
		return true
	}

	return false
}

// SetBlockHeader gets a reference to the given Web3BlockHeader and assigns it to the BlockHeader field.
func (o *WatchBlocksV1Progress) SetBlockHeader(v Web3BlockHeader) {
	o.BlockHeader = &v
}

// GetBlockData returns the BlockData field value if set, zero value otherwise.
func (o *WatchBlocksV1Progress) GetBlockData() WatchBlocksV1BlockData {
	if o == nil || IsNil(o.BlockData) {
		var ret WatchBlocksV1BlockData
		return ret
	}
	return *o.BlockData
}

// GetBlockDataOk returns a tuple with the BlockData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchBlocksV1Progress) GetBlockDataOk() (*WatchBlocksV1BlockData, bool) {
	if o == nil || IsNil(o.BlockData) {
		return nil, false
	}
	return o.BlockData, true
}

// HasBlockData returns a boolean if a field has been set.
func (o *WatchBlocksV1Progress) HasBlockData() bool {
	if o != nil && !IsNil(o.BlockData) {
		return true
	}

	return false
}

// SetBlockData gets a reference to the given WatchBlocksV1BlockData and assigns it to the BlockData field.
func (o *WatchBlocksV1Progress) SetBlockData(v WatchBlocksV1BlockData) {
	o.BlockData = &v
}

func (o WatchBlocksV1Progress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchBlocksV1Progress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockHeader) {
		toSerialize["blockHeader"] = o.BlockHeader
	}
	if !IsNil(o.BlockData) {
		toSerialize["blockData"] = o.BlockData
	}
	return toSerialize, nil
}

type NullableWatchBlocksV1Progress struct {
	value *WatchBlocksV1Progress
	isSet bool
}

func (v NullableWatchBlocksV1Progress) Get() *WatchBlocksV1Progress {
	return v.value
}

func (v *NullableWatchBlocksV1Progress) Set(val *WatchBlocksV1Progress) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchBlocksV1Progress) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchBlocksV1Progress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchBlocksV1Progress(val *WatchBlocksV1Progress) *NullableWatchBlocksV1Progress {
	return &NullableWatchBlocksV1Progress{value: val, isSet: true}
}

func (v NullableWatchBlocksV1Progress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchBlocksV1Progress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


