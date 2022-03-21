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
)

// TradesData struct for TradesData
type TradesData struct {
	Pagination *TradesDataPagination `json:"pagination,omitempty"`
	List []TradesDataList `json:"list,omitempty"`
}

// NewTradesData instantiates a new TradesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradesData() *TradesData {
	this := TradesData{}
	return &this
}

// NewTradesDataWithDefaults instantiates a new TradesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradesDataWithDefaults() *TradesData {
	this := TradesData{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *TradesData) GetPagination() TradesDataPagination {
	if o == nil || o.Pagination == nil {
		var ret TradesDataPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesData) GetPaginationOk() (*TradesDataPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *TradesData) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given TradesDataPagination and assigns it to the Pagination field.
func (o *TradesData) SetPagination(v TradesDataPagination) {
	o.Pagination = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *TradesData) GetList() []TradesDataList {
	if o == nil || o.List == nil {
		var ret []TradesDataList
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesData) GetListOk() ([]TradesDataList, bool) {
	if o == nil || o.List == nil {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *TradesData) HasList() bool {
	if o != nil && o.List != nil {
		return true
	}

	return false
}

// SetList gets a reference to the given []TradesDataList and assigns it to the List field.
func (o *TradesData) SetList(v []TradesDataList) {
	o.List = v
}

func (o TradesData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	if o.List != nil {
		toSerialize["list"] = o.List
	}
	return json.Marshal(toSerialize)
}

type NullableTradesData struct {
	value *TradesData
	isSet bool
}

func (v NullableTradesData) Get() *TradesData {
	return v.value
}

func (v *NullableTradesData) Set(val *TradesData) {
	v.value = val
	v.isSet = true
}

func (v NullableTradesData) IsSet() bool {
	return v.isSet
}

func (v *NullableTradesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradesData(val *TradesData) *NullableTradesData {
	return &NullableTradesData{value: val, isSet: true}
}

func (v NullableTradesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


