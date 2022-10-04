// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// CommonGetCountryListReader is a Reader for the CommonGetCountryList structure.
type CommonGetCountryListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommonGetCountryListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommonGetCountryListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCommonGetCountryListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCommonGetCountryListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCommonGetCountryListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCommonGetCountryListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCommonGetCountryListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCommonGetCountryListOK creates a CommonGetCountryListOK with default headers values
func NewCommonGetCountryListOK() *CommonGetCountryListOK {
	return &CommonGetCountryListOK{}
}

/*
CommonGetCountryListOK describes a response with status code 200, with default header values.

Success
*/
type CommonGetCountryListOK struct {
	Payload []*models.CountryListDto
}

// IsSuccess returns true when this common get country list o k response has a 2xx status code
func (o *CommonGetCountryListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this common get country list o k response has a 3xx status code
func (o *CommonGetCountryListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this common get country list o k response has a 4xx status code
func (o *CommonGetCountryListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this common get country list o k response has a 5xx status code
func (o *CommonGetCountryListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this common get country list o k response a status code equal to that given
func (o *CommonGetCountryListOK) IsCode(code int) bool {
	return code == 200
}

func (o *CommonGetCountryListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/countries][%d] commonGetCountryListOK  %+v", 200, o.Payload)
}

func (o *CommonGetCountryListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/countries][%d] commonGetCountryListOK  %+v", 200, o.Payload)
}

func (o *CommonGetCountryListOK) GetPayload() []*models.CountryListDto {
	return o.Payload
}

func (o *CommonGetCountryListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommonGetCountryListBadRequest creates a CommonGetCountryListBadRequest with default headers values
func NewCommonGetCountryListBadRequest() *CommonGetCountryListBadRequest {
	return &CommonGetCountryListBadRequest{}
}

/*
CommonGetCountryListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CommonGetCountryListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this common get country list bad request response has a 2xx status code
func (o *CommonGetCountryListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this common get country list bad request response has a 3xx status code
func (o *CommonGetCountryListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this common get country list bad request response has a 4xx status code
func (o *CommonGetCountryListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this common get country list bad request response has a 5xx status code
func (o *CommonGetCountryListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this common get country list bad request response a status code equal to that given
func (o *CommonGetCountryListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CommonGetCountryListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/countries][%d] commonGetCountryListBadRequest  %+v", 400, o.Payload)
}

func (o *CommonGetCountryListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/countries][%d] commonGetCountryListBadRequest  %+v", 400, o.Payload)
}

func (o *CommonGetCountryListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CommonGetCountryListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommonGetCountryListUnauthorized creates a CommonGetCountryListUnauthorized with default headers values
func NewCommonGetCountryListUnauthorized() *CommonGetCountryListUnauthorized {
	return &CommonGetCountryListUnauthorized{}
}

/*
CommonGetCountryListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CommonGetCountryListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this common get country list unauthorized response has a 2xx status code
func (o *CommonGetCountryListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this common get country list unauthorized response has a 3xx status code
func (o *CommonGetCountryListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this common get country list unauthorized response has a 4xx status code
func (o *CommonGetCountryListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this common get country list unauthorized response has a 5xx status code
func (o *CommonGetCountryListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this common get country list unauthorized response a status code equal to that given
func (o *CommonGetCountryListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CommonGetCountryListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/countries][%d] commonGetCountryListUnauthorized  %+v", 401, o.Payload)
}

func (o *CommonGetCountryListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/countries][%d] commonGetCountryListUnauthorized  %+v", 401, o.Payload)
}

func (o *CommonGetCountryListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CommonGetCountryListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommonGetCountryListForbidden creates a CommonGetCountryListForbidden with default headers values
func NewCommonGetCountryListForbidden() *CommonGetCountryListForbidden {
	return &CommonGetCountryListForbidden{}
}

/*
CommonGetCountryListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CommonGetCountryListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this common get country list forbidden response has a 2xx status code
func (o *CommonGetCountryListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this common get country list forbidden response has a 3xx status code
func (o *CommonGetCountryListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this common get country list forbidden response has a 4xx status code
func (o *CommonGetCountryListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this common get country list forbidden response has a 5xx status code
func (o *CommonGetCountryListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this common get country list forbidden response a status code equal to that given
func (o *CommonGetCountryListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CommonGetCountryListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/countries][%d] commonGetCountryListForbidden  %+v", 403, o.Payload)
}

func (o *CommonGetCountryListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/countries][%d] commonGetCountryListForbidden  %+v", 403, o.Payload)
}

func (o *CommonGetCountryListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CommonGetCountryListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommonGetCountryListNotFound creates a CommonGetCountryListNotFound with default headers values
func NewCommonGetCountryListNotFound() *CommonGetCountryListNotFound {
	return &CommonGetCountryListNotFound{}
}

/*
CommonGetCountryListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CommonGetCountryListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this common get country list not found response has a 2xx status code
func (o *CommonGetCountryListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this common get country list not found response has a 3xx status code
func (o *CommonGetCountryListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this common get country list not found response has a 4xx status code
func (o *CommonGetCountryListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this common get country list not found response has a 5xx status code
func (o *CommonGetCountryListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this common get country list not found response a status code equal to that given
func (o *CommonGetCountryListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CommonGetCountryListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/countries][%d] commonGetCountryListNotFound  %+v", 404, o.Payload)
}

func (o *CommonGetCountryListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/countries][%d] commonGetCountryListNotFound  %+v", 404, o.Payload)
}

func (o *CommonGetCountryListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CommonGetCountryListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommonGetCountryListInternalServerError creates a CommonGetCountryListInternalServerError with default headers values
func NewCommonGetCountryListInternalServerError() *CommonGetCountryListInternalServerError {
	return &CommonGetCountryListInternalServerError{}
}

/*
CommonGetCountryListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CommonGetCountryListInternalServerError struct {
}

// IsSuccess returns true when this common get country list internal server error response has a 2xx status code
func (o *CommonGetCountryListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this common get country list internal server error response has a 3xx status code
func (o *CommonGetCountryListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this common get country list internal server error response has a 4xx status code
func (o *CommonGetCountryListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this common get country list internal server error response has a 5xx status code
func (o *CommonGetCountryListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this common get country list internal server error response a status code equal to that given
func (o *CommonGetCountryListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CommonGetCountryListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/countries][%d] commonGetCountryListInternalServerError ", 500)
}

func (o *CommonGetCountryListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/countries][%d] commonGetCountryListInternalServerError ", 500)
}

func (o *CommonGetCountryListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
