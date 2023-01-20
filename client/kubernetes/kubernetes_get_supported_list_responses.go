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

// KubernetesGetSupportedListReader is a Reader for the KubernetesGetSupportedList structure.
type KubernetesGetSupportedListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetSupportedListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetSupportedListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetSupportedListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetSupportedListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetSupportedListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetSupportedListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetSupportedListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetSupportedListOK creates a KubernetesGetSupportedListOK with default headers values
func NewKubernetesGetSupportedListOK() *KubernetesGetSupportedListOK {
	return &KubernetesGetSupportedListOK{}
}

/*
KubernetesGetSupportedListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetSupportedListOK struct {
	Payload []*models.KubernetesVersionListDto
}

// IsSuccess returns true when this kubernetes get supported list o k response has a 2xx status code
func (o *KubernetesGetSupportedListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get supported list o k response has a 3xx status code
func (o *KubernetesGetSupportedListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get supported list o k response has a 4xx status code
func (o *KubernetesGetSupportedListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get supported list o k response has a 5xx status code
func (o *KubernetesGetSupportedListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get supported list o k response a status code equal to that given
func (o *KubernetesGetSupportedListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kubernetes get supported list o k response
func (o *KubernetesGetSupportedListOK) Code() int {
	return 200
}

func (o *KubernetesGetSupportedListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/supported/list][%d] kubernetesGetSupportedListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetSupportedListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/supported/list][%d] kubernetesGetSupportedListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetSupportedListOK) GetPayload() []*models.KubernetesVersionListDto {
	return o.Payload
}

func (o *KubernetesGetSupportedListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetSupportedListBadRequest creates a KubernetesGetSupportedListBadRequest with default headers values
func NewKubernetesGetSupportedListBadRequest() *KubernetesGetSupportedListBadRequest {
	return &KubernetesGetSupportedListBadRequest{}
}

/*
KubernetesGetSupportedListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetSupportedListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes get supported list bad request response has a 2xx status code
func (o *KubernetesGetSupportedListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get supported list bad request response has a 3xx status code
func (o *KubernetesGetSupportedListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get supported list bad request response has a 4xx status code
func (o *KubernetesGetSupportedListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get supported list bad request response has a 5xx status code
func (o *KubernetesGetSupportedListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get supported list bad request response a status code equal to that given
func (o *KubernetesGetSupportedListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the kubernetes get supported list bad request response
func (o *KubernetesGetSupportedListBadRequest) Code() int {
	return 400
}

func (o *KubernetesGetSupportedListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/supported/list][%d] kubernetesGetSupportedListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetSupportedListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/supported/list][%d] kubernetesGetSupportedListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetSupportedListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesGetSupportedListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetSupportedListUnauthorized creates a KubernetesGetSupportedListUnauthorized with default headers values
func NewKubernetesGetSupportedListUnauthorized() *KubernetesGetSupportedListUnauthorized {
	return &KubernetesGetSupportedListUnauthorized{}
}

/*
KubernetesGetSupportedListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetSupportedListUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes get supported list unauthorized response has a 2xx status code
func (o *KubernetesGetSupportedListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get supported list unauthorized response has a 3xx status code
func (o *KubernetesGetSupportedListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get supported list unauthorized response has a 4xx status code
func (o *KubernetesGetSupportedListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get supported list unauthorized response has a 5xx status code
func (o *KubernetesGetSupportedListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get supported list unauthorized response a status code equal to that given
func (o *KubernetesGetSupportedListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the kubernetes get supported list unauthorized response
func (o *KubernetesGetSupportedListUnauthorized) Code() int {
	return 401
}

func (o *KubernetesGetSupportedListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/supported/list][%d] kubernetesGetSupportedListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetSupportedListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/supported/list][%d] kubernetesGetSupportedListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetSupportedListUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesGetSupportedListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetSupportedListForbidden creates a KubernetesGetSupportedListForbidden with default headers values
func NewKubernetesGetSupportedListForbidden() *KubernetesGetSupportedListForbidden {
	return &KubernetesGetSupportedListForbidden{}
}

/*
KubernetesGetSupportedListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetSupportedListForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes get supported list forbidden response has a 2xx status code
func (o *KubernetesGetSupportedListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get supported list forbidden response has a 3xx status code
func (o *KubernetesGetSupportedListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get supported list forbidden response has a 4xx status code
func (o *KubernetesGetSupportedListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get supported list forbidden response has a 5xx status code
func (o *KubernetesGetSupportedListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get supported list forbidden response a status code equal to that given
func (o *KubernetesGetSupportedListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the kubernetes get supported list forbidden response
func (o *KubernetesGetSupportedListForbidden) Code() int {
	return 403
}

func (o *KubernetesGetSupportedListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/supported/list][%d] kubernetesGetSupportedListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetSupportedListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/supported/list][%d] kubernetesGetSupportedListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetSupportedListForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesGetSupportedListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetSupportedListNotFound creates a KubernetesGetSupportedListNotFound with default headers values
func NewKubernetesGetSupportedListNotFound() *KubernetesGetSupportedListNotFound {
	return &KubernetesGetSupportedListNotFound{}
}

/*
KubernetesGetSupportedListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetSupportedListNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes get supported list not found response has a 2xx status code
func (o *KubernetesGetSupportedListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get supported list not found response has a 3xx status code
func (o *KubernetesGetSupportedListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get supported list not found response has a 4xx status code
func (o *KubernetesGetSupportedListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get supported list not found response has a 5xx status code
func (o *KubernetesGetSupportedListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get supported list not found response a status code equal to that given
func (o *KubernetesGetSupportedListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the kubernetes get supported list not found response
func (o *KubernetesGetSupportedListNotFound) Code() int {
	return 404
}

func (o *KubernetesGetSupportedListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/supported/list][%d] kubernetesGetSupportedListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetSupportedListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/supported/list][%d] kubernetesGetSupportedListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetSupportedListNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesGetSupportedListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetSupportedListInternalServerError creates a KubernetesGetSupportedListInternalServerError with default headers values
func NewKubernetesGetSupportedListInternalServerError() *KubernetesGetSupportedListInternalServerError {
	return &KubernetesGetSupportedListInternalServerError{}
}

/*
KubernetesGetSupportedListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetSupportedListInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get supported list internal server error response has a 2xx status code
func (o *KubernetesGetSupportedListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get supported list internal server error response has a 3xx status code
func (o *KubernetesGetSupportedListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get supported list internal server error response has a 4xx status code
func (o *KubernetesGetSupportedListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get supported list internal server error response has a 5xx status code
func (o *KubernetesGetSupportedListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get supported list internal server error response a status code equal to that given
func (o *KubernetesGetSupportedListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the kubernetes get supported list internal server error response
func (o *KubernetesGetSupportedListInternalServerError) Code() int {
	return 500
}

func (o *KubernetesGetSupportedListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/supported/list][%d] kubernetesGetSupportedListInternalServerError ", 500)
}

func (o *KubernetesGetSupportedListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/supported/list][%d] kubernetesGetSupportedListInternalServerError ", 500)
}

func (o *KubernetesGetSupportedListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
