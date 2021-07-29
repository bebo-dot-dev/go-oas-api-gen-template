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

// UserCredentials struct for UserCredentials
type UserCredentials struct {
	// the username of an account
	Username string `json:"username"`
	// the password of an account
	Password string `json:"password"`
}

// NewUserCredentials instantiates a new UserCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCredentials(username string, password string) *UserCredentials {
	this := UserCredentials{}
	this.Username = username
	this.Password = password
	return &this
}

// NewUserCredentialsWithDefaults instantiates a new UserCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCredentialsWithDefaults() *UserCredentials {
	this := UserCredentials{}
	return &this
}

// GetUsername returns the Username field value
func (o *UserCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserCredentials) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserCredentials) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *UserCredentials) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserCredentials) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserCredentials) SetPassword(v string) {
	o.Password = v
}

func (o UserCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableUserCredentials struct {
	value *UserCredentials
	isSet bool
}

func (v NullableUserCredentials) Get() *UserCredentials {
	return v.value
}

func (v *NullableUserCredentials) Set(val *UserCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCredentials(val *UserCredentials) *NullableUserCredentials {
	return &NullableUserCredentials{value: val, isSet: true}
}

func (v NullableUserCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
