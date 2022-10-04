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

// AlertingIntegrationsListReader is a Reader for the AlertingIntegrationsList structure.
type AlertingIntegrationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertingIntegrationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertingIntegrationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertingIntegrationsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAlertingIntegrationsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertingIntegrationsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertingIntegrationsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAlertingIntegrationsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertingIntegrationsListOK creates a AlertingIntegrationsListOK with default headers values
func NewAlertingIntegrationsListOK() *AlertingIntegrationsListOK {
	return &AlertingIntegrationsListOK{}
}

/*
AlertingIntegrationsListOK describes a response with status code 200, with default header values.

Success
*/
type AlertingIntegrationsListOK struct {
	Payload []*models.AlertingIntegrationsListDto
}

// IsSuccess returns true when this alerting integrations list o k response has a 2xx status code
func (o *AlertingIntegrationsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alerting integrations list o k response has a 3xx status code
func (o *AlertingIntegrationsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting integrations list o k response has a 4xx status code
func (o *AlertingIntegrationsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting integrations list o k response has a 5xx status code
func (o *AlertingIntegrationsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting integrations list o k response a status code equal to that given
func (o *AlertingIntegrationsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *AlertingIntegrationsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingIntegrations/{alertingProfileId}][%d] alertingIntegrationsListOK  %+v", 200, o.Payload)
}

func (o *AlertingIntegrationsListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingIntegrations/{alertingProfileId}][%d] alertingIntegrationsListOK  %+v", 200, o.Payload)
}

func (o *AlertingIntegrationsListOK) GetPayload() []*models.AlertingIntegrationsListDto {
	return o.Payload
}

func (o *AlertingIntegrationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsListBadRequest creates a AlertingIntegrationsListBadRequest with default headers values
func NewAlertingIntegrationsListBadRequest() *AlertingIntegrationsListBadRequest {
	return &AlertingIntegrationsListBadRequest{}
}

/*
AlertingIntegrationsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AlertingIntegrationsListBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this alerting integrations list bad request response has a 2xx status code
func (o *AlertingIntegrationsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting integrations list bad request response has a 3xx status code
func (o *AlertingIntegrationsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting integrations list bad request response has a 4xx status code
func (o *AlertingIntegrationsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting integrations list bad request response has a 5xx status code
func (o *AlertingIntegrationsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting integrations list bad request response a status code equal to that given
func (o *AlertingIntegrationsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AlertingIntegrationsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingIntegrations/{alertingProfileId}][%d] alertingIntegrationsListBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingIntegrationsListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingIntegrations/{alertingProfileId}][%d] alertingIntegrationsListBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingIntegrationsListBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AlertingIntegrationsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsListUnauthorized creates a AlertingIntegrationsListUnauthorized with default headers values
func NewAlertingIntegrationsListUnauthorized() *AlertingIntegrationsListUnauthorized {
	return &AlertingIntegrationsListUnauthorized{}
}

/*
AlertingIntegrationsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AlertingIntegrationsListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this alerting integrations list unauthorized response has a 2xx status code
func (o *AlertingIntegrationsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting integrations list unauthorized response has a 3xx status code
func (o *AlertingIntegrationsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting integrations list unauthorized response has a 4xx status code
func (o *AlertingIntegrationsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting integrations list unauthorized response has a 5xx status code
func (o *AlertingIntegrationsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting integrations list unauthorized response a status code equal to that given
func (o *AlertingIntegrationsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AlertingIntegrationsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingIntegrations/{alertingProfileId}][%d] alertingIntegrationsListUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingIntegrationsListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingIntegrations/{alertingProfileId}][%d] alertingIntegrationsListUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingIntegrationsListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AlertingIntegrationsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsListForbidden creates a AlertingIntegrationsListForbidden with default headers values
func NewAlertingIntegrationsListForbidden() *AlertingIntegrationsListForbidden {
	return &AlertingIntegrationsListForbidden{}
}

/*
AlertingIntegrationsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AlertingIntegrationsListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this alerting integrations list forbidden response has a 2xx status code
func (o *AlertingIntegrationsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting integrations list forbidden response has a 3xx status code
func (o *AlertingIntegrationsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting integrations list forbidden response has a 4xx status code
func (o *AlertingIntegrationsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting integrations list forbidden response has a 5xx status code
func (o *AlertingIntegrationsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting integrations list forbidden response a status code equal to that given
func (o *AlertingIntegrationsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AlertingIntegrationsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingIntegrations/{alertingProfileId}][%d] alertingIntegrationsListForbidden  %+v", 403, o.Payload)
}

func (o *AlertingIntegrationsListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingIntegrations/{alertingProfileId}][%d] alertingIntegrationsListForbidden  %+v", 403, o.Payload)
}

func (o *AlertingIntegrationsListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AlertingIntegrationsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsListNotFound creates a AlertingIntegrationsListNotFound with default headers values
func NewAlertingIntegrationsListNotFound() *AlertingIntegrationsListNotFound {
	return &AlertingIntegrationsListNotFound{}
}

/*
AlertingIntegrationsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AlertingIntegrationsListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this alerting integrations list not found response has a 2xx status code
func (o *AlertingIntegrationsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting integrations list not found response has a 3xx status code
func (o *AlertingIntegrationsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting integrations list not found response has a 4xx status code
func (o *AlertingIntegrationsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting integrations list not found response has a 5xx status code
func (o *AlertingIntegrationsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting integrations list not found response a status code equal to that given
func (o *AlertingIntegrationsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AlertingIntegrationsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingIntegrations/{alertingProfileId}][%d] alertingIntegrationsListNotFound  %+v", 404, o.Payload)
}

func (o *AlertingIntegrationsListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingIntegrations/{alertingProfileId}][%d] alertingIntegrationsListNotFound  %+v", 404, o.Payload)
}

func (o *AlertingIntegrationsListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AlertingIntegrationsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingIntegrationsListInternalServerError creates a AlertingIntegrationsListInternalServerError with default headers values
func NewAlertingIntegrationsListInternalServerError() *AlertingIntegrationsListInternalServerError {
	return &AlertingIntegrationsListInternalServerError{}
}

/*
AlertingIntegrationsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AlertingIntegrationsListInternalServerError struct {
}

// IsSuccess returns true when this alerting integrations list internal server error response has a 2xx status code
func (o *AlertingIntegrationsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting integrations list internal server error response has a 3xx status code
func (o *AlertingIntegrationsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting integrations list internal server error response has a 4xx status code
func (o *AlertingIntegrationsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting integrations list internal server error response has a 5xx status code
func (o *AlertingIntegrationsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this alerting integrations list internal server error response a status code equal to that given
func (o *AlertingIntegrationsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AlertingIntegrationsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingIntegrations/{alertingProfileId}][%d] alertingIntegrationsListInternalServerError ", 500)
}

func (o *AlertingIntegrationsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingIntegrations/{alertingProfileId}][%d] alertingIntegrationsListInternalServerError ", 500)
}

func (o *AlertingIntegrationsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
