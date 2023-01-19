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

// AzureZonesReader is a Reader for the AzureZones structure.
type AzureZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAzureZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAzureZonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAzureZonesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAzureZonesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAzureZonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAzureZonesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAzureZonesOK creates a AzureZonesOK with default headers values
func NewAzureZonesOK() *AzureZonesOK {
	return &AzureZonesOK{}
}

/*
AzureZonesOK describes a response with status code 200, with default header values.

Success
*/
type AzureZonesOK struct {
	Payload *models.AzResult
}

// IsSuccess returns true when this azure zones o k response has a 2xx status code
func (o *AzureZonesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure zones o k response has a 3xx status code
func (o *AzureZonesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure zones o k response has a 4xx status code
func (o *AzureZonesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure zones o k response has a 5xx status code
func (o *AzureZonesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this azure zones o k response a status code equal to that given
func (o *AzureZonesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the azure zones o k response
func (o *AzureZonesOK) Code() int {
	return 200
}

func (o *AzureZonesOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/zones][%d] azureZonesOK  %+v", 200, o.Payload)
}

func (o *AzureZonesOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/zones][%d] azureZonesOK  %+v", 200, o.Payload)
}

func (o *AzureZonesOK) GetPayload() *models.AzResult {
	return o.Payload
}

func (o *AzureZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureZonesBadRequest creates a AzureZonesBadRequest with default headers values
func NewAzureZonesBadRequest() *AzureZonesBadRequest {
	return &AzureZonesBadRequest{}
}

/*
AzureZonesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AzureZonesBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this azure zones bad request response has a 2xx status code
func (o *AzureZonesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure zones bad request response has a 3xx status code
func (o *AzureZonesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure zones bad request response has a 4xx status code
func (o *AzureZonesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure zones bad request response has a 5xx status code
func (o *AzureZonesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this azure zones bad request response a status code equal to that given
func (o *AzureZonesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the azure zones bad request response
func (o *AzureZonesBadRequest) Code() int {
	return 400
}

func (o *AzureZonesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/zones][%d] azureZonesBadRequest  %+v", 400, o.Payload)
}

func (o *AzureZonesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/zones][%d] azureZonesBadRequest  %+v", 400, o.Payload)
}

func (o *AzureZonesBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzureZonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureZonesUnauthorized creates a AzureZonesUnauthorized with default headers values
func NewAzureZonesUnauthorized() *AzureZonesUnauthorized {
	return &AzureZonesUnauthorized{}
}

/*
AzureZonesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AzureZonesUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this azure zones unauthorized response has a 2xx status code
func (o *AzureZonesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure zones unauthorized response has a 3xx status code
func (o *AzureZonesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure zones unauthorized response has a 4xx status code
func (o *AzureZonesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure zones unauthorized response has a 5xx status code
func (o *AzureZonesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this azure zones unauthorized response a status code equal to that given
func (o *AzureZonesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the azure zones unauthorized response
func (o *AzureZonesUnauthorized) Code() int {
	return 401
}

func (o *AzureZonesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/zones][%d] azureZonesUnauthorized  %+v", 401, o.Payload)
}

func (o *AzureZonesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/zones][%d] azureZonesUnauthorized  %+v", 401, o.Payload)
}

func (o *AzureZonesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzureZonesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureZonesForbidden creates a AzureZonesForbidden with default headers values
func NewAzureZonesForbidden() *AzureZonesForbidden {
	return &AzureZonesForbidden{}
}

/*
AzureZonesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AzureZonesForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this azure zones forbidden response has a 2xx status code
func (o *AzureZonesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure zones forbidden response has a 3xx status code
func (o *AzureZonesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure zones forbidden response has a 4xx status code
func (o *AzureZonesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure zones forbidden response has a 5xx status code
func (o *AzureZonesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this azure zones forbidden response a status code equal to that given
func (o *AzureZonesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the azure zones forbidden response
func (o *AzureZonesForbidden) Code() int {
	return 403
}

func (o *AzureZonesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/zones][%d] azureZonesForbidden  %+v", 403, o.Payload)
}

func (o *AzureZonesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/zones][%d] azureZonesForbidden  %+v", 403, o.Payload)
}

func (o *AzureZonesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzureZonesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureZonesNotFound creates a AzureZonesNotFound with default headers values
func NewAzureZonesNotFound() *AzureZonesNotFound {
	return &AzureZonesNotFound{}
}

/*
AzureZonesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AzureZonesNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this azure zones not found response has a 2xx status code
func (o *AzureZonesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure zones not found response has a 3xx status code
func (o *AzureZonesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure zones not found response has a 4xx status code
func (o *AzureZonesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure zones not found response has a 5xx status code
func (o *AzureZonesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this azure zones not found response a status code equal to that given
func (o *AzureZonesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the azure zones not found response
func (o *AzureZonesNotFound) Code() int {
	return 404
}

func (o *AzureZonesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/zones][%d] azureZonesNotFound  %+v", 404, o.Payload)
}

func (o *AzureZonesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/zones][%d] azureZonesNotFound  %+v", 404, o.Payload)
}

func (o *AzureZonesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzureZonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureZonesInternalServerError creates a AzureZonesInternalServerError with default headers values
func NewAzureZonesInternalServerError() *AzureZonesInternalServerError {
	return &AzureZonesInternalServerError{}
}

/*
AzureZonesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AzureZonesInternalServerError struct {
}

// IsSuccess returns true when this azure zones internal server error response has a 2xx status code
func (o *AzureZonesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure zones internal server error response has a 3xx status code
func (o *AzureZonesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure zones internal server error response has a 4xx status code
func (o *AzureZonesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure zones internal server error response has a 5xx status code
func (o *AzureZonesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this azure zones internal server error response a status code equal to that given
func (o *AzureZonesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the azure zones internal server error response
func (o *AzureZonesInternalServerError) Code() int {
	return 500
}

func (o *AzureZonesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/zones][%d] azureZonesInternalServerError ", 500)
}

func (o *AzureZonesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/zones][%d] azureZonesInternalServerError ", 500)
}

func (o *AzureZonesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
