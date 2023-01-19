// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// StandAloneActionsInstancePasswordReader is a Reader for the StandAloneActionsInstancePassword structure.
type StandAloneActionsInstancePasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneActionsInstancePasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneActionsInstancePasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneActionsInstancePasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneActionsInstancePasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneActionsInstancePasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneActionsInstancePasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneActionsInstancePasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneActionsInstancePasswordOK creates a StandAloneActionsInstancePasswordOK with default headers values
func NewStandAloneActionsInstancePasswordOK() *StandAloneActionsInstancePasswordOK {
	return &StandAloneActionsInstancePasswordOK{}
}

/*
StandAloneActionsInstancePasswordOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneActionsInstancePasswordOK struct {
	Payload string
}

// IsSuccess returns true when this stand alone actions instance password o k response has a 2xx status code
func (o *StandAloneActionsInstancePasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone actions instance password o k response has a 3xx status code
func (o *StandAloneActionsInstancePasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions instance password o k response has a 4xx status code
func (o *StandAloneActionsInstancePasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions instance password o k response has a 5xx status code
func (o *StandAloneActionsInstancePasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions instance password o k response a status code equal to that given
func (o *StandAloneActionsInstancePasswordOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stand alone actions instance password o k response
func (o *StandAloneActionsInstancePasswordOK) Code() int {
	return 200
}

func (o *StandAloneActionsInstancePasswordOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/password][%d] standAloneActionsInstancePasswordOK  %+v", 200, o.Payload)
}

func (o *StandAloneActionsInstancePasswordOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/password][%d] standAloneActionsInstancePasswordOK  %+v", 200, o.Payload)
}

func (o *StandAloneActionsInstancePasswordOK) GetPayload() string {
	return o.Payload
}

func (o *StandAloneActionsInstancePasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsInstancePasswordBadRequest creates a StandAloneActionsInstancePasswordBadRequest with default headers values
func NewStandAloneActionsInstancePasswordBadRequest() *StandAloneActionsInstancePasswordBadRequest {
	return &StandAloneActionsInstancePasswordBadRequest{}
}

/*
StandAloneActionsInstancePasswordBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneActionsInstancePasswordBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone actions instance password bad request response has a 2xx status code
func (o *StandAloneActionsInstancePasswordBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions instance password bad request response has a 3xx status code
func (o *StandAloneActionsInstancePasswordBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions instance password bad request response has a 4xx status code
func (o *StandAloneActionsInstancePasswordBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions instance password bad request response has a 5xx status code
func (o *StandAloneActionsInstancePasswordBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions instance password bad request response a status code equal to that given
func (o *StandAloneActionsInstancePasswordBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stand alone actions instance password bad request response
func (o *StandAloneActionsInstancePasswordBadRequest) Code() int {
	return 400
}

func (o *StandAloneActionsInstancePasswordBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/password][%d] standAloneActionsInstancePasswordBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsInstancePasswordBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/password][%d] standAloneActionsInstancePasswordBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsInstancePasswordBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsInstancePasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsInstancePasswordUnauthorized creates a StandAloneActionsInstancePasswordUnauthorized with default headers values
func NewStandAloneActionsInstancePasswordUnauthorized() *StandAloneActionsInstancePasswordUnauthorized {
	return &StandAloneActionsInstancePasswordUnauthorized{}
}

/*
StandAloneActionsInstancePasswordUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneActionsInstancePasswordUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone actions instance password unauthorized response has a 2xx status code
func (o *StandAloneActionsInstancePasswordUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions instance password unauthorized response has a 3xx status code
func (o *StandAloneActionsInstancePasswordUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions instance password unauthorized response has a 4xx status code
func (o *StandAloneActionsInstancePasswordUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions instance password unauthorized response has a 5xx status code
func (o *StandAloneActionsInstancePasswordUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions instance password unauthorized response a status code equal to that given
func (o *StandAloneActionsInstancePasswordUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stand alone actions instance password unauthorized response
func (o *StandAloneActionsInstancePasswordUnauthorized) Code() int {
	return 401
}

func (o *StandAloneActionsInstancePasswordUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/password][%d] standAloneActionsInstancePasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneActionsInstancePasswordUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/password][%d] standAloneActionsInstancePasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneActionsInstancePasswordUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsInstancePasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsInstancePasswordForbidden creates a StandAloneActionsInstancePasswordForbidden with default headers values
func NewStandAloneActionsInstancePasswordForbidden() *StandAloneActionsInstancePasswordForbidden {
	return &StandAloneActionsInstancePasswordForbidden{}
}

/*
StandAloneActionsInstancePasswordForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneActionsInstancePasswordForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone actions instance password forbidden response has a 2xx status code
func (o *StandAloneActionsInstancePasswordForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions instance password forbidden response has a 3xx status code
func (o *StandAloneActionsInstancePasswordForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions instance password forbidden response has a 4xx status code
func (o *StandAloneActionsInstancePasswordForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions instance password forbidden response has a 5xx status code
func (o *StandAloneActionsInstancePasswordForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions instance password forbidden response a status code equal to that given
func (o *StandAloneActionsInstancePasswordForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stand alone actions instance password forbidden response
func (o *StandAloneActionsInstancePasswordForbidden) Code() int {
	return 403
}

func (o *StandAloneActionsInstancePasswordForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/password][%d] standAloneActionsInstancePasswordForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneActionsInstancePasswordForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/password][%d] standAloneActionsInstancePasswordForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneActionsInstancePasswordForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsInstancePasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsInstancePasswordNotFound creates a StandAloneActionsInstancePasswordNotFound with default headers values
func NewStandAloneActionsInstancePasswordNotFound() *StandAloneActionsInstancePasswordNotFound {
	return &StandAloneActionsInstancePasswordNotFound{}
}

/*
StandAloneActionsInstancePasswordNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneActionsInstancePasswordNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone actions instance password not found response has a 2xx status code
func (o *StandAloneActionsInstancePasswordNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions instance password not found response has a 3xx status code
func (o *StandAloneActionsInstancePasswordNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions instance password not found response has a 4xx status code
func (o *StandAloneActionsInstancePasswordNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions instance password not found response has a 5xx status code
func (o *StandAloneActionsInstancePasswordNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions instance password not found response a status code equal to that given
func (o *StandAloneActionsInstancePasswordNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stand alone actions instance password not found response
func (o *StandAloneActionsInstancePasswordNotFound) Code() int {
	return 404
}

func (o *StandAloneActionsInstancePasswordNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/password][%d] standAloneActionsInstancePasswordNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneActionsInstancePasswordNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/password][%d] standAloneActionsInstancePasswordNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneActionsInstancePasswordNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsInstancePasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsInstancePasswordInternalServerError creates a StandAloneActionsInstancePasswordInternalServerError with default headers values
func NewStandAloneActionsInstancePasswordInternalServerError() *StandAloneActionsInstancePasswordInternalServerError {
	return &StandAloneActionsInstancePasswordInternalServerError{}
}

/*
StandAloneActionsInstancePasswordInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneActionsInstancePasswordInternalServerError struct {
}

// IsSuccess returns true when this stand alone actions instance password internal server error response has a 2xx status code
func (o *StandAloneActionsInstancePasswordInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions instance password internal server error response has a 3xx status code
func (o *StandAloneActionsInstancePasswordInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions instance password internal server error response has a 4xx status code
func (o *StandAloneActionsInstancePasswordInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions instance password internal server error response has a 5xx status code
func (o *StandAloneActionsInstancePasswordInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone actions instance password internal server error response a status code equal to that given
func (o *StandAloneActionsInstancePasswordInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stand alone actions instance password internal server error response
func (o *StandAloneActionsInstancePasswordInternalServerError) Code() int {
	return 500
}

func (o *StandAloneActionsInstancePasswordInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/password][%d] standAloneActionsInstancePasswordInternalServerError ", 500)
}

func (o *StandAloneActionsInstancePasswordInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/password][%d] standAloneActionsInstancePasswordInternalServerError ", 500)
}

func (o *StandAloneActionsInstancePasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
