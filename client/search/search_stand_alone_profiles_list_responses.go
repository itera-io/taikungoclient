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

// SearchStandAloneProfilesListReader is a Reader for the SearchStandAloneProfilesList structure.
type SearchStandAloneProfilesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchStandAloneProfilesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchStandAloneProfilesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchStandAloneProfilesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchStandAloneProfilesListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchStandAloneProfilesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchStandAloneProfilesListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchStandAloneProfilesListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchStandAloneProfilesListOK creates a SearchStandAloneProfilesListOK with default headers values
func NewSearchStandAloneProfilesListOK() *SearchStandAloneProfilesListOK {
	return &SearchStandAloneProfilesListOK{}
}

/*
SearchStandAloneProfilesListOK describes a response with status code 200, with default header values.

Success
*/
type SearchStandAloneProfilesListOK struct {
	Payload *SearchStandAloneProfilesListOKBody
}

// IsSuccess returns true when this search stand alone profiles list o k response has a 2xx status code
func (o *SearchStandAloneProfilesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search stand alone profiles list o k response has a 3xx status code
func (o *SearchStandAloneProfilesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search stand alone profiles list o k response has a 4xx status code
func (o *SearchStandAloneProfilesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search stand alone profiles list o k response has a 5xx status code
func (o *SearchStandAloneProfilesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search stand alone profiles list o k response a status code equal to that given
func (o *SearchStandAloneProfilesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchStandAloneProfilesListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/stand-alone-profiles][%d] searchStandAloneProfilesListOK  %+v", 200, o.Payload)
}

func (o *SearchStandAloneProfilesListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/stand-alone-profiles][%d] searchStandAloneProfilesListOK  %+v", 200, o.Payload)
}

func (o *SearchStandAloneProfilesListOK) GetPayload() *SearchStandAloneProfilesListOKBody {
	return o.Payload
}

func (o *SearchStandAloneProfilesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchStandAloneProfilesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchStandAloneProfilesListBadRequest creates a SearchStandAloneProfilesListBadRequest with default headers values
func NewSearchStandAloneProfilesListBadRequest() *SearchStandAloneProfilesListBadRequest {
	return &SearchStandAloneProfilesListBadRequest{}
}

/*
SearchStandAloneProfilesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchStandAloneProfilesListBadRequest struct {
	Payload []*SearchStandAloneProfilesListBadRequestBodyItems0
}

// IsSuccess returns true when this search stand alone profiles list bad request response has a 2xx status code
func (o *SearchStandAloneProfilesListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search stand alone profiles list bad request response has a 3xx status code
func (o *SearchStandAloneProfilesListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search stand alone profiles list bad request response has a 4xx status code
func (o *SearchStandAloneProfilesListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search stand alone profiles list bad request response has a 5xx status code
func (o *SearchStandAloneProfilesListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search stand alone profiles list bad request response a status code equal to that given
func (o *SearchStandAloneProfilesListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchStandAloneProfilesListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/stand-alone-profiles][%d] searchStandAloneProfilesListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchStandAloneProfilesListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/stand-alone-profiles][%d] searchStandAloneProfilesListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchStandAloneProfilesListBadRequest) GetPayload() []*SearchStandAloneProfilesListBadRequestBodyItems0 {
	return o.Payload
}

func (o *SearchStandAloneProfilesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchStandAloneProfilesListUnauthorized creates a SearchStandAloneProfilesListUnauthorized with default headers values
func NewSearchStandAloneProfilesListUnauthorized() *SearchStandAloneProfilesListUnauthorized {
	return &SearchStandAloneProfilesListUnauthorized{}
}

/*
SearchStandAloneProfilesListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchStandAloneProfilesListUnauthorized struct {
	Payload *SearchStandAloneProfilesListUnauthorizedBody
}

// IsSuccess returns true when this search stand alone profiles list unauthorized response has a 2xx status code
func (o *SearchStandAloneProfilesListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search stand alone profiles list unauthorized response has a 3xx status code
func (o *SearchStandAloneProfilesListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search stand alone profiles list unauthorized response has a 4xx status code
func (o *SearchStandAloneProfilesListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search stand alone profiles list unauthorized response has a 5xx status code
func (o *SearchStandAloneProfilesListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search stand alone profiles list unauthorized response a status code equal to that given
func (o *SearchStandAloneProfilesListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchStandAloneProfilesListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/stand-alone-profiles][%d] searchStandAloneProfilesListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchStandAloneProfilesListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/stand-alone-profiles][%d] searchStandAloneProfilesListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchStandAloneProfilesListUnauthorized) GetPayload() *SearchStandAloneProfilesListUnauthorizedBody {
	return o.Payload
}

func (o *SearchStandAloneProfilesListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchStandAloneProfilesListUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchStandAloneProfilesListForbidden creates a SearchStandAloneProfilesListForbidden with default headers values
func NewSearchStandAloneProfilesListForbidden() *SearchStandAloneProfilesListForbidden {
	return &SearchStandAloneProfilesListForbidden{}
}

/*
SearchStandAloneProfilesListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchStandAloneProfilesListForbidden struct {
	Payload *SearchStandAloneProfilesListForbiddenBody
}

// IsSuccess returns true when this search stand alone profiles list forbidden response has a 2xx status code
func (o *SearchStandAloneProfilesListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search stand alone profiles list forbidden response has a 3xx status code
func (o *SearchStandAloneProfilesListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search stand alone profiles list forbidden response has a 4xx status code
func (o *SearchStandAloneProfilesListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search stand alone profiles list forbidden response has a 5xx status code
func (o *SearchStandAloneProfilesListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search stand alone profiles list forbidden response a status code equal to that given
func (o *SearchStandAloneProfilesListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchStandAloneProfilesListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/stand-alone-profiles][%d] searchStandAloneProfilesListForbidden  %+v", 403, o.Payload)
}

func (o *SearchStandAloneProfilesListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/stand-alone-profiles][%d] searchStandAloneProfilesListForbidden  %+v", 403, o.Payload)
}

func (o *SearchStandAloneProfilesListForbidden) GetPayload() *SearchStandAloneProfilesListForbiddenBody {
	return o.Payload
}

func (o *SearchStandAloneProfilesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchStandAloneProfilesListForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchStandAloneProfilesListNotFound creates a SearchStandAloneProfilesListNotFound with default headers values
func NewSearchStandAloneProfilesListNotFound() *SearchStandAloneProfilesListNotFound {
	return &SearchStandAloneProfilesListNotFound{}
}

/*
SearchStandAloneProfilesListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchStandAloneProfilesListNotFound struct {
	Payload *SearchStandAloneProfilesListNotFoundBody
}

// IsSuccess returns true when this search stand alone profiles list not found response has a 2xx status code
func (o *SearchStandAloneProfilesListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search stand alone profiles list not found response has a 3xx status code
func (o *SearchStandAloneProfilesListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search stand alone profiles list not found response has a 4xx status code
func (o *SearchStandAloneProfilesListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search stand alone profiles list not found response has a 5xx status code
func (o *SearchStandAloneProfilesListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search stand alone profiles list not found response a status code equal to that given
func (o *SearchStandAloneProfilesListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchStandAloneProfilesListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/stand-alone-profiles][%d] searchStandAloneProfilesListNotFound  %+v", 404, o.Payload)
}

func (o *SearchStandAloneProfilesListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/stand-alone-profiles][%d] searchStandAloneProfilesListNotFound  %+v", 404, o.Payload)
}

func (o *SearchStandAloneProfilesListNotFound) GetPayload() *SearchStandAloneProfilesListNotFoundBody {
	return o.Payload
}

func (o *SearchStandAloneProfilesListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchStandAloneProfilesListNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchStandAloneProfilesListInternalServerError creates a SearchStandAloneProfilesListInternalServerError with default headers values
func NewSearchStandAloneProfilesListInternalServerError() *SearchStandAloneProfilesListInternalServerError {
	return &SearchStandAloneProfilesListInternalServerError{}
}

/*
SearchStandAloneProfilesListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchStandAloneProfilesListInternalServerError struct {
}

// IsSuccess returns true when this search stand alone profiles list internal server error response has a 2xx status code
func (o *SearchStandAloneProfilesListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search stand alone profiles list internal server error response has a 3xx status code
func (o *SearchStandAloneProfilesListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search stand alone profiles list internal server error response has a 4xx status code
func (o *SearchStandAloneProfilesListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search stand alone profiles list internal server error response has a 5xx status code
func (o *SearchStandAloneProfilesListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search stand alone profiles list internal server error response a status code equal to that given
func (o *SearchStandAloneProfilesListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchStandAloneProfilesListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/stand-alone-profiles][%d] searchStandAloneProfilesListInternalServerError ", 500)
}

func (o *SearchStandAloneProfilesListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/stand-alone-profiles][%d] searchStandAloneProfilesListInternalServerError ", 500)
}

func (o *SearchStandAloneProfilesListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
SearchStandAloneProfilesListBadRequestBodyItems0 search stand alone profiles list bad request body items0
swagger:model SearchStandAloneProfilesListBadRequestBodyItems0
*/
type SearchStandAloneProfilesListBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this search stand alone profiles list bad request body items0
func (o *SearchStandAloneProfilesListBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search stand alone profiles list bad request body items0 based on context it is used
func (o *SearchStandAloneProfilesListBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchStandAloneProfilesListBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchStandAloneProfilesListBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res SearchStandAloneProfilesListBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchStandAloneProfilesListBody search stand alone profiles list body
swagger:model SearchStandAloneProfilesListBody
*/
type SearchStandAloneProfilesListBody struct {

	// limit
	Limit int32 `json:"limit,omitempty"`

	// offset
	Offset int32 `json:"offset,omitempty"`

	// search term
	SearchTerm string `json:"searchTerm,omitempty"`
}

// Validate validates this search stand alone profiles list body
func (o *SearchStandAloneProfilesListBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search stand alone profiles list body based on context it is used
func (o *SearchStandAloneProfilesListBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchStandAloneProfilesListBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchStandAloneProfilesListBody) UnmarshalBinary(b []byte) error {
	var res SearchStandAloneProfilesListBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchStandAloneProfilesListForbiddenBody search stand alone profiles list forbidden body
swagger:model SearchStandAloneProfilesListForbiddenBody
*/
type SearchStandAloneProfilesListForbiddenBody struct {

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

// Validate validates this search stand alone profiles list forbidden body
func (o *SearchStandAloneProfilesListForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search stand alone profiles list forbidden body based on context it is used
func (o *SearchStandAloneProfilesListForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchStandAloneProfilesListForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchStandAloneProfilesListForbiddenBody) UnmarshalBinary(b []byte) error {
	var res SearchStandAloneProfilesListForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchStandAloneProfilesListNotFoundBody search stand alone profiles list not found body
swagger:model SearchStandAloneProfilesListNotFoundBody
*/
type SearchStandAloneProfilesListNotFoundBody struct {

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

// Validate validates this search stand alone profiles list not found body
func (o *SearchStandAloneProfilesListNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search stand alone profiles list not found body based on context it is used
func (o *SearchStandAloneProfilesListNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchStandAloneProfilesListNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchStandAloneProfilesListNotFoundBody) UnmarshalBinary(b []byte) error {
	var res SearchStandAloneProfilesListNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchStandAloneProfilesListOKBody search stand alone profiles list o k body
swagger:model SearchStandAloneProfilesListOKBody
*/
type SearchStandAloneProfilesListOKBody struct {

	// data
	Data []*SearchStandAloneProfilesListOKBodyDataItems0 `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this search stand alone profiles list o k body
func (o *SearchStandAloneProfilesListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchStandAloneProfilesListOKBody) validateData(formats strfmt.Registry) error {
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
					return ve.ValidateName("searchStandAloneProfilesListOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("searchStandAloneProfilesListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this search stand alone profiles list o k body based on the context it is used
func (o *SearchStandAloneProfilesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchStandAloneProfilesListOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("searchStandAloneProfilesListOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("searchStandAloneProfilesListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SearchStandAloneProfilesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchStandAloneProfilesListOKBody) UnmarshalBinary(b []byte) error {
	var res SearchStandAloneProfilesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchStandAloneProfilesListOKBodyDataItems0 search stand alone profiles list o k body data items0
swagger:model SearchStandAloneProfilesListOKBodyDataItems0
*/
type SearchStandAloneProfilesListOKBodyDataItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`
}

// Validate validates this search stand alone profiles list o k body data items0
func (o *SearchStandAloneProfilesListOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search stand alone profiles list o k body data items0 based on context it is used
func (o *SearchStandAloneProfilesListOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchStandAloneProfilesListOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchStandAloneProfilesListOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res SearchStandAloneProfilesListOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchStandAloneProfilesListUnauthorizedBody search stand alone profiles list unauthorized body
swagger:model SearchStandAloneProfilesListUnauthorizedBody
*/
type SearchStandAloneProfilesListUnauthorizedBody struct {

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

// Validate validates this search stand alone profiles list unauthorized body
func (o *SearchStandAloneProfilesListUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search stand alone profiles list unauthorized body based on context it is used
func (o *SearchStandAloneProfilesListUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchStandAloneProfilesListUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchStandAloneProfilesListUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res SearchStandAloneProfilesListUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
