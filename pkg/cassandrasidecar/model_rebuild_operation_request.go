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

// RebuildOperationRequest rebuilds data by streaming from other nodes, 
type RebuildOperationRequest struct {
	Type string `json:"type"`
	// specific keyspace to rebuild, if not specified, all keyspaces are rebuilt 
	Keyspace *string `json:"keyspace,omitempty"`
	// name of DC from which to select sources for streaming, by default, pick any DC 
	SourceDC *string `json:"sourceDC,omitempty"`
	// rebuild specific token ranges 
	SpecificTokens *[]TokenRange `json:"specificTokens,omitempty"`
	// specify hosts that this node should stream from when specificTokens are used 
	SpecificSources *[]string `json:"specificSources,omitempty"`
}

// NewRebuildOperationRequest instantiates a new RebuildOperationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRebuildOperationRequest(type_ string, ) *RebuildOperationRequest {
	this := RebuildOperationRequest{}
	this.Type = type_
	return &this
}

// NewRebuildOperationRequestWithDefaults instantiates a new RebuildOperationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRebuildOperationRequestWithDefaults() *RebuildOperationRequest {
	this := RebuildOperationRequest{}
	return &this
}

// GetType returns the Type field value
func (o *RebuildOperationRequest) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RebuildOperationRequest) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RebuildOperationRequest) SetType(v string) {
	o.Type = v
}

// GetKeyspace returns the Keyspace field value if set, zero value otherwise.
func (o *RebuildOperationRequest) GetKeyspace() string {
	if o == nil || o.Keyspace == nil {
		var ret string
		return ret
	}
	return *o.Keyspace
}

// GetKeyspaceOk returns a tuple with the Keyspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebuildOperationRequest) GetKeyspaceOk() (*string, bool) {
	if o == nil || o.Keyspace == nil {
		return nil, false
	}
	return o.Keyspace, true
}

// HasKeyspace returns a boolean if a field has been set.
func (o *RebuildOperationRequest) HasKeyspace() bool {
	if o != nil && o.Keyspace != nil {
		return true
	}

	return false
}

// SetKeyspace gets a reference to the given string and assigns it to the Keyspace field.
func (o *RebuildOperationRequest) SetKeyspace(v string) {
	o.Keyspace = &v
}

// GetSourceDC returns the SourceDC field value if set, zero value otherwise.
func (o *RebuildOperationRequest) GetSourceDC() string {
	if o == nil || o.SourceDC == nil {
		var ret string
		return ret
	}
	return *o.SourceDC
}

// GetSourceDCOk returns a tuple with the SourceDC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebuildOperationRequest) GetSourceDCOk() (*string, bool) {
	if o == nil || o.SourceDC == nil {
		return nil, false
	}
	return o.SourceDC, true
}

// HasSourceDC returns a boolean if a field has been set.
func (o *RebuildOperationRequest) HasSourceDC() bool {
	if o != nil && o.SourceDC != nil {
		return true
	}

	return false
}

// SetSourceDC gets a reference to the given string and assigns it to the SourceDC field.
func (o *RebuildOperationRequest) SetSourceDC(v string) {
	o.SourceDC = &v
}

// GetSpecificTokens returns the SpecificTokens field value if set, zero value otherwise.
func (o *RebuildOperationRequest) GetSpecificTokens() []TokenRange {
	if o == nil || o.SpecificTokens == nil {
		var ret []TokenRange
		return ret
	}
	return *o.SpecificTokens
}

// GetSpecificTokensOk returns a tuple with the SpecificTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebuildOperationRequest) GetSpecificTokensOk() (*[]TokenRange, bool) {
	if o == nil || o.SpecificTokens == nil {
		return nil, false
	}
	return o.SpecificTokens, true
}

// HasSpecificTokens returns a boolean if a field has been set.
func (o *RebuildOperationRequest) HasSpecificTokens() bool {
	if o != nil && o.SpecificTokens != nil {
		return true
	}

	return false
}

// SetSpecificTokens gets a reference to the given []TokenRange and assigns it to the SpecificTokens field.
func (o *RebuildOperationRequest) SetSpecificTokens(v []TokenRange) {
	o.SpecificTokens = &v
}

// GetSpecificSources returns the SpecificSources field value if set, zero value otherwise.
func (o *RebuildOperationRequest) GetSpecificSources() []string {
	if o == nil || o.SpecificSources == nil {
		var ret []string
		return ret
	}
	return *o.SpecificSources
}

// GetSpecificSourcesOk returns a tuple with the SpecificSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebuildOperationRequest) GetSpecificSourcesOk() (*[]string, bool) {
	if o == nil || o.SpecificSources == nil {
		return nil, false
	}
	return o.SpecificSources, true
}

// HasSpecificSources returns a boolean if a field has been set.
func (o *RebuildOperationRequest) HasSpecificSources() bool {
	if o != nil && o.SpecificSources != nil {
		return true
	}

	return false
}

// SetSpecificSources gets a reference to the given []string and assigns it to the SpecificSources field.
func (o *RebuildOperationRequest) SetSpecificSources(v []string) {
	o.SpecificSources = &v
}

func (o RebuildOperationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Keyspace != nil {
		toSerialize["keyspace"] = o.Keyspace
	}
	if o.SourceDC != nil {
		toSerialize["sourceDC"] = o.SourceDC
	}
	if o.SpecificTokens != nil {
		toSerialize["specificTokens"] = o.SpecificTokens
	}
	if o.SpecificSources != nil {
		toSerialize["specificSources"] = o.SpecificSources
	}
	return json.Marshal(toSerialize)
}

type NullableRebuildOperationRequest struct {
	value *RebuildOperationRequest
	isSet bool
}

func (v NullableRebuildOperationRequest) Get() *RebuildOperationRequest {
	return v.value
}

func (v *NullableRebuildOperationRequest) Set(val *RebuildOperationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRebuildOperationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRebuildOperationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRebuildOperationRequest(val *RebuildOperationRequest) *NullableRebuildOperationRequest {
	return &NullableRebuildOperationRequest{value: val, isSet: true}
}

func (v NullableRebuildOperationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRebuildOperationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
