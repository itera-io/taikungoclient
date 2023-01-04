// Code generated by go-swagger; DO NOT EDIT.

package google_cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// GoogleCloudBillingAccountListReader is a Reader for the GoogleCloudBillingAccountList structure.
type GoogleCloudBillingAccountListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GoogleCloudBillingAccountListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGoogleCloudBillingAccountListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGoogleCloudBillingAccountListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGoogleCloudBillingAccountListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGoogleCloudBillingAccountListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGoogleCloudBillingAccountListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGoogleCloudBillingAccountListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGoogleCloudBillingAccountListOK creates a GoogleCloudBillingAccountListOK with default headers values
func NewGoogleCloudBillingAccountListOK() *GoogleCloudBillingAccountListOK {
	return &GoogleCloudBillingAccountListOK{}
}

/*
GoogleCloudBillingAccountListOK describes a response with status code 200, with default header values.

Success
*/
type GoogleCloudBillingAccountListOK struct {
	Payload []*models.CommonStringBasedDropdownDto
}

// IsSuccess returns true when this google cloud billing account list o k response has a 2xx status code
func (o *GoogleCloudBillingAccountListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this google cloud billing account list o k response has a 3xx status code
func (o *GoogleCloudBillingAccountListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud billing account list o k response has a 4xx status code
func (o *GoogleCloudBillingAccountListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this google cloud billing account list o k response has a 5xx status code
func (o *GoogleCloudBillingAccountListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud billing account list o k response a status code equal to that given
func (o *GoogleCloudBillingAccountListOK) IsCode(code int) bool {
	return code == 200
}

func (o *GoogleCloudBillingAccountListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/billing-accounts][%d] googleCloudBillingAccountListOK  %+v", 200, o.Payload)
}

func (o *GoogleCloudBillingAccountListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/billing-accounts][%d] googleCloudBillingAccountListOK  %+v", 200, o.Payload)
}

func (o *GoogleCloudBillingAccountListOK) GetPayload() []*models.CommonStringBasedDropdownDto {
	return o.Payload
}

func (o *GoogleCloudBillingAccountListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudBillingAccountListBadRequest creates a GoogleCloudBillingAccountListBadRequest with default headers values
func NewGoogleCloudBillingAccountListBadRequest() *GoogleCloudBillingAccountListBadRequest {
	return &GoogleCloudBillingAccountListBadRequest{}
}

/*
GoogleCloudBillingAccountListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GoogleCloudBillingAccountListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this google cloud billing account list bad request response has a 2xx status code
func (o *GoogleCloudBillingAccountListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud billing account list bad request response has a 3xx status code
func (o *GoogleCloudBillingAccountListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud billing account list bad request response has a 4xx status code
func (o *GoogleCloudBillingAccountListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this google cloud billing account list bad request response has a 5xx status code
func (o *GoogleCloudBillingAccountListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud billing account list bad request response a status code equal to that given
func (o *GoogleCloudBillingAccountListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GoogleCloudBillingAccountListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/billing-accounts][%d] googleCloudBillingAccountListBadRequest  %+v", 400, o.Payload)
}

func (o *GoogleCloudBillingAccountListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/billing-accounts][%d] googleCloudBillingAccountListBadRequest  %+v", 400, o.Payload)
}

func (o *GoogleCloudBillingAccountListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *GoogleCloudBillingAccountListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudBillingAccountListUnauthorized creates a GoogleCloudBillingAccountListUnauthorized with default headers values
func NewGoogleCloudBillingAccountListUnauthorized() *GoogleCloudBillingAccountListUnauthorized {
	return &GoogleCloudBillingAccountListUnauthorized{}
}

/*
GoogleCloudBillingAccountListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GoogleCloudBillingAccountListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this google cloud billing account list unauthorized response has a 2xx status code
func (o *GoogleCloudBillingAccountListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud billing account list unauthorized response has a 3xx status code
func (o *GoogleCloudBillingAccountListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud billing account list unauthorized response has a 4xx status code
func (o *GoogleCloudBillingAccountListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this google cloud billing account list unauthorized response has a 5xx status code
func (o *GoogleCloudBillingAccountListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud billing account list unauthorized response a status code equal to that given
func (o *GoogleCloudBillingAccountListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GoogleCloudBillingAccountListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/billing-accounts][%d] googleCloudBillingAccountListUnauthorized  %+v", 401, o.Payload)
}

func (o *GoogleCloudBillingAccountListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/billing-accounts][%d] googleCloudBillingAccountListUnauthorized  %+v", 401, o.Payload)
}

func (o *GoogleCloudBillingAccountListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GoogleCloudBillingAccountListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudBillingAccountListForbidden creates a GoogleCloudBillingAccountListForbidden with default headers values
func NewGoogleCloudBillingAccountListForbidden() *GoogleCloudBillingAccountListForbidden {
	return &GoogleCloudBillingAccountListForbidden{}
}

/*
GoogleCloudBillingAccountListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GoogleCloudBillingAccountListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this google cloud billing account list forbidden response has a 2xx status code
func (o *GoogleCloudBillingAccountListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud billing account list forbidden response has a 3xx status code
func (o *GoogleCloudBillingAccountListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud billing account list forbidden response has a 4xx status code
func (o *GoogleCloudBillingAccountListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this google cloud billing account list forbidden response has a 5xx status code
func (o *GoogleCloudBillingAccountListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud billing account list forbidden response a status code equal to that given
func (o *GoogleCloudBillingAccountListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GoogleCloudBillingAccountListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/billing-accounts][%d] googleCloudBillingAccountListForbidden  %+v", 403, o.Payload)
}

func (o *GoogleCloudBillingAccountListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/billing-accounts][%d] googleCloudBillingAccountListForbidden  %+v", 403, o.Payload)
}

func (o *GoogleCloudBillingAccountListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GoogleCloudBillingAccountListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudBillingAccountListNotFound creates a GoogleCloudBillingAccountListNotFound with default headers values
func NewGoogleCloudBillingAccountListNotFound() *GoogleCloudBillingAccountListNotFound {
	return &GoogleCloudBillingAccountListNotFound{}
}

/*
GoogleCloudBillingAccountListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GoogleCloudBillingAccountListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this google cloud billing account list not found response has a 2xx status code
func (o *GoogleCloudBillingAccountListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud billing account list not found response has a 3xx status code
func (o *GoogleCloudBillingAccountListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud billing account list not found response has a 4xx status code
func (o *GoogleCloudBillingAccountListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this google cloud billing account list not found response has a 5xx status code
func (o *GoogleCloudBillingAccountListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud billing account list not found response a status code equal to that given
func (o *GoogleCloudBillingAccountListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GoogleCloudBillingAccountListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/billing-accounts][%d] googleCloudBillingAccountListNotFound  %+v", 404, o.Payload)
}

func (o *GoogleCloudBillingAccountListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/billing-accounts][%d] googleCloudBillingAccountListNotFound  %+v", 404, o.Payload)
}

func (o *GoogleCloudBillingAccountListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GoogleCloudBillingAccountListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudBillingAccountListInternalServerError creates a GoogleCloudBillingAccountListInternalServerError with default headers values
func NewGoogleCloudBillingAccountListInternalServerError() *GoogleCloudBillingAccountListInternalServerError {
	return &GoogleCloudBillingAccountListInternalServerError{}
}

/*
GoogleCloudBillingAccountListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GoogleCloudBillingAccountListInternalServerError struct {
}

// IsSuccess returns true when this google cloud billing account list internal server error response has a 2xx status code
func (o *GoogleCloudBillingAccountListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud billing account list internal server error response has a 3xx status code
func (o *GoogleCloudBillingAccountListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud billing account list internal server error response has a 4xx status code
func (o *GoogleCloudBillingAccountListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this google cloud billing account list internal server error response has a 5xx status code
func (o *GoogleCloudBillingAccountListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this google cloud billing account list internal server error response a status code equal to that given
func (o *GoogleCloudBillingAccountListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GoogleCloudBillingAccountListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/billing-accounts][%d] googleCloudBillingAccountListInternalServerError ", 500)
}

func (o *GoogleCloudBillingAccountListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/billing-accounts][%d] googleCloudBillingAccountListInternalServerError ", 500)
}

func (o *GoogleCloudBillingAccountListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
