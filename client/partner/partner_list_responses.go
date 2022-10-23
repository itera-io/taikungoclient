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

// PartnerListReader is a Reader for the PartnerList structure.
type PartnerListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPartnerListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPartnerListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPartnerListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPartnerListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPartnerListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPartnerListOK creates a PartnerListOK with default headers values
func NewPartnerListOK() *PartnerListOK {
	return &PartnerListOK{}
}

/*
PartnerListOK describes a response with status code 200, with default header values.

Success
*/
type PartnerListOK struct {
	Payload *models.PartnersList
}

// IsSuccess returns true when this partner list o k response has a 2xx status code
func (o *PartnerListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner list o k response has a 3xx status code
func (o *PartnerListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner list o k response has a 4xx status code
func (o *PartnerListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner list o k response has a 5xx status code
func (o *PartnerListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner list o k response a status code equal to that given
func (o *PartnerListOK) IsCode(code int) bool {
	return code == 200
}

func (o *PartnerListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner][%d] partnerListOK  %+v", 200, o.Payload)
}

func (o *PartnerListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner][%d] partnerListOK  %+v", 200, o.Payload)
}

func (o *PartnerListOK) GetPayload() *models.PartnersList {
	return o.Payload
}

func (o *PartnerListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PartnersList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerListBadRequest creates a PartnerListBadRequest with default headers values
func NewPartnerListBadRequest() *PartnerListBadRequest {
	return &PartnerListBadRequest{}
}

/*
PartnerListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PartnerListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this partner list bad request response has a 2xx status code
func (o *PartnerListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner list bad request response has a 3xx status code
func (o *PartnerListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner list bad request response has a 4xx status code
func (o *PartnerListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner list bad request response has a 5xx status code
func (o *PartnerListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this partner list bad request response a status code equal to that given
func (o *PartnerListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PartnerListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner][%d] partnerListBadRequest  %+v", 400, o.Payload)
}

func (o *PartnerListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner][%d] partnerListBadRequest  %+v", 400, o.Payload)
}

func (o *PartnerListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *PartnerListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerListUnauthorized creates a PartnerListUnauthorized with default headers values
func NewPartnerListUnauthorized() *PartnerListUnauthorized {
	return &PartnerListUnauthorized{}
}

/*
PartnerListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PartnerListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this partner list unauthorized response has a 2xx status code
func (o *PartnerListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner list unauthorized response has a 3xx status code
func (o *PartnerListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner list unauthorized response has a 4xx status code
func (o *PartnerListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner list unauthorized response has a 5xx status code
func (o *PartnerListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this partner list unauthorized response a status code equal to that given
func (o *PartnerListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PartnerListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner][%d] partnerListUnauthorized  %+v", 401, o.Payload)
}

func (o *PartnerListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner][%d] partnerListUnauthorized  %+v", 401, o.Payload)
}

func (o *PartnerListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PartnerListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerListForbidden creates a PartnerListForbidden with default headers values
func NewPartnerListForbidden() *PartnerListForbidden {
	return &PartnerListForbidden{}
}

/*
PartnerListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PartnerListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this partner list forbidden response has a 2xx status code
func (o *PartnerListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner list forbidden response has a 3xx status code
func (o *PartnerListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner list forbidden response has a 4xx status code
func (o *PartnerListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner list forbidden response has a 5xx status code
func (o *PartnerListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this partner list forbidden response a status code equal to that given
func (o *PartnerListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PartnerListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner][%d] partnerListForbidden  %+v", 403, o.Payload)
}

func (o *PartnerListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner][%d] partnerListForbidden  %+v", 403, o.Payload)
}

func (o *PartnerListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PartnerListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerListNotFound creates a PartnerListNotFound with default headers values
func NewPartnerListNotFound() *PartnerListNotFound {
	return &PartnerListNotFound{}
}

/*
PartnerListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PartnerListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this partner list not found response has a 2xx status code
func (o *PartnerListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner list not found response has a 3xx status code
func (o *PartnerListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner list not found response has a 4xx status code
func (o *PartnerListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner list not found response has a 5xx status code
func (o *PartnerListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this partner list not found response a status code equal to that given
func (o *PartnerListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PartnerListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner][%d] partnerListNotFound  %+v", 404, o.Payload)
}

func (o *PartnerListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner][%d] partnerListNotFound  %+v", 404, o.Payload)
}

func (o *PartnerListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PartnerListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerListInternalServerError creates a PartnerListInternalServerError with default headers values
func NewPartnerListInternalServerError() *PartnerListInternalServerError {
	return &PartnerListInternalServerError{}
}

/*
PartnerListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PartnerListInternalServerError struct {
}

// IsSuccess returns true when this partner list internal server error response has a 2xx status code
func (o *PartnerListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner list internal server error response has a 3xx status code
func (o *PartnerListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner list internal server error response has a 4xx status code
func (o *PartnerListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner list internal server error response has a 5xx status code
func (o *PartnerListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this partner list internal server error response a status code equal to that given
func (o *PartnerListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PartnerListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner][%d] partnerListInternalServerError ", 500)
}

func (o *PartnerListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner][%d] partnerListInternalServerError ", 500)
}

func (o *PartnerListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
