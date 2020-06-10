# DecommissionOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | 
**Force** | Pointer to **bool** | forcibly decommission a node, even if by doing so there will not be enough replicas for responding to requests, this option is relevant only for Cassandra 4.x and it is not in use for lower versions, defaults to false.  | [optional] 

## Methods

### NewDecommissionOperationRequest

`func NewDecommissionOperationRequest(type_ string, ) *DecommissionOperationRequest`

NewDecommissionOperationRequest instantiates a new DecommissionOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecommissionOperationRequestWithDefaults

`func NewDecommissionOperationRequestWithDefaults() *DecommissionOperationRequest`

NewDecommissionOperationRequestWithDefaults instantiates a new DecommissionOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DecommissionOperationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DecommissionOperationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DecommissionOperationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetForce

`func (o *DecommissionOperationRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *DecommissionOperationRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *DecommissionOperationRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *DecommissionOperationRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


