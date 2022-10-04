// Code generated by go-swagger; DO NOT EDIT.

package ops_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// OpsCredentialsMakeDefaultReader is a Reader for the OpsCredentialsMakeDefault structure.
type OpsCredentialsMakeDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpsCredentialsMakeDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpsCredentialsMakeDefaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpsCredentialsMakeDefaultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpsCredentialsMakeDefaultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpsCredentialsMakeDefaultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpsCredentialsMakeDefaultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpsCredentialsMakeDefaultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpsCredentialsMakeDefaultOK creates a OpsCredentialsMakeDefaultOK with default headers values
func NewOpsCredentialsMakeDefaultOK() *OpsCredentialsMakeDefaultOK {
	return &OpsCredentialsMakeDefaultOK{}
}

/*
OpsCredentialsMakeDefaultOK describes a response with status code 200, with default header values.

Success
*/
type OpsCredentialsMakeDefaultOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this ops credentials make default o k response has a 2xx status code
func (o *OpsCredentialsMakeDefaultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ops credentials make default o k response has a 3xx status code
func (o *OpsCredentialsMakeDefaultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials make default o k response has a 4xx status code
func (o *OpsCredentialsMakeDefaultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ops credentials make default o k response has a 5xx status code
func (o *OpsCredentialsMakeDefaultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials make default o k response a status code equal to that given
func (o *OpsCredentialsMakeDefaultOK) IsCode(code int) bool {
	return code == 200
}

func (o *OpsCredentialsMakeDefaultOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/makedefault][%d] opsCredentialsMakeDefaultOK  %+v", 200, o.Payload)
}

func (o *OpsCredentialsMakeDefaultOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/makedefault][%d] opsCredentialsMakeDefaultOK  %+v", 200, o.Payload)
}

func (o *OpsCredentialsMakeDefaultOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *OpsCredentialsMakeDefaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsMakeDefaultBadRequest creates a OpsCredentialsMakeDefaultBadRequest with default headers values
func NewOpsCredentialsMakeDefaultBadRequest() *OpsCredentialsMakeDefaultBadRequest {
	return &OpsCredentialsMakeDefaultBadRequest{}
}

/*
OpsCredentialsMakeDefaultBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpsCredentialsMakeDefaultBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this ops credentials make default bad request response has a 2xx status code
func (o *OpsCredentialsMakeDefaultBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials make default bad request response has a 3xx status code
func (o *OpsCredentialsMakeDefaultBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials make default bad request response has a 4xx status code
func (o *OpsCredentialsMakeDefaultBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ops credentials make default bad request response has a 5xx status code
func (o *OpsCredentialsMakeDefaultBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials make default bad request response a status code equal to that given
func (o *OpsCredentialsMakeDefaultBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OpsCredentialsMakeDefaultBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/makedefault][%d] opsCredentialsMakeDefaultBadRequest  %+v", 400, o.Payload)
}

func (o *OpsCredentialsMakeDefaultBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/makedefault][%d] opsCredentialsMakeDefaultBadRequest  %+v", 400, o.Payload)
}

func (o *OpsCredentialsMakeDefaultBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OpsCredentialsMakeDefaultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsMakeDefaultUnauthorized creates a OpsCredentialsMakeDefaultUnauthorized with default headers values
func NewOpsCredentialsMakeDefaultUnauthorized() *OpsCredentialsMakeDefaultUnauthorized {
	return &OpsCredentialsMakeDefaultUnauthorized{}
}

/*
OpsCredentialsMakeDefaultUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpsCredentialsMakeDefaultUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ops credentials make default unauthorized response has a 2xx status code
func (o *OpsCredentialsMakeDefaultUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials make default unauthorized response has a 3xx status code
func (o *OpsCredentialsMakeDefaultUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials make default unauthorized response has a 4xx status code
func (o *OpsCredentialsMakeDefaultUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ops credentials make default unauthorized response has a 5xx status code
func (o *OpsCredentialsMakeDefaultUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials make default unauthorized response a status code equal to that given
func (o *OpsCredentialsMakeDefaultUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OpsCredentialsMakeDefaultUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/makedefault][%d] opsCredentialsMakeDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *OpsCredentialsMakeDefaultUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/makedefault][%d] opsCredentialsMakeDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *OpsCredentialsMakeDefaultUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpsCredentialsMakeDefaultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsMakeDefaultForbidden creates a OpsCredentialsMakeDefaultForbidden with default headers values
func NewOpsCredentialsMakeDefaultForbidden() *OpsCredentialsMakeDefaultForbidden {
	return &OpsCredentialsMakeDefaultForbidden{}
}

/*
OpsCredentialsMakeDefaultForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpsCredentialsMakeDefaultForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ops credentials make default forbidden response has a 2xx status code
func (o *OpsCredentialsMakeDefaultForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials make default forbidden response has a 3xx status code
func (o *OpsCredentialsMakeDefaultForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials make default forbidden response has a 4xx status code
func (o *OpsCredentialsMakeDefaultForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ops credentials make default forbidden response has a 5xx status code
func (o *OpsCredentialsMakeDefaultForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials make default forbidden response a status code equal to that given
func (o *OpsCredentialsMakeDefaultForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OpsCredentialsMakeDefaultForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/makedefault][%d] opsCredentialsMakeDefaultForbidden  %+v", 403, o.Payload)
}

func (o *OpsCredentialsMakeDefaultForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/makedefault][%d] opsCredentialsMakeDefaultForbidden  %+v", 403, o.Payload)
}

func (o *OpsCredentialsMakeDefaultForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpsCredentialsMakeDefaultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsMakeDefaultNotFound creates a OpsCredentialsMakeDefaultNotFound with default headers values
func NewOpsCredentialsMakeDefaultNotFound() *OpsCredentialsMakeDefaultNotFound {
	return &OpsCredentialsMakeDefaultNotFound{}
}

/*
OpsCredentialsMakeDefaultNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpsCredentialsMakeDefaultNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ops credentials make default not found response has a 2xx status code
func (o *OpsCredentialsMakeDefaultNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials make default not found response has a 3xx status code
func (o *OpsCredentialsMakeDefaultNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials make default not found response has a 4xx status code
func (o *OpsCredentialsMakeDefaultNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ops credentials make default not found response has a 5xx status code
func (o *OpsCredentialsMakeDefaultNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials make default not found response a status code equal to that given
func (o *OpsCredentialsMakeDefaultNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OpsCredentialsMakeDefaultNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/makedefault][%d] opsCredentialsMakeDefaultNotFound  %+v", 404, o.Payload)
}

func (o *OpsCredentialsMakeDefaultNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/makedefault][%d] opsCredentialsMakeDefaultNotFound  %+v", 404, o.Payload)
}

func (o *OpsCredentialsMakeDefaultNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpsCredentialsMakeDefaultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsMakeDefaultInternalServerError creates a OpsCredentialsMakeDefaultInternalServerError with default headers values
func NewOpsCredentialsMakeDefaultInternalServerError() *OpsCredentialsMakeDefaultInternalServerError {
	return &OpsCredentialsMakeDefaultInternalServerError{}
}

/*
OpsCredentialsMakeDefaultInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpsCredentialsMakeDefaultInternalServerError struct {
}

// IsSuccess returns true when this ops credentials make default internal server error response has a 2xx status code
func (o *OpsCredentialsMakeDefaultInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials make default internal server error response has a 3xx status code
func (o *OpsCredentialsMakeDefaultInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials make default internal server error response has a 4xx status code
func (o *OpsCredentialsMakeDefaultInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ops credentials make default internal server error response has a 5xx status code
func (o *OpsCredentialsMakeDefaultInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ops credentials make default internal server error response a status code equal to that given
func (o *OpsCredentialsMakeDefaultInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OpsCredentialsMakeDefaultInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/makedefault][%d] opsCredentialsMakeDefaultInternalServerError ", 500)
}

func (o *OpsCredentialsMakeDefaultInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/makedefault][%d] opsCredentialsMakeDefaultInternalServerError ", 500)
}

func (o *OpsCredentialsMakeDefaultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
