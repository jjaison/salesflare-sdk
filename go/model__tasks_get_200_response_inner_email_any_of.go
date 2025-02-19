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

// checks if the TasksGet200ResponseInnerEmailAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TasksGet200ResponseInnerEmailAnyOf{}

// TasksGet200ResponseInnerEmailAnyOf struct for TasksGet200ResponseInnerEmailAnyOf
type TasksGet200ResponseInnerEmailAnyOf struct {
	Id int32 `json:"id"`
	EmailMessageId string `json:"email_message_id"`
	ServiceMessageId string `json:"service_message_id"`
	DataSource int32 `json:"data_source"`
	Subject string `json:"subject"`
}

type _TasksGet200ResponseInnerEmailAnyOf TasksGet200ResponseInnerEmailAnyOf

// NewTasksGet200ResponseInnerEmailAnyOf instantiates a new TasksGet200ResponseInnerEmailAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTasksGet200ResponseInnerEmailAnyOf(id int32, emailMessageId string, serviceMessageId string, dataSource int32, subject string) *TasksGet200ResponseInnerEmailAnyOf {
	this := TasksGet200ResponseInnerEmailAnyOf{}
	this.Id = id
	this.EmailMessageId = emailMessageId
	this.ServiceMessageId = serviceMessageId
	this.DataSource = dataSource
	this.Subject = subject
	return &this
}

// NewTasksGet200ResponseInnerEmailAnyOfWithDefaults instantiates a new TasksGet200ResponseInnerEmailAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTasksGet200ResponseInnerEmailAnyOfWithDefaults() *TasksGet200ResponseInnerEmailAnyOf {
	this := TasksGet200ResponseInnerEmailAnyOf{}
	return &this
}

// GetId returns the Id field value
func (o *TasksGet200ResponseInnerEmailAnyOf) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TasksGet200ResponseInnerEmailAnyOf) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TasksGet200ResponseInnerEmailAnyOf) SetId(v int32) {
	o.Id = v
}

// GetEmailMessageId returns the EmailMessageId field value
func (o *TasksGet200ResponseInnerEmailAnyOf) GetEmailMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailMessageId
}

// GetEmailMessageIdOk returns a tuple with the EmailMessageId field value
// and a boolean to check if the value has been set.
func (o *TasksGet200ResponseInnerEmailAnyOf) GetEmailMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailMessageId, true
}

// SetEmailMessageId sets field value
func (o *TasksGet200ResponseInnerEmailAnyOf) SetEmailMessageId(v string) {
	o.EmailMessageId = v
}

// GetServiceMessageId returns the ServiceMessageId field value
func (o *TasksGet200ResponseInnerEmailAnyOf) GetServiceMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceMessageId
}

// GetServiceMessageIdOk returns a tuple with the ServiceMessageId field value
// and a boolean to check if the value has been set.
func (o *TasksGet200ResponseInnerEmailAnyOf) GetServiceMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceMessageId, true
}

// SetServiceMessageId sets field value
func (o *TasksGet200ResponseInnerEmailAnyOf) SetServiceMessageId(v string) {
	o.ServiceMessageId = v
}

// GetDataSource returns the DataSource field value
func (o *TasksGet200ResponseInnerEmailAnyOf) GetDataSource() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TasksGet200ResponseInnerEmailAnyOf) GetDataSourceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TasksGet200ResponseInnerEmailAnyOf) SetDataSource(v int32) {
	o.DataSource = v
}

// GetSubject returns the Subject field value
func (o *TasksGet200ResponseInnerEmailAnyOf) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *TasksGet200ResponseInnerEmailAnyOf) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *TasksGet200ResponseInnerEmailAnyOf) SetSubject(v string) {
	o.Subject = v
}

func (o TasksGet200ResponseInnerEmailAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TasksGet200ResponseInnerEmailAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["email_message_id"] = o.EmailMessageId
	toSerialize["service_message_id"] = o.ServiceMessageId
	toSerialize["data_source"] = o.DataSource
	toSerialize["subject"] = o.Subject
	return toSerialize, nil
}

func (o *TasksGet200ResponseInnerEmailAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"email_message_id",
		"service_message_id",
		"data_source",
		"subject",
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

	varTasksGet200ResponseInnerEmailAnyOf := _TasksGet200ResponseInnerEmailAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTasksGet200ResponseInnerEmailAnyOf)

	if err != nil {
		return err
	}

	*o = TasksGet200ResponseInnerEmailAnyOf(varTasksGet200ResponseInnerEmailAnyOf)

	return err
}

type NullableTasksGet200ResponseInnerEmailAnyOf struct {
	value *TasksGet200ResponseInnerEmailAnyOf
	isSet bool
}

func (v NullableTasksGet200ResponseInnerEmailAnyOf) Get() *TasksGet200ResponseInnerEmailAnyOf {
	return v.value
}

func (v *NullableTasksGet200ResponseInnerEmailAnyOf) Set(val *TasksGet200ResponseInnerEmailAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTasksGet200ResponseInnerEmailAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTasksGet200ResponseInnerEmailAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTasksGet200ResponseInnerEmailAnyOf(val *TasksGet200ResponseInnerEmailAnyOf) *NullableTasksGet200ResponseInnerEmailAnyOf {
	return &NullableTasksGet200ResponseInnerEmailAnyOf{value: val, isSet: true}
}

func (v NullableTasksGet200ResponseInnerEmailAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTasksGet200ResponseInnerEmailAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


