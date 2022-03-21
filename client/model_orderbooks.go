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

// Orderbooks struct for Orderbooks
type Orderbooks struct {
	// 0 Nomal end
	Status *int32 `json:"status,omitempty"`
	Data *OrderbooksData `json:"data,omitempty"`
	Responsetime *time.Time `json:"responsetime,omitempty"`
}

// NewOrderbooks instantiates a new Orderbooks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderbooks() *Orderbooks {
	this := Orderbooks{}
	return &this
}

// NewOrderbooksWithDefaults instantiates a new Orderbooks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderbooksWithDefaults() *Orderbooks {
	this := Orderbooks{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Orderbooks) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Orderbooks) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Orderbooks) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *Orderbooks) SetStatus(v int32) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Orderbooks) GetData() OrderbooksData {
	if o == nil || o.Data == nil {
		var ret OrderbooksData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Orderbooks) GetDataOk() (*OrderbooksData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Orderbooks) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given OrderbooksData and assigns it to the Data field.
func (o *Orderbooks) SetData(v OrderbooksData) {
	o.Data = &v
}

// GetResponsetime returns the Responsetime field value if set, zero value otherwise.
func (o *Orderbooks) GetResponsetime() time.Time {
	if o == nil || o.Responsetime == nil {
		var ret time.Time
		return ret
	}
	return *o.Responsetime
}

// GetResponsetimeOk returns a tuple with the Responsetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Orderbooks) GetResponsetimeOk() (*time.Time, bool) {
	if o == nil || o.Responsetime == nil {
		return nil, false
	}
	return o.Responsetime, true
}

// HasResponsetime returns a boolean if a field has been set.
func (o *Orderbooks) HasResponsetime() bool {
	if o != nil && o.Responsetime != nil {
		return true
	}

	return false
}

// SetResponsetime gets a reference to the given time.Time and assigns it to the Responsetime field.
func (o *Orderbooks) SetResponsetime(v time.Time) {
	o.Responsetime = &v
}

func (o Orderbooks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Responsetime != nil {
		toSerialize["responsetime"] = o.Responsetime
	}
	return json.Marshal(toSerialize)
}

type NullableOrderbooks struct {
	value *Orderbooks
	isSet bool
}

func (v NullableOrderbooks) Get() *Orderbooks {
	return v.value
}

func (v *NullableOrderbooks) Set(val *Orderbooks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderbooks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderbooks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderbooks(val *Orderbooks) *NullableOrderbooks {
	return &NullableOrderbooks{value: val, isSet: true}
}

func (v NullableOrderbooks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderbooks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


