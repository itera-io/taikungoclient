// Code generated by go-swagger; DO NOT EDIT.

package showback_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ShowbackRulesUpdateReader is a Reader for the ShowbackRulesUpdate structure.
type ShowbackRulesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackRulesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackRulesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackRulesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackRulesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackRulesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackRulesUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackRulesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackRulesUpdateOK creates a ShowbackRulesUpdateOK with default headers values
func NewShowbackRulesUpdateOK() *ShowbackRulesUpdateOK {
	return &ShowbackRulesUpdateOK{}
}

/*
ShowbackRulesUpdateOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackRulesUpdateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this showback rules update o k response has a 2xx status code
func (o *ShowbackRulesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this showback rules update o k response has a 3xx status code
func (o *ShowbackRulesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules update o k response has a 4xx status code
func (o *ShowbackRulesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback rules update o k response has a 5xx status code
func (o *ShowbackRulesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules update o k response a status code equal to that given
func (o *ShowbackRulesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ShowbackRulesUpdateOK) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/update][%d] showbackRulesUpdateOK  %+v", 200, o.Payload)
}

func (o *ShowbackRulesUpdateOK) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/update][%d] showbackRulesUpdateOK  %+v", 200, o.Payload)
}

func (o *ShowbackRulesUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ShowbackRulesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesUpdateBadRequest creates a ShowbackRulesUpdateBadRequest with default headers values
func NewShowbackRulesUpdateBadRequest() *ShowbackRulesUpdateBadRequest {
	return &ShowbackRulesUpdateBadRequest{}
}

/*
ShowbackRulesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackRulesUpdateBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this showback rules update bad request response has a 2xx status code
func (o *ShowbackRulesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules update bad request response has a 3xx status code
func (o *ShowbackRulesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules update bad request response has a 4xx status code
func (o *ShowbackRulesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules update bad request response has a 5xx status code
func (o *ShowbackRulesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules update bad request response a status code equal to that given
func (o *ShowbackRulesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ShowbackRulesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/update][%d] showbackRulesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackRulesUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/update][%d] showbackRulesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackRulesUpdateBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ShowbackRulesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesUpdateUnauthorized creates a ShowbackRulesUpdateUnauthorized with default headers values
func NewShowbackRulesUpdateUnauthorized() *ShowbackRulesUpdateUnauthorized {
	return &ShowbackRulesUpdateUnauthorized{}
}

/*
ShowbackRulesUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackRulesUpdateUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this showback rules update unauthorized response has a 2xx status code
func (o *ShowbackRulesUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules update unauthorized response has a 3xx status code
func (o *ShowbackRulesUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules update unauthorized response has a 4xx status code
func (o *ShowbackRulesUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules update unauthorized response has a 5xx status code
func (o *ShowbackRulesUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules update unauthorized response a status code equal to that given
func (o *ShowbackRulesUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ShowbackRulesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/update][%d] showbackRulesUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackRulesUpdateUnauthorized) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/update][%d] showbackRulesUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackRulesUpdateUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ShowbackRulesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesUpdateForbidden creates a ShowbackRulesUpdateForbidden with default headers values
func NewShowbackRulesUpdateForbidden() *ShowbackRulesUpdateForbidden {
	return &ShowbackRulesUpdateForbidden{}
}

/*
ShowbackRulesUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackRulesUpdateForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this showback rules update forbidden response has a 2xx status code
func (o *ShowbackRulesUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules update forbidden response has a 3xx status code
func (o *ShowbackRulesUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules update forbidden response has a 4xx status code
func (o *ShowbackRulesUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules update forbidden response has a 5xx status code
func (o *ShowbackRulesUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules update forbidden response a status code equal to that given
func (o *ShowbackRulesUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ShowbackRulesUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/update][%d] showbackRulesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackRulesUpdateForbidden) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/update][%d] showbackRulesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackRulesUpdateForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ShowbackRulesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesUpdateNotFound creates a ShowbackRulesUpdateNotFound with default headers values
func NewShowbackRulesUpdateNotFound() *ShowbackRulesUpdateNotFound {
	return &ShowbackRulesUpdateNotFound{}
}

/*
ShowbackRulesUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackRulesUpdateNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this showback rules update not found response has a 2xx status code
func (o *ShowbackRulesUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules update not found response has a 3xx status code
func (o *ShowbackRulesUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules update not found response has a 4xx status code
func (o *ShowbackRulesUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules update not found response has a 5xx status code
func (o *ShowbackRulesUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules update not found response a status code equal to that given
func (o *ShowbackRulesUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ShowbackRulesUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/update][%d] showbackRulesUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackRulesUpdateNotFound) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/update][%d] showbackRulesUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackRulesUpdateNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ShowbackRulesUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesUpdateInternalServerError creates a ShowbackRulesUpdateInternalServerError with default headers values
func NewShowbackRulesUpdateInternalServerError() *ShowbackRulesUpdateInternalServerError {
	return &ShowbackRulesUpdateInternalServerError{}
}

/*
ShowbackRulesUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackRulesUpdateInternalServerError struct {
}

// IsSuccess returns true when this showback rules update internal server error response has a 2xx status code
func (o *ShowbackRulesUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules update internal server error response has a 3xx status code
func (o *ShowbackRulesUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules update internal server error response has a 4xx status code
func (o *ShowbackRulesUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback rules update internal server error response has a 5xx status code
func (o *ShowbackRulesUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this showback rules update internal server error response a status code equal to that given
func (o *ShowbackRulesUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ShowbackRulesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/update][%d] showbackRulesUpdateInternalServerError ", 500)
}

func (o *ShowbackRulesUpdateInternalServerError) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/update][%d] showbackRulesUpdateInternalServerError ", 500)
}

func (o *ShowbackRulesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
