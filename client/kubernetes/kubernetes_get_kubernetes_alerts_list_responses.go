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

// KubernetesGetKubernetesAlertsListReader is a Reader for the KubernetesGetKubernetesAlertsList structure.
type KubernetesGetKubernetesAlertsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetKubernetesAlertsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetKubernetesAlertsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetKubernetesAlertsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetKubernetesAlertsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetKubernetesAlertsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetKubernetesAlertsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetKubernetesAlertsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetKubernetesAlertsListOK creates a KubernetesGetKubernetesAlertsListOK with default headers values
func NewKubernetesGetKubernetesAlertsListOK() *KubernetesGetKubernetesAlertsListOK {
	return &KubernetesGetKubernetesAlertsListOK{}
}

/*
KubernetesGetKubernetesAlertsListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetKubernetesAlertsListOK struct {
	Payload *models.KubernetesAlertList
}

// IsSuccess returns true when this kubernetes get kubernetes alerts list o k response has a 2xx status code
func (o *KubernetesGetKubernetesAlertsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get kubernetes alerts list o k response has a 3xx status code
func (o *KubernetesGetKubernetesAlertsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get kubernetes alerts list o k response has a 4xx status code
func (o *KubernetesGetKubernetesAlertsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get kubernetes alerts list o k response has a 5xx status code
func (o *KubernetesGetKubernetesAlertsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get kubernetes alerts list o k response a status code equal to that given
func (o *KubernetesGetKubernetesAlertsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetKubernetesAlertsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/alerts][%d] kubernetesGetKubernetesAlertsListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetKubernetesAlertsListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/alerts][%d] kubernetesGetKubernetesAlertsListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetKubernetesAlertsListOK) GetPayload() *models.KubernetesAlertList {
	return o.Payload
}

func (o *KubernetesGetKubernetesAlertsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesAlertList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubernetesAlertsListBadRequest creates a KubernetesGetKubernetesAlertsListBadRequest with default headers values
func NewKubernetesGetKubernetesAlertsListBadRequest() *KubernetesGetKubernetesAlertsListBadRequest {
	return &KubernetesGetKubernetesAlertsListBadRequest{}
}

/*
KubernetesGetKubernetesAlertsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetKubernetesAlertsListBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get kubernetes alerts list bad request response has a 2xx status code
func (o *KubernetesGetKubernetesAlertsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get kubernetes alerts list bad request response has a 3xx status code
func (o *KubernetesGetKubernetesAlertsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get kubernetes alerts list bad request response has a 4xx status code
func (o *KubernetesGetKubernetesAlertsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get kubernetes alerts list bad request response has a 5xx status code
func (o *KubernetesGetKubernetesAlertsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get kubernetes alerts list bad request response a status code equal to that given
func (o *KubernetesGetKubernetesAlertsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetKubernetesAlertsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/alerts][%d] kubernetesGetKubernetesAlertsListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetKubernetesAlertsListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/alerts][%d] kubernetesGetKubernetesAlertsListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetKubernetesAlertsListBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetKubernetesAlertsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubernetesAlertsListUnauthorized creates a KubernetesGetKubernetesAlertsListUnauthorized with default headers values
func NewKubernetesGetKubernetesAlertsListUnauthorized() *KubernetesGetKubernetesAlertsListUnauthorized {
	return &KubernetesGetKubernetesAlertsListUnauthorized{}
}

/*
KubernetesGetKubernetesAlertsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetKubernetesAlertsListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get kubernetes alerts list unauthorized response has a 2xx status code
func (o *KubernetesGetKubernetesAlertsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get kubernetes alerts list unauthorized response has a 3xx status code
func (o *KubernetesGetKubernetesAlertsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get kubernetes alerts list unauthorized response has a 4xx status code
func (o *KubernetesGetKubernetesAlertsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get kubernetes alerts list unauthorized response has a 5xx status code
func (o *KubernetesGetKubernetesAlertsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get kubernetes alerts list unauthorized response a status code equal to that given
func (o *KubernetesGetKubernetesAlertsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetKubernetesAlertsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/alerts][%d] kubernetesGetKubernetesAlertsListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetKubernetesAlertsListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/alerts][%d] kubernetesGetKubernetesAlertsListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetKubernetesAlertsListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetKubernetesAlertsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubernetesAlertsListForbidden creates a KubernetesGetKubernetesAlertsListForbidden with default headers values
func NewKubernetesGetKubernetesAlertsListForbidden() *KubernetesGetKubernetesAlertsListForbidden {
	return &KubernetesGetKubernetesAlertsListForbidden{}
}

/*
KubernetesGetKubernetesAlertsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetKubernetesAlertsListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get kubernetes alerts list forbidden response has a 2xx status code
func (o *KubernetesGetKubernetesAlertsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get kubernetes alerts list forbidden response has a 3xx status code
func (o *KubernetesGetKubernetesAlertsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get kubernetes alerts list forbidden response has a 4xx status code
func (o *KubernetesGetKubernetesAlertsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get kubernetes alerts list forbidden response has a 5xx status code
func (o *KubernetesGetKubernetesAlertsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get kubernetes alerts list forbidden response a status code equal to that given
func (o *KubernetesGetKubernetesAlertsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetKubernetesAlertsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/alerts][%d] kubernetesGetKubernetesAlertsListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetKubernetesAlertsListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/alerts][%d] kubernetesGetKubernetesAlertsListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetKubernetesAlertsListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetKubernetesAlertsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubernetesAlertsListNotFound creates a KubernetesGetKubernetesAlertsListNotFound with default headers values
func NewKubernetesGetKubernetesAlertsListNotFound() *KubernetesGetKubernetesAlertsListNotFound {
	return &KubernetesGetKubernetesAlertsListNotFound{}
}

/*
KubernetesGetKubernetesAlertsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetKubernetesAlertsListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get kubernetes alerts list not found response has a 2xx status code
func (o *KubernetesGetKubernetesAlertsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get kubernetes alerts list not found response has a 3xx status code
func (o *KubernetesGetKubernetesAlertsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get kubernetes alerts list not found response has a 4xx status code
func (o *KubernetesGetKubernetesAlertsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get kubernetes alerts list not found response has a 5xx status code
func (o *KubernetesGetKubernetesAlertsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get kubernetes alerts list not found response a status code equal to that given
func (o *KubernetesGetKubernetesAlertsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetKubernetesAlertsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/alerts][%d] kubernetesGetKubernetesAlertsListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetKubernetesAlertsListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/alerts][%d] kubernetesGetKubernetesAlertsListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetKubernetesAlertsListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetKubernetesAlertsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubernetesAlertsListInternalServerError creates a KubernetesGetKubernetesAlertsListInternalServerError with default headers values
func NewKubernetesGetKubernetesAlertsListInternalServerError() *KubernetesGetKubernetesAlertsListInternalServerError {
	return &KubernetesGetKubernetesAlertsListInternalServerError{}
}

/*
KubernetesGetKubernetesAlertsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetKubernetesAlertsListInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get kubernetes alerts list internal server error response has a 2xx status code
func (o *KubernetesGetKubernetesAlertsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get kubernetes alerts list internal server error response has a 3xx status code
func (o *KubernetesGetKubernetesAlertsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get kubernetes alerts list internal server error response has a 4xx status code
func (o *KubernetesGetKubernetesAlertsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get kubernetes alerts list internal server error response has a 5xx status code
func (o *KubernetesGetKubernetesAlertsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get kubernetes alerts list internal server error response a status code equal to that given
func (o *KubernetesGetKubernetesAlertsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetKubernetesAlertsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/alerts][%d] kubernetesGetKubernetesAlertsListInternalServerError ", 500)
}

func (o *KubernetesGetKubernetesAlertsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/alerts][%d] kubernetesGetKubernetesAlertsListInternalServerError ", 500)
}

func (o *KubernetesGetKubernetesAlertsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
