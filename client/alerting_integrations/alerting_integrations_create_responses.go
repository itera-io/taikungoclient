// Code generated by go-swagger; DO NOT EDIT.

package alerting_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AlertingIntegrationsCreateReader is a Reader for the AlertingIntegrationsCreate structure.
type AlertingIntegrationsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertingIntegrationsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertingIntegrationsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertingIntegrationsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAlertingIntegrationsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertingIntegrationsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertingIntegrationsCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAlertingIntegrationsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertingIntegrationsCreateOK creates a AlertingIntegrationsCreateOK with default headers values
func NewAlertingIntegrationsCreateOK() *AlertingIntegrationsCreateOK {
	return &AlertingIntegrationsCreateOK{}
}

/* AlertingIntegrationsCreateOK describes a response with status code 200, with default header values.

Success
*/
type AlertingIntegrationsCreateOK struct {
	Payload models.Unit
}

func (o *AlertingIntegrationsCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingIntegrations/create][%d] alertingIntegrationsCreateOK  %+v", 200, o.Payload)
}
func (o *AlertingIntegrationsCreateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AlertingIntegrationsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsCreateBadRequest creates a AlertingIntegrationsCreateBadRequest with default headers values
func NewAlertingIntegrationsCreateBadRequest() *AlertingIntegrationsCreateBadRequest {
	return &AlertingIntegrationsCreateBadRequest{}
}

/* AlertingIntegrationsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AlertingIntegrationsCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AlertingIntegrationsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingIntegrations/create][%d] alertingIntegrationsCreateBadRequest  %+v", 400, o.Payload)
}
func (o *AlertingIntegrationsCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AlertingIntegrationsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsCreateUnauthorized creates a AlertingIntegrationsCreateUnauthorized with default headers values
func NewAlertingIntegrationsCreateUnauthorized() *AlertingIntegrationsCreateUnauthorized {
	return &AlertingIntegrationsCreateUnauthorized{}
}

/* AlertingIntegrationsCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AlertingIntegrationsCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AlertingIntegrationsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingIntegrations/create][%d] alertingIntegrationsCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *AlertingIntegrationsCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingIntegrationsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsCreateForbidden creates a AlertingIntegrationsCreateForbidden with default headers values
func NewAlertingIntegrationsCreateForbidden() *AlertingIntegrationsCreateForbidden {
	return &AlertingIntegrationsCreateForbidden{}
}

/* AlertingIntegrationsCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AlertingIntegrationsCreateForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AlertingIntegrationsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingIntegrations/create][%d] alertingIntegrationsCreateForbidden  %+v", 403, o.Payload)
}
func (o *AlertingIntegrationsCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingIntegrationsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsCreateNotFound creates a AlertingIntegrationsCreateNotFound with default headers values
func NewAlertingIntegrationsCreateNotFound() *AlertingIntegrationsCreateNotFound {
	return &AlertingIntegrationsCreateNotFound{}
}

/* AlertingIntegrationsCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AlertingIntegrationsCreateNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AlertingIntegrationsCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingIntegrations/create][%d] alertingIntegrationsCreateNotFound  %+v", 404, o.Payload)
}
func (o *AlertingIntegrationsCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingIntegrationsCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsCreateInternalServerError creates a AlertingIntegrationsCreateInternalServerError with default headers values
func NewAlertingIntegrationsCreateInternalServerError() *AlertingIntegrationsCreateInternalServerError {
	return &AlertingIntegrationsCreateInternalServerError{}
}

/* AlertingIntegrationsCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AlertingIntegrationsCreateInternalServerError struct {
}

func (o *AlertingIntegrationsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingIntegrations/create][%d] alertingIntegrationsCreateInternalServerError ", 500)
}

func (o *AlertingIntegrationsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
