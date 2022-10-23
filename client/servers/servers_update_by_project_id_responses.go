// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ServersUpdateByProjectIDReader is a Reader for the ServersUpdateByProjectID structure.
type ServersUpdateByProjectIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersUpdateByProjectIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServersUpdateByProjectIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServersUpdateByProjectIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServersUpdateByProjectIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServersUpdateByProjectIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServersUpdateByProjectIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServersUpdateByProjectIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServersUpdateByProjectIDOK creates a ServersUpdateByProjectIDOK with default headers values
func NewServersUpdateByProjectIDOK() *ServersUpdateByProjectIDOK {
	return &ServersUpdateByProjectIDOK{}
}

/*
ServersUpdateByProjectIDOK describes a response with status code 200, with default header values.

Success
*/
type ServersUpdateByProjectIDOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this servers update by project Id o k response has a 2xx status code
func (o *ServersUpdateByProjectIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this servers update by project Id o k response has a 3xx status code
func (o *ServersUpdateByProjectIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers update by project Id o k response has a 4xx status code
func (o *ServersUpdateByProjectIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this servers update by project Id o k response has a 5xx status code
func (o *ServersUpdateByProjectIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this servers update by project Id o k response a status code equal to that given
func (o *ServersUpdateByProjectIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServersUpdateByProjectIDOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Servers/update/{projectId}][%d] serversUpdateByProjectIdOK  %+v", 200, o.Payload)
}

func (o *ServersUpdateByProjectIDOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Servers/update/{projectId}][%d] serversUpdateByProjectIdOK  %+v", 200, o.Payload)
}

func (o *ServersUpdateByProjectIDOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ServersUpdateByProjectIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersUpdateByProjectIDBadRequest creates a ServersUpdateByProjectIDBadRequest with default headers values
func NewServersUpdateByProjectIDBadRequest() *ServersUpdateByProjectIDBadRequest {
	return &ServersUpdateByProjectIDBadRequest{}
}

/*
ServersUpdateByProjectIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServersUpdateByProjectIDBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this servers update by project Id bad request response has a 2xx status code
func (o *ServersUpdateByProjectIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers update by project Id bad request response has a 3xx status code
func (o *ServersUpdateByProjectIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers update by project Id bad request response has a 4xx status code
func (o *ServersUpdateByProjectIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers update by project Id bad request response has a 5xx status code
func (o *ServersUpdateByProjectIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this servers update by project Id bad request response a status code equal to that given
func (o *ServersUpdateByProjectIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ServersUpdateByProjectIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Servers/update/{projectId}][%d] serversUpdateByProjectIdBadRequest  %+v", 400, o.Payload)
}

func (o *ServersUpdateByProjectIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Servers/update/{projectId}][%d] serversUpdateByProjectIdBadRequest  %+v", 400, o.Payload)
}

func (o *ServersUpdateByProjectIDBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *ServersUpdateByProjectIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersUpdateByProjectIDUnauthorized creates a ServersUpdateByProjectIDUnauthorized with default headers values
func NewServersUpdateByProjectIDUnauthorized() *ServersUpdateByProjectIDUnauthorized {
	return &ServersUpdateByProjectIDUnauthorized{}
}

/*
ServersUpdateByProjectIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServersUpdateByProjectIDUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this servers update by project Id unauthorized response has a 2xx status code
func (o *ServersUpdateByProjectIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers update by project Id unauthorized response has a 3xx status code
func (o *ServersUpdateByProjectIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers update by project Id unauthorized response has a 4xx status code
func (o *ServersUpdateByProjectIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers update by project Id unauthorized response has a 5xx status code
func (o *ServersUpdateByProjectIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this servers update by project Id unauthorized response a status code equal to that given
func (o *ServersUpdateByProjectIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ServersUpdateByProjectIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Servers/update/{projectId}][%d] serversUpdateByProjectIdUnauthorized  %+v", 401, o.Payload)
}

func (o *ServersUpdateByProjectIDUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Servers/update/{projectId}][%d] serversUpdateByProjectIdUnauthorized  %+v", 401, o.Payload)
}

func (o *ServersUpdateByProjectIDUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ServersUpdateByProjectIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersUpdateByProjectIDForbidden creates a ServersUpdateByProjectIDForbidden with default headers values
func NewServersUpdateByProjectIDForbidden() *ServersUpdateByProjectIDForbidden {
	return &ServersUpdateByProjectIDForbidden{}
}

/*
ServersUpdateByProjectIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServersUpdateByProjectIDForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this servers update by project Id forbidden response has a 2xx status code
func (o *ServersUpdateByProjectIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers update by project Id forbidden response has a 3xx status code
func (o *ServersUpdateByProjectIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers update by project Id forbidden response has a 4xx status code
func (o *ServersUpdateByProjectIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers update by project Id forbidden response has a 5xx status code
func (o *ServersUpdateByProjectIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this servers update by project Id forbidden response a status code equal to that given
func (o *ServersUpdateByProjectIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ServersUpdateByProjectIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Servers/update/{projectId}][%d] serversUpdateByProjectIdForbidden  %+v", 403, o.Payload)
}

func (o *ServersUpdateByProjectIDForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Servers/update/{projectId}][%d] serversUpdateByProjectIdForbidden  %+v", 403, o.Payload)
}

func (o *ServersUpdateByProjectIDForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ServersUpdateByProjectIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersUpdateByProjectIDNotFound creates a ServersUpdateByProjectIDNotFound with default headers values
func NewServersUpdateByProjectIDNotFound() *ServersUpdateByProjectIDNotFound {
	return &ServersUpdateByProjectIDNotFound{}
}

/*
ServersUpdateByProjectIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServersUpdateByProjectIDNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this servers update by project Id not found response has a 2xx status code
func (o *ServersUpdateByProjectIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers update by project Id not found response has a 3xx status code
func (o *ServersUpdateByProjectIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers update by project Id not found response has a 4xx status code
func (o *ServersUpdateByProjectIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers update by project Id not found response has a 5xx status code
func (o *ServersUpdateByProjectIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this servers update by project Id not found response a status code equal to that given
func (o *ServersUpdateByProjectIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ServersUpdateByProjectIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Servers/update/{projectId}][%d] serversUpdateByProjectIdNotFound  %+v", 404, o.Payload)
}

func (o *ServersUpdateByProjectIDNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Servers/update/{projectId}][%d] serversUpdateByProjectIdNotFound  %+v", 404, o.Payload)
}

func (o *ServersUpdateByProjectIDNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ServersUpdateByProjectIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersUpdateByProjectIDInternalServerError creates a ServersUpdateByProjectIDInternalServerError with default headers values
func NewServersUpdateByProjectIDInternalServerError() *ServersUpdateByProjectIDInternalServerError {
	return &ServersUpdateByProjectIDInternalServerError{}
}

/*
ServersUpdateByProjectIDInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ServersUpdateByProjectIDInternalServerError struct {
}

// IsSuccess returns true when this servers update by project Id internal server error response has a 2xx status code
func (o *ServersUpdateByProjectIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers update by project Id internal server error response has a 3xx status code
func (o *ServersUpdateByProjectIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers update by project Id internal server error response has a 4xx status code
func (o *ServersUpdateByProjectIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this servers update by project Id internal server error response has a 5xx status code
func (o *ServersUpdateByProjectIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this servers update by project Id internal server error response a status code equal to that given
func (o *ServersUpdateByProjectIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ServersUpdateByProjectIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Servers/update/{projectId}][%d] serversUpdateByProjectIdInternalServerError ", 500)
}

func (o *ServersUpdateByProjectIDInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Servers/update/{projectId}][%d] serversUpdateByProjectIdInternalServerError ", 500)
}

func (o *ServersUpdateByProjectIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
