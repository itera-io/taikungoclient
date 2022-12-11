// Code generated by go-swagger; DO NOT EDIT.

package partner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// PartnerDetailsReader is a Reader for the PartnerDetails structure.
type PartnerDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPartnerDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPartnerDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPartnerDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPartnerDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPartnerDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPartnerDetailsOK creates a PartnerDetailsOK with default headers values
func NewPartnerDetailsOK() *PartnerDetailsOK {
	return &PartnerDetailsOK{}
}

/*
PartnerDetailsOK describes a response with status code 200, with default header values.

Success
*/
type PartnerDetailsOK struct {
	Payload *models.PartnerDetailsDto
}

// IsSuccess returns true when this partner details o k response has a 2xx status code
func (o *PartnerDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner details o k response has a 3xx status code
func (o *PartnerDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner details o k response has a 4xx status code
func (o *PartnerDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner details o k response has a 5xx status code
func (o *PartnerDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner details o k response a status code equal to that given
func (o *PartnerDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PartnerDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/details][%d] partnerDetailsOK  %+v", 200, o.Payload)
}

func (o *PartnerDetailsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/details][%d] partnerDetailsOK  %+v", 200, o.Payload)
}

func (o *PartnerDetailsOK) GetPayload() *models.PartnerDetailsDto {
	return o.Payload
}

func (o *PartnerDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PartnerDetailsDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerDetailsBadRequest creates a PartnerDetailsBadRequest with default headers values
func NewPartnerDetailsBadRequest() *PartnerDetailsBadRequest {
	return &PartnerDetailsBadRequest{}
}

/*
PartnerDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PartnerDetailsBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this partner details bad request response has a 2xx status code
func (o *PartnerDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner details bad request response has a 3xx status code
func (o *PartnerDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner details bad request response has a 4xx status code
func (o *PartnerDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner details bad request response has a 5xx status code
func (o *PartnerDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this partner details bad request response a status code equal to that given
func (o *PartnerDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PartnerDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/details][%d] partnerDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *PartnerDetailsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/details][%d] partnerDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *PartnerDetailsBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerDetailsUnauthorized creates a PartnerDetailsUnauthorized with default headers values
func NewPartnerDetailsUnauthorized() *PartnerDetailsUnauthorized {
	return &PartnerDetailsUnauthorized{}
}

/*
PartnerDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PartnerDetailsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this partner details unauthorized response has a 2xx status code
func (o *PartnerDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner details unauthorized response has a 3xx status code
func (o *PartnerDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner details unauthorized response has a 4xx status code
func (o *PartnerDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner details unauthorized response has a 5xx status code
func (o *PartnerDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this partner details unauthorized response a status code equal to that given
func (o *PartnerDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PartnerDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/details][%d] partnerDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *PartnerDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/details][%d] partnerDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *PartnerDetailsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PartnerDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerDetailsForbidden creates a PartnerDetailsForbidden with default headers values
func NewPartnerDetailsForbidden() *PartnerDetailsForbidden {
	return &PartnerDetailsForbidden{}
}

/*
PartnerDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PartnerDetailsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this partner details forbidden response has a 2xx status code
func (o *PartnerDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner details forbidden response has a 3xx status code
func (o *PartnerDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner details forbidden response has a 4xx status code
func (o *PartnerDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner details forbidden response has a 5xx status code
func (o *PartnerDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this partner details forbidden response a status code equal to that given
func (o *PartnerDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PartnerDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/details][%d] partnerDetailsForbidden  %+v", 403, o.Payload)
}

func (o *PartnerDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/details][%d] partnerDetailsForbidden  %+v", 403, o.Payload)
}

func (o *PartnerDetailsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PartnerDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerDetailsNotFound creates a PartnerDetailsNotFound with default headers values
func NewPartnerDetailsNotFound() *PartnerDetailsNotFound {
	return &PartnerDetailsNotFound{}
}

/*
PartnerDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PartnerDetailsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this partner details not found response has a 2xx status code
func (o *PartnerDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner details not found response has a 3xx status code
func (o *PartnerDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner details not found response has a 4xx status code
func (o *PartnerDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner details not found response has a 5xx status code
func (o *PartnerDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this partner details not found response a status code equal to that given
func (o *PartnerDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PartnerDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/details][%d] partnerDetailsNotFound  %+v", 404, o.Payload)
}

func (o *PartnerDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/details][%d] partnerDetailsNotFound  %+v", 404, o.Payload)
}

func (o *PartnerDetailsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PartnerDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerDetailsInternalServerError creates a PartnerDetailsInternalServerError with default headers values
func NewPartnerDetailsInternalServerError() *PartnerDetailsInternalServerError {
	return &PartnerDetailsInternalServerError{}
}

/*
PartnerDetailsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PartnerDetailsInternalServerError struct {
}

// IsSuccess returns true when this partner details internal server error response has a 2xx status code
func (o *PartnerDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner details internal server error response has a 3xx status code
func (o *PartnerDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner details internal server error response has a 4xx status code
func (o *PartnerDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner details internal server error response has a 5xx status code
func (o *PartnerDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this partner details internal server error response a status code equal to that given
func (o *PartnerDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PartnerDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/details][%d] partnerDetailsInternalServerError ", 500)
}

func (o *PartnerDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/details][%d] partnerDetailsInternalServerError ", 500)
}

func (o *PartnerDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
