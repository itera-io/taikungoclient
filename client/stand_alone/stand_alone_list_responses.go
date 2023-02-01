// Code generated by go-swagger; DO NOT EDIT.

package stand_alone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// StandAloneListReader is a Reader for the StandAloneList structure.
type StandAloneListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneListOK creates a StandAloneListOK with default headers values
func NewStandAloneListOK() *StandAloneListOK {
	return &StandAloneListOK{}
}

/*
StandAloneListOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneListOK struct {
	Payload *models.StandaloneVmsList
}

// IsSuccess returns true when this stand alone list o k response has a 2xx status code
func (o *StandAloneListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone list o k response has a 3xx status code
func (o *StandAloneListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone list o k response has a 4xx status code
func (o *StandAloneListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone list o k response has a 5xx status code
func (o *StandAloneListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone list o k response a status code equal to that given
func (o *StandAloneListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stand alone list o k response
func (o *StandAloneListOK) Code() int {
	return 200
}

func (o *StandAloneListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone][%d] standAloneListOK  %+v", 200, o.Payload)
}

func (o *StandAloneListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone][%d] standAloneListOK  %+v", 200, o.Payload)
}

func (o *StandAloneListOK) GetPayload() *models.StandaloneVmsList {
	return o.Payload
}

func (o *StandAloneListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandaloneVmsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneListBadRequest creates a StandAloneListBadRequest with default headers values
func NewStandAloneListBadRequest() *StandAloneListBadRequest {
	return &StandAloneListBadRequest{}
}

/*
StandAloneListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone list bad request response has a 2xx status code
func (o *StandAloneListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone list bad request response has a 3xx status code
func (o *StandAloneListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone list bad request response has a 4xx status code
func (o *StandAloneListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone list bad request response has a 5xx status code
func (o *StandAloneListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone list bad request response a status code equal to that given
func (o *StandAloneListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stand alone list bad request response
func (o *StandAloneListBadRequest) Code() int {
	return 400
}

func (o *StandAloneListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone][%d] standAloneListBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone][%d] standAloneListBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneListUnauthorized creates a StandAloneListUnauthorized with default headers values
func NewStandAloneListUnauthorized() *StandAloneListUnauthorized {
	return &StandAloneListUnauthorized{}
}

/*
StandAloneListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone list unauthorized response has a 2xx status code
func (o *StandAloneListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone list unauthorized response has a 3xx status code
func (o *StandAloneListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone list unauthorized response has a 4xx status code
func (o *StandAloneListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone list unauthorized response has a 5xx status code
func (o *StandAloneListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone list unauthorized response a status code equal to that given
func (o *StandAloneListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stand alone list unauthorized response
func (o *StandAloneListUnauthorized) Code() int {
	return 401
}

func (o *StandAloneListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone][%d] standAloneListUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone][%d] standAloneListUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneListForbidden creates a StandAloneListForbidden with default headers values
func NewStandAloneListForbidden() *StandAloneListForbidden {
	return &StandAloneListForbidden{}
}

/*
StandAloneListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone list forbidden response has a 2xx status code
func (o *StandAloneListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone list forbidden response has a 3xx status code
func (o *StandAloneListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone list forbidden response has a 4xx status code
func (o *StandAloneListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone list forbidden response has a 5xx status code
func (o *StandAloneListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone list forbidden response a status code equal to that given
func (o *StandAloneListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stand alone list forbidden response
func (o *StandAloneListForbidden) Code() int {
	return 403
}

func (o *StandAloneListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone][%d] standAloneListForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone][%d] standAloneListForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneListNotFound creates a StandAloneListNotFound with default headers values
func NewStandAloneListNotFound() *StandAloneListNotFound {
	return &StandAloneListNotFound{}
}

/*
StandAloneListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone list not found response has a 2xx status code
func (o *StandAloneListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone list not found response has a 3xx status code
func (o *StandAloneListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone list not found response has a 4xx status code
func (o *StandAloneListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone list not found response has a 5xx status code
func (o *StandAloneListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone list not found response a status code equal to that given
func (o *StandAloneListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stand alone list not found response
func (o *StandAloneListNotFound) Code() int {
	return 404
}

func (o *StandAloneListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone][%d] standAloneListNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone][%d] standAloneListNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneListInternalServerError creates a StandAloneListInternalServerError with default headers values
func NewStandAloneListInternalServerError() *StandAloneListInternalServerError {
	return &StandAloneListInternalServerError{}
}

/*
StandAloneListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneListInternalServerError struct {
}

// IsSuccess returns true when this stand alone list internal server error response has a 2xx status code
func (o *StandAloneListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone list internal server error response has a 3xx status code
func (o *StandAloneListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone list internal server error response has a 4xx status code
func (o *StandAloneListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone list internal server error response has a 5xx status code
func (o *StandAloneListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone list internal server error response a status code equal to that given
func (o *StandAloneListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stand alone list internal server error response
func (o *StandAloneListInternalServerError) Code() int {
	return 500
}

func (o *StandAloneListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone][%d] standAloneListInternalServerError ", 500)
}

func (o *StandAloneListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone][%d] standAloneListInternalServerError ", 500)
}

func (o *StandAloneListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
