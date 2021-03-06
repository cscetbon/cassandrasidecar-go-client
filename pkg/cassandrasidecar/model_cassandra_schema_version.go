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

// CassandraSchemaVersion struct for CassandraSchemaVersion
type CassandraSchemaVersion struct {
	// UUID representing a version of Cassandra node as reported by its StorageServiceMBean#getSchemaVersion
	SchemaVersion *string `json:"schemaVersion,omitempty"`
}

// NewCassandraSchemaVersion instantiates a new CassandraSchemaVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCassandraSchemaVersion() *CassandraSchemaVersion {
	this := CassandraSchemaVersion{}
	return &this
}

// NewCassandraSchemaVersionWithDefaults instantiates a new CassandraSchemaVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCassandraSchemaVersionWithDefaults() *CassandraSchemaVersion {
	this := CassandraSchemaVersion{}
	return &this
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *CassandraSchemaVersion) GetSchemaVersion() string {
	if o == nil || o.SchemaVersion == nil {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CassandraSchemaVersion) GetSchemaVersionOk() (*string, bool) {
	if o == nil || o.SchemaVersion == nil {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *CassandraSchemaVersion) HasSchemaVersion() bool {
	if o != nil && o.SchemaVersion != nil {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *CassandraSchemaVersion) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

func (o CassandraSchemaVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SchemaVersion != nil {
		toSerialize["schemaVersion"] = o.SchemaVersion
	}
	return json.Marshal(toSerialize)
}

type NullableCassandraSchemaVersion struct {
	value *CassandraSchemaVersion
	isSet bool
}

func (v NullableCassandraSchemaVersion) Get() *CassandraSchemaVersion {
	return v.value
}

func (v *NullableCassandraSchemaVersion) Set(val *CassandraSchemaVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableCassandraSchemaVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableCassandraSchemaVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCassandraSchemaVersion(val *CassandraSchemaVersion) *NullableCassandraSchemaVersion {
	return &NullableCassandraSchemaVersion{value: val, isSet: true}
}

func (v NullableCassandraSchemaVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCassandraSchemaVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
