# DataRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **int32** | quantified value of bandwidth, an integer  | [optional] 
**Unit** | Pointer to **string** | unit of &#39;data bandwidth&#39;  | 

## Methods

### NewDataRate

`func NewDataRate(unit string, ) *DataRate`

NewDataRate instantiates a new DataRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataRateWithDefaults

`func NewDataRateWithDefaults() *DataRate`

NewDataRateWithDefaults instantiates a new DataRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DataRate) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataRate) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataRate) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *DataRate) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUnit

`func (o *DataRate) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *DataRate) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *DataRate) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


