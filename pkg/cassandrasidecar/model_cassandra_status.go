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

// CassandraStatus struct for CassandraStatus
type CassandraStatus struct {
	// State of a Cassandra node as reported from StorageServiceMBean#getOperationMode
	NodeState *string `json:"nodeState,omitempty"`
}

// NewCassandraStatus instantiates a new CassandraStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCassandraStatus() *CassandraStatus {
	this := CassandraStatus{}
	return &this
}

// NewCassandraStatusWithDefaults instantiates a new CassandraStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCassandraStatusWithDefaults() *CassandraStatus {
	this := CassandraStatus{}
	return &this
}

// GetNodeState returns the NodeState field value if set, zero value otherwise.
func (o *CassandraStatus) GetNodeState() string {
	if o == nil || o.NodeState == nil {
		var ret string
		return ret
	}
	return *o.NodeState
}

// GetNodeStateOk returns a tuple with the NodeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CassandraStatus) GetNodeStateOk() (*string, bool) {
	if o == nil || o.NodeState == nil {
		return nil, false
	}
	return o.NodeState, true
}

// HasNodeState returns a boolean if a field has been set.
func (o *CassandraStatus) HasNodeState() bool {
	if o != nil && o.NodeState != nil {
		return true
	}

	return false
}

// SetNodeState gets a reference to the given string and assigns it to the NodeState field.
func (o *CassandraStatus) SetNodeState(v string) {
	o.NodeState = &v
}

func (o CassandraStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeState != nil {
		toSerialize["nodeState"] = o.NodeState
	}
	return json.Marshal(toSerialize)
}

type NullableCassandraStatus struct {
	value *CassandraStatus
	isSet bool
}

func (v NullableCassandraStatus) Get() *CassandraStatus {
	return v.value
}

func (v *NullableCassandraStatus) Set(val *CassandraStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCassandraStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCassandraStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCassandraStatus(val *CassandraStatus) *NullableCassandraStatus {
	return &NullableCassandraStatus{value: val, isSet: true}
}

func (v NullableCassandraStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCassandraStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
