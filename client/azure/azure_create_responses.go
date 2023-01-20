// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AzureCreateReader is a Reader for the AzureCreate structure.
type AzureCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAzureCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAzureCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAzureCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAzureCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAzureCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAzureCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAzureCreateOK creates a AzureCreateOK with default headers values
func NewAzureCreateOK() *AzureCreateOK {
	return &AzureCreateOK{}
}

/*
AzureCreateOK describes a response with status code 200, with default header values.

Success
*/
type AzureCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this azure create o k response has a 2xx status code
func (o *AzureCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure create o k response has a 3xx status code
func (o *AzureCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure create o k response has a 4xx status code
func (o *AzureCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure create o k response has a 5xx status code
func (o *AzureCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this azure create o k response a status code equal to that given
func (o *AzureCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the azure create o k response
func (o *AzureCreateOK) Code() int {
	return 200
}

func (o *AzureCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/create][%d] azureCreateOK  %+v", 200, o.Payload)
}

func (o *AzureCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/create][%d] azureCreateOK  %+v", 200, o.Payload)
}

func (o *AzureCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *AzureCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureCreateBadRequest creates a AzureCreateBadRequest with default headers values
func NewAzureCreateBadRequest() *AzureCreateBadRequest {
	return &AzureCreateBadRequest{}
}

/*
AzureCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AzureCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this azure create bad request response has a 2xx status code
func (o *AzureCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure create bad request response has a 3xx status code
func (o *AzureCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure create bad request response has a 4xx status code
func (o *AzureCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure create bad request response has a 5xx status code
func (o *AzureCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this azure create bad request response a status code equal to that given
func (o *AzureCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the azure create bad request response
func (o *AzureCreateBadRequest) Code() int {
	return 400
}

func (o *AzureCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/create][%d] azureCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AzureCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/create][%d] azureCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AzureCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *AzureCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureCreateUnauthorized creates a AzureCreateUnauthorized with default headers values
func NewAzureCreateUnauthorized() *AzureCreateUnauthorized {
	return &AzureCreateUnauthorized{}
}

/*
AzureCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AzureCreateUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this azure create unauthorized response has a 2xx status code
func (o *AzureCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure create unauthorized response has a 3xx status code
func (o *AzureCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure create unauthorized response has a 4xx status code
func (o *AzureCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure create unauthorized response has a 5xx status code
func (o *AzureCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this azure create unauthorized response a status code equal to that given
func (o *AzureCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the azure create unauthorized response
func (o *AzureCreateUnauthorized) Code() int {
	return 401
}

func (o *AzureCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/create][%d] azureCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AzureCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/create][%d] azureCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AzureCreateUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *AzureCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureCreateForbidden creates a AzureCreateForbidden with default headers values
func NewAzureCreateForbidden() *AzureCreateForbidden {
	return &AzureCreateForbidden{}
}

/*
AzureCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AzureCreateForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this azure create forbidden response has a 2xx status code
func (o *AzureCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure create forbidden response has a 3xx status code
func (o *AzureCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure create forbidden response has a 4xx status code
func (o *AzureCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure create forbidden response has a 5xx status code
func (o *AzureCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this azure create forbidden response a status code equal to that given
func (o *AzureCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the azure create forbidden response
func (o *AzureCreateForbidden) Code() int {
	return 403
}

func (o *AzureCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/create][%d] azureCreateForbidden  %+v", 403, o.Payload)
}

func (o *AzureCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/create][%d] azureCreateForbidden  %+v", 403, o.Payload)
}

func (o *AzureCreateForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *AzureCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureCreateNotFound creates a AzureCreateNotFound with default headers values
func NewAzureCreateNotFound() *AzureCreateNotFound {
	return &AzureCreateNotFound{}
}

/*
AzureCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AzureCreateNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this azure create not found response has a 2xx status code
func (o *AzureCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure create not found response has a 3xx status code
func (o *AzureCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure create not found response has a 4xx status code
func (o *AzureCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure create not found response has a 5xx status code
func (o *AzureCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this azure create not found response a status code equal to that given
func (o *AzureCreateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the azure create not found response
func (o *AzureCreateNotFound) Code() int {
	return 404
}

func (o *AzureCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/create][%d] azureCreateNotFound  %+v", 404, o.Payload)
}

func (o *AzureCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/create][%d] azureCreateNotFound  %+v", 404, o.Payload)
}

func (o *AzureCreateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *AzureCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureCreateInternalServerError creates a AzureCreateInternalServerError with default headers values
func NewAzureCreateInternalServerError() *AzureCreateInternalServerError {
	return &AzureCreateInternalServerError{}
}

/*
AzureCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AzureCreateInternalServerError struct {
}

// IsSuccess returns true when this azure create internal server error response has a 2xx status code
func (o *AzureCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure create internal server error response has a 3xx status code
func (o *AzureCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure create internal server error response has a 4xx status code
func (o *AzureCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure create internal server error response has a 5xx status code
func (o *AzureCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this azure create internal server error response a status code equal to that given
func (o *AzureCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the azure create internal server error response
func (o *AzureCreateInternalServerError) Code() int {
	return 500
}

func (o *AzureCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/create][%d] azureCreateInternalServerError ", 500)
}

func (o *AzureCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/create][%d] azureCreateInternalServerError ", 500)
}

func (o *AzureCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
