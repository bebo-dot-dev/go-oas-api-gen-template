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

// PingResponse struct for PingResponse
type PingResponse struct {
	Hostname string `json:"hostname"`
}

// NewPingResponse instantiates a new PingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingResponse(hostname string) *PingResponse {
	this := PingResponse{}
	this.Hostname = hostname
	return &this
}

// NewPingResponseWithDefaults instantiates a new PingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingResponseWithDefaults() *PingResponse {
	this := PingResponse{}
	return &this
}

// GetHostname returns the Hostname field value
func (o *PingResponse) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *PingResponse) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *PingResponse) SetHostname(v string) {
	o.Hostname = v
}

func (o PingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	return json.Marshal(toSerialize)
}

type NullablePingResponse struct {
	value *PingResponse
	isSet bool
}

func (v NullablePingResponse) Get() *PingResponse {
	return v.value
}

func (v *NullablePingResponse) Set(val *PingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingResponse(val *PingResponse) *NullablePingResponse {
	return &NullablePingResponse{value: val, isSet: true}
}

func (v NullablePingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
