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

// DecommissionOperationRequest decommissions a Cassandra node 
type DecommissionOperationRequest struct {
	Type string `json:"type"`
	// forcibly decommission a node, even if by doing so there will not be enough replicas for responding to requests, this option is relevant only for Cassandra 4.x and it is not in use for lower versions, defaults to false. 
	Force *bool `json:"force,omitempty"`
}

// NewDecommissionOperationRequest instantiates a new DecommissionOperationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecommissionOperationRequest(type_ string, ) *DecommissionOperationRequest {
	this := DecommissionOperationRequest{}
	this.Type = type_
	return &this
}

// NewDecommissionOperationRequestWithDefaults instantiates a new DecommissionOperationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecommissionOperationRequestWithDefaults() *DecommissionOperationRequest {
	this := DecommissionOperationRequest{}
	return &this
}

// GetType returns the Type field value
func (o *DecommissionOperationRequest) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DecommissionOperationRequest) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DecommissionOperationRequest) SetType(v string) {
	o.Type = v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *DecommissionOperationRequest) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecommissionOperationRequest) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *DecommissionOperationRequest) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *DecommissionOperationRequest) SetForce(v bool) {
	o.Force = &v
}

func (o DecommissionOperationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	return json.Marshal(toSerialize)
}

type NullableDecommissionOperationRequest struct {
	value *DecommissionOperationRequest
	isSet bool
}

func (v NullableDecommissionOperationRequest) Get() *DecommissionOperationRequest {
	return v.value
}

func (v *NullableDecommissionOperationRequest) Set(val *DecommissionOperationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDecommissionOperationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDecommissionOperationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecommissionOperationRequest(val *DecommissionOperationRequest) *NullableDecommissionOperationRequest {
	return &NullableDecommissionOperationRequest{value: val, isSet: true}
}

func (v NullableDecommissionOperationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecommissionOperationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
