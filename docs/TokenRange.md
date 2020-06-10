# TokenRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** | starting token of a range  | 
**End** | Pointer to **string** | ending token of a range  | 

## Methods

### NewTokenRange

`func NewTokenRange(start string, end string, ) *TokenRange`

NewTokenRange instantiates a new TokenRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRangeWithDefaults

`func NewTokenRangeWithDefaults() *TokenRange`

NewTokenRangeWithDefaults instantiates a new TokenRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *TokenRange) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TokenRange) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TokenRange) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *TokenRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TokenRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TokenRange) SetEnd(v string)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


