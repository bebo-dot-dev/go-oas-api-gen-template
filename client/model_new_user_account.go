/*
 * User API
 *
 * An API for creating and authenticating system users
 *
 * API version: 0.1.0
 * Contact: joe@bebo.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NewUserAccount struct for NewUserAccount
type NewUserAccount struct {
	// the id of an account
	Id int32 `json:"id"`
	// describes whether an account was created
	Created bool `json:"created"`
}

// NewNewUserAccount instantiates a new NewUserAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewUserAccount(id int32, created bool) *NewUserAccount {
	this := NewUserAccount{}
	this.Id = id
	this.Created = created
	return &this
}

// NewNewUserAccountWithDefaults instantiates a new NewUserAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewUserAccountWithDefaults() *NewUserAccount {
	this := NewUserAccount{}
	return &this
}

// GetId returns the Id field value
func (o *NewUserAccount) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NewUserAccount) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NewUserAccount) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *NewUserAccount) GetCreated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *NewUserAccount) GetCreatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *NewUserAccount) SetCreated(v bool) {
	o.Created = v
}

func (o NewUserAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created"] = o.Created
	}
	return json.Marshal(toSerialize)
}

type NullableNewUserAccount struct {
	value *NewUserAccount
	isSet bool
}

func (v NullableNewUserAccount) Get() *NewUserAccount {
	return v.value
}

func (v *NullableNewUserAccount) Set(val *NewUserAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUserAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUserAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUserAccount(val *NewUserAccount) *NullableNewUserAccount {
	return &NullableNewUserAccount{value: val, isSet: true}
}

func (v NullableNewUserAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUserAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
