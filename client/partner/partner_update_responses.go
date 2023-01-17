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

// PartnerUpdateReader is a Reader for the PartnerUpdate structure.
type PartnerUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPartnerUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPartnerUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPartnerUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPartnerUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPartnerUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPartnerUpdateOK creates a PartnerUpdateOK with default headers values
func NewPartnerUpdateOK() *PartnerUpdateOK {
	return &PartnerUpdateOK{}
}

/*
PartnerUpdateOK describes a response with status code 200, with default header values.

Success
*/
type PartnerUpdateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this partner update o k response has a 2xx status code
func (o *PartnerUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner update o k response has a 3xx status code
func (o *PartnerUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner update o k response has a 4xx status code
func (o *PartnerUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner update o k response has a 5xx status code
func (o *PartnerUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner update o k response a status code equal to that given
func (o *PartnerUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner update o k response
func (o *PartnerUpdateOK) Code() int {
	return 200
}

func (o *PartnerUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Partner/update/{id}][%d] partnerUpdateOK  %+v", 200, o.Payload)
}

func (o *PartnerUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Partner/update/{id}][%d] partnerUpdateOK  %+v", 200, o.Payload)
}

func (o *PartnerUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *PartnerUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerUpdateBadRequest creates a PartnerUpdateBadRequest with default headers values
func NewPartnerUpdateBadRequest() *PartnerUpdateBadRequest {
	return &PartnerUpdateBadRequest{}
}

/*
PartnerUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PartnerUpdateBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this partner update bad request response has a 2xx status code
func (o *PartnerUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner update bad request response has a 3xx status code
func (o *PartnerUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner update bad request response has a 4xx status code
func (o *PartnerUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner update bad request response has a 5xx status code
func (o *PartnerUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this partner update bad request response a status code equal to that given
func (o *PartnerUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the partner update bad request response
func (o *PartnerUpdateBadRequest) Code() int {
	return 400
}

func (o *PartnerUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Partner/update/{id}][%d] partnerUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *PartnerUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Partner/update/{id}][%d] partnerUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *PartnerUpdateBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PartnerUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerUpdateUnauthorized creates a PartnerUpdateUnauthorized with default headers values
func NewPartnerUpdateUnauthorized() *PartnerUpdateUnauthorized {
	return &PartnerUpdateUnauthorized{}
}

/*
PartnerUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PartnerUpdateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this partner update unauthorized response has a 2xx status code
func (o *PartnerUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner update unauthorized response has a 3xx status code
func (o *PartnerUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner update unauthorized response has a 4xx status code
func (o *PartnerUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner update unauthorized response has a 5xx status code
func (o *PartnerUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this partner update unauthorized response a status code equal to that given
func (o *PartnerUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the partner update unauthorized response
func (o *PartnerUpdateUnauthorized) Code() int {
	return 401
}

func (o *PartnerUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Partner/update/{id}][%d] partnerUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *PartnerUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Partner/update/{id}][%d] partnerUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *PartnerUpdateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PartnerUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerUpdateForbidden creates a PartnerUpdateForbidden with default headers values
func NewPartnerUpdateForbidden() *PartnerUpdateForbidden {
	return &PartnerUpdateForbidden{}
}

/*
PartnerUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PartnerUpdateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this partner update forbidden response has a 2xx status code
func (o *PartnerUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner update forbidden response has a 3xx status code
func (o *PartnerUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner update forbidden response has a 4xx status code
func (o *PartnerUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner update forbidden response has a 5xx status code
func (o *PartnerUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this partner update forbidden response a status code equal to that given
func (o *PartnerUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the partner update forbidden response
func (o *PartnerUpdateForbidden) Code() int {
	return 403
}

func (o *PartnerUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Partner/update/{id}][%d] partnerUpdateForbidden  %+v", 403, o.Payload)
}

func (o *PartnerUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Partner/update/{id}][%d] partnerUpdateForbidden  %+v", 403, o.Payload)
}

func (o *PartnerUpdateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PartnerUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerUpdateNotFound creates a PartnerUpdateNotFound with default headers values
func NewPartnerUpdateNotFound() *PartnerUpdateNotFound {
	return &PartnerUpdateNotFound{}
}

/*
PartnerUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PartnerUpdateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this partner update not found response has a 2xx status code
func (o *PartnerUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner update not found response has a 3xx status code
func (o *PartnerUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner update not found response has a 4xx status code
func (o *PartnerUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner update not found response has a 5xx status code
func (o *PartnerUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this partner update not found response a status code equal to that given
func (o *PartnerUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the partner update not found response
func (o *PartnerUpdateNotFound) Code() int {
	return 404
}

func (o *PartnerUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Partner/update/{id}][%d] partnerUpdateNotFound  %+v", 404, o.Payload)
}

func (o *PartnerUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Partner/update/{id}][%d] partnerUpdateNotFound  %+v", 404, o.Payload)
}

func (o *PartnerUpdateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PartnerUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerUpdateInternalServerError creates a PartnerUpdateInternalServerError with default headers values
func NewPartnerUpdateInternalServerError() *PartnerUpdateInternalServerError {
	return &PartnerUpdateInternalServerError{}
}

/*
PartnerUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PartnerUpdateInternalServerError struct {
}

// IsSuccess returns true when this partner update internal server error response has a 2xx status code
func (o *PartnerUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner update internal server error response has a 3xx status code
func (o *PartnerUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner update internal server error response has a 4xx status code
func (o *PartnerUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner update internal server error response has a 5xx status code
func (o *PartnerUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this partner update internal server error response a status code equal to that given
func (o *PartnerUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the partner update internal server error response
func (o *PartnerUpdateInternalServerError) Code() int {
	return 500
}

func (o *PartnerUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Partner/update/{id}][%d] partnerUpdateInternalServerError ", 500)
}

func (o *PartnerUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Partner/update/{id}][%d] partnerUpdateInternalServerError ", 500)
}

func (o *PartnerUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
