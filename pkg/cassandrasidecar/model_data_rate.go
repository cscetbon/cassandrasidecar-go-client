/*
 * Instaclustr Cassandra Sidecar
 *
 * REST API for Cassandra Sidecar from Instaclustr
 *
 * API version: 1.1.0
 * Contact: support@instaclustr.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cassandrasidecar

import (
	"encoding/json"
)

// DataRate bandwidth used during uploads 
type DataRate struct {
	// quantified value of bandwidth, an integer 
	Value *int32 `json:"value,omitempty"`
	// unit of 'data bandwidth' 
	Unit string `json:"unit"`
}

// NewDataRate instantiates a new DataRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataRate(unit string, ) *DataRate {
	this := DataRate{}
	this.Unit = unit
	return &this
}

// NewDataRateWithDefaults instantiates a new DataRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataRateWithDefaults() *DataRate {
	this := DataRate{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DataRate) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRate) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DataRate) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *DataRate) SetValue(v int32) {
	o.Value = &v
}

// GetUnit returns the Unit field value
func (o *DataRate) GetUnit() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *DataRate) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *DataRate) SetUnit(v string) {
	o.Unit = v
}

func (o DataRate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableDataRate struct {
	value *DataRate
	isSet bool
}

func (v NullableDataRate) Get() *DataRate {
	return v.value
}

func (v *NullableDataRate) Set(val *DataRate) {
	v.value = val
	v.isSet = true
}

func (v NullableDataRate) IsSet() bool {
	return v.isSet
}

func (v *NullableDataRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataRate(val *DataRate) *NullableDataRate {
	return &NullableDataRate{value: val, isSet: true}
}

func (v NullableDataRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
