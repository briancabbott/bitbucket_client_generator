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

// IssueAttachment An issue file attachment's meta data. Note this does not contain the file's actual contents.
type IssueAttachment struct {
	Links *BranchingModelSettingsLinks `json:"links,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewIssueAttachment instantiates a new IssueAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueAttachment() *IssueAttachment {
	this := IssueAttachment{}
	return &this
}

// NewIssueAttachmentWithDefaults instantiates a new IssueAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueAttachmentWithDefaults() *IssueAttachment {
	this := IssueAttachment{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IssueAttachment) GetLinks() BranchingModelSettingsLinks {
	if o == nil || o.Links == nil {
		var ret BranchingModelSettingsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttachment) GetLinksOk() (*BranchingModelSettingsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IssueAttachment) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given BranchingModelSettingsLinks and assigns it to the Links field.
func (o *IssueAttachment) SetLinks(v BranchingModelSettingsLinks) {
	o.Links = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IssueAttachment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttachment) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IssueAttachment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IssueAttachment) SetName(v string) {
	o.Name = &v
}

func (o IssueAttachment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableIssueAttachment struct {
	value *IssueAttachment
	isSet bool
}

func (v NullableIssueAttachment) Get() *IssueAttachment {
	return v.value
}

func (v *NullableIssueAttachment) Set(val *IssueAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueAttachment(val *IssueAttachment) *NullableIssueAttachment {
	return &NullableIssueAttachment{value: val, isSet: true}
}

func (v NullableIssueAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


