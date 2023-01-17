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

// OpsCredentialsListReader is a Reader for the OpsCredentialsList structure.
type OpsCredentialsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpsCredentialsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpsCredentialsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpsCredentialsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpsCredentialsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpsCredentialsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpsCredentialsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpsCredentialsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpsCredentialsListOK creates a OpsCredentialsListOK with default headers values
func NewOpsCredentialsListOK() *OpsCredentialsListOK {
	return &OpsCredentialsListOK{}
}

/*
OpsCredentialsListOK describes a response with status code 200, with default header values.

Success
*/
type OpsCredentialsListOK struct {
	Payload *models.OperationCredentials
}

// IsSuccess returns true when this ops credentials list o k response has a 2xx status code
func (o *OpsCredentialsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ops credentials list o k response has a 3xx status code
func (o *OpsCredentialsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials list o k response has a 4xx status code
func (o *OpsCredentialsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ops credentials list o k response has a 5xx status code
func (o *OpsCredentialsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials list o k response a status code equal to that given
func (o *OpsCredentialsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ops credentials list o k response
func (o *OpsCredentialsListOK) Code() int {
	return 200
}

func (o *OpsCredentialsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OpsCredentials/list][%d] opsCredentialsListOK  %+v", 200, o.Payload)
}

func (o *OpsCredentialsListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/OpsCredentials/list][%d] opsCredentialsListOK  %+v", 200, o.Payload)
}

func (o *OpsCredentialsListOK) GetPayload() *models.OperationCredentials {
	return o.Payload
}

func (o *OpsCredentialsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OperationCredentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsListBadRequest creates a OpsCredentialsListBadRequest with default headers values
func NewOpsCredentialsListBadRequest() *OpsCredentialsListBadRequest {
	return &OpsCredentialsListBadRequest{}
}

/*
OpsCredentialsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpsCredentialsListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ops credentials list bad request response has a 2xx status code
func (o *OpsCredentialsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials list bad request response has a 3xx status code
func (o *OpsCredentialsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials list bad request response has a 4xx status code
func (o *OpsCredentialsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ops credentials list bad request response has a 5xx status code
func (o *OpsCredentialsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials list bad request response a status code equal to that given
func (o *OpsCredentialsListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the ops credentials list bad request response
func (o *OpsCredentialsListBadRequest) Code() int {
	return 400
}

func (o *OpsCredentialsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OpsCredentials/list][%d] opsCredentialsListBadRequest  %+v", 400, o.Payload)
}

func (o *OpsCredentialsListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/OpsCredentials/list][%d] opsCredentialsListBadRequest  %+v", 400, o.Payload)
}

func (o *OpsCredentialsListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpsCredentialsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsListUnauthorized creates a OpsCredentialsListUnauthorized with default headers values
func NewOpsCredentialsListUnauthorized() *OpsCredentialsListUnauthorized {
	return &OpsCredentialsListUnauthorized{}
}

/*
OpsCredentialsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpsCredentialsListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ops credentials list unauthorized response has a 2xx status code
func (o *OpsCredentialsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials list unauthorized response has a 3xx status code
func (o *OpsCredentialsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials list unauthorized response has a 4xx status code
func (o *OpsCredentialsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ops credentials list unauthorized response has a 5xx status code
func (o *OpsCredentialsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials list unauthorized response a status code equal to that given
func (o *OpsCredentialsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the ops credentials list unauthorized response
func (o *OpsCredentialsListUnauthorized) Code() int {
	return 401
}

func (o *OpsCredentialsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OpsCredentials/list][%d] opsCredentialsListUnauthorized  %+v", 401, o.Payload)
}

func (o *OpsCredentialsListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/OpsCredentials/list][%d] opsCredentialsListUnauthorized  %+v", 401, o.Payload)
}

func (o *OpsCredentialsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpsCredentialsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsListForbidden creates a OpsCredentialsListForbidden with default headers values
func NewOpsCredentialsListForbidden() *OpsCredentialsListForbidden {
	return &OpsCredentialsListForbidden{}
}

/*
OpsCredentialsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpsCredentialsListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ops credentials list forbidden response has a 2xx status code
func (o *OpsCredentialsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials list forbidden response has a 3xx status code
func (o *OpsCredentialsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials list forbidden response has a 4xx status code
func (o *OpsCredentialsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ops credentials list forbidden response has a 5xx status code
func (o *OpsCredentialsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials list forbidden response a status code equal to that given
func (o *OpsCredentialsListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the ops credentials list forbidden response
func (o *OpsCredentialsListForbidden) Code() int {
	return 403
}

func (o *OpsCredentialsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OpsCredentials/list][%d] opsCredentialsListForbidden  %+v", 403, o.Payload)
}

func (o *OpsCredentialsListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/OpsCredentials/list][%d] opsCredentialsListForbidden  %+v", 403, o.Payload)
}

func (o *OpsCredentialsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpsCredentialsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsListNotFound creates a OpsCredentialsListNotFound with default headers values
func NewOpsCredentialsListNotFound() *OpsCredentialsListNotFound {
	return &OpsCredentialsListNotFound{}
}

/*
OpsCredentialsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpsCredentialsListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ops credentials list not found response has a 2xx status code
func (o *OpsCredentialsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials list not found response has a 3xx status code
func (o *OpsCredentialsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials list not found response has a 4xx status code
func (o *OpsCredentialsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ops credentials list not found response has a 5xx status code
func (o *OpsCredentialsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials list not found response a status code equal to that given
func (o *OpsCredentialsListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the ops credentials list not found response
func (o *OpsCredentialsListNotFound) Code() int {
	return 404
}

func (o *OpsCredentialsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OpsCredentials/list][%d] opsCredentialsListNotFound  %+v", 404, o.Payload)
}

func (o *OpsCredentialsListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/OpsCredentials/list][%d] opsCredentialsListNotFound  %+v", 404, o.Payload)
}

func (o *OpsCredentialsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpsCredentialsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsListInternalServerError creates a OpsCredentialsListInternalServerError with default headers values
func NewOpsCredentialsListInternalServerError() *OpsCredentialsListInternalServerError {
	return &OpsCredentialsListInternalServerError{}
}

/*
OpsCredentialsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpsCredentialsListInternalServerError struct {
}

// IsSuccess returns true when this ops credentials list internal server error response has a 2xx status code
func (o *OpsCredentialsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials list internal server error response has a 3xx status code
func (o *OpsCredentialsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials list internal server error response has a 4xx status code
func (o *OpsCredentialsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ops credentials list internal server error response has a 5xx status code
func (o *OpsCredentialsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ops credentials list internal server error response a status code equal to that given
func (o *OpsCredentialsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the ops credentials list internal server error response
func (o *OpsCredentialsListInternalServerError) Code() int {
	return 500
}

func (o *OpsCredentialsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OpsCredentials/list][%d] opsCredentialsListInternalServerError ", 500)
}

func (o *OpsCredentialsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/OpsCredentials/list][%d] opsCredentialsListInternalServerError ", 500)
}

func (o *OpsCredentialsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
