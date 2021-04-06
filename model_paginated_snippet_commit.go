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

// PaginatedSnippetCommit A paginated list of snippet commits.
type PaginatedSnippetCommit struct {
	// Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Size *int32 `json:"size,omitempty"`
	// Page number of the current results. This is an optional element that is not provided in all responses.
	Page *int32 `json:"page,omitempty"`
	// Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Pagelen *int32 `json:"pagelen,omitempty"`
	// Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Next *string `json:"next,omitempty"`
	// Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Previous *string `json:"previous,omitempty"`
	Values *[]SnippetCommit `json:"values,omitempty"`
}

// NewPaginatedSnippetCommit instantiates a new PaginatedSnippetCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSnippetCommit() *PaginatedSnippetCommit {
	this := PaginatedSnippetCommit{}
	return &this
}

// NewPaginatedSnippetCommitWithDefaults instantiates a new PaginatedSnippetCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSnippetCommitWithDefaults() *PaginatedSnippetCommit {
	this := PaginatedSnippetCommit{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PaginatedSnippetCommit) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedSnippetCommit) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PaginatedSnippetCommit) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *PaginatedSnippetCommit) SetSize(v int32) {
	o.Size = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PaginatedSnippetCommit) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedSnippetCommit) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PaginatedSnippetCommit) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *PaginatedSnippetCommit) SetPage(v int32) {
	o.Page = &v
}

// GetPagelen returns the Pagelen field value if set, zero value otherwise.
func (o *PaginatedSnippetCommit) GetPagelen() int32 {
	if o == nil || o.Pagelen == nil {
		var ret int32
		return ret
	}
	return *o.Pagelen
}

// GetPagelenOk returns a tuple with the Pagelen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedSnippetCommit) GetPagelenOk() (*int32, bool) {
	if o == nil || o.Pagelen == nil {
		return nil, false
	}
	return o.Pagelen, true
}

// HasPagelen returns a boolean if a field has been set.
func (o *PaginatedSnippetCommit) HasPagelen() bool {
	if o != nil && o.Pagelen != nil {
		return true
	}

	return false
}

// SetPagelen gets a reference to the given int32 and assigns it to the Pagelen field.
func (o *PaginatedSnippetCommit) SetPagelen(v int32) {
	o.Pagelen = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginatedSnippetCommit) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedSnippetCommit) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedSnippetCommit) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *PaginatedSnippetCommit) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *PaginatedSnippetCommit) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedSnippetCommit) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedSnippetCommit) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *PaginatedSnippetCommit) SetPrevious(v string) {
	o.Previous = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *PaginatedSnippetCommit) GetValues() []SnippetCommit {
	if o == nil || o.Values == nil {
		var ret []SnippetCommit
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedSnippetCommit) GetValuesOk() (*[]SnippetCommit, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *PaginatedSnippetCommit) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []SnippetCommit and assigns it to the Values field.
func (o *PaginatedSnippetCommit) SetValues(v []SnippetCommit) {
	o.Values = &v
}

func (o PaginatedSnippetCommit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Pagelen != nil {
		toSerialize["pagelen"] = o.Pagelen
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSnippetCommit struct {
	value *PaginatedSnippetCommit
	isSet bool
}

func (v NullablePaginatedSnippetCommit) Get() *PaginatedSnippetCommit {
	return v.value
}

func (v *NullablePaginatedSnippetCommit) Set(val *PaginatedSnippetCommit) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSnippetCommit) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSnippetCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSnippetCommit(val *PaginatedSnippetCommit) *NullablePaginatedSnippetCommit {
	return &NullablePaginatedSnippetCommit{value: val, isSet: true}
}

func (v NullablePaginatedSnippetCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSnippetCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


