// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// StandAloneProfileDropdownListReader is a Reader for the StandAloneProfileDropdownList structure.
type StandAloneProfileDropdownListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneProfileDropdownListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneProfileDropdownListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneProfileDropdownListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneProfileDropdownListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneProfileDropdownListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneProfileDropdownListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneProfileDropdownListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneProfileDropdownListOK creates a StandAloneProfileDropdownListOK with default headers values
func NewStandAloneProfileDropdownListOK() *StandAloneProfileDropdownListOK {
	return &StandAloneProfileDropdownListOK{}
}

/*
StandAloneProfileDropdownListOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneProfileDropdownListOK struct {
	Payload []*models.CommonDropdownDto
}

// IsSuccess returns true when this stand alone profile dropdown list o k response has a 2xx status code
func (o *StandAloneProfileDropdownListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone profile dropdown list o k response has a 3xx status code
func (o *StandAloneProfileDropdownListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile dropdown list o k response has a 4xx status code
func (o *StandAloneProfileDropdownListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone profile dropdown list o k response has a 5xx status code
func (o *StandAloneProfileDropdownListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile dropdown list o k response a status code equal to that given
func (o *StandAloneProfileDropdownListOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneProfileDropdownListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile/list][%d] standAloneProfileDropdownListOK  %+v", 200, o.Payload)
}

func (o *StandAloneProfileDropdownListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile/list][%d] standAloneProfileDropdownListOK  %+v", 200, o.Payload)
}

func (o *StandAloneProfileDropdownListOK) GetPayload() []*models.CommonDropdownDto {
	return o.Payload
}

func (o *StandAloneProfileDropdownListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileDropdownListBadRequest creates a StandAloneProfileDropdownListBadRequest with default headers values
func NewStandAloneProfileDropdownListBadRequest() *StandAloneProfileDropdownListBadRequest {
	return &StandAloneProfileDropdownListBadRequest{}
}

/*
StandAloneProfileDropdownListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneProfileDropdownListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this stand alone profile dropdown list bad request response has a 2xx status code
func (o *StandAloneProfileDropdownListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile dropdown list bad request response has a 3xx status code
func (o *StandAloneProfileDropdownListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile dropdown list bad request response has a 4xx status code
func (o *StandAloneProfileDropdownListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone profile dropdown list bad request response has a 5xx status code
func (o *StandAloneProfileDropdownListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile dropdown list bad request response a status code equal to that given
func (o *StandAloneProfileDropdownListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneProfileDropdownListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile/list][%d] standAloneProfileDropdownListBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneProfileDropdownListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile/list][%d] standAloneProfileDropdownListBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneProfileDropdownListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *StandAloneProfileDropdownListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileDropdownListUnauthorized creates a StandAloneProfileDropdownListUnauthorized with default headers values
func NewStandAloneProfileDropdownListUnauthorized() *StandAloneProfileDropdownListUnauthorized {
	return &StandAloneProfileDropdownListUnauthorized{}
}

/*
StandAloneProfileDropdownListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneProfileDropdownListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone profile dropdown list unauthorized response has a 2xx status code
func (o *StandAloneProfileDropdownListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile dropdown list unauthorized response has a 3xx status code
func (o *StandAloneProfileDropdownListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile dropdown list unauthorized response has a 4xx status code
func (o *StandAloneProfileDropdownListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone profile dropdown list unauthorized response has a 5xx status code
func (o *StandAloneProfileDropdownListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile dropdown list unauthorized response a status code equal to that given
func (o *StandAloneProfileDropdownListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneProfileDropdownListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile/list][%d] standAloneProfileDropdownListUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneProfileDropdownListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile/list][%d] standAloneProfileDropdownListUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneProfileDropdownListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneProfileDropdownListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileDropdownListForbidden creates a StandAloneProfileDropdownListForbidden with default headers values
func NewStandAloneProfileDropdownListForbidden() *StandAloneProfileDropdownListForbidden {
	return &StandAloneProfileDropdownListForbidden{}
}

/*
StandAloneProfileDropdownListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneProfileDropdownListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone profile dropdown list forbidden response has a 2xx status code
func (o *StandAloneProfileDropdownListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile dropdown list forbidden response has a 3xx status code
func (o *StandAloneProfileDropdownListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile dropdown list forbidden response has a 4xx status code
func (o *StandAloneProfileDropdownListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone profile dropdown list forbidden response has a 5xx status code
func (o *StandAloneProfileDropdownListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile dropdown list forbidden response a status code equal to that given
func (o *StandAloneProfileDropdownListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneProfileDropdownListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile/list][%d] standAloneProfileDropdownListForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneProfileDropdownListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile/list][%d] standAloneProfileDropdownListForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneProfileDropdownListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneProfileDropdownListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileDropdownListNotFound creates a StandAloneProfileDropdownListNotFound with default headers values
func NewStandAloneProfileDropdownListNotFound() *StandAloneProfileDropdownListNotFound {
	return &StandAloneProfileDropdownListNotFound{}
}

/*
StandAloneProfileDropdownListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneProfileDropdownListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone profile dropdown list not found response has a 2xx status code
func (o *StandAloneProfileDropdownListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile dropdown list not found response has a 3xx status code
func (o *StandAloneProfileDropdownListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile dropdown list not found response has a 4xx status code
func (o *StandAloneProfileDropdownListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone profile dropdown list not found response has a 5xx status code
func (o *StandAloneProfileDropdownListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile dropdown list not found response a status code equal to that given
func (o *StandAloneProfileDropdownListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneProfileDropdownListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile/list][%d] standAloneProfileDropdownListNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneProfileDropdownListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile/list][%d] standAloneProfileDropdownListNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneProfileDropdownListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneProfileDropdownListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileDropdownListInternalServerError creates a StandAloneProfileDropdownListInternalServerError with default headers values
func NewStandAloneProfileDropdownListInternalServerError() *StandAloneProfileDropdownListInternalServerError {
	return &StandAloneProfileDropdownListInternalServerError{}
}

/*
StandAloneProfileDropdownListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneProfileDropdownListInternalServerError struct {
}

// IsSuccess returns true when this stand alone profile dropdown list internal server error response has a 2xx status code
func (o *StandAloneProfileDropdownListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile dropdown list internal server error response has a 3xx status code
func (o *StandAloneProfileDropdownListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile dropdown list internal server error response has a 4xx status code
func (o *StandAloneProfileDropdownListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone profile dropdown list internal server error response has a 5xx status code
func (o *StandAloneProfileDropdownListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone profile dropdown list internal server error response a status code equal to that given
func (o *StandAloneProfileDropdownListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneProfileDropdownListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile/list][%d] standAloneProfileDropdownListInternalServerError ", 500)
}

func (o *StandAloneProfileDropdownListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile/list][%d] standAloneProfileDropdownListInternalServerError ", 500)
}

func (o *StandAloneProfileDropdownListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
