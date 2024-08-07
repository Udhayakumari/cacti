/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
)

// checks if the GatewayOptionsWallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayOptionsWallet{}

// GatewayOptionsWallet struct for GatewayOptionsWallet
type GatewayOptionsWallet struct {
	Keychain *FabricSigningCredential `json:"keychain,omitempty"`
	Json *string `json:"json,omitempty"`
}

// NewGatewayOptionsWallet instantiates a new GatewayOptionsWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayOptionsWallet() *GatewayOptionsWallet {
	this := GatewayOptionsWallet{}
	return &this
}

// NewGatewayOptionsWalletWithDefaults instantiates a new GatewayOptionsWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayOptionsWalletWithDefaults() *GatewayOptionsWallet {
	this := GatewayOptionsWallet{}
	return &this
}

// GetKeychain returns the Keychain field value if set, zero value otherwise.
func (o *GatewayOptionsWallet) GetKeychain() FabricSigningCredential {
	if o == nil || IsNil(o.Keychain) {
		var ret FabricSigningCredential
		return ret
	}
	return *o.Keychain
}

// GetKeychainOk returns a tuple with the Keychain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayOptionsWallet) GetKeychainOk() (*FabricSigningCredential, bool) {
	if o == nil || IsNil(o.Keychain) {
		return nil, false
	}
	return o.Keychain, true
}

// HasKeychain returns a boolean if a field has been set.
func (o *GatewayOptionsWallet) HasKeychain() bool {
	if o != nil && !IsNil(o.Keychain) {
		return true
	}

	return false
}

// SetKeychain gets a reference to the given FabricSigningCredential and assigns it to the Keychain field.
func (o *GatewayOptionsWallet) SetKeychain(v FabricSigningCredential) {
	o.Keychain = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayOptionsWallet) GetJson() string {
	if o == nil || IsNil(o.Json) {
		var ret string
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayOptionsWallet) GetJsonOk() (*string, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayOptionsWallet) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given string and assigns it to the Json field.
func (o *GatewayOptionsWallet) SetJson(v string) {
	o.Json = &v
}

func (o GatewayOptionsWallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayOptionsWallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Keychain) {
		toSerialize["keychain"] = o.Keychain
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	return toSerialize, nil
}

type NullableGatewayOptionsWallet struct {
	value *GatewayOptionsWallet
	isSet bool
}

func (v NullableGatewayOptionsWallet) Get() *GatewayOptionsWallet {
	return v.value
}

func (v *NullableGatewayOptionsWallet) Set(val *GatewayOptionsWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayOptionsWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayOptionsWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayOptionsWallet(val *GatewayOptionsWallet) *NullableGatewayOptionsWallet {
	return &NullableGatewayOptionsWallet{value: val, isSet: true}
}

func (v NullableGatewayOptionsWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayOptionsWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

