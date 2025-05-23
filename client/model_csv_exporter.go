/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CsvExporter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CsvExporter{}

// CsvExporter struct for CsvExporter
type CsvExporter struct {
	FileName string `json:"fileName"`
	ContentType string `json:"contentType"`
	Content string `json:"content"`
}

type _CsvExporter CsvExporter

// NewCsvExporter instantiates a new CsvExporter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsvExporter(fileName string, contentType string, content string) *CsvExporter {
	this := CsvExporter{}
	this.FileName = fileName
	this.ContentType = contentType
	this.Content = content
	return &this
}

// NewCsvExporterWithDefaults instantiates a new CsvExporter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsvExporterWithDefaults() *CsvExporter {
	this := CsvExporter{}
	return &this
}

// GetFileName returns the FileName field value
func (o *CsvExporter) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *CsvExporter) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *CsvExporter) SetFileName(v string) {
	o.FileName = v
}

// GetContentType returns the ContentType field value
func (o *CsvExporter) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *CsvExporter) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *CsvExporter) SetContentType(v string) {
	o.ContentType = v
}

// GetContent returns the Content field value
func (o *CsvExporter) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CsvExporter) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CsvExporter) SetContent(v string) {
	o.Content = v
}

func (o CsvExporter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CsvExporter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fileName"] = o.FileName
	toSerialize["contentType"] = o.ContentType
	toSerialize["content"] = o.Content
	return toSerialize, nil
}

func (o *CsvExporter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fileName",
		"contentType",
		"content",
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

	varCsvExporter := _CsvExporter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCsvExporter)

	if err != nil {
		return err
	}

	*o = CsvExporter(varCsvExporter)

	return err
}

type NullableCsvExporter struct {
	value *CsvExporter
	isSet bool
}

func (v NullableCsvExporter) Get() *CsvExporter {
	return v.value
}

func (v *NullableCsvExporter) Set(val *CsvExporter) {
	v.value = val
	v.isSet = true
}

func (v NullableCsvExporter) IsSet() bool {
	return v.isSet
}

func (v *NullableCsvExporter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsvExporter(val *CsvExporter) *NullableCsvExporter {
	return &NullableCsvExporter{value: val, isSet: true}
}

func (v NullableCsvExporter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsvExporter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


