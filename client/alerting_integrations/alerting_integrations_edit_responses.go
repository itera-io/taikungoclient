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

// AlertingIntegrationsEditReader is a Reader for the AlertingIntegrationsEdit structure.
type AlertingIntegrationsEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertingIntegrationsEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertingIntegrationsEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertingIntegrationsEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAlertingIntegrationsEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertingIntegrationsEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertingIntegrationsEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAlertingIntegrationsEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertingIntegrationsEditOK creates a AlertingIntegrationsEditOK with default headers values
func NewAlertingIntegrationsEditOK() *AlertingIntegrationsEditOK {
	return &AlertingIntegrationsEditOK{}
}

/* AlertingIntegrationsEditOK describes a response with status code 200, with default header values.

Success
*/
type AlertingIntegrationsEditOK struct {
	Payload models.Unit
}

func (o *AlertingIntegrationsEditOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditOK  %+v", 200, o.Payload)
}
func (o *AlertingIntegrationsEditOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AlertingIntegrationsEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsEditBadRequest creates a AlertingIntegrationsEditBadRequest with default headers values
func NewAlertingIntegrationsEditBadRequest() *AlertingIntegrationsEditBadRequest {
	return &AlertingIntegrationsEditBadRequest{}
}

/* AlertingIntegrationsEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AlertingIntegrationsEditBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AlertingIntegrationsEditBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditBadRequest  %+v", 400, o.Payload)
}
func (o *AlertingIntegrationsEditBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AlertingIntegrationsEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsEditUnauthorized creates a AlertingIntegrationsEditUnauthorized with default headers values
func NewAlertingIntegrationsEditUnauthorized() *AlertingIntegrationsEditUnauthorized {
	return &AlertingIntegrationsEditUnauthorized{}
}

/* AlertingIntegrationsEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AlertingIntegrationsEditUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AlertingIntegrationsEditUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditUnauthorized  %+v", 401, o.Payload)
}
func (o *AlertingIntegrationsEditUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingIntegrationsEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsEditForbidden creates a AlertingIntegrationsEditForbidden with default headers values
func NewAlertingIntegrationsEditForbidden() *AlertingIntegrationsEditForbidden {
	return &AlertingIntegrationsEditForbidden{}
}

/* AlertingIntegrationsEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AlertingIntegrationsEditForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AlertingIntegrationsEditForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditForbidden  %+v", 403, o.Payload)
}
func (o *AlertingIntegrationsEditForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingIntegrationsEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsEditNotFound creates a AlertingIntegrationsEditNotFound with default headers values
func NewAlertingIntegrationsEditNotFound() *AlertingIntegrationsEditNotFound {
	return &AlertingIntegrationsEditNotFound{}
}

/* AlertingIntegrationsEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AlertingIntegrationsEditNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AlertingIntegrationsEditNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditNotFound  %+v", 404, o.Payload)
}
func (o *AlertingIntegrationsEditNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingIntegrationsEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsEditInternalServerError creates a AlertingIntegrationsEditInternalServerError with default headers values
func NewAlertingIntegrationsEditInternalServerError() *AlertingIntegrationsEditInternalServerError {
	return &AlertingIntegrationsEditInternalServerError{}
}

/* AlertingIntegrationsEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AlertingIntegrationsEditInternalServerError struct {
}

func (o *AlertingIntegrationsEditInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditInternalServerError ", 500)
}

func (o *AlertingIntegrationsEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
