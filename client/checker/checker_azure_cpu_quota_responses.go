// Code generated by go-swagger; DO NOT EDIT.

package checker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// CheckerAzureCPUQuotaReader is a Reader for the CheckerAzureCPUQuota structure.
type CheckerAzureCPUQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerAzureCPUQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerAzureCPUQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerAzureCPUQuotaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerAzureCPUQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerAzureCPUQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerAzureCPUQuotaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerAzureCPUQuotaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerAzureCPUQuotaOK creates a CheckerAzureCPUQuotaOK with default headers values
func NewCheckerAzureCPUQuotaOK() *CheckerAzureCPUQuotaOK {
	return &CheckerAzureCPUQuotaOK{}
}

/*
CheckerAzureCPUQuotaOK describes a response with status code 200, with default header values.

Success
*/
type CheckerAzureCPUQuotaOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this checker azure Cpu quota o k response has a 2xx status code
func (o *CheckerAzureCPUQuotaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker azure Cpu quota o k response has a 3xx status code
func (o *CheckerAzureCPUQuotaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker azure Cpu quota o k response has a 4xx status code
func (o *CheckerAzureCPUQuotaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker azure Cpu quota o k response has a 5xx status code
func (o *CheckerAzureCPUQuotaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker azure Cpu quota o k response a status code equal to that given
func (o *CheckerAzureCPUQuotaOK) IsCode(code int) bool {
	return code == 200
}

func (o *CheckerAzureCPUQuotaOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaOK  %+v", 200, o.Payload)
}

func (o *CheckerAzureCPUQuotaOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaOK  %+v", 200, o.Payload)
}

func (o *CheckerAzureCPUQuotaOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CheckerAzureCPUQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAzureCPUQuotaBadRequest creates a CheckerAzureCPUQuotaBadRequest with default headers values
func NewCheckerAzureCPUQuotaBadRequest() *CheckerAzureCPUQuotaBadRequest {
	return &CheckerAzureCPUQuotaBadRequest{}
}

/*
CheckerAzureCPUQuotaBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerAzureCPUQuotaBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this checker azure Cpu quota bad request response has a 2xx status code
func (o *CheckerAzureCPUQuotaBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker azure Cpu quota bad request response has a 3xx status code
func (o *CheckerAzureCPUQuotaBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker azure Cpu quota bad request response has a 4xx status code
func (o *CheckerAzureCPUQuotaBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker azure Cpu quota bad request response has a 5xx status code
func (o *CheckerAzureCPUQuotaBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker azure Cpu quota bad request response a status code equal to that given
func (o *CheckerAzureCPUQuotaBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CheckerAzureCPUQuotaBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerAzureCPUQuotaBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerAzureCPUQuotaBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CheckerAzureCPUQuotaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAzureCPUQuotaUnauthorized creates a CheckerAzureCPUQuotaUnauthorized with default headers values
func NewCheckerAzureCPUQuotaUnauthorized() *CheckerAzureCPUQuotaUnauthorized {
	return &CheckerAzureCPUQuotaUnauthorized{}
}

/*
CheckerAzureCPUQuotaUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerAzureCPUQuotaUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this checker azure Cpu quota unauthorized response has a 2xx status code
func (o *CheckerAzureCPUQuotaUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker azure Cpu quota unauthorized response has a 3xx status code
func (o *CheckerAzureCPUQuotaUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker azure Cpu quota unauthorized response has a 4xx status code
func (o *CheckerAzureCPUQuotaUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker azure Cpu quota unauthorized response has a 5xx status code
func (o *CheckerAzureCPUQuotaUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker azure Cpu quota unauthorized response a status code equal to that given
func (o *CheckerAzureCPUQuotaUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CheckerAzureCPUQuotaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerAzureCPUQuotaUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerAzureCPUQuotaUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CheckerAzureCPUQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAzureCPUQuotaForbidden creates a CheckerAzureCPUQuotaForbidden with default headers values
func NewCheckerAzureCPUQuotaForbidden() *CheckerAzureCPUQuotaForbidden {
	return &CheckerAzureCPUQuotaForbidden{}
}

/*
CheckerAzureCPUQuotaForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerAzureCPUQuotaForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this checker azure Cpu quota forbidden response has a 2xx status code
func (o *CheckerAzureCPUQuotaForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker azure Cpu quota forbidden response has a 3xx status code
func (o *CheckerAzureCPUQuotaForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker azure Cpu quota forbidden response has a 4xx status code
func (o *CheckerAzureCPUQuotaForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker azure Cpu quota forbidden response has a 5xx status code
func (o *CheckerAzureCPUQuotaForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker azure Cpu quota forbidden response a status code equal to that given
func (o *CheckerAzureCPUQuotaForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CheckerAzureCPUQuotaForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaForbidden  %+v", 403, o.Payload)
}

func (o *CheckerAzureCPUQuotaForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaForbidden  %+v", 403, o.Payload)
}

func (o *CheckerAzureCPUQuotaForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CheckerAzureCPUQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAzureCPUQuotaNotFound creates a CheckerAzureCPUQuotaNotFound with default headers values
func NewCheckerAzureCPUQuotaNotFound() *CheckerAzureCPUQuotaNotFound {
	return &CheckerAzureCPUQuotaNotFound{}
}

/*
CheckerAzureCPUQuotaNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerAzureCPUQuotaNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this checker azure Cpu quota not found response has a 2xx status code
func (o *CheckerAzureCPUQuotaNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker azure Cpu quota not found response has a 3xx status code
func (o *CheckerAzureCPUQuotaNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker azure Cpu quota not found response has a 4xx status code
func (o *CheckerAzureCPUQuotaNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker azure Cpu quota not found response has a 5xx status code
func (o *CheckerAzureCPUQuotaNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker azure Cpu quota not found response a status code equal to that given
func (o *CheckerAzureCPUQuotaNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CheckerAzureCPUQuotaNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaNotFound  %+v", 404, o.Payload)
}

func (o *CheckerAzureCPUQuotaNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaNotFound  %+v", 404, o.Payload)
}

func (o *CheckerAzureCPUQuotaNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CheckerAzureCPUQuotaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAzureCPUQuotaInternalServerError creates a CheckerAzureCPUQuotaInternalServerError with default headers values
func NewCheckerAzureCPUQuotaInternalServerError() *CheckerAzureCPUQuotaInternalServerError {
	return &CheckerAzureCPUQuotaInternalServerError{}
}

/*
CheckerAzureCPUQuotaInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerAzureCPUQuotaInternalServerError struct {
}

// IsSuccess returns true when this checker azure Cpu quota internal server error response has a 2xx status code
func (o *CheckerAzureCPUQuotaInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker azure Cpu quota internal server error response has a 3xx status code
func (o *CheckerAzureCPUQuotaInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker azure Cpu quota internal server error response has a 4xx status code
func (o *CheckerAzureCPUQuotaInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker azure Cpu quota internal server error response has a 5xx status code
func (o *CheckerAzureCPUQuotaInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker azure Cpu quota internal server error response a status code equal to that given
func (o *CheckerAzureCPUQuotaInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CheckerAzureCPUQuotaInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaInternalServerError ", 500)
}

func (o *CheckerAzureCPUQuotaInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaInternalServerError ", 500)
}

func (o *CheckerAzureCPUQuotaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
