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

// AzureListReader is a Reader for the AzureList structure.
type AzureListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAzureListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAzureListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAzureListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAzureListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAzureListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAzureListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAzureListOK creates a AzureListOK with default headers values
func NewAzureListOK() *AzureListOK {
	return &AzureListOK{}
}

/* AzureListOK describes a response with status code 200, with default header values.

Success
*/
type AzureListOK struct {
	Payload *models.AzureCredentialList
}

func (o *AzureListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/list][%d] azureListOK  %+v", 200, o.Payload)
}
func (o *AzureListOK) GetPayload() *models.AzureCredentialList {
	return o.Payload
}

func (o *AzureListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureCredentialList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureListBadRequest creates a AzureListBadRequest with default headers values
func NewAzureListBadRequest() *AzureListBadRequest {
	return &AzureListBadRequest{}
}

/* AzureListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AzureListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AzureListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/list][%d] azureListBadRequest  %+v", 400, o.Payload)
}
func (o *AzureListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AzureListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureListUnauthorized creates a AzureListUnauthorized with default headers values
func NewAzureListUnauthorized() *AzureListUnauthorized {
	return &AzureListUnauthorized{}
}

/* AzureListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AzureListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AzureListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/list][%d] azureListUnauthorized  %+v", 401, o.Payload)
}
func (o *AzureListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzureListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureListForbidden creates a AzureListForbidden with default headers values
func NewAzureListForbidden() *AzureListForbidden {
	return &AzureListForbidden{}
}

/* AzureListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AzureListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AzureListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/list][%d] azureListForbidden  %+v", 403, o.Payload)
}
func (o *AzureListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzureListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureListNotFound creates a AzureListNotFound with default headers values
func NewAzureListNotFound() *AzureListNotFound {
	return &AzureListNotFound{}
}

/* AzureListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AzureListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AzureListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/list][%d] azureListNotFound  %+v", 404, o.Payload)
}
func (o *AzureListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzureListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureListInternalServerError creates a AzureListInternalServerError with default headers values
func NewAzureListInternalServerError() *AzureListInternalServerError {
	return &AzureListInternalServerError{}
}

/* AzureListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AzureListInternalServerError struct {
}

func (o *AzureListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/list][%d] azureListInternalServerError ", 500)
}

func (o *AzureListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}