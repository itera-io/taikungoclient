// Code generated by go-swagger; DO NOT EDIT.

package checker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// CheckerOpenstackTaikunLbImageReader is a Reader for the CheckerOpenstackTaikunLbImage structure.
type CheckerOpenstackTaikunLbImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerOpenstackTaikunLbImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerOpenstackTaikunLbImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerOpenstackTaikunLbImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerOpenstackTaikunLbImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerOpenstackTaikunLbImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerOpenstackTaikunLbImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerOpenstackTaikunLbImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerOpenstackTaikunLbImageOK creates a CheckerOpenstackTaikunLbImageOK with default headers values
func NewCheckerOpenstackTaikunLbImageOK() *CheckerOpenstackTaikunLbImageOK {
	return &CheckerOpenstackTaikunLbImageOK{}
}

/*
CheckerOpenstackTaikunLbImageOK describes a response with status code 200, with default header values.

Success
*/
type CheckerOpenstackTaikunLbImageOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this checker openstack taikun lb image o k response has a 2xx status code
func (o *CheckerOpenstackTaikunLbImageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker openstack taikun lb image o k response has a 3xx status code
func (o *CheckerOpenstackTaikunLbImageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker openstack taikun lb image o k response has a 4xx status code
func (o *CheckerOpenstackTaikunLbImageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker openstack taikun lb image o k response has a 5xx status code
func (o *CheckerOpenstackTaikunLbImageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker openstack taikun lb image o k response a status code equal to that given
func (o *CheckerOpenstackTaikunLbImageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the checker openstack taikun lb image o k response
func (o *CheckerOpenstackTaikunLbImageOK) Code() int {
	return 200
}

func (o *CheckerOpenstackTaikunLbImageOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageOK  %+v", 200, o.Payload)
}

func (o *CheckerOpenstackTaikunLbImageOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageOK  %+v", 200, o.Payload)
}

func (o *CheckerOpenstackTaikunLbImageOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CheckerOpenstackTaikunLbImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackTaikunLbImageBadRequest creates a CheckerOpenstackTaikunLbImageBadRequest with default headers values
func NewCheckerOpenstackTaikunLbImageBadRequest() *CheckerOpenstackTaikunLbImageBadRequest {
	return &CheckerOpenstackTaikunLbImageBadRequest{}
}

/*
CheckerOpenstackTaikunLbImageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerOpenstackTaikunLbImageBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this checker openstack taikun lb image bad request response has a 2xx status code
func (o *CheckerOpenstackTaikunLbImageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker openstack taikun lb image bad request response has a 3xx status code
func (o *CheckerOpenstackTaikunLbImageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker openstack taikun lb image bad request response has a 4xx status code
func (o *CheckerOpenstackTaikunLbImageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker openstack taikun lb image bad request response has a 5xx status code
func (o *CheckerOpenstackTaikunLbImageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker openstack taikun lb image bad request response a status code equal to that given
func (o *CheckerOpenstackTaikunLbImageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the checker openstack taikun lb image bad request response
func (o *CheckerOpenstackTaikunLbImageBadRequest) Code() int {
	return 400
}

func (o *CheckerOpenstackTaikunLbImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerOpenstackTaikunLbImageBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerOpenstackTaikunLbImageBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerOpenstackTaikunLbImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackTaikunLbImageUnauthorized creates a CheckerOpenstackTaikunLbImageUnauthorized with default headers values
func NewCheckerOpenstackTaikunLbImageUnauthorized() *CheckerOpenstackTaikunLbImageUnauthorized {
	return &CheckerOpenstackTaikunLbImageUnauthorized{}
}

/*
CheckerOpenstackTaikunLbImageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerOpenstackTaikunLbImageUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this checker openstack taikun lb image unauthorized response has a 2xx status code
func (o *CheckerOpenstackTaikunLbImageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker openstack taikun lb image unauthorized response has a 3xx status code
func (o *CheckerOpenstackTaikunLbImageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker openstack taikun lb image unauthorized response has a 4xx status code
func (o *CheckerOpenstackTaikunLbImageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker openstack taikun lb image unauthorized response has a 5xx status code
func (o *CheckerOpenstackTaikunLbImageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker openstack taikun lb image unauthorized response a status code equal to that given
func (o *CheckerOpenstackTaikunLbImageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the checker openstack taikun lb image unauthorized response
func (o *CheckerOpenstackTaikunLbImageUnauthorized) Code() int {
	return 401
}

func (o *CheckerOpenstackTaikunLbImageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerOpenstackTaikunLbImageUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerOpenstackTaikunLbImageUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerOpenstackTaikunLbImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackTaikunLbImageForbidden creates a CheckerOpenstackTaikunLbImageForbidden with default headers values
func NewCheckerOpenstackTaikunLbImageForbidden() *CheckerOpenstackTaikunLbImageForbidden {
	return &CheckerOpenstackTaikunLbImageForbidden{}
}

/*
CheckerOpenstackTaikunLbImageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerOpenstackTaikunLbImageForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this checker openstack taikun lb image forbidden response has a 2xx status code
func (o *CheckerOpenstackTaikunLbImageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker openstack taikun lb image forbidden response has a 3xx status code
func (o *CheckerOpenstackTaikunLbImageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker openstack taikun lb image forbidden response has a 4xx status code
func (o *CheckerOpenstackTaikunLbImageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker openstack taikun lb image forbidden response has a 5xx status code
func (o *CheckerOpenstackTaikunLbImageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker openstack taikun lb image forbidden response a status code equal to that given
func (o *CheckerOpenstackTaikunLbImageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the checker openstack taikun lb image forbidden response
func (o *CheckerOpenstackTaikunLbImageForbidden) Code() int {
	return 403
}

func (o *CheckerOpenstackTaikunLbImageForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageForbidden  %+v", 403, o.Payload)
}

func (o *CheckerOpenstackTaikunLbImageForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageForbidden  %+v", 403, o.Payload)
}

func (o *CheckerOpenstackTaikunLbImageForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerOpenstackTaikunLbImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackTaikunLbImageNotFound creates a CheckerOpenstackTaikunLbImageNotFound with default headers values
func NewCheckerOpenstackTaikunLbImageNotFound() *CheckerOpenstackTaikunLbImageNotFound {
	return &CheckerOpenstackTaikunLbImageNotFound{}
}

/*
CheckerOpenstackTaikunLbImageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerOpenstackTaikunLbImageNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this checker openstack taikun lb image not found response has a 2xx status code
func (o *CheckerOpenstackTaikunLbImageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker openstack taikun lb image not found response has a 3xx status code
func (o *CheckerOpenstackTaikunLbImageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker openstack taikun lb image not found response has a 4xx status code
func (o *CheckerOpenstackTaikunLbImageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker openstack taikun lb image not found response has a 5xx status code
func (o *CheckerOpenstackTaikunLbImageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker openstack taikun lb image not found response a status code equal to that given
func (o *CheckerOpenstackTaikunLbImageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the checker openstack taikun lb image not found response
func (o *CheckerOpenstackTaikunLbImageNotFound) Code() int {
	return 404
}

func (o *CheckerOpenstackTaikunLbImageNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageNotFound  %+v", 404, o.Payload)
}

func (o *CheckerOpenstackTaikunLbImageNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageNotFound  %+v", 404, o.Payload)
}

func (o *CheckerOpenstackTaikunLbImageNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerOpenstackTaikunLbImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackTaikunLbImageInternalServerError creates a CheckerOpenstackTaikunLbImageInternalServerError with default headers values
func NewCheckerOpenstackTaikunLbImageInternalServerError() *CheckerOpenstackTaikunLbImageInternalServerError {
	return &CheckerOpenstackTaikunLbImageInternalServerError{}
}

/*
CheckerOpenstackTaikunLbImageInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerOpenstackTaikunLbImageInternalServerError struct {
}

// IsSuccess returns true when this checker openstack taikun lb image internal server error response has a 2xx status code
func (o *CheckerOpenstackTaikunLbImageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker openstack taikun lb image internal server error response has a 3xx status code
func (o *CheckerOpenstackTaikunLbImageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker openstack taikun lb image internal server error response has a 4xx status code
func (o *CheckerOpenstackTaikunLbImageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker openstack taikun lb image internal server error response has a 5xx status code
func (o *CheckerOpenstackTaikunLbImageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker openstack taikun lb image internal server error response a status code equal to that given
func (o *CheckerOpenstackTaikunLbImageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the checker openstack taikun lb image internal server error response
func (o *CheckerOpenstackTaikunLbImageInternalServerError) Code() int {
	return 500
}

func (o *CheckerOpenstackTaikunLbImageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageInternalServerError ", 500)
}

func (o *CheckerOpenstackTaikunLbImageInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageInternalServerError ", 500)
}

func (o *CheckerOpenstackTaikunLbImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
