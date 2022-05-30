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

/* CheckerAzureCPUQuotaOK describes a response with status code 200, with default header values.

Success
*/
type CheckerAzureCPUQuotaOK struct {
	Payload models.Unit
}

func (o *CheckerAzureCPUQuotaOK) Error() string {
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

/* CheckerAzureCPUQuotaBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerAzureCPUQuotaBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CheckerAzureCPUQuotaBadRequest) Error() string {
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

/* CheckerAzureCPUQuotaUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerAzureCPUQuotaUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CheckerAzureCPUQuotaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaUnauthorized  %+v", 401, o.Payload)
}
func (o *CheckerAzureCPUQuotaUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerAzureCPUQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAzureCPUQuotaForbidden creates a CheckerAzureCPUQuotaForbidden with default headers values
func NewCheckerAzureCPUQuotaForbidden() *CheckerAzureCPUQuotaForbidden {
	return &CheckerAzureCPUQuotaForbidden{}
}

/* CheckerAzureCPUQuotaForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerAzureCPUQuotaForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CheckerAzureCPUQuotaForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaForbidden  %+v", 403, o.Payload)
}
func (o *CheckerAzureCPUQuotaForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerAzureCPUQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAzureCPUQuotaNotFound creates a CheckerAzureCPUQuotaNotFound with default headers values
func NewCheckerAzureCPUQuotaNotFound() *CheckerAzureCPUQuotaNotFound {
	return &CheckerAzureCPUQuotaNotFound{}
}

/* CheckerAzureCPUQuotaNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerAzureCPUQuotaNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CheckerAzureCPUQuotaNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaNotFound  %+v", 404, o.Payload)
}
func (o *CheckerAzureCPUQuotaNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerAzureCPUQuotaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAzureCPUQuotaInternalServerError creates a CheckerAzureCPUQuotaInternalServerError with default headers values
func NewCheckerAzureCPUQuotaInternalServerError() *CheckerAzureCPUQuotaInternalServerError {
	return &CheckerAzureCPUQuotaInternalServerError{}
}

/* CheckerAzureCPUQuotaInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerAzureCPUQuotaInternalServerError struct {
}

func (o *CheckerAzureCPUQuotaInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure/quota/cpu][%d] checkerAzureCpuQuotaInternalServerError ", 500)
}

func (o *CheckerAzureCPUQuotaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}