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

// ServersDetailsReader is a Reader for the ServersDetails structure.
type ServersDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServersDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServersDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServersDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServersDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServersDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServersDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServersDetailsOK creates a ServersDetailsOK with default headers values
func NewServersDetailsOK() *ServersDetailsOK {
	return &ServersDetailsOK{}
}

/* ServersDetailsOK describes a response with status code 200, with default header values.

Success
*/
type ServersDetailsOK struct {
	Payload *models.ServersListForDetails
}

func (o *ServersDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers/{projectId}][%d] serversDetailsOK  %+v", 200, o.Payload)
}
func (o *ServersDetailsOK) GetPayload() *models.ServersListForDetails {
	return o.Payload
}

func (o *ServersDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServersListForDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersDetailsBadRequest creates a ServersDetailsBadRequest with default headers values
func NewServersDetailsBadRequest() *ServersDetailsBadRequest {
	return &ServersDetailsBadRequest{}
}

/* ServersDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServersDetailsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ServersDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers/{projectId}][%d] serversDetailsBadRequest  %+v", 400, o.Payload)
}
func (o *ServersDetailsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ServersDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersDetailsUnauthorized creates a ServersDetailsUnauthorized with default headers values
func NewServersDetailsUnauthorized() *ServersDetailsUnauthorized {
	return &ServersDetailsUnauthorized{}
}

/* ServersDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServersDetailsUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ServersDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers/{projectId}][%d] serversDetailsUnauthorized  %+v", 401, o.Payload)
}
func (o *ServersDetailsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ServersDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersDetailsForbidden creates a ServersDetailsForbidden with default headers values
func NewServersDetailsForbidden() *ServersDetailsForbidden {
	return &ServersDetailsForbidden{}
}

/* ServersDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServersDetailsForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ServersDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers/{projectId}][%d] serversDetailsForbidden  %+v", 403, o.Payload)
}
func (o *ServersDetailsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ServersDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersDetailsNotFound creates a ServersDetailsNotFound with default headers values
func NewServersDetailsNotFound() *ServersDetailsNotFound {
	return &ServersDetailsNotFound{}
}

/* ServersDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServersDetailsNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ServersDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers/{projectId}][%d] serversDetailsNotFound  %+v", 404, o.Payload)
}
func (o *ServersDetailsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ServersDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersDetailsInternalServerError creates a ServersDetailsInternalServerError with default headers values
func NewServersDetailsInternalServerError() *ServersDetailsInternalServerError {
	return &ServersDetailsInternalServerError{}
}

/* ServersDetailsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ServersDetailsInternalServerError struct {
}

func (o *ServersDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers/{projectId}][%d] serversDetailsInternalServerError ", 500)
}

func (o *ServersDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
