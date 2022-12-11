// Code generated by go-swagger; DO NOT EDIT.

package project_quotas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectQuotasListReader is a Reader for the ProjectQuotasList structure.
type ProjectQuotasListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectQuotasListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectQuotasListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectQuotasListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectQuotasListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectQuotasListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectQuotasListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectQuotasListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectQuotasListOK creates a ProjectQuotasListOK with default headers values
func NewProjectQuotasListOK() *ProjectQuotasListOK {
	return &ProjectQuotasListOK{}
}

/*
ProjectQuotasListOK describes a response with status code 200, with default header values.

Success
*/
type ProjectQuotasListOK struct {
	Payload *models.ProjectQuotaList
}

// IsSuccess returns true when this project quotas list o k response has a 2xx status code
func (o *ProjectQuotasListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project quotas list o k response has a 3xx status code
func (o *ProjectQuotasListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project quotas list o k response has a 4xx status code
func (o *ProjectQuotasListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project quotas list o k response has a 5xx status code
func (o *ProjectQuotasListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project quotas list o k response a status code equal to that given
func (o *ProjectQuotasListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectQuotasListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectQuotas][%d] projectQuotasListOK  %+v", 200, o.Payload)
}

func (o *ProjectQuotasListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectQuotas][%d] projectQuotasListOK  %+v", 200, o.Payload)
}

func (o *ProjectQuotasListOK) GetPayload() *models.ProjectQuotaList {
	return o.Payload
}

func (o *ProjectQuotasListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectQuotaList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectQuotasListBadRequest creates a ProjectQuotasListBadRequest with default headers values
func NewProjectQuotasListBadRequest() *ProjectQuotasListBadRequest {
	return &ProjectQuotasListBadRequest{}
}

/*
ProjectQuotasListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectQuotasListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this project quotas list bad request response has a 2xx status code
func (o *ProjectQuotasListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project quotas list bad request response has a 3xx status code
func (o *ProjectQuotasListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project quotas list bad request response has a 4xx status code
func (o *ProjectQuotasListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project quotas list bad request response has a 5xx status code
func (o *ProjectQuotasListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project quotas list bad request response a status code equal to that given
func (o *ProjectQuotasListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectQuotasListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectQuotas][%d] projectQuotasListBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectQuotasListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectQuotas][%d] projectQuotasListBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectQuotasListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectQuotasListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectQuotasListUnauthorized creates a ProjectQuotasListUnauthorized with default headers values
func NewProjectQuotasListUnauthorized() *ProjectQuotasListUnauthorized {
	return &ProjectQuotasListUnauthorized{}
}

/*
ProjectQuotasListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectQuotasListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project quotas list unauthorized response has a 2xx status code
func (o *ProjectQuotasListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project quotas list unauthorized response has a 3xx status code
func (o *ProjectQuotasListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project quotas list unauthorized response has a 4xx status code
func (o *ProjectQuotasListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project quotas list unauthorized response has a 5xx status code
func (o *ProjectQuotasListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project quotas list unauthorized response a status code equal to that given
func (o *ProjectQuotasListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectQuotasListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectQuotas][%d] projectQuotasListUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectQuotasListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectQuotas][%d] projectQuotasListUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectQuotasListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectQuotasListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectQuotasListForbidden creates a ProjectQuotasListForbidden with default headers values
func NewProjectQuotasListForbidden() *ProjectQuotasListForbidden {
	return &ProjectQuotasListForbidden{}
}

/*
ProjectQuotasListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectQuotasListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project quotas list forbidden response has a 2xx status code
func (o *ProjectQuotasListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project quotas list forbidden response has a 3xx status code
func (o *ProjectQuotasListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project quotas list forbidden response has a 4xx status code
func (o *ProjectQuotasListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project quotas list forbidden response has a 5xx status code
func (o *ProjectQuotasListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project quotas list forbidden response a status code equal to that given
func (o *ProjectQuotasListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectQuotasListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectQuotas][%d] projectQuotasListForbidden  %+v", 403, o.Payload)
}

func (o *ProjectQuotasListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectQuotas][%d] projectQuotasListForbidden  %+v", 403, o.Payload)
}

func (o *ProjectQuotasListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectQuotasListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectQuotasListNotFound creates a ProjectQuotasListNotFound with default headers values
func NewProjectQuotasListNotFound() *ProjectQuotasListNotFound {
	return &ProjectQuotasListNotFound{}
}

/*
ProjectQuotasListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectQuotasListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project quotas list not found response has a 2xx status code
func (o *ProjectQuotasListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project quotas list not found response has a 3xx status code
func (o *ProjectQuotasListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project quotas list not found response has a 4xx status code
func (o *ProjectQuotasListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project quotas list not found response has a 5xx status code
func (o *ProjectQuotasListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project quotas list not found response a status code equal to that given
func (o *ProjectQuotasListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectQuotasListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectQuotas][%d] projectQuotasListNotFound  %+v", 404, o.Payload)
}

func (o *ProjectQuotasListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectQuotas][%d] projectQuotasListNotFound  %+v", 404, o.Payload)
}

func (o *ProjectQuotasListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectQuotasListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectQuotasListInternalServerError creates a ProjectQuotasListInternalServerError with default headers values
func NewProjectQuotasListInternalServerError() *ProjectQuotasListInternalServerError {
	return &ProjectQuotasListInternalServerError{}
}

/*
ProjectQuotasListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectQuotasListInternalServerError struct {
}

// IsSuccess returns true when this project quotas list internal server error response has a 2xx status code
func (o *ProjectQuotasListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project quotas list internal server error response has a 3xx status code
func (o *ProjectQuotasListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project quotas list internal server error response has a 4xx status code
func (o *ProjectQuotasListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project quotas list internal server error response has a 5xx status code
func (o *ProjectQuotasListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project quotas list internal server error response a status code equal to that given
func (o *ProjectQuotasListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectQuotasListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectQuotas][%d] projectQuotasListInternalServerError ", 500)
}

func (o *ProjectQuotasListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectQuotas][%d] projectQuotasListInternalServerError ", 500)
}

func (o *ProjectQuotasListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
