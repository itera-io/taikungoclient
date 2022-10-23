// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SearchPodsListReader is a Reader for the SearchPodsList structure.
type SearchPodsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchPodsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchPodsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchPodsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchPodsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchPodsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchPodsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchPodsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchPodsListOK creates a SearchPodsListOK with default headers values
func NewSearchPodsListOK() *SearchPodsListOK {
	return &SearchPodsListOK{}
}

/*
SearchPodsListOK describes a response with status code 200, with default header values.

Success
*/
type SearchPodsListOK struct {
	Payload *SearchPodsListOKBody
}

// IsSuccess returns true when this search pods list o k response has a 2xx status code
func (o *SearchPodsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search pods list o k response has a 3xx status code
func (o *SearchPodsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search pods list o k response has a 4xx status code
func (o *SearchPodsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search pods list o k response has a 5xx status code
func (o *SearchPodsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search pods list o k response a status code equal to that given
func (o *SearchPodsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchPodsListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pods][%d] searchPodsListOK  %+v", 200, o.Payload)
}

func (o *SearchPodsListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pods][%d] searchPodsListOK  %+v", 200, o.Payload)
}

func (o *SearchPodsListOK) GetPayload() *SearchPodsListOKBody {
	return o.Payload
}

func (o *SearchPodsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchPodsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPodsListBadRequest creates a SearchPodsListBadRequest with default headers values
func NewSearchPodsListBadRequest() *SearchPodsListBadRequest {
	return &SearchPodsListBadRequest{}
}

/*
SearchPodsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchPodsListBadRequest struct {
	Payload []*SearchPodsListBadRequestBodyItems0
}

// IsSuccess returns true when this search pods list bad request response has a 2xx status code
func (o *SearchPodsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search pods list bad request response has a 3xx status code
func (o *SearchPodsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search pods list bad request response has a 4xx status code
func (o *SearchPodsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search pods list bad request response has a 5xx status code
func (o *SearchPodsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search pods list bad request response a status code equal to that given
func (o *SearchPodsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchPodsListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pods][%d] searchPodsListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchPodsListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pods][%d] searchPodsListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchPodsListBadRequest) GetPayload() []*SearchPodsListBadRequestBodyItems0 {
	return o.Payload
}

func (o *SearchPodsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPodsListUnauthorized creates a SearchPodsListUnauthorized with default headers values
func NewSearchPodsListUnauthorized() *SearchPodsListUnauthorized {
	return &SearchPodsListUnauthorized{}
}

/*
SearchPodsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchPodsListUnauthorized struct {
	Payload *SearchPodsListUnauthorizedBody
}

// IsSuccess returns true when this search pods list unauthorized response has a 2xx status code
func (o *SearchPodsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search pods list unauthorized response has a 3xx status code
func (o *SearchPodsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search pods list unauthorized response has a 4xx status code
func (o *SearchPodsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search pods list unauthorized response has a 5xx status code
func (o *SearchPodsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search pods list unauthorized response a status code equal to that given
func (o *SearchPodsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchPodsListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pods][%d] searchPodsListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchPodsListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pods][%d] searchPodsListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchPodsListUnauthorized) GetPayload() *SearchPodsListUnauthorizedBody {
	return o.Payload
}

func (o *SearchPodsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchPodsListUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPodsListForbidden creates a SearchPodsListForbidden with default headers values
func NewSearchPodsListForbidden() *SearchPodsListForbidden {
	return &SearchPodsListForbidden{}
}

/*
SearchPodsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchPodsListForbidden struct {
	Payload *SearchPodsListForbiddenBody
}

// IsSuccess returns true when this search pods list forbidden response has a 2xx status code
func (o *SearchPodsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search pods list forbidden response has a 3xx status code
func (o *SearchPodsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search pods list forbidden response has a 4xx status code
func (o *SearchPodsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search pods list forbidden response has a 5xx status code
func (o *SearchPodsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search pods list forbidden response a status code equal to that given
func (o *SearchPodsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchPodsListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pods][%d] searchPodsListForbidden  %+v", 403, o.Payload)
}

func (o *SearchPodsListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pods][%d] searchPodsListForbidden  %+v", 403, o.Payload)
}

func (o *SearchPodsListForbidden) GetPayload() *SearchPodsListForbiddenBody {
	return o.Payload
}

func (o *SearchPodsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchPodsListForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPodsListNotFound creates a SearchPodsListNotFound with default headers values
func NewSearchPodsListNotFound() *SearchPodsListNotFound {
	return &SearchPodsListNotFound{}
}

/*
SearchPodsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchPodsListNotFound struct {
	Payload *SearchPodsListNotFoundBody
}

// IsSuccess returns true when this search pods list not found response has a 2xx status code
func (o *SearchPodsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search pods list not found response has a 3xx status code
func (o *SearchPodsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search pods list not found response has a 4xx status code
func (o *SearchPodsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search pods list not found response has a 5xx status code
func (o *SearchPodsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search pods list not found response a status code equal to that given
func (o *SearchPodsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchPodsListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pods][%d] searchPodsListNotFound  %+v", 404, o.Payload)
}

func (o *SearchPodsListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pods][%d] searchPodsListNotFound  %+v", 404, o.Payload)
}

func (o *SearchPodsListNotFound) GetPayload() *SearchPodsListNotFoundBody {
	return o.Payload
}

func (o *SearchPodsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchPodsListNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPodsListInternalServerError creates a SearchPodsListInternalServerError with default headers values
func NewSearchPodsListInternalServerError() *SearchPodsListInternalServerError {
	return &SearchPodsListInternalServerError{}
}

/*
SearchPodsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchPodsListInternalServerError struct {
}

// IsSuccess returns true when this search pods list internal server error response has a 2xx status code
func (o *SearchPodsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search pods list internal server error response has a 3xx status code
func (o *SearchPodsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search pods list internal server error response has a 4xx status code
func (o *SearchPodsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search pods list internal server error response has a 5xx status code
func (o *SearchPodsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search pods list internal server error response a status code equal to that given
func (o *SearchPodsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchPodsListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pods][%d] searchPodsListInternalServerError ", 500)
}

func (o *SearchPodsListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pods][%d] searchPodsListInternalServerError ", 500)
}

func (o *SearchPodsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
SearchPodsListBadRequestBodyItems0 search pods list bad request body items0
swagger:model SearchPodsListBadRequestBodyItems0
*/
type SearchPodsListBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this search pods list bad request body items0
func (o *SearchPodsListBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search pods list bad request body items0 based on context it is used
func (o *SearchPodsListBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchPodsListBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchPodsListBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res SearchPodsListBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchPodsListBody search pods list body
swagger:model SearchPodsListBody
*/
type SearchPodsListBody struct {

	// limit
	Limit int32 `json:"limit,omitempty"`

	// offset
	Offset int32 `json:"offset,omitempty"`

	// search term
	SearchTerm string `json:"searchTerm,omitempty"`
}

// Validate validates this search pods list body
func (o *SearchPodsListBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search pods list body based on context it is used
func (o *SearchPodsListBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchPodsListBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchPodsListBody) UnmarshalBinary(b []byte) error {
	var res SearchPodsListBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchPodsListForbiddenBody search pods list forbidden body
swagger:model SearchPodsListForbiddenBody
*/
type SearchPodsListForbiddenBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this search pods list forbidden body
func (o *SearchPodsListForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search pods list forbidden body based on context it is used
func (o *SearchPodsListForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchPodsListForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchPodsListForbiddenBody) UnmarshalBinary(b []byte) error {
	var res SearchPodsListForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchPodsListNotFoundBody search pods list not found body
swagger:model SearchPodsListNotFoundBody
*/
type SearchPodsListNotFoundBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this search pods list not found body
func (o *SearchPodsListNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search pods list not found body based on context it is used
func (o *SearchPodsListNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchPodsListNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchPodsListNotFoundBody) UnmarshalBinary(b []byte) error {
	var res SearchPodsListNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchPodsListOKBody search pods list o k body
swagger:model SearchPodsListOKBody
*/
type SearchPodsListOKBody struct {

	// data
	Data []*SearchPodsListOKBodyDataItems0 `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this search pods list o k body
func (o *SearchPodsListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchPodsListOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("searchPodsListOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("searchPodsListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this search pods list o k body based on the context it is used
func (o *SearchPodsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchPodsListOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("searchPodsListOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("searchPodsListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SearchPodsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchPodsListOKBody) UnmarshalBinary(b []byte) error {
	var res SearchPodsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchPodsListOKBodyDataItems0 search pods list o k body data items0
swagger:model SearchPodsListOKBodyDataItems0
*/
type SearchPodsListOKBodyDataItems0 struct {

	// metadata name
	MetadataName string `json:"metadataName,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this search pods list o k body data items0
func (o *SearchPodsListOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search pods list o k body data items0 based on context it is used
func (o *SearchPodsListOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchPodsListOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchPodsListOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res SearchPodsListOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchPodsListUnauthorizedBody search pods list unauthorized body
swagger:model SearchPodsListUnauthorizedBody
*/
type SearchPodsListUnauthorizedBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this search pods list unauthorized body
func (o *SearchPodsListUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search pods list unauthorized body based on context it is used
func (o *SearchPodsListUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchPodsListUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchPodsListUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res SearchPodsListUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
