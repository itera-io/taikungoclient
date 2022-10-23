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

// AzurePublishersReader is a Reader for the AzurePublishers structure.
type AzurePublishersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzurePublishersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAzurePublishersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAzurePublishersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAzurePublishersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAzurePublishersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAzurePublishersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAzurePublishersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAzurePublishersOK creates a AzurePublishersOK with default headers values
func NewAzurePublishersOK() *AzurePublishersOK {
	return &AzurePublishersOK{}
}

/*
AzurePublishersOK describes a response with status code 200, with default header values.

Success
*/
type AzurePublishersOK struct {
	Payload *models.AzurePublishersList
}

// IsSuccess returns true when this azure publishers o k response has a 2xx status code
func (o *AzurePublishersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure publishers o k response has a 3xx status code
func (o *AzurePublishersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure publishers o k response has a 4xx status code
func (o *AzurePublishersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure publishers o k response has a 5xx status code
func (o *AzurePublishersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this azure publishers o k response a status code equal to that given
func (o *AzurePublishersOK) IsCode(code int) bool {
	return code == 200
}

func (o *AzurePublishersOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersOK  %+v", 200, o.Payload)
}

func (o *AzurePublishersOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersOK  %+v", 200, o.Payload)
}

func (o *AzurePublishersOK) GetPayload() *models.AzurePublishersList {
	return o.Payload
}

func (o *AzurePublishersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzurePublishersList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzurePublishersBadRequest creates a AzurePublishersBadRequest with default headers values
func NewAzurePublishersBadRequest() *AzurePublishersBadRequest {
	return &AzurePublishersBadRequest{}
}

/*
AzurePublishersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AzurePublishersBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this azure publishers bad request response has a 2xx status code
func (o *AzurePublishersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure publishers bad request response has a 3xx status code
func (o *AzurePublishersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure publishers bad request response has a 4xx status code
func (o *AzurePublishersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure publishers bad request response has a 5xx status code
func (o *AzurePublishersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this azure publishers bad request response a status code equal to that given
func (o *AzurePublishersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AzurePublishersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersBadRequest  %+v", 400, o.Payload)
}

func (o *AzurePublishersBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersBadRequest  %+v", 400, o.Payload)
}

func (o *AzurePublishersBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *AzurePublishersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzurePublishersUnauthorized creates a AzurePublishersUnauthorized with default headers values
func NewAzurePublishersUnauthorized() *AzurePublishersUnauthorized {
	return &AzurePublishersUnauthorized{}
}

/*
AzurePublishersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AzurePublishersUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this azure publishers unauthorized response has a 2xx status code
func (o *AzurePublishersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure publishers unauthorized response has a 3xx status code
func (o *AzurePublishersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure publishers unauthorized response has a 4xx status code
func (o *AzurePublishersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure publishers unauthorized response has a 5xx status code
func (o *AzurePublishersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this azure publishers unauthorized response a status code equal to that given
func (o *AzurePublishersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AzurePublishersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersUnauthorized  %+v", 401, o.Payload)
}

func (o *AzurePublishersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersUnauthorized  %+v", 401, o.Payload)
}

func (o *AzurePublishersUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzurePublishersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzurePublishersForbidden creates a AzurePublishersForbidden with default headers values
func NewAzurePublishersForbidden() *AzurePublishersForbidden {
	return &AzurePublishersForbidden{}
}

/*
AzurePublishersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AzurePublishersForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this azure publishers forbidden response has a 2xx status code
func (o *AzurePublishersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure publishers forbidden response has a 3xx status code
func (o *AzurePublishersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure publishers forbidden response has a 4xx status code
func (o *AzurePublishersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure publishers forbidden response has a 5xx status code
func (o *AzurePublishersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this azure publishers forbidden response a status code equal to that given
func (o *AzurePublishersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AzurePublishersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersForbidden  %+v", 403, o.Payload)
}

func (o *AzurePublishersForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersForbidden  %+v", 403, o.Payload)
}

func (o *AzurePublishersForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzurePublishersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzurePublishersNotFound creates a AzurePublishersNotFound with default headers values
func NewAzurePublishersNotFound() *AzurePublishersNotFound {
	return &AzurePublishersNotFound{}
}

/*
AzurePublishersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AzurePublishersNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this azure publishers not found response has a 2xx status code
func (o *AzurePublishersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure publishers not found response has a 3xx status code
func (o *AzurePublishersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure publishers not found response has a 4xx status code
func (o *AzurePublishersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure publishers not found response has a 5xx status code
func (o *AzurePublishersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this azure publishers not found response a status code equal to that given
func (o *AzurePublishersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AzurePublishersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersNotFound  %+v", 404, o.Payload)
}

func (o *AzurePublishersNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersNotFound  %+v", 404, o.Payload)
}

func (o *AzurePublishersNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzurePublishersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzurePublishersInternalServerError creates a AzurePublishersInternalServerError with default headers values
func NewAzurePublishersInternalServerError() *AzurePublishersInternalServerError {
	return &AzurePublishersInternalServerError{}
}

/*
AzurePublishersInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AzurePublishersInternalServerError struct {
}

// IsSuccess returns true when this azure publishers internal server error response has a 2xx status code
func (o *AzurePublishersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure publishers internal server error response has a 3xx status code
func (o *AzurePublishersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure publishers internal server error response has a 4xx status code
func (o *AzurePublishersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure publishers internal server error response has a 5xx status code
func (o *AzurePublishersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this azure publishers internal server error response a status code equal to that given
func (o *AzurePublishersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AzurePublishersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersInternalServerError ", 500)
}

func (o *AzurePublishersInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersInternalServerError ", 500)
}

func (o *AzurePublishersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
