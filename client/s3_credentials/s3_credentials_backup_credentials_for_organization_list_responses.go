// Code generated by go-swagger; DO NOT EDIT.

package s3_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// S3CredentialsBackupCredentialsForOrganizationListReader is a Reader for the S3CredentialsBackupCredentialsForOrganizationList structure.
type S3CredentialsBackupCredentialsForOrganizationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3CredentialsBackupCredentialsForOrganizationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3CredentialsBackupCredentialsForOrganizationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewS3CredentialsBackupCredentialsForOrganizationListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewS3CredentialsBackupCredentialsForOrganizationListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewS3CredentialsBackupCredentialsForOrganizationListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewS3CredentialsBackupCredentialsForOrganizationListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewS3CredentialsBackupCredentialsForOrganizationListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewS3CredentialsBackupCredentialsForOrganizationListOK creates a S3CredentialsBackupCredentialsForOrganizationListOK with default headers values
func NewS3CredentialsBackupCredentialsForOrganizationListOK() *S3CredentialsBackupCredentialsForOrganizationListOK {
	return &S3CredentialsBackupCredentialsForOrganizationListOK{}
}

/*
S3CredentialsBackupCredentialsForOrganizationListOK describes a response with status code 200, with default header values.

Success
*/
type S3CredentialsBackupCredentialsForOrganizationListOK struct {
	Payload []*S3CredentialsBackupCredentialsForOrganizationListOKBodyItems0
}

// IsSuccess returns true when this s3 credentials backup credentials for organization list o k response has a 2xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 credentials backup credentials for organization list o k response has a 3xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials backup credentials for organization list o k response has a 4xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 credentials backup credentials for organization list o k response has a 5xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials backup credentials for organization list o k response a status code equal to that given
func (o *S3CredentialsBackupCredentialsForOrganizationListOK) IsCode(code int) bool {
	return code == 200
}

func (o *S3CredentialsBackupCredentialsForOrganizationListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials][%d] s3CredentialsBackupCredentialsForOrganizationListOK  %+v", 200, o.Payload)
}

func (o *S3CredentialsBackupCredentialsForOrganizationListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials][%d] s3CredentialsBackupCredentialsForOrganizationListOK  %+v", 200, o.Payload)
}

func (o *S3CredentialsBackupCredentialsForOrganizationListOK) GetPayload() []*S3CredentialsBackupCredentialsForOrganizationListOKBodyItems0 {
	return o.Payload
}

func (o *S3CredentialsBackupCredentialsForOrganizationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsBackupCredentialsForOrganizationListBadRequest creates a S3CredentialsBackupCredentialsForOrganizationListBadRequest with default headers values
func NewS3CredentialsBackupCredentialsForOrganizationListBadRequest() *S3CredentialsBackupCredentialsForOrganizationListBadRequest {
	return &S3CredentialsBackupCredentialsForOrganizationListBadRequest{}
}

/*
S3CredentialsBackupCredentialsForOrganizationListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type S3CredentialsBackupCredentialsForOrganizationListBadRequest struct {
	Payload []*S3CredentialsBackupCredentialsForOrganizationListBadRequestBodyItems0
}

// IsSuccess returns true when this s3 credentials backup credentials for organization list bad request response has a 2xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials backup credentials for organization list bad request response has a 3xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials backup credentials for organization list bad request response has a 4xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials backup credentials for organization list bad request response has a 5xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials backup credentials for organization list bad request response a status code equal to that given
func (o *S3CredentialsBackupCredentialsForOrganizationListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *S3CredentialsBackupCredentialsForOrganizationListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials][%d] s3CredentialsBackupCredentialsForOrganizationListBadRequest  %+v", 400, o.Payload)
}

func (o *S3CredentialsBackupCredentialsForOrganizationListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials][%d] s3CredentialsBackupCredentialsForOrganizationListBadRequest  %+v", 400, o.Payload)
}

func (o *S3CredentialsBackupCredentialsForOrganizationListBadRequest) GetPayload() []*S3CredentialsBackupCredentialsForOrganizationListBadRequestBodyItems0 {
	return o.Payload
}

func (o *S3CredentialsBackupCredentialsForOrganizationListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsBackupCredentialsForOrganizationListUnauthorized creates a S3CredentialsBackupCredentialsForOrganizationListUnauthorized with default headers values
func NewS3CredentialsBackupCredentialsForOrganizationListUnauthorized() *S3CredentialsBackupCredentialsForOrganizationListUnauthorized {
	return &S3CredentialsBackupCredentialsForOrganizationListUnauthorized{}
}

/*
S3CredentialsBackupCredentialsForOrganizationListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type S3CredentialsBackupCredentialsForOrganizationListUnauthorized struct {
	Payload *S3CredentialsBackupCredentialsForOrganizationListUnauthorizedBody
}

// IsSuccess returns true when this s3 credentials backup credentials for organization list unauthorized response has a 2xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials backup credentials for organization list unauthorized response has a 3xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials backup credentials for organization list unauthorized response has a 4xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials backup credentials for organization list unauthorized response has a 5xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials backup credentials for organization list unauthorized response a status code equal to that given
func (o *S3CredentialsBackupCredentialsForOrganizationListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *S3CredentialsBackupCredentialsForOrganizationListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials][%d] s3CredentialsBackupCredentialsForOrganizationListUnauthorized  %+v", 401, o.Payload)
}

func (o *S3CredentialsBackupCredentialsForOrganizationListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials][%d] s3CredentialsBackupCredentialsForOrganizationListUnauthorized  %+v", 401, o.Payload)
}

func (o *S3CredentialsBackupCredentialsForOrganizationListUnauthorized) GetPayload() *S3CredentialsBackupCredentialsForOrganizationListUnauthorizedBody {
	return o.Payload
}

func (o *S3CredentialsBackupCredentialsForOrganizationListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(S3CredentialsBackupCredentialsForOrganizationListUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsBackupCredentialsForOrganizationListForbidden creates a S3CredentialsBackupCredentialsForOrganizationListForbidden with default headers values
func NewS3CredentialsBackupCredentialsForOrganizationListForbidden() *S3CredentialsBackupCredentialsForOrganizationListForbidden {
	return &S3CredentialsBackupCredentialsForOrganizationListForbidden{}
}

/*
S3CredentialsBackupCredentialsForOrganizationListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type S3CredentialsBackupCredentialsForOrganizationListForbidden struct {
	Payload *S3CredentialsBackupCredentialsForOrganizationListForbiddenBody
}

// IsSuccess returns true when this s3 credentials backup credentials for organization list forbidden response has a 2xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials backup credentials for organization list forbidden response has a 3xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials backup credentials for organization list forbidden response has a 4xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials backup credentials for organization list forbidden response has a 5xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials backup credentials for organization list forbidden response a status code equal to that given
func (o *S3CredentialsBackupCredentialsForOrganizationListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *S3CredentialsBackupCredentialsForOrganizationListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials][%d] s3CredentialsBackupCredentialsForOrganizationListForbidden  %+v", 403, o.Payload)
}

func (o *S3CredentialsBackupCredentialsForOrganizationListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials][%d] s3CredentialsBackupCredentialsForOrganizationListForbidden  %+v", 403, o.Payload)
}

func (o *S3CredentialsBackupCredentialsForOrganizationListForbidden) GetPayload() *S3CredentialsBackupCredentialsForOrganizationListForbiddenBody {
	return o.Payload
}

func (o *S3CredentialsBackupCredentialsForOrganizationListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(S3CredentialsBackupCredentialsForOrganizationListForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsBackupCredentialsForOrganizationListNotFound creates a S3CredentialsBackupCredentialsForOrganizationListNotFound with default headers values
func NewS3CredentialsBackupCredentialsForOrganizationListNotFound() *S3CredentialsBackupCredentialsForOrganizationListNotFound {
	return &S3CredentialsBackupCredentialsForOrganizationListNotFound{}
}

/*
S3CredentialsBackupCredentialsForOrganizationListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type S3CredentialsBackupCredentialsForOrganizationListNotFound struct {
	Payload *S3CredentialsBackupCredentialsForOrganizationListNotFoundBody
}

// IsSuccess returns true when this s3 credentials backup credentials for organization list not found response has a 2xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials backup credentials for organization list not found response has a 3xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials backup credentials for organization list not found response has a 4xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials backup credentials for organization list not found response has a 5xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials backup credentials for organization list not found response a status code equal to that given
func (o *S3CredentialsBackupCredentialsForOrganizationListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *S3CredentialsBackupCredentialsForOrganizationListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials][%d] s3CredentialsBackupCredentialsForOrganizationListNotFound  %+v", 404, o.Payload)
}

func (o *S3CredentialsBackupCredentialsForOrganizationListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials][%d] s3CredentialsBackupCredentialsForOrganizationListNotFound  %+v", 404, o.Payload)
}

func (o *S3CredentialsBackupCredentialsForOrganizationListNotFound) GetPayload() *S3CredentialsBackupCredentialsForOrganizationListNotFoundBody {
	return o.Payload
}

func (o *S3CredentialsBackupCredentialsForOrganizationListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(S3CredentialsBackupCredentialsForOrganizationListNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsBackupCredentialsForOrganizationListInternalServerError creates a S3CredentialsBackupCredentialsForOrganizationListInternalServerError with default headers values
func NewS3CredentialsBackupCredentialsForOrganizationListInternalServerError() *S3CredentialsBackupCredentialsForOrganizationListInternalServerError {
	return &S3CredentialsBackupCredentialsForOrganizationListInternalServerError{}
}

/*
S3CredentialsBackupCredentialsForOrganizationListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type S3CredentialsBackupCredentialsForOrganizationListInternalServerError struct {
}

// IsSuccess returns true when this s3 credentials backup credentials for organization list internal server error response has a 2xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials backup credentials for organization list internal server error response has a 3xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials backup credentials for organization list internal server error response has a 4xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 credentials backup credentials for organization list internal server error response has a 5xx status code
func (o *S3CredentialsBackupCredentialsForOrganizationListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this s3 credentials backup credentials for organization list internal server error response a status code equal to that given
func (o *S3CredentialsBackupCredentialsForOrganizationListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *S3CredentialsBackupCredentialsForOrganizationListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials][%d] s3CredentialsBackupCredentialsForOrganizationListInternalServerError ", 500)
}

func (o *S3CredentialsBackupCredentialsForOrganizationListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials][%d] s3CredentialsBackupCredentialsForOrganizationListInternalServerError ", 500)
}

func (o *S3CredentialsBackupCredentialsForOrganizationListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
S3CredentialsBackupCredentialsForOrganizationListBadRequestBodyItems0 s3 credentials backup credentials for organization list bad request body items0
swagger:model S3CredentialsBackupCredentialsForOrganizationListBadRequestBodyItems0
*/
type S3CredentialsBackupCredentialsForOrganizationListBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this s3 credentials backup credentials for organization list bad request body items0
func (o *S3CredentialsBackupCredentialsForOrganizationListBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s3 credentials backup credentials for organization list bad request body items0 based on context it is used
func (o *S3CredentialsBackupCredentialsForOrganizationListBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *S3CredentialsBackupCredentialsForOrganizationListBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *S3CredentialsBackupCredentialsForOrganizationListBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res S3CredentialsBackupCredentialsForOrganizationListBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
S3CredentialsBackupCredentialsForOrganizationListForbiddenBody s3 credentials backup credentials for organization list forbidden body
swagger:model S3CredentialsBackupCredentialsForOrganizationListForbiddenBody
*/
type S3CredentialsBackupCredentialsForOrganizationListForbiddenBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this s3 credentials backup credentials for organization list forbidden body
func (o *S3CredentialsBackupCredentialsForOrganizationListForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s3 credentials backup credentials for organization list forbidden body based on context it is used
func (o *S3CredentialsBackupCredentialsForOrganizationListForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *S3CredentialsBackupCredentialsForOrganizationListForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *S3CredentialsBackupCredentialsForOrganizationListForbiddenBody) UnmarshalBinary(b []byte) error {
	var res S3CredentialsBackupCredentialsForOrganizationListForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
S3CredentialsBackupCredentialsForOrganizationListNotFoundBody s3 credentials backup credentials for organization list not found body
swagger:model S3CredentialsBackupCredentialsForOrganizationListNotFoundBody
*/
type S3CredentialsBackupCredentialsForOrganizationListNotFoundBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this s3 credentials backup credentials for organization list not found body
func (o *S3CredentialsBackupCredentialsForOrganizationListNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s3 credentials backup credentials for organization list not found body based on context it is used
func (o *S3CredentialsBackupCredentialsForOrganizationListNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *S3CredentialsBackupCredentialsForOrganizationListNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *S3CredentialsBackupCredentialsForOrganizationListNotFoundBody) UnmarshalBinary(b []byte) error {
	var res S3CredentialsBackupCredentialsForOrganizationListNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
S3CredentialsBackupCredentialsForOrganizationListOKBodyItems0 s3 credentials backup credentials for organization list o k body items0
swagger:model S3CredentialsBackupCredentialsForOrganizationListOKBodyItems0
*/
type S3CredentialsBackupCredentialsForOrganizationListOKBodyItems0 struct {

	// backup credential Id
	BackupCredentialID int32 `json:"backupCredentialId,omitempty"`

	// is default
	IsDefault bool `json:"isDefault"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this s3 credentials backup credentials for organization list o k body items0
func (o *S3CredentialsBackupCredentialsForOrganizationListOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s3 credentials backup credentials for organization list o k body items0 based on context it is used
func (o *S3CredentialsBackupCredentialsForOrganizationListOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *S3CredentialsBackupCredentialsForOrganizationListOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *S3CredentialsBackupCredentialsForOrganizationListOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res S3CredentialsBackupCredentialsForOrganizationListOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
S3CredentialsBackupCredentialsForOrganizationListUnauthorizedBody s3 credentials backup credentials for organization list unauthorized body
swagger:model S3CredentialsBackupCredentialsForOrganizationListUnauthorizedBody
*/
type S3CredentialsBackupCredentialsForOrganizationListUnauthorizedBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this s3 credentials backup credentials for organization list unauthorized body
func (o *S3CredentialsBackupCredentialsForOrganizationListUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s3 credentials backup credentials for organization list unauthorized body based on context it is used
func (o *S3CredentialsBackupCredentialsForOrganizationListUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *S3CredentialsBackupCredentialsForOrganizationListUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *S3CredentialsBackupCredentialsForOrganizationListUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res S3CredentialsBackupCredentialsForOrganizationListUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
