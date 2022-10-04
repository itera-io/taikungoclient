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

// AzureLocationsReader is a Reader for the AzureLocations structure.
type AzureLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAzureLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAzureLocationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAzureLocationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAzureLocationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAzureLocationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAzureLocationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAzureLocationsOK creates a AzureLocationsOK with default headers values
func NewAzureLocationsOK() *AzureLocationsOK {
	return &AzureLocationsOK{}
}

/*
AzureLocationsOK describes a response with status code 200, with default header values.

Success
*/
type AzureLocationsOK struct {
	Payload []string
}

// IsSuccess returns true when this azure locations o k response has a 2xx status code
func (o *AzureLocationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure locations o k response has a 3xx status code
func (o *AzureLocationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure locations o k response has a 4xx status code
func (o *AzureLocationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure locations o k response has a 5xx status code
func (o *AzureLocationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this azure locations o k response a status code equal to that given
func (o *AzureLocationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *AzureLocationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/locations][%d] azureLocationsOK  %+v", 200, o.Payload)
}

func (o *AzureLocationsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/locations][%d] azureLocationsOK  %+v", 200, o.Payload)
}

func (o *AzureLocationsOK) GetPayload() []string {
	return o.Payload
}

func (o *AzureLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureLocationsBadRequest creates a AzureLocationsBadRequest with default headers values
func NewAzureLocationsBadRequest() *AzureLocationsBadRequest {
	return &AzureLocationsBadRequest{}
}

/*
AzureLocationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AzureLocationsBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this azure locations bad request response has a 2xx status code
func (o *AzureLocationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure locations bad request response has a 3xx status code
func (o *AzureLocationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure locations bad request response has a 4xx status code
func (o *AzureLocationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure locations bad request response has a 5xx status code
func (o *AzureLocationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this azure locations bad request response a status code equal to that given
func (o *AzureLocationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AzureLocationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/locations][%d] azureLocationsBadRequest  %+v", 400, o.Payload)
}

func (o *AzureLocationsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/locations][%d] azureLocationsBadRequest  %+v", 400, o.Payload)
}

func (o *AzureLocationsBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AzureLocationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureLocationsUnauthorized creates a AzureLocationsUnauthorized with default headers values
func NewAzureLocationsUnauthorized() *AzureLocationsUnauthorized {
	return &AzureLocationsUnauthorized{}
}

/*
AzureLocationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AzureLocationsUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this azure locations unauthorized response has a 2xx status code
func (o *AzureLocationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure locations unauthorized response has a 3xx status code
func (o *AzureLocationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure locations unauthorized response has a 4xx status code
func (o *AzureLocationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure locations unauthorized response has a 5xx status code
func (o *AzureLocationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this azure locations unauthorized response a status code equal to that given
func (o *AzureLocationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AzureLocationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/locations][%d] azureLocationsUnauthorized  %+v", 401, o.Payload)
}

func (o *AzureLocationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/locations][%d] azureLocationsUnauthorized  %+v", 401, o.Payload)
}

func (o *AzureLocationsUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AzureLocationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureLocationsForbidden creates a AzureLocationsForbidden with default headers values
func NewAzureLocationsForbidden() *AzureLocationsForbidden {
	return &AzureLocationsForbidden{}
}

/*
AzureLocationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AzureLocationsForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this azure locations forbidden response has a 2xx status code
func (o *AzureLocationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure locations forbidden response has a 3xx status code
func (o *AzureLocationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure locations forbidden response has a 4xx status code
func (o *AzureLocationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure locations forbidden response has a 5xx status code
func (o *AzureLocationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this azure locations forbidden response a status code equal to that given
func (o *AzureLocationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AzureLocationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/locations][%d] azureLocationsForbidden  %+v", 403, o.Payload)
}

func (o *AzureLocationsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/locations][%d] azureLocationsForbidden  %+v", 403, o.Payload)
}

func (o *AzureLocationsForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AzureLocationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureLocationsNotFound creates a AzureLocationsNotFound with default headers values
func NewAzureLocationsNotFound() *AzureLocationsNotFound {
	return &AzureLocationsNotFound{}
}

/*
AzureLocationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AzureLocationsNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this azure locations not found response has a 2xx status code
func (o *AzureLocationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure locations not found response has a 3xx status code
func (o *AzureLocationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure locations not found response has a 4xx status code
func (o *AzureLocationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure locations not found response has a 5xx status code
func (o *AzureLocationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this azure locations not found response a status code equal to that given
func (o *AzureLocationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AzureLocationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/locations][%d] azureLocationsNotFound  %+v", 404, o.Payload)
}

func (o *AzureLocationsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/locations][%d] azureLocationsNotFound  %+v", 404, o.Payload)
}

func (o *AzureLocationsNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AzureLocationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureLocationsInternalServerError creates a AzureLocationsInternalServerError with default headers values
func NewAzureLocationsInternalServerError() *AzureLocationsInternalServerError {
	return &AzureLocationsInternalServerError{}
}

/*
AzureLocationsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AzureLocationsInternalServerError struct {
}

// IsSuccess returns true when this azure locations internal server error response has a 2xx status code
func (o *AzureLocationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure locations internal server error response has a 3xx status code
func (o *AzureLocationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure locations internal server error response has a 4xx status code
func (o *AzureLocationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure locations internal server error response has a 5xx status code
func (o *AzureLocationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this azure locations internal server error response a status code equal to that given
func (o *AzureLocationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AzureLocationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/locations][%d] azureLocationsInternalServerError ", 500)
}

func (o *AzureLocationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/locations][%d] azureLocationsInternalServerError ", 500)
}

func (o *AzureLocationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
