// Code generated by go-swagger; DO NOT EDIT.

package kubespray

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// KubesprayListReader is a Reader for the KubesprayList structure.
type KubesprayListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubesprayListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubesprayListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubesprayListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubesprayListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubesprayListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubesprayListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubesprayListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubesprayListOK creates a KubesprayListOK with default headers values
func NewKubesprayListOK() *KubesprayListOK {
	return &KubesprayListOK{}
}

/*
KubesprayListOK describes a response with status code 200, with default header values.

Success
*/
type KubesprayListOK struct {
	Payload *models.Kubesprays
}

// IsSuccess returns true when this kubespray list o k response has a 2xx status code
func (o *KubesprayListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubespray list o k response has a 3xx status code
func (o *KubesprayListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubespray list o k response has a 4xx status code
func (o *KubesprayListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubespray list o k response has a 5xx status code
func (o *KubesprayListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubespray list o k response a status code equal to that given
func (o *KubesprayListOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubesprayListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubespray][%d] kubesprayListOK  %+v", 200, o.Payload)
}

func (o *KubesprayListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubespray][%d] kubesprayListOK  %+v", 200, o.Payload)
}

func (o *KubesprayListOK) GetPayload() *models.Kubesprays {
	return o.Payload
}

func (o *KubesprayListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Kubesprays)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubesprayListBadRequest creates a KubesprayListBadRequest with default headers values
func NewKubesprayListBadRequest() *KubesprayListBadRequest {
	return &KubesprayListBadRequest{}
}

/*
KubesprayListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubesprayListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kubespray list bad request response has a 2xx status code
func (o *KubesprayListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubespray list bad request response has a 3xx status code
func (o *KubesprayListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubespray list bad request response has a 4xx status code
func (o *KubesprayListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubespray list bad request response has a 5xx status code
func (o *KubesprayListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubespray list bad request response a status code equal to that given
func (o *KubesprayListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubesprayListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubespray][%d] kubesprayListBadRequest  %+v", 400, o.Payload)
}

func (o *KubesprayListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubespray][%d] kubesprayListBadRequest  %+v", 400, o.Payload)
}

func (o *KubesprayListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubesprayListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubesprayListUnauthorized creates a KubesprayListUnauthorized with default headers values
func NewKubesprayListUnauthorized() *KubesprayListUnauthorized {
	return &KubesprayListUnauthorized{}
}

/*
KubesprayListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubesprayListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubespray list unauthorized response has a 2xx status code
func (o *KubesprayListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubespray list unauthorized response has a 3xx status code
func (o *KubesprayListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubespray list unauthorized response has a 4xx status code
func (o *KubesprayListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubespray list unauthorized response has a 5xx status code
func (o *KubesprayListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubespray list unauthorized response a status code equal to that given
func (o *KubesprayListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubesprayListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubespray][%d] kubesprayListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubesprayListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubespray][%d] kubesprayListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubesprayListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubesprayListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubesprayListForbidden creates a KubesprayListForbidden with default headers values
func NewKubesprayListForbidden() *KubesprayListForbidden {
	return &KubesprayListForbidden{}
}

/*
KubesprayListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubesprayListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubespray list forbidden response has a 2xx status code
func (o *KubesprayListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubespray list forbidden response has a 3xx status code
func (o *KubesprayListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubespray list forbidden response has a 4xx status code
func (o *KubesprayListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubespray list forbidden response has a 5xx status code
func (o *KubesprayListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubespray list forbidden response a status code equal to that given
func (o *KubesprayListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubesprayListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubespray][%d] kubesprayListForbidden  %+v", 403, o.Payload)
}

func (o *KubesprayListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubespray][%d] kubesprayListForbidden  %+v", 403, o.Payload)
}

func (o *KubesprayListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubesprayListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubesprayListNotFound creates a KubesprayListNotFound with default headers values
func NewKubesprayListNotFound() *KubesprayListNotFound {
	return &KubesprayListNotFound{}
}

/*
KubesprayListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubesprayListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubespray list not found response has a 2xx status code
func (o *KubesprayListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubespray list not found response has a 3xx status code
func (o *KubesprayListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubespray list not found response has a 4xx status code
func (o *KubesprayListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubespray list not found response has a 5xx status code
func (o *KubesprayListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubespray list not found response a status code equal to that given
func (o *KubesprayListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubesprayListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubespray][%d] kubesprayListNotFound  %+v", 404, o.Payload)
}

func (o *KubesprayListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubespray][%d] kubesprayListNotFound  %+v", 404, o.Payload)
}

func (o *KubesprayListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubesprayListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubesprayListInternalServerError creates a KubesprayListInternalServerError with default headers values
func NewKubesprayListInternalServerError() *KubesprayListInternalServerError {
	return &KubesprayListInternalServerError{}
}

/*
KubesprayListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubesprayListInternalServerError struct {
}

// IsSuccess returns true when this kubespray list internal server error response has a 2xx status code
func (o *KubesprayListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubespray list internal server error response has a 3xx status code
func (o *KubesprayListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubespray list internal server error response has a 4xx status code
func (o *KubesprayListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubespray list internal server error response has a 5xx status code
func (o *KubesprayListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubespray list internal server error response a status code equal to that given
func (o *KubesprayListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubesprayListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubespray][%d] kubesprayListInternalServerError ", 500)
}

func (o *KubesprayListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubespray][%d] kubesprayListInternalServerError ", 500)
}

func (o *KubesprayListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
