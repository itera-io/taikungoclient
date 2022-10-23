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

// StandAloneProfileEditReader is a Reader for the StandAloneProfileEdit structure.
type StandAloneProfileEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneProfileEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneProfileEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneProfileEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneProfileEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneProfileEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneProfileEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneProfileEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneProfileEditOK creates a StandAloneProfileEditOK with default headers values
func NewStandAloneProfileEditOK() *StandAloneProfileEditOK {
	return &StandAloneProfileEditOK{}
}

/*
StandAloneProfileEditOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneProfileEditOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone profile edit o k response has a 2xx status code
func (o *StandAloneProfileEditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone profile edit o k response has a 3xx status code
func (o *StandAloneProfileEditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile edit o k response has a 4xx status code
func (o *StandAloneProfileEditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone profile edit o k response has a 5xx status code
func (o *StandAloneProfileEditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile edit o k response a status code equal to that given
func (o *StandAloneProfileEditOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneProfileEditOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/edit][%d] standAloneProfileEditOK  %+v", 200, o.Payload)
}

func (o *StandAloneProfileEditOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/edit][%d] standAloneProfileEditOK  %+v", 200, o.Payload)
}

func (o *StandAloneProfileEditOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneProfileEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileEditBadRequest creates a StandAloneProfileEditBadRequest with default headers values
func NewStandAloneProfileEditBadRequest() *StandAloneProfileEditBadRequest {
	return &StandAloneProfileEditBadRequest{}
}

/*
StandAloneProfileEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneProfileEditBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this stand alone profile edit bad request response has a 2xx status code
func (o *StandAloneProfileEditBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile edit bad request response has a 3xx status code
func (o *StandAloneProfileEditBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile edit bad request response has a 4xx status code
func (o *StandAloneProfileEditBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone profile edit bad request response has a 5xx status code
func (o *StandAloneProfileEditBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile edit bad request response a status code equal to that given
func (o *StandAloneProfileEditBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneProfileEditBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/edit][%d] standAloneProfileEditBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneProfileEditBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/edit][%d] standAloneProfileEditBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneProfileEditBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *StandAloneProfileEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileEditUnauthorized creates a StandAloneProfileEditUnauthorized with default headers values
func NewStandAloneProfileEditUnauthorized() *StandAloneProfileEditUnauthorized {
	return &StandAloneProfileEditUnauthorized{}
}

/*
StandAloneProfileEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneProfileEditUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone profile edit unauthorized response has a 2xx status code
func (o *StandAloneProfileEditUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile edit unauthorized response has a 3xx status code
func (o *StandAloneProfileEditUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile edit unauthorized response has a 4xx status code
func (o *StandAloneProfileEditUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone profile edit unauthorized response has a 5xx status code
func (o *StandAloneProfileEditUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile edit unauthorized response a status code equal to that given
func (o *StandAloneProfileEditUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneProfileEditUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/edit][%d] standAloneProfileEditUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneProfileEditUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/edit][%d] standAloneProfileEditUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneProfileEditUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneProfileEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileEditForbidden creates a StandAloneProfileEditForbidden with default headers values
func NewStandAloneProfileEditForbidden() *StandAloneProfileEditForbidden {
	return &StandAloneProfileEditForbidden{}
}

/*
StandAloneProfileEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneProfileEditForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone profile edit forbidden response has a 2xx status code
func (o *StandAloneProfileEditForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile edit forbidden response has a 3xx status code
func (o *StandAloneProfileEditForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile edit forbidden response has a 4xx status code
func (o *StandAloneProfileEditForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone profile edit forbidden response has a 5xx status code
func (o *StandAloneProfileEditForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile edit forbidden response a status code equal to that given
func (o *StandAloneProfileEditForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneProfileEditForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/edit][%d] standAloneProfileEditForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneProfileEditForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/edit][%d] standAloneProfileEditForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneProfileEditForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneProfileEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileEditNotFound creates a StandAloneProfileEditNotFound with default headers values
func NewStandAloneProfileEditNotFound() *StandAloneProfileEditNotFound {
	return &StandAloneProfileEditNotFound{}
}

/*
StandAloneProfileEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneProfileEditNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone profile edit not found response has a 2xx status code
func (o *StandAloneProfileEditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile edit not found response has a 3xx status code
func (o *StandAloneProfileEditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile edit not found response has a 4xx status code
func (o *StandAloneProfileEditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone profile edit not found response has a 5xx status code
func (o *StandAloneProfileEditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile edit not found response a status code equal to that given
func (o *StandAloneProfileEditNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneProfileEditNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/edit][%d] standAloneProfileEditNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneProfileEditNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/edit][%d] standAloneProfileEditNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneProfileEditNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneProfileEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileEditInternalServerError creates a StandAloneProfileEditInternalServerError with default headers values
func NewStandAloneProfileEditInternalServerError() *StandAloneProfileEditInternalServerError {
	return &StandAloneProfileEditInternalServerError{}
}

/*
StandAloneProfileEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneProfileEditInternalServerError struct {
}

// IsSuccess returns true when this stand alone profile edit internal server error response has a 2xx status code
func (o *StandAloneProfileEditInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile edit internal server error response has a 3xx status code
func (o *StandAloneProfileEditInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile edit internal server error response has a 4xx status code
func (o *StandAloneProfileEditInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone profile edit internal server error response has a 5xx status code
func (o *StandAloneProfileEditInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone profile edit internal server error response a status code equal to that given
func (o *StandAloneProfileEditInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneProfileEditInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/edit][%d] standAloneProfileEditInternalServerError ", 500)
}

func (o *StandAloneProfileEditInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/edit][%d] standAloneProfileEditInternalServerError ", 500)
}

func (o *StandAloneProfileEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
