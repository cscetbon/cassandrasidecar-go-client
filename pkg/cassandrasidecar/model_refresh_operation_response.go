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
	"time"
)

// RefreshOperationResponse struct for RefreshOperationResponse
type RefreshOperationResponse struct {
	Type string `json:"type"`
	// unique identifier of an operation, a random id is assigned to each operation after a request is submitted, from caller's perspective, an id is sent back as a response to his request so he can further query state of that operation, referencing id, by operations/{id} endpoint 
	Id string `json:"id"`
	// state of an operation, operation might be in various states, PENDING - this operation is pending for being submitted. RUNNING - this operation is actively doing its job, COMPLETED - this operation has finished successfully, CANCELLED - this operation was interrupted while being run, FAILED - this operation has finished errorneously 
	State string `json:"state"`
	// float from 0.0 to 1.0, 1.0 telling that operation is completed, either successfully or with errors. 
	Progress float32 `json:"progress"`
	// timestamp telling when this operation was created on Sidecar's side 
	CreationTime time.Time `json:"creationTime"`
	// timestamp telling when this operation was started by Sidecar, if an operation is created, it does not necessarily mean that it will be started right away, in most cases it is the case but if e.g. ExecutorService is full on its working thread, an execution of an operation is postponed and start time is updated only after that 
	StartTime *time.Time `json:"startTime,omitempty"`
	// timestamp telling when an operation has finished, irrelevant of its result, an operation can be failed and it would still have this field populated. 
	CompletionTime *time.Time `json:"completionTime,omitempty"`
	// This field contains serialized java.lang.Throwable in case this operation has failed 
	FailureCause *map[string]interface{} `json:"failureCause,omitempty"`
	// keyspace to refresh 
	Keyspace string `json:"keyspace"`
	// table to refresh 
	Table string `json:"table"`
}

// NewRefreshOperationResponse instantiates a new RefreshOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshOperationResponse(type_ string, id string, state string, progress float32, creationTime time.Time, keyspace string, table string, ) *RefreshOperationResponse {
	this := RefreshOperationResponse{}
	this.Type = type_
	this.Id = id
	this.State = state
	this.Progress = progress
	this.CreationTime = creationTime
	this.Keyspace = keyspace
	this.Table = table
	return &this
}

// NewRefreshOperationResponseWithDefaults instantiates a new RefreshOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshOperationResponseWithDefaults() *RefreshOperationResponse {
	this := RefreshOperationResponse{}
	return &this
}

// GetType returns the Type field value
func (o *RefreshOperationResponse) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RefreshOperationResponse) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RefreshOperationResponse) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *RefreshOperationResponse) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RefreshOperationResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RefreshOperationResponse) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *RefreshOperationResponse) GetState() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *RefreshOperationResponse) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *RefreshOperationResponse) SetState(v string) {
	o.State = v
}

// GetProgress returns the Progress field value
func (o *RefreshOperationResponse) GetProgress() float32 {
	if o == nil  {
		var ret float32
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *RefreshOperationResponse) GetProgressOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *RefreshOperationResponse) SetProgress(v float32) {
	o.Progress = v
}

// GetCreationTime returns the CreationTime field value
func (o *RefreshOperationResponse) GetCreationTime() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *RefreshOperationResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *RefreshOperationResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *RefreshOperationResponse) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshOperationResponse) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *RefreshOperationResponse) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *RefreshOperationResponse) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *RefreshOperationResponse) GetCompletionTime() time.Time {
	if o == nil || o.CompletionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshOperationResponse) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *RefreshOperationResponse) HasCompletionTime() bool {
	if o != nil && o.CompletionTime != nil {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given time.Time and assigns it to the CompletionTime field.
func (o *RefreshOperationResponse) SetCompletionTime(v time.Time) {
	o.CompletionTime = &v
}

// GetFailureCause returns the FailureCause field value if set, zero value otherwise.
func (o *RefreshOperationResponse) GetFailureCause() map[string]interface{} {
	if o == nil || o.FailureCause == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.FailureCause
}

// GetFailureCauseOk returns a tuple with the FailureCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshOperationResponse) GetFailureCauseOk() (*map[string]interface{}, bool) {
	if o == nil || o.FailureCause == nil {
		return nil, false
	}
	return o.FailureCause, true
}

// HasFailureCause returns a boolean if a field has been set.
func (o *RefreshOperationResponse) HasFailureCause() bool {
	if o != nil && o.FailureCause != nil {
		return true
	}

	return false
}

// SetFailureCause gets a reference to the given map[string]interface{} and assigns it to the FailureCause field.
func (o *RefreshOperationResponse) SetFailureCause(v map[string]interface{}) {
	o.FailureCause = &v
}

// GetKeyspace returns the Keyspace field value
func (o *RefreshOperationResponse) GetKeyspace() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Keyspace
}

// GetKeyspaceOk returns a tuple with the Keyspace field value
// and a boolean to check if the value has been set.
func (o *RefreshOperationResponse) GetKeyspaceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Keyspace, true
}

// SetKeyspace sets field value
func (o *RefreshOperationResponse) SetKeyspace(v string) {
	o.Keyspace = v
}

// GetTable returns the Table field value
func (o *RefreshOperationResponse) GetTable() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Table
}

// GetTableOk returns a tuple with the Table field value
// and a boolean to check if the value has been set.
func (o *RefreshOperationResponse) GetTableOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Table, true
}

// SetTable sets field value
func (o *RefreshOperationResponse) SetTable(v string) {
	o.Table = v
}

func (o RefreshOperationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["progress"] = o.Progress
	}
	if true {
		toSerialize["creationTime"] = o.CreationTime
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.CompletionTime != nil {
		toSerialize["completionTime"] = o.CompletionTime
	}
	if o.FailureCause != nil {
		toSerialize["failureCause"] = o.FailureCause
	}
	if true {
		toSerialize["keyspace"] = o.Keyspace
	}
	if true {
		toSerialize["table"] = o.Table
	}
	return json.Marshal(toSerialize)
}

type NullableRefreshOperationResponse struct {
	value *RefreshOperationResponse
	isSet bool
}

func (v NullableRefreshOperationResponse) Get() *RefreshOperationResponse {
	return v.value
}

func (v *NullableRefreshOperationResponse) Set(val *RefreshOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshOperationResponse(val *RefreshOperationResponse) *NullableRefreshOperationResponse {
	return &NullableRefreshOperationResponse{value: val, isSet: true}
}

func (v NullableRefreshOperationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
