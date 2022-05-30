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

// AzureOffersReader is a Reader for the AzureOffers structure.
type AzureOffersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureOffersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAzureOffersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAzureOffersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAzureOffersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAzureOffersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAzureOffersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAzureOffersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAzureOffersOK creates a AzureOffersOK with default headers values
func NewAzureOffersOK() *AzureOffersOK {
	return &AzureOffersOK{}
}

/* AzureOffersOK describes a response with status code 200, with default header values.

Success
*/
type AzureOffersOK struct {
	Payload *models.AzureOffersList
}

func (o *AzureOffersOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/offers/{cloudId}/{publisher}][%d] azureOffersOK  %+v", 200, o.Payload)
}
func (o *AzureOffersOK) GetPayload() *models.AzureOffersList {
	return o.Payload
}

func (o *AzureOffersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureOffersList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureOffersBadRequest creates a AzureOffersBadRequest with default headers values
func NewAzureOffersBadRequest() *AzureOffersBadRequest {
	return &AzureOffersBadRequest{}
}

/* AzureOffersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AzureOffersBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AzureOffersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/offers/{cloudId}/{publisher}][%d] azureOffersBadRequest  %+v", 400, o.Payload)
}
func (o *AzureOffersBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AzureOffersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureOffersUnauthorized creates a AzureOffersUnauthorized with default headers values
func NewAzureOffersUnauthorized() *AzureOffersUnauthorized {
	return &AzureOffersUnauthorized{}
}

/* AzureOffersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AzureOffersUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AzureOffersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/offers/{cloudId}/{publisher}][%d] azureOffersUnauthorized  %+v", 401, o.Payload)
}
func (o *AzureOffersUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzureOffersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureOffersForbidden creates a AzureOffersForbidden with default headers values
func NewAzureOffersForbidden() *AzureOffersForbidden {
	return &AzureOffersForbidden{}
}

/* AzureOffersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AzureOffersForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AzureOffersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/offers/{cloudId}/{publisher}][%d] azureOffersForbidden  %+v", 403, o.Payload)
}
func (o *AzureOffersForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzureOffersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureOffersNotFound creates a AzureOffersNotFound with default headers values
func NewAzureOffersNotFound() *AzureOffersNotFound {
	return &AzureOffersNotFound{}
}

/* AzureOffersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AzureOffersNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AzureOffersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/offers/{cloudId}/{publisher}][%d] azureOffersNotFound  %+v", 404, o.Payload)
}
func (o *AzureOffersNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AzureOffersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureOffersInternalServerError creates a AzureOffersInternalServerError with default headers values
func NewAzureOffersInternalServerError() *AzureOffersInternalServerError {
	return &AzureOffersInternalServerError{}
}

/* AzureOffersInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AzureOffersInternalServerError struct {
}

func (o *AzureOffersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/offers/{cloudId}/{publisher}][%d] azureOffersInternalServerError ", 500)
}

func (o *AzureOffersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}