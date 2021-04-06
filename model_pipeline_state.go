/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bitbucket_client

import (
	"encoding/json"
)

// PipelineState struct for PipelineState
type PipelineState struct {
	Object
}

// NewPipelineState instantiates a new PipelineState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineState(type_ string) *PipelineState {
	this := PipelineState{}
	this.Type = type_
	return &this
}

// NewPipelineStateWithDefaults instantiates a new PipelineState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStateWithDefaults() *PipelineState {
	this := PipelineState{}
	return &this
}

func (o PipelineState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineState struct {
	value *PipelineState
	isSet bool
}

func (v NullablePipelineState) Get() *PipelineState {
	return v.value
}

func (v *NullablePipelineState) Set(val *PipelineState) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineState) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineState(val *PipelineState) *NullablePipelineState {
	return &NullablePipelineState{value: val, isSet: true}
}

func (v NullablePipelineState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


