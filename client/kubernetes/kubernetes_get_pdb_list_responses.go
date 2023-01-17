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

// KubernetesGetPdbListReader is a Reader for the KubernetesGetPdbList structure.
type KubernetesGetPdbListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetPdbListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetPdbListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetPdbListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetPdbListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetPdbListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetPdbListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetPdbListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetPdbListOK creates a KubernetesGetPdbListOK with default headers values
func NewKubernetesGetPdbListOK() *KubernetesGetPdbListOK {
	return &KubernetesGetPdbListOK{}
}

/*
KubernetesGetPdbListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetPdbListOK struct {
	Payload *models.PodDisruptions
}

// IsSuccess returns true when this kubernetes get pdb list o k response has a 2xx status code
func (o *KubernetesGetPdbListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get pdb list o k response has a 3xx status code
func (o *KubernetesGetPdbListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pdb list o k response has a 4xx status code
func (o *KubernetesGetPdbListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get pdb list o k response has a 5xx status code
func (o *KubernetesGetPdbListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pdb list o k response a status code equal to that given
func (o *KubernetesGetPdbListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kubernetes get pdb list o k response
func (o *KubernetesGetPdbListOK) Code() int {
	return 200
}

func (o *KubernetesGetPdbListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pdb][%d] kubernetesGetPdbListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetPdbListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pdb][%d] kubernetesGetPdbListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetPdbListOK) GetPayload() *models.PodDisruptions {
	return o.Payload
}

func (o *KubernetesGetPdbListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodDisruptions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPdbListBadRequest creates a KubernetesGetPdbListBadRequest with default headers values
func NewKubernetesGetPdbListBadRequest() *KubernetesGetPdbListBadRequest {
	return &KubernetesGetPdbListBadRequest{}
}

/*
KubernetesGetPdbListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetPdbListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get pdb list bad request response has a 2xx status code
func (o *KubernetesGetPdbListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pdb list bad request response has a 3xx status code
func (o *KubernetesGetPdbListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pdb list bad request response has a 4xx status code
func (o *KubernetesGetPdbListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get pdb list bad request response has a 5xx status code
func (o *KubernetesGetPdbListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pdb list bad request response a status code equal to that given
func (o *KubernetesGetPdbListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the kubernetes get pdb list bad request response
func (o *KubernetesGetPdbListBadRequest) Code() int {
	return 400
}

func (o *KubernetesGetPdbListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pdb][%d] kubernetesGetPdbListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetPdbListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pdb][%d] kubernetesGetPdbListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetPdbListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetPdbListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPdbListUnauthorized creates a KubernetesGetPdbListUnauthorized with default headers values
func NewKubernetesGetPdbListUnauthorized() *KubernetesGetPdbListUnauthorized {
	return &KubernetesGetPdbListUnauthorized{}
}

/*
KubernetesGetPdbListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetPdbListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get pdb list unauthorized response has a 2xx status code
func (o *KubernetesGetPdbListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pdb list unauthorized response has a 3xx status code
func (o *KubernetesGetPdbListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pdb list unauthorized response has a 4xx status code
func (o *KubernetesGetPdbListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get pdb list unauthorized response has a 5xx status code
func (o *KubernetesGetPdbListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pdb list unauthorized response a status code equal to that given
func (o *KubernetesGetPdbListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the kubernetes get pdb list unauthorized response
func (o *KubernetesGetPdbListUnauthorized) Code() int {
	return 401
}

func (o *KubernetesGetPdbListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pdb][%d] kubernetesGetPdbListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetPdbListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pdb][%d] kubernetesGetPdbListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetPdbListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetPdbListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPdbListForbidden creates a KubernetesGetPdbListForbidden with default headers values
func NewKubernetesGetPdbListForbidden() *KubernetesGetPdbListForbidden {
	return &KubernetesGetPdbListForbidden{}
}

/*
KubernetesGetPdbListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetPdbListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get pdb list forbidden response has a 2xx status code
func (o *KubernetesGetPdbListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pdb list forbidden response has a 3xx status code
func (o *KubernetesGetPdbListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pdb list forbidden response has a 4xx status code
func (o *KubernetesGetPdbListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get pdb list forbidden response has a 5xx status code
func (o *KubernetesGetPdbListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pdb list forbidden response a status code equal to that given
func (o *KubernetesGetPdbListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the kubernetes get pdb list forbidden response
func (o *KubernetesGetPdbListForbidden) Code() int {
	return 403
}

func (o *KubernetesGetPdbListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pdb][%d] kubernetesGetPdbListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetPdbListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pdb][%d] kubernetesGetPdbListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetPdbListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetPdbListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPdbListNotFound creates a KubernetesGetPdbListNotFound with default headers values
func NewKubernetesGetPdbListNotFound() *KubernetesGetPdbListNotFound {
	return &KubernetesGetPdbListNotFound{}
}

/*
KubernetesGetPdbListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetPdbListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get pdb list not found response has a 2xx status code
func (o *KubernetesGetPdbListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pdb list not found response has a 3xx status code
func (o *KubernetesGetPdbListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pdb list not found response has a 4xx status code
func (o *KubernetesGetPdbListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get pdb list not found response has a 5xx status code
func (o *KubernetesGetPdbListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pdb list not found response a status code equal to that given
func (o *KubernetesGetPdbListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the kubernetes get pdb list not found response
func (o *KubernetesGetPdbListNotFound) Code() int {
	return 404
}

func (o *KubernetesGetPdbListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pdb][%d] kubernetesGetPdbListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetPdbListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pdb][%d] kubernetesGetPdbListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetPdbListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetPdbListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPdbListInternalServerError creates a KubernetesGetPdbListInternalServerError with default headers values
func NewKubernetesGetPdbListInternalServerError() *KubernetesGetPdbListInternalServerError {
	return &KubernetesGetPdbListInternalServerError{}
}

/*
KubernetesGetPdbListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetPdbListInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get pdb list internal server error response has a 2xx status code
func (o *KubernetesGetPdbListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pdb list internal server error response has a 3xx status code
func (o *KubernetesGetPdbListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pdb list internal server error response has a 4xx status code
func (o *KubernetesGetPdbListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get pdb list internal server error response has a 5xx status code
func (o *KubernetesGetPdbListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get pdb list internal server error response a status code equal to that given
func (o *KubernetesGetPdbListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the kubernetes get pdb list internal server error response
func (o *KubernetesGetPdbListInternalServerError) Code() int {
	return 500
}

func (o *KubernetesGetPdbListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pdb][%d] kubernetesGetPdbListInternalServerError ", 500)
}

func (o *KubernetesGetPdbListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pdb][%d] kubernetesGetPdbListInternalServerError ", 500)
}

func (o *KubernetesGetPdbListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
