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

// PullrequestMergeParameters The metadata that describes a pull request merge.
type PullrequestMergeParameters struct {
	Type string `json:"type"`
	// The commit message that will be used on the resulting commit.
	Message *string `json:"message,omitempty"`
	// Whether the source branch should be deleted. If this is not provided, we fallback to the value used when the pull request was created, which defaults to False
	CloseSourceBranch *bool `json:"close_source_branch,omitempty"`
	// The merge strategy that will be used to merge the pull request.
	MergeStrategy *string `json:"merge_strategy,omitempty"`
}

// NewPullrequestMergeParameters instantiates a new PullrequestMergeParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullrequestMergeParameters(type_ string) *PullrequestMergeParameters {
	this := PullrequestMergeParameters{}
	this.Type = type_
	var mergeStrategy string = "merge_commit"
	this.MergeStrategy = &mergeStrategy
	return &this
}

// NewPullrequestMergeParametersWithDefaults instantiates a new PullrequestMergeParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullrequestMergeParametersWithDefaults() *PullrequestMergeParameters {
	this := PullrequestMergeParameters{}
	var mergeStrategy string = "merge_commit"
	this.MergeStrategy = &mergeStrategy
	return &this
}

// GetType returns the Type field value
func (o *PullrequestMergeParameters) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PullrequestMergeParameters) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PullrequestMergeParameters) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PullrequestMergeParameters) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullrequestMergeParameters) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PullrequestMergeParameters) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PullrequestMergeParameters) SetMessage(v string) {
	o.Message = &v
}

// GetCloseSourceBranch returns the CloseSourceBranch field value if set, zero value otherwise.
func (o *PullrequestMergeParameters) GetCloseSourceBranch() bool {
	if o == nil || o.CloseSourceBranch == nil {
		var ret bool
		return ret
	}
	return *o.CloseSourceBranch
}

// GetCloseSourceBranchOk returns a tuple with the CloseSourceBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullrequestMergeParameters) GetCloseSourceBranchOk() (*bool, bool) {
	if o == nil || o.CloseSourceBranch == nil {
		return nil, false
	}
	return o.CloseSourceBranch, true
}

// HasCloseSourceBranch returns a boolean if a field has been set.
func (o *PullrequestMergeParameters) HasCloseSourceBranch() bool {
	if o != nil && o.CloseSourceBranch != nil {
		return true
	}

	return false
}

// SetCloseSourceBranch gets a reference to the given bool and assigns it to the CloseSourceBranch field.
func (o *PullrequestMergeParameters) SetCloseSourceBranch(v bool) {
	o.CloseSourceBranch = &v
}

// GetMergeStrategy returns the MergeStrategy field value if set, zero value otherwise.
func (o *PullrequestMergeParameters) GetMergeStrategy() string {
	if o == nil || o.MergeStrategy == nil {
		var ret string
		return ret
	}
	return *o.MergeStrategy
}

// GetMergeStrategyOk returns a tuple with the MergeStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullrequestMergeParameters) GetMergeStrategyOk() (*string, bool) {
	if o == nil || o.MergeStrategy == nil {
		return nil, false
	}
	return o.MergeStrategy, true
}

// HasMergeStrategy returns a boolean if a field has been set.
func (o *PullrequestMergeParameters) HasMergeStrategy() bool {
	if o != nil && o.MergeStrategy != nil {
		return true
	}

	return false
}

// SetMergeStrategy gets a reference to the given string and assigns it to the MergeStrategy field.
func (o *PullrequestMergeParameters) SetMergeStrategy(v string) {
	o.MergeStrategy = &v
}

func (o PullrequestMergeParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.CloseSourceBranch != nil {
		toSerialize["close_source_branch"] = o.CloseSourceBranch
	}
	if o.MergeStrategy != nil {
		toSerialize["merge_strategy"] = o.MergeStrategy
	}
	return json.Marshal(toSerialize)
}

type NullablePullrequestMergeParameters struct {
	value *PullrequestMergeParameters
	isSet bool
}

func (v NullablePullrequestMergeParameters) Get() *PullrequestMergeParameters {
	return v.value
}

func (v *NullablePullrequestMergeParameters) Set(val *PullrequestMergeParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePullrequestMergeParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePullrequestMergeParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullrequestMergeParameters(val *PullrequestMergeParameters) *NullablePullrequestMergeParameters {
	return &NullablePullrequestMergeParameters{value: val, isSet: true}
}

func (v NullablePullrequestMergeParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullrequestMergeParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


