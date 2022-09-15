// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// KubernetesDeleteEventsReader is a Reader for the KubernetesDeleteEvents structure.
type KubernetesDeleteEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDeleteEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDeleteEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDeleteEventsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDeleteEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDeleteEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDeleteEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDeleteEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDeleteEventsOK creates a KubernetesDeleteEventsOK with default headers values
func NewKubernetesDeleteEventsOK() *KubernetesDeleteEventsOK {
	return &KubernetesDeleteEventsOK{}
}

/*
KubernetesDeleteEventsOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDeleteEventsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kubernetes delete events o k response has a 2xx status code
func (o *KubernetesDeleteEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes delete events o k response has a 3xx status code
func (o *KubernetesDeleteEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes delete events o k response has a 4xx status code
func (o *KubernetesDeleteEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes delete events o k response has a 5xx status code
func (o *KubernetesDeleteEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes delete events o k response a status code equal to that given
func (o *KubernetesDeleteEventsOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesDeleteEventsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/removeevents][%d] kubernetesDeleteEventsOK  %+v", 200, o.Payload)
}

func (o *KubernetesDeleteEventsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/removeevents][%d] kubernetesDeleteEventsOK  %+v", 200, o.Payload)
}

func (o *KubernetesDeleteEventsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesDeleteEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDeleteEventsBadRequest creates a KubernetesDeleteEventsBadRequest with default headers values
func NewKubernetesDeleteEventsBadRequest() *KubernetesDeleteEventsBadRequest {
	return &KubernetesDeleteEventsBadRequest{}
}

/*
KubernetesDeleteEventsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDeleteEventsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kubernetes delete events bad request response has a 2xx status code
func (o *KubernetesDeleteEventsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes delete events bad request response has a 3xx status code
func (o *KubernetesDeleteEventsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes delete events bad request response has a 4xx status code
func (o *KubernetesDeleteEventsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes delete events bad request response has a 5xx status code
func (o *KubernetesDeleteEventsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes delete events bad request response a status code equal to that given
func (o *KubernetesDeleteEventsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesDeleteEventsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/removeevents][%d] kubernetesDeleteEventsBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDeleteEventsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/removeevents][%d] kubernetesDeleteEventsBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDeleteEventsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesDeleteEventsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDeleteEventsUnauthorized creates a KubernetesDeleteEventsUnauthorized with default headers values
func NewKubernetesDeleteEventsUnauthorized() *KubernetesDeleteEventsUnauthorized {
	return &KubernetesDeleteEventsUnauthorized{}
}

/*
KubernetesDeleteEventsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDeleteEventsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes delete events unauthorized response has a 2xx status code
func (o *KubernetesDeleteEventsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes delete events unauthorized response has a 3xx status code
func (o *KubernetesDeleteEventsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes delete events unauthorized response has a 4xx status code
func (o *KubernetesDeleteEventsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes delete events unauthorized response has a 5xx status code
func (o *KubernetesDeleteEventsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes delete events unauthorized response a status code equal to that given
func (o *KubernetesDeleteEventsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesDeleteEventsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/removeevents][%d] kubernetesDeleteEventsUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDeleteEventsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/removeevents][%d] kubernetesDeleteEventsUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDeleteEventsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDeleteEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDeleteEventsForbidden creates a KubernetesDeleteEventsForbidden with default headers values
func NewKubernetesDeleteEventsForbidden() *KubernetesDeleteEventsForbidden {
	return &KubernetesDeleteEventsForbidden{}
}

/*
KubernetesDeleteEventsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDeleteEventsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes delete events forbidden response has a 2xx status code
func (o *KubernetesDeleteEventsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes delete events forbidden response has a 3xx status code
func (o *KubernetesDeleteEventsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes delete events forbidden response has a 4xx status code
func (o *KubernetesDeleteEventsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes delete events forbidden response has a 5xx status code
func (o *KubernetesDeleteEventsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes delete events forbidden response a status code equal to that given
func (o *KubernetesDeleteEventsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesDeleteEventsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/removeevents][%d] kubernetesDeleteEventsForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDeleteEventsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/removeevents][%d] kubernetesDeleteEventsForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDeleteEventsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDeleteEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDeleteEventsNotFound creates a KubernetesDeleteEventsNotFound with default headers values
func NewKubernetesDeleteEventsNotFound() *KubernetesDeleteEventsNotFound {
	return &KubernetesDeleteEventsNotFound{}
}

/*
KubernetesDeleteEventsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDeleteEventsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes delete events not found response has a 2xx status code
func (o *KubernetesDeleteEventsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes delete events not found response has a 3xx status code
func (o *KubernetesDeleteEventsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes delete events not found response has a 4xx status code
func (o *KubernetesDeleteEventsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes delete events not found response has a 5xx status code
func (o *KubernetesDeleteEventsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes delete events not found response a status code equal to that given
func (o *KubernetesDeleteEventsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesDeleteEventsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/removeevents][%d] kubernetesDeleteEventsNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDeleteEventsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/removeevents][%d] kubernetesDeleteEventsNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDeleteEventsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDeleteEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDeleteEventsInternalServerError creates a KubernetesDeleteEventsInternalServerError with default headers values
func NewKubernetesDeleteEventsInternalServerError() *KubernetesDeleteEventsInternalServerError {
	return &KubernetesDeleteEventsInternalServerError{}
}

/*
KubernetesDeleteEventsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDeleteEventsInternalServerError struct {
}

// IsSuccess returns true when this kubernetes delete events internal server error response has a 2xx status code
func (o *KubernetesDeleteEventsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes delete events internal server error response has a 3xx status code
func (o *KubernetesDeleteEventsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes delete events internal server error response has a 4xx status code
func (o *KubernetesDeleteEventsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes delete events internal server error response has a 5xx status code
func (o *KubernetesDeleteEventsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes delete events internal server error response a status code equal to that given
func (o *KubernetesDeleteEventsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesDeleteEventsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/removeevents][%d] kubernetesDeleteEventsInternalServerError ", 500)
}

func (o *KubernetesDeleteEventsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/removeevents][%d] kubernetesDeleteEventsInternalServerError ", 500)
}

func (o *KubernetesDeleteEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
