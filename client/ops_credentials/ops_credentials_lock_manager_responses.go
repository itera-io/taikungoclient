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

// OpsCredentialsLockManagerReader is a Reader for the OpsCredentialsLockManager structure.
type OpsCredentialsLockManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpsCredentialsLockManagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpsCredentialsLockManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpsCredentialsLockManagerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpsCredentialsLockManagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpsCredentialsLockManagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpsCredentialsLockManagerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpsCredentialsLockManagerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpsCredentialsLockManagerOK creates a OpsCredentialsLockManagerOK with default headers values
func NewOpsCredentialsLockManagerOK() *OpsCredentialsLockManagerOK {
	return &OpsCredentialsLockManagerOK{}
}

/*
OpsCredentialsLockManagerOK describes a response with status code 200, with default header values.

Success
*/
type OpsCredentialsLockManagerOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this ops credentials lock manager o k response has a 2xx status code
func (o *OpsCredentialsLockManagerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ops credentials lock manager o k response has a 3xx status code
func (o *OpsCredentialsLockManagerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials lock manager o k response has a 4xx status code
func (o *OpsCredentialsLockManagerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ops credentials lock manager o k response has a 5xx status code
func (o *OpsCredentialsLockManagerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials lock manager o k response a status code equal to that given
func (o *OpsCredentialsLockManagerOK) IsCode(code int) bool {
	return code == 200
}

func (o *OpsCredentialsLockManagerOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/lockmanager][%d] opsCredentialsLockManagerOK  %+v", 200, o.Payload)
}

func (o *OpsCredentialsLockManagerOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/lockmanager][%d] opsCredentialsLockManagerOK  %+v", 200, o.Payload)
}

func (o *OpsCredentialsLockManagerOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *OpsCredentialsLockManagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsLockManagerBadRequest creates a OpsCredentialsLockManagerBadRequest with default headers values
func NewOpsCredentialsLockManagerBadRequest() *OpsCredentialsLockManagerBadRequest {
	return &OpsCredentialsLockManagerBadRequest{}
}

/*
OpsCredentialsLockManagerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpsCredentialsLockManagerBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this ops credentials lock manager bad request response has a 2xx status code
func (o *OpsCredentialsLockManagerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials lock manager bad request response has a 3xx status code
func (o *OpsCredentialsLockManagerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials lock manager bad request response has a 4xx status code
func (o *OpsCredentialsLockManagerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ops credentials lock manager bad request response has a 5xx status code
func (o *OpsCredentialsLockManagerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials lock manager bad request response a status code equal to that given
func (o *OpsCredentialsLockManagerBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OpsCredentialsLockManagerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/lockmanager][%d] opsCredentialsLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *OpsCredentialsLockManagerBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/lockmanager][%d] opsCredentialsLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *OpsCredentialsLockManagerBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *OpsCredentialsLockManagerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsLockManagerUnauthorized creates a OpsCredentialsLockManagerUnauthorized with default headers values
func NewOpsCredentialsLockManagerUnauthorized() *OpsCredentialsLockManagerUnauthorized {
	return &OpsCredentialsLockManagerUnauthorized{}
}

/*
OpsCredentialsLockManagerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpsCredentialsLockManagerUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ops credentials lock manager unauthorized response has a 2xx status code
func (o *OpsCredentialsLockManagerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials lock manager unauthorized response has a 3xx status code
func (o *OpsCredentialsLockManagerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials lock manager unauthorized response has a 4xx status code
func (o *OpsCredentialsLockManagerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ops credentials lock manager unauthorized response has a 5xx status code
func (o *OpsCredentialsLockManagerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials lock manager unauthorized response a status code equal to that given
func (o *OpsCredentialsLockManagerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OpsCredentialsLockManagerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/lockmanager][%d] opsCredentialsLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *OpsCredentialsLockManagerUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/lockmanager][%d] opsCredentialsLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *OpsCredentialsLockManagerUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpsCredentialsLockManagerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsLockManagerForbidden creates a OpsCredentialsLockManagerForbidden with default headers values
func NewOpsCredentialsLockManagerForbidden() *OpsCredentialsLockManagerForbidden {
	return &OpsCredentialsLockManagerForbidden{}
}

/*
OpsCredentialsLockManagerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpsCredentialsLockManagerForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ops credentials lock manager forbidden response has a 2xx status code
func (o *OpsCredentialsLockManagerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials lock manager forbidden response has a 3xx status code
func (o *OpsCredentialsLockManagerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials lock manager forbidden response has a 4xx status code
func (o *OpsCredentialsLockManagerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ops credentials lock manager forbidden response has a 5xx status code
func (o *OpsCredentialsLockManagerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials lock manager forbidden response a status code equal to that given
func (o *OpsCredentialsLockManagerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OpsCredentialsLockManagerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/lockmanager][%d] opsCredentialsLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *OpsCredentialsLockManagerForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/lockmanager][%d] opsCredentialsLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *OpsCredentialsLockManagerForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpsCredentialsLockManagerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsLockManagerNotFound creates a OpsCredentialsLockManagerNotFound with default headers values
func NewOpsCredentialsLockManagerNotFound() *OpsCredentialsLockManagerNotFound {
	return &OpsCredentialsLockManagerNotFound{}
}

/*
OpsCredentialsLockManagerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpsCredentialsLockManagerNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ops credentials lock manager not found response has a 2xx status code
func (o *OpsCredentialsLockManagerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials lock manager not found response has a 3xx status code
func (o *OpsCredentialsLockManagerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials lock manager not found response has a 4xx status code
func (o *OpsCredentialsLockManagerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ops credentials lock manager not found response has a 5xx status code
func (o *OpsCredentialsLockManagerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ops credentials lock manager not found response a status code equal to that given
func (o *OpsCredentialsLockManagerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OpsCredentialsLockManagerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/lockmanager][%d] opsCredentialsLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *OpsCredentialsLockManagerNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/lockmanager][%d] opsCredentialsLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *OpsCredentialsLockManagerNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpsCredentialsLockManagerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpsCredentialsLockManagerInternalServerError creates a OpsCredentialsLockManagerInternalServerError with default headers values
func NewOpsCredentialsLockManagerInternalServerError() *OpsCredentialsLockManagerInternalServerError {
	return &OpsCredentialsLockManagerInternalServerError{}
}

/*
OpsCredentialsLockManagerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpsCredentialsLockManagerInternalServerError struct {
}

// IsSuccess returns true when this ops credentials lock manager internal server error response has a 2xx status code
func (o *OpsCredentialsLockManagerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ops credentials lock manager internal server error response has a 3xx status code
func (o *OpsCredentialsLockManagerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ops credentials lock manager internal server error response has a 4xx status code
func (o *OpsCredentialsLockManagerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ops credentials lock manager internal server error response has a 5xx status code
func (o *OpsCredentialsLockManagerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ops credentials lock manager internal server error response a status code equal to that given
func (o *OpsCredentialsLockManagerInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OpsCredentialsLockManagerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/lockmanager][%d] opsCredentialsLockManagerInternalServerError ", 500)
}

func (o *OpsCredentialsLockManagerInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpsCredentials/lockmanager][%d] opsCredentialsLockManagerInternalServerError ", 500)
}

func (o *OpsCredentialsLockManagerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
