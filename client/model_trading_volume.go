/*
GMO Coin APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: dev@tricoro.tech
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// TradingVolume struct for TradingVolume
type TradingVolume struct {
	// 0 Nomal end
	Status *int32 `json:"status,omitempty"`
	Data NullableInterface{} `json:"data,omitempty"`
	Responsetime *time.Time `json:"responsetime,omitempty"`
}

// NewTradingVolume instantiates a new TradingVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingVolume() *TradingVolume {
	this := TradingVolume{}
	return &this
}

// NewTradingVolumeWithDefaults instantiates a new TradingVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingVolumeWithDefaults() *TradingVolume {
	this := TradingVolume{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TradingVolume) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingVolume) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TradingVolume) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *TradingVolume) SetStatus(v int32) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TradingVolume) GetData() interface{} {
	if o == nil || o.Data.Get() == nil {
		var ret interface{}
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TradingVolume) GetDataOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *TradingVolume) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableInterface{} and assigns it to the Data field.
func (o *TradingVolume) SetData(v interface{}) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *TradingVolume) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *TradingVolume) UnsetData() {
	o.Data.Unset()
}

// GetResponsetime returns the Responsetime field value if set, zero value otherwise.
func (o *TradingVolume) GetResponsetime() time.Time {
	if o == nil || o.Responsetime == nil {
		var ret time.Time
		return ret
	}
	return *o.Responsetime
}

// GetResponsetimeOk returns a tuple with the Responsetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingVolume) GetResponsetimeOk() (*time.Time, bool) {
	if o == nil || o.Responsetime == nil {
		return nil, false
	}
	return o.Responsetime, true
}

// HasResponsetime returns a boolean if a field has been set.
func (o *TradingVolume) HasResponsetime() bool {
	if o != nil && o.Responsetime != nil {
		return true
	}

	return false
}

// SetResponsetime gets a reference to the given time.Time and assigns it to the Responsetime field.
func (o *TradingVolume) SetResponsetime(v time.Time) {
	o.Responsetime = &v
}

func (o TradingVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.Responsetime != nil {
		toSerialize["responsetime"] = o.Responsetime
	}
	return json.Marshal(toSerialize)
}

type NullableTradingVolume struct {
	value *TradingVolume
	isSet bool
}

func (v NullableTradingVolume) Get() *TradingVolume {
	return v.value
}

func (v *NullableTradingVolume) Set(val *TradingVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingVolume(val *TradingVolume) *NullableTradingVolume {
	return &NullableTradingVolume{value: val, isSet: true}
}

func (v NullableTradingVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


