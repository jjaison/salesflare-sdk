/*
app.salesflare.com

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package salesflare

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the V1MPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1MPost200Response{}

// V1MPost200Response struct for V1MPost200Response
type V1MPost200Response struct {
	Success bool `json:"success"`
}

type _V1MPost200Response V1MPost200Response

// NewV1MPost200Response instantiates a new V1MPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1MPost200Response(success bool) *V1MPost200Response {
	this := V1MPost200Response{}
	this.Success = success
	return &this
}

// NewV1MPost200ResponseWithDefaults instantiates a new V1MPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1MPost200ResponseWithDefaults() *V1MPost200Response {
	this := V1MPost200Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *V1MPost200Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *V1MPost200Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *V1MPost200Response) SetSuccess(v bool) {
	o.Success = v
}

func (o V1MPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1MPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

func (o *V1MPost200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV1MPost200Response := _V1MPost200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1MPost200Response)

	if err != nil {
		return err
	}

	*o = V1MPost200Response(varV1MPost200Response)

	return err
}

type NullableV1MPost200Response struct {
	value *V1MPost200Response
	isSet bool
}

func (v NullableV1MPost200Response) Get() *V1MPost200Response {
	return v.value
}

func (v *NullableV1MPost200Response) Set(val *V1MPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1MPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1MPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1MPost200Response(val *V1MPost200Response) *NullableV1MPost200Response {
	return &NullableV1MPost200Response{value: val, isSet: true}
}

func (v NullableV1MPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1MPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


