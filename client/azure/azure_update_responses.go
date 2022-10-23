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

// AzureUpdateReader is a Reader for the AzureUpdate structure.
type AzureUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAzureUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAzureUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAzureUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAzureUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAzureUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAzureUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAzureUpdateOK creates a AzureUpdateOK with default headers values
func NewAzureUpdateOK() *AzureUpdateOK {
	return &AzureUpdateOK{}
}

/*
AzureUpdateOK describes a response with status code 200, with default header values.

Success
*/
type AzureUpdateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this azure update o k response has a 2xx status code
func (o *AzureUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure update o k response has a 3xx status code
func (o *AzureUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure update o k response has a 4xx status code
func (o *AzureUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure update o k response has a 5xx status code
func (o *AzureUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this azure update o k response a status code equal to that given
func (o *AzureUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *AzureUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/update][%d] azureUpdateOK  %+v", 200, o.Payload)
}

func (o *AzureUpdateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/update][%d] azureUpdateOK  %+v", 200, o.Payload)
}

func (o *AzureUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AzureUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureUpdateBadRequest creates a AzureUpdateBadRequest with default headers values
func NewAzureUpdateBadRequest() *AzureUpdateBadRequest {
	return &AzureUpdateBadRequest{}
}

/*
AzureUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AzureUpdateBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this azure update bad request response has a 2xx status code
func (o *AzureUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure update bad request response has a 3xx status code
func (o *AzureUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure update bad request response has a 4xx status code
func (o *AzureUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure update bad request response has a 5xx status code
func (o *AzureUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this azure update bad request response a status code equal to that given
func (o *AzureUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AzureUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/update][%d] azureUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *AzureUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/update][%d] azureUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *AzureUpdateBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *AzureUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureUpdateUnauthorized creates a AzureUpdateUnauthorized with default headers values
func NewAzureUpdateUnauthorized() *AzureUpdateUnauthorized {
	return &AzureUpdateUnauthorized{}
}

/*
AzureUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AzureUpdateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this azure update unauthorized response has a 2xx status code
func (o *AzureUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure update unauthorized response has a 3xx status code
func (o *AzureUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure update unauthorized response has a 4xx status code
func (o *AzureUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure update unauthorized response has a 5xx status code
func (o *AzureUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this azure update unauthorized response a status code equal to that given
func (o *AzureUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AzureUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/update][%d] azureUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *AzureUpdateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/update][%d] azureUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *AzureUpdateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzureUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureUpdateForbidden creates a AzureUpdateForbidden with default headers values
func NewAzureUpdateForbidden() *AzureUpdateForbidden {
	return &AzureUpdateForbidden{}
}

/*
AzureUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AzureUpdateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this azure update forbidden response has a 2xx status code
func (o *AzureUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure update forbidden response has a 3xx status code
func (o *AzureUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure update forbidden response has a 4xx status code
func (o *AzureUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure update forbidden response has a 5xx status code
func (o *AzureUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this azure update forbidden response a status code equal to that given
func (o *AzureUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AzureUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/update][%d] azureUpdateForbidden  %+v", 403, o.Payload)
}

func (o *AzureUpdateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/update][%d] azureUpdateForbidden  %+v", 403, o.Payload)
}

func (o *AzureUpdateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzureUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureUpdateNotFound creates a AzureUpdateNotFound with default headers values
func NewAzureUpdateNotFound() *AzureUpdateNotFound {
	return &AzureUpdateNotFound{}
}

/*
AzureUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AzureUpdateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this azure update not found response has a 2xx status code
func (o *AzureUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure update not found response has a 3xx status code
func (o *AzureUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure update not found response has a 4xx status code
func (o *AzureUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure update not found response has a 5xx status code
func (o *AzureUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this azure update not found response a status code equal to that given
func (o *AzureUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AzureUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/update][%d] azureUpdateNotFound  %+v", 404, o.Payload)
}

func (o *AzureUpdateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/update][%d] azureUpdateNotFound  %+v", 404, o.Payload)
}

func (o *AzureUpdateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzureUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureUpdateInternalServerError creates a AzureUpdateInternalServerError with default headers values
func NewAzureUpdateInternalServerError() *AzureUpdateInternalServerError {
	return &AzureUpdateInternalServerError{}
}

/*
AzureUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AzureUpdateInternalServerError struct {
}

// IsSuccess returns true when this azure update internal server error response has a 2xx status code
func (o *AzureUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure update internal server error response has a 3xx status code
func (o *AzureUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure update internal server error response has a 4xx status code
func (o *AzureUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure update internal server error response has a 5xx status code
func (o *AzureUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this azure update internal server error response a status code equal to that given
func (o *AzureUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AzureUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/update][%d] azureUpdateInternalServerError ", 500)
}

func (o *AzureUpdateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Azure/update][%d] azureUpdateInternalServerError ", 500)
}

func (o *AzureUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
