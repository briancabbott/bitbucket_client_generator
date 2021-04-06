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

// BranchingModelSettingsBranchTypes struct for BranchingModelSettingsBranchTypes
type BranchingModelSettingsBranchTypes struct {
	// Whether the branch type is enabled or not. A disabled branch type may contain an invalid `prefix`.
	Enabled *bool `json:"enabled,omitempty"`
	// The kind of the branch type.
	Kind string `json:"kind"`
	// The prefix for this branch type. A branch with this prefix will be classified as per `kind`. The `prefix` of an enabled branch type must be a valid branch prefix.Additionally, it cannot be blank, empty or `null`. The `prefix` for a disabled branch type can be empty or invalid.
	Prefix *string `json:"prefix,omitempty"`
}

// NewBranchingModelSettingsBranchTypes instantiates a new BranchingModelSettingsBranchTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranchingModelSettingsBranchTypes(kind string) *BranchingModelSettingsBranchTypes {
	this := BranchingModelSettingsBranchTypes{}
	this.Kind = kind
	return &this
}

// NewBranchingModelSettingsBranchTypesWithDefaults instantiates a new BranchingModelSettingsBranchTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchingModelSettingsBranchTypesWithDefaults() *BranchingModelSettingsBranchTypes {
	this := BranchingModelSettingsBranchTypes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BranchingModelSettingsBranchTypes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchingModelSettingsBranchTypes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BranchingModelSettingsBranchTypes) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BranchingModelSettingsBranchTypes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetKind returns the Kind field value
func (o *BranchingModelSettingsBranchTypes) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *BranchingModelSettingsBranchTypes) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *BranchingModelSettingsBranchTypes) SetKind(v string) {
	o.Kind = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *BranchingModelSettingsBranchTypes) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchingModelSettingsBranchTypes) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *BranchingModelSettingsBranchTypes) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *BranchingModelSettingsBranchTypes) SetPrefix(v string) {
	o.Prefix = &v
}

func (o BranchingModelSettingsBranchTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	return json.Marshal(toSerialize)
}

type NullableBranchingModelSettingsBranchTypes struct {
	value *BranchingModelSettingsBranchTypes
	isSet bool
}

func (v NullableBranchingModelSettingsBranchTypes) Get() *BranchingModelSettingsBranchTypes {
	return v.value
}

func (v *NullableBranchingModelSettingsBranchTypes) Set(val *BranchingModelSettingsBranchTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableBranchingModelSettingsBranchTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableBranchingModelSettingsBranchTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranchingModelSettingsBranchTypes(val *BranchingModelSettingsBranchTypes) *NullableBranchingModelSettingsBranchTypes {
	return &NullableBranchingModelSettingsBranchTypes{value: val, isSet: true}
}

func (v NullableBranchingModelSettingsBranchTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranchingModelSettingsBranchTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


