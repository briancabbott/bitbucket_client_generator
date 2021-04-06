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
	"time"
)

// ReportAnnotation struct for ReportAnnotation
type ReportAnnotation struct {
	Object
	// ID of the annotation provided by the annotation creator. It can be used to identify the annotation as an alternative to it's generated uuid. It is not used by Bitbucket, but only by the annotation creator for updating or deleting this specific annotation. Needs to be unique.
	ExternalId *string `json:"external_id,omitempty"`
	// The UUID that can be used to identify the annotation.
	Uuid *string `json:"uuid,omitempty"`
	// The type of the report.
	AnnotationType *string `json:"annotation_type,omitempty"`
	// The path of the file on which this annotation should be placed. This is the path of the file relative to the git repository. If no path is provided, then it will appear in the overview modal on all pull requests where the tip of the branch is the given commit, regardless of which files were modified.
	Path *string `json:"path,omitempty"`
	// The line number that the annotation should belong to. If no line number is provided, then it will default to 0 and in a pull request it will appear at the top of the file specified by the path field.
	Line *int32 `json:"line,omitempty"`
	// The message to display to users.
	Summary *string `json:"summary,omitempty"`
	// The details to show to users when clicking on the annotation.
	Details *string `json:"details,omitempty"`
	// The state of the report. May be set to PENDING and later updated.
	Result *string `json:"result,omitempty"`
	// The severity of the annotation.
	Severity *string `json:"severity,omitempty"`
	// A URL linking to the annotation in an external tool.
	Link *string `json:"link,omitempty"`
	// The timestamp when the report was created.
	CreatedOn *time.Time `json:"created_on,omitempty"`
	// The timestamp when the report was updated.
	UpdatedOn *time.Time `json:"updated_on,omitempty"`
}

// NewReportAnnotation instantiates a new ReportAnnotation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportAnnotation(type_ string) *ReportAnnotation {
	this := ReportAnnotation{}
	this.Type = type_
	return &this
}

// NewReportAnnotationWithDefaults instantiates a new ReportAnnotation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportAnnotationWithDefaults() *ReportAnnotation {
	this := ReportAnnotation{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ReportAnnotation) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportAnnotation) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ReportAnnotation) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ReportAnnotation) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ReportAnnotation) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportAnnotation) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ReportAnnotation) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ReportAnnotation) SetUuid(v string) {
	o.Uuid = &v
}

// GetAnnotationType returns the AnnotationType field value if set, zero value otherwise.
func (o *ReportAnnotation) GetAnnotationType() string {
	if o == nil || o.AnnotationType == nil {
		var ret string
		return ret
	}
	return *o.AnnotationType
}

// GetAnnotationTypeOk returns a tuple with the AnnotationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportAnnotation) GetAnnotationTypeOk() (*string, bool) {
	if o == nil || o.AnnotationType == nil {
		return nil, false
	}
	return o.AnnotationType, true
}

// HasAnnotationType returns a boolean if a field has been set.
func (o *ReportAnnotation) HasAnnotationType() bool {
	if o != nil && o.AnnotationType != nil {
		return true
	}

	return false
}

// SetAnnotationType gets a reference to the given string and assigns it to the AnnotationType field.
func (o *ReportAnnotation) SetAnnotationType(v string) {
	o.AnnotationType = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ReportAnnotation) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportAnnotation) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ReportAnnotation) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ReportAnnotation) SetPath(v string) {
	o.Path = &v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *ReportAnnotation) GetLine() int32 {
	if o == nil || o.Line == nil {
		var ret int32
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportAnnotation) GetLineOk() (*int32, bool) {
	if o == nil || o.Line == nil {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *ReportAnnotation) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given int32 and assigns it to the Line field.
func (o *ReportAnnotation) SetLine(v int32) {
	o.Line = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ReportAnnotation) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportAnnotation) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ReportAnnotation) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *ReportAnnotation) SetSummary(v string) {
	o.Summary = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ReportAnnotation) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportAnnotation) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ReportAnnotation) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ReportAnnotation) SetDetails(v string) {
	o.Details = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReportAnnotation) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportAnnotation) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReportAnnotation) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *ReportAnnotation) SetResult(v string) {
	o.Result = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ReportAnnotation) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportAnnotation) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ReportAnnotation) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ReportAnnotation) SetSeverity(v string) {
	o.Severity = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *ReportAnnotation) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportAnnotation) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *ReportAnnotation) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *ReportAnnotation) SetLink(v string) {
	o.Link = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *ReportAnnotation) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportAnnotation) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *ReportAnnotation) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *ReportAnnotation) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *ReportAnnotation) GetUpdatedOn() time.Time {
	if o == nil || o.UpdatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportAnnotation) GetUpdatedOnOk() (*time.Time, bool) {
	if o == nil || o.UpdatedOn == nil {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *ReportAnnotation) HasUpdatedOn() bool {
	if o != nil && o.UpdatedOn != nil {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given time.Time and assigns it to the UpdatedOn field.
func (o *ReportAnnotation) SetUpdatedOn(v time.Time) {
	o.UpdatedOn = &v
}

func (o ReportAnnotation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.AnnotationType != nil {
		toSerialize["annotation_type"] = o.AnnotationType
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Line != nil {
		toSerialize["line"] = o.Line
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.CreatedOn != nil {
		toSerialize["created_on"] = o.CreatedOn
	}
	if o.UpdatedOn != nil {
		toSerialize["updated_on"] = o.UpdatedOn
	}
	return json.Marshal(toSerialize)
}

type NullableReportAnnotation struct {
	value *ReportAnnotation
	isSet bool
}

func (v NullableReportAnnotation) Get() *ReportAnnotation {
	return v.value
}

func (v *NullableReportAnnotation) Set(val *ReportAnnotation) {
	v.value = val
	v.isSet = true
}

func (v NullableReportAnnotation) IsSet() bool {
	return v.isSet
}

func (v *NullableReportAnnotation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportAnnotation(val *ReportAnnotation) *NullableReportAnnotation {
	return &NullableReportAnnotation{value: val, isSet: true}
}

func (v NullableReportAnnotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportAnnotation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


