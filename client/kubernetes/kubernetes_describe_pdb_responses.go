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

// KubernetesDescribePdbReader is a Reader for the KubernetesDescribePdb structure.
type KubernetesDescribePdbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribePdbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribePdbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribePdbBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribePdbUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribePdbForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribePdbNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribePdbInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribePdbOK creates a KubernetesDescribePdbOK with default headers values
func NewKubernetesDescribePdbOK() *KubernetesDescribePdbOK {
	return &KubernetesDescribePdbOK{}
}

/*
KubernetesDescribePdbOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribePdbOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe pdb o k response has a 2xx status code
func (o *KubernetesDescribePdbOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe pdb o k response has a 3xx status code
func (o *KubernetesDescribePdbOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pdb o k response has a 4xx status code
func (o *KubernetesDescribePdbOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe pdb o k response has a 5xx status code
func (o *KubernetesDescribePdbOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pdb o k response a status code equal to that given
func (o *KubernetesDescribePdbOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesDescribePdbOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pdb][%d] kubernetesDescribePdbOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribePdbOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pdb][%d] kubernetesDescribePdbOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribePdbOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribePdbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePdbBadRequest creates a KubernetesDescribePdbBadRequest with default headers values
func NewKubernetesDescribePdbBadRequest() *KubernetesDescribePdbBadRequest {
	return &KubernetesDescribePdbBadRequest{}
}

/*
KubernetesDescribePdbBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribePdbBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this kubernetes describe pdb bad request response has a 2xx status code
func (o *KubernetesDescribePdbBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pdb bad request response has a 3xx status code
func (o *KubernetesDescribePdbBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pdb bad request response has a 4xx status code
func (o *KubernetesDescribePdbBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe pdb bad request response has a 5xx status code
func (o *KubernetesDescribePdbBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pdb bad request response a status code equal to that given
func (o *KubernetesDescribePdbBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesDescribePdbBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pdb][%d] kubernetesDescribePdbBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribePdbBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pdb][%d] kubernetesDescribePdbBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribePdbBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *KubernetesDescribePdbBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePdbUnauthorized creates a KubernetesDescribePdbUnauthorized with default headers values
func NewKubernetesDescribePdbUnauthorized() *KubernetesDescribePdbUnauthorized {
	return &KubernetesDescribePdbUnauthorized{}
}

/*
KubernetesDescribePdbUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribePdbUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes describe pdb unauthorized response has a 2xx status code
func (o *KubernetesDescribePdbUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pdb unauthorized response has a 3xx status code
func (o *KubernetesDescribePdbUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pdb unauthorized response has a 4xx status code
func (o *KubernetesDescribePdbUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe pdb unauthorized response has a 5xx status code
func (o *KubernetesDescribePdbUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pdb unauthorized response a status code equal to that given
func (o *KubernetesDescribePdbUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesDescribePdbUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pdb][%d] kubernetesDescribePdbUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribePdbUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pdb][%d] kubernetesDescribePdbUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribePdbUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesDescribePdbUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePdbForbidden creates a KubernetesDescribePdbForbidden with default headers values
func NewKubernetesDescribePdbForbidden() *KubernetesDescribePdbForbidden {
	return &KubernetesDescribePdbForbidden{}
}

/*
KubernetesDescribePdbForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribePdbForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes describe pdb forbidden response has a 2xx status code
func (o *KubernetesDescribePdbForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pdb forbidden response has a 3xx status code
func (o *KubernetesDescribePdbForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pdb forbidden response has a 4xx status code
func (o *KubernetesDescribePdbForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe pdb forbidden response has a 5xx status code
func (o *KubernetesDescribePdbForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pdb forbidden response a status code equal to that given
func (o *KubernetesDescribePdbForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesDescribePdbForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pdb][%d] kubernetesDescribePdbForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribePdbForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pdb][%d] kubernetesDescribePdbForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribePdbForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesDescribePdbForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePdbNotFound creates a KubernetesDescribePdbNotFound with default headers values
func NewKubernetesDescribePdbNotFound() *KubernetesDescribePdbNotFound {
	return &KubernetesDescribePdbNotFound{}
}

/*
KubernetesDescribePdbNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribePdbNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes describe pdb not found response has a 2xx status code
func (o *KubernetesDescribePdbNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pdb not found response has a 3xx status code
func (o *KubernetesDescribePdbNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pdb not found response has a 4xx status code
func (o *KubernetesDescribePdbNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe pdb not found response has a 5xx status code
func (o *KubernetesDescribePdbNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pdb not found response a status code equal to that given
func (o *KubernetesDescribePdbNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesDescribePdbNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pdb][%d] kubernetesDescribePdbNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribePdbNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pdb][%d] kubernetesDescribePdbNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribePdbNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesDescribePdbNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePdbInternalServerError creates a KubernetesDescribePdbInternalServerError with default headers values
func NewKubernetesDescribePdbInternalServerError() *KubernetesDescribePdbInternalServerError {
	return &KubernetesDescribePdbInternalServerError{}
}

/*
KubernetesDescribePdbInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribePdbInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe pdb internal server error response has a 2xx status code
func (o *KubernetesDescribePdbInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pdb internal server error response has a 3xx status code
func (o *KubernetesDescribePdbInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pdb internal server error response has a 4xx status code
func (o *KubernetesDescribePdbInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe pdb internal server error response has a 5xx status code
func (o *KubernetesDescribePdbInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe pdb internal server error response a status code equal to that given
func (o *KubernetesDescribePdbInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesDescribePdbInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pdb][%d] kubernetesDescribePdbInternalServerError ", 500)
}

func (o *KubernetesDescribePdbInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pdb][%d] kubernetesDescribePdbInternalServerError ", 500)
}

func (o *KubernetesDescribePdbInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
