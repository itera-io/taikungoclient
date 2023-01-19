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

/*
AlertingIntegrationsEditOK describes a response with status code 200, with default header values.

Success
*/
type AlertingIntegrationsEditOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this alerting integrations edit o k response has a 2xx status code
func (o *AlertingIntegrationsEditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alerting integrations edit o k response has a 3xx status code
func (o *AlertingIntegrationsEditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting integrations edit o k response has a 4xx status code
func (o *AlertingIntegrationsEditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting integrations edit o k response has a 5xx status code
func (o *AlertingIntegrationsEditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting integrations edit o k response a status code equal to that given
func (o *AlertingIntegrationsEditOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the alerting integrations edit o k response
func (o *AlertingIntegrationsEditOK) Code() int {
	return 200
}

func (o *AlertingIntegrationsEditOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditOK  %+v", 200, o.Payload)
}

func (o *AlertingIntegrationsEditOK) String() string {
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

/*
AlertingIntegrationsEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AlertingIntegrationsEditBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting integrations edit bad request response has a 2xx status code
func (o *AlertingIntegrationsEditBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting integrations edit bad request response has a 3xx status code
func (o *AlertingIntegrationsEditBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting integrations edit bad request response has a 4xx status code
func (o *AlertingIntegrationsEditBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting integrations edit bad request response has a 5xx status code
func (o *AlertingIntegrationsEditBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting integrations edit bad request response a status code equal to that given
func (o *AlertingIntegrationsEditBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the alerting integrations edit bad request response
func (o *AlertingIntegrationsEditBadRequest) Code() int {
	return 400
}

func (o *AlertingIntegrationsEditBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingIntegrationsEditBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingIntegrationsEditBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingIntegrationsEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

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

/*
AlertingIntegrationsEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AlertingIntegrationsEditUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting integrations edit unauthorized response has a 2xx status code
func (o *AlertingIntegrationsEditUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting integrations edit unauthorized response has a 3xx status code
func (o *AlertingIntegrationsEditUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting integrations edit unauthorized response has a 4xx status code
func (o *AlertingIntegrationsEditUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting integrations edit unauthorized response has a 5xx status code
func (o *AlertingIntegrationsEditUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting integrations edit unauthorized response a status code equal to that given
func (o *AlertingIntegrationsEditUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the alerting integrations edit unauthorized response
func (o *AlertingIntegrationsEditUnauthorized) Code() int {
	return 401
}

func (o *AlertingIntegrationsEditUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingIntegrationsEditUnauthorized) String() string {
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

/*
AlertingIntegrationsEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AlertingIntegrationsEditForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting integrations edit forbidden response has a 2xx status code
func (o *AlertingIntegrationsEditForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting integrations edit forbidden response has a 3xx status code
func (o *AlertingIntegrationsEditForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting integrations edit forbidden response has a 4xx status code
func (o *AlertingIntegrationsEditForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting integrations edit forbidden response has a 5xx status code
func (o *AlertingIntegrationsEditForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting integrations edit forbidden response a status code equal to that given
func (o *AlertingIntegrationsEditForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the alerting integrations edit forbidden response
func (o *AlertingIntegrationsEditForbidden) Code() int {
	return 403
}

func (o *AlertingIntegrationsEditForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditForbidden  %+v", 403, o.Payload)
}

func (o *AlertingIntegrationsEditForbidden) String() string {
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

/*
AlertingIntegrationsEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AlertingIntegrationsEditNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting integrations edit not found response has a 2xx status code
func (o *AlertingIntegrationsEditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting integrations edit not found response has a 3xx status code
func (o *AlertingIntegrationsEditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting integrations edit not found response has a 4xx status code
func (o *AlertingIntegrationsEditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting integrations edit not found response has a 5xx status code
func (o *AlertingIntegrationsEditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting integrations edit not found response a status code equal to that given
func (o *AlertingIntegrationsEditNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the alerting integrations edit not found response
func (o *AlertingIntegrationsEditNotFound) Code() int {
	return 404
}

func (o *AlertingIntegrationsEditNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditNotFound  %+v", 404, o.Payload)
}

func (o *AlertingIntegrationsEditNotFound) String() string {
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

/*
AlertingIntegrationsEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AlertingIntegrationsEditInternalServerError struct {
}

// IsSuccess returns true when this alerting integrations edit internal server error response has a 2xx status code
func (o *AlertingIntegrationsEditInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting integrations edit internal server error response has a 3xx status code
func (o *AlertingIntegrationsEditInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting integrations edit internal server error response has a 4xx status code
func (o *AlertingIntegrationsEditInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting integrations edit internal server error response has a 5xx status code
func (o *AlertingIntegrationsEditInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this alerting integrations edit internal server error response a status code equal to that given
func (o *AlertingIntegrationsEditInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the alerting integrations edit internal server error response
func (o *AlertingIntegrationsEditInternalServerError) Code() int {
	return 500
}

func (o *AlertingIntegrationsEditInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditInternalServerError ", 500)
}

func (o *AlertingIntegrationsEditInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingIntegrations/edit][%d] alertingIntegrationsEditInternalServerError ", 500)
}

func (o *AlertingIntegrationsEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
