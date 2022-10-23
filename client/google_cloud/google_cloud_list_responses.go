// Code generated by go-swagger; DO NOT EDIT.

package google_cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GoogleCloudListReader is a Reader for the GoogleCloudList structure.
type GoogleCloudListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GoogleCloudListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGoogleCloudListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGoogleCloudListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGoogleCloudListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGoogleCloudListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGoogleCloudListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGoogleCloudListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGoogleCloudListOK creates a GoogleCloudListOK with default headers values
func NewGoogleCloudListOK() *GoogleCloudListOK {
	return &GoogleCloudListOK{}
}

/*
GoogleCloudListOK describes a response with status code 200, with default header values.

Success
*/
type GoogleCloudListOK struct {
	Payload *GoogleCloudListOKBody
}

// IsSuccess returns true when this google cloud list o k response has a 2xx status code
func (o *GoogleCloudListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this google cloud list o k response has a 3xx status code
func (o *GoogleCloudListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud list o k response has a 4xx status code
func (o *GoogleCloudListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this google cloud list o k response has a 5xx status code
func (o *GoogleCloudListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud list o k response a status code equal to that given
func (o *GoogleCloudListOK) IsCode(code int) bool {
	return code == 200
}

func (o *GoogleCloudListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/GoogleCloud/list][%d] googleCloudListOK  %+v", 200, o.Payload)
}

func (o *GoogleCloudListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/GoogleCloud/list][%d] googleCloudListOK  %+v", 200, o.Payload)
}

func (o *GoogleCloudListOK) GetPayload() *GoogleCloudListOKBody {
	return o.Payload
}

func (o *GoogleCloudListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GoogleCloudListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudListBadRequest creates a GoogleCloudListBadRequest with default headers values
func NewGoogleCloudListBadRequest() *GoogleCloudListBadRequest {
	return &GoogleCloudListBadRequest{}
}

/*
GoogleCloudListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GoogleCloudListBadRequest struct {
	Payload []*GoogleCloudListBadRequestBodyItems0
}

// IsSuccess returns true when this google cloud list bad request response has a 2xx status code
func (o *GoogleCloudListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud list bad request response has a 3xx status code
func (o *GoogleCloudListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud list bad request response has a 4xx status code
func (o *GoogleCloudListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this google cloud list bad request response has a 5xx status code
func (o *GoogleCloudListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud list bad request response a status code equal to that given
func (o *GoogleCloudListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GoogleCloudListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/GoogleCloud/list][%d] googleCloudListBadRequest  %+v", 400, o.Payload)
}

func (o *GoogleCloudListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/GoogleCloud/list][%d] googleCloudListBadRequest  %+v", 400, o.Payload)
}

func (o *GoogleCloudListBadRequest) GetPayload() []*GoogleCloudListBadRequestBodyItems0 {
	return o.Payload
}

func (o *GoogleCloudListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudListUnauthorized creates a GoogleCloudListUnauthorized with default headers values
func NewGoogleCloudListUnauthorized() *GoogleCloudListUnauthorized {
	return &GoogleCloudListUnauthorized{}
}

/*
GoogleCloudListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GoogleCloudListUnauthorized struct {
	Payload *GoogleCloudListUnauthorizedBody
}

// IsSuccess returns true when this google cloud list unauthorized response has a 2xx status code
func (o *GoogleCloudListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud list unauthorized response has a 3xx status code
func (o *GoogleCloudListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud list unauthorized response has a 4xx status code
func (o *GoogleCloudListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this google cloud list unauthorized response has a 5xx status code
func (o *GoogleCloudListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud list unauthorized response a status code equal to that given
func (o *GoogleCloudListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GoogleCloudListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/GoogleCloud/list][%d] googleCloudListUnauthorized  %+v", 401, o.Payload)
}

func (o *GoogleCloudListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/GoogleCloud/list][%d] googleCloudListUnauthorized  %+v", 401, o.Payload)
}

func (o *GoogleCloudListUnauthorized) GetPayload() *GoogleCloudListUnauthorizedBody {
	return o.Payload
}

func (o *GoogleCloudListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GoogleCloudListUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudListForbidden creates a GoogleCloudListForbidden with default headers values
func NewGoogleCloudListForbidden() *GoogleCloudListForbidden {
	return &GoogleCloudListForbidden{}
}

/*
GoogleCloudListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GoogleCloudListForbidden struct {
	Payload *GoogleCloudListForbiddenBody
}

// IsSuccess returns true when this google cloud list forbidden response has a 2xx status code
func (o *GoogleCloudListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud list forbidden response has a 3xx status code
func (o *GoogleCloudListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud list forbidden response has a 4xx status code
func (o *GoogleCloudListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this google cloud list forbidden response has a 5xx status code
func (o *GoogleCloudListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud list forbidden response a status code equal to that given
func (o *GoogleCloudListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GoogleCloudListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/GoogleCloud/list][%d] googleCloudListForbidden  %+v", 403, o.Payload)
}

func (o *GoogleCloudListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/GoogleCloud/list][%d] googleCloudListForbidden  %+v", 403, o.Payload)
}

func (o *GoogleCloudListForbidden) GetPayload() *GoogleCloudListForbiddenBody {
	return o.Payload
}

func (o *GoogleCloudListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GoogleCloudListForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudListNotFound creates a GoogleCloudListNotFound with default headers values
func NewGoogleCloudListNotFound() *GoogleCloudListNotFound {
	return &GoogleCloudListNotFound{}
}

/*
GoogleCloudListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GoogleCloudListNotFound struct {
	Payload *GoogleCloudListNotFoundBody
}

// IsSuccess returns true when this google cloud list not found response has a 2xx status code
func (o *GoogleCloudListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud list not found response has a 3xx status code
func (o *GoogleCloudListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud list not found response has a 4xx status code
func (o *GoogleCloudListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this google cloud list not found response has a 5xx status code
func (o *GoogleCloudListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud list not found response a status code equal to that given
func (o *GoogleCloudListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GoogleCloudListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/GoogleCloud/list][%d] googleCloudListNotFound  %+v", 404, o.Payload)
}

func (o *GoogleCloudListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/GoogleCloud/list][%d] googleCloudListNotFound  %+v", 404, o.Payload)
}

func (o *GoogleCloudListNotFound) GetPayload() *GoogleCloudListNotFoundBody {
	return o.Payload
}

func (o *GoogleCloudListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GoogleCloudListNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudListInternalServerError creates a GoogleCloudListInternalServerError with default headers values
func NewGoogleCloudListInternalServerError() *GoogleCloudListInternalServerError {
	return &GoogleCloudListInternalServerError{}
}

/*
GoogleCloudListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GoogleCloudListInternalServerError struct {
}

// IsSuccess returns true when this google cloud list internal server error response has a 2xx status code
func (o *GoogleCloudListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud list internal server error response has a 3xx status code
func (o *GoogleCloudListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud list internal server error response has a 4xx status code
func (o *GoogleCloudListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this google cloud list internal server error response has a 5xx status code
func (o *GoogleCloudListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this google cloud list internal server error response a status code equal to that given
func (o *GoogleCloudListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GoogleCloudListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/GoogleCloud/list][%d] googleCloudListInternalServerError ", 500)
}

func (o *GoogleCloudListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/GoogleCloud/list][%d] googleCloudListInternalServerError ", 500)
}

func (o *GoogleCloudListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
GoogleCloudListBadRequestBodyItems0 google cloud list bad request body items0
swagger:model GoogleCloudListBadRequestBodyItems0
*/
type GoogleCloudListBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this google cloud list bad request body items0
func (o *GoogleCloudListBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this google cloud list bad request body items0 based on context it is used
func (o *GoogleCloudListBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GoogleCloudListBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GoogleCloudListBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res GoogleCloudListBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GoogleCloudListForbiddenBody google cloud list forbidden body
swagger:model GoogleCloudListForbiddenBody
*/
type GoogleCloudListForbiddenBody struct {

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

// Validate validates this google cloud list forbidden body
func (o *GoogleCloudListForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this google cloud list forbidden body based on context it is used
func (o *GoogleCloudListForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GoogleCloudListForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GoogleCloudListForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GoogleCloudListForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GoogleCloudListNotFoundBody google cloud list not found body
swagger:model GoogleCloudListNotFoundBody
*/
type GoogleCloudListNotFoundBody struct {

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

// Validate validates this google cloud list not found body
func (o *GoogleCloudListNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this google cloud list not found body based on context it is used
func (o *GoogleCloudListNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GoogleCloudListNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GoogleCloudListNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GoogleCloudListNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GoogleCloudListOKBody google cloud list o k body
swagger:model GoogleCloudListOKBody
*/
type GoogleCloudListOKBody struct {

	// data
	Data []*GoogleCloudListOKBodyDataItems0 `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this google cloud list o k body
func (o *GoogleCloudListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GoogleCloudListOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("googleCloudListOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("googleCloudListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this google cloud list o k body based on the context it is used
func (o *GoogleCloudListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GoogleCloudListOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("googleCloudListOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("googleCloudListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GoogleCloudListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GoogleCloudListOKBody) UnmarshalBinary(b []byte) error {
	var res GoogleCloudListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GoogleCloudListOKBodyDataItems0 google cloud list o k body data items0
swagger:model GoogleCloudListOKBodyDataItems0
*/
type GoogleCloudListOKBodyDataItems0 struct {

	// billing account Id
	BillingAccountID string `json:"billingAccountId,omitempty"`

	// billing account name
	BillingAccountName string `json:"billingAccountName,omitempty"`

	// continent name
	ContinentName string `json:"continentName,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// folder Id
	FolderID string `json:"folderId,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is default
	IsDefault bool `json:"isDefault"`

	// is locked
	IsLocked bool `json:"isLocked"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// partner logo
	PartnerLogo string `json:"partnerLogo,omitempty"`

	// partner name
	PartnerName string `json:"partnerName,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// projects
	Projects []*GoogleCloudListOKBodyDataItems0ProjectsItems0 `json:"projects"`

	// region
	Region string `json:"region,omitempty"`

	// zones
	Zones []string `json:"zones"`
}

// Validate validates this google cloud list o k body data items0
func (o *GoogleCloudListOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GoogleCloudListOKBodyDataItems0) validateProjects(formats strfmt.Registry) error {
	if swag.IsZero(o.Projects) { // not required
		return nil
	}

	for i := 0; i < len(o.Projects); i++ {
		if swag.IsZero(o.Projects[i]) { // not required
			continue
		}

		if o.Projects[i] != nil {
			if err := o.Projects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this google cloud list o k body data items0 based on the context it is used
func (o *GoogleCloudListOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateProjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GoogleCloudListOKBodyDataItems0) contextValidateProjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Projects); i++ {

		if o.Projects[i] != nil {
			if err := o.Projects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GoogleCloudListOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GoogleCloudListOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res GoogleCloudListOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GoogleCloudListOKBodyDataItems0ProjectsItems0 google cloud list o k body data items0 projects items0
swagger:model GoogleCloudListOKBodyDataItems0ProjectsItems0
*/
type GoogleCloudListOKBodyDataItems0ProjectsItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this google cloud list o k body data items0 projects items0
func (o *GoogleCloudListOKBodyDataItems0ProjectsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this google cloud list o k body data items0 projects items0 based on context it is used
func (o *GoogleCloudListOKBodyDataItems0ProjectsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GoogleCloudListOKBodyDataItems0ProjectsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GoogleCloudListOKBodyDataItems0ProjectsItems0) UnmarshalBinary(b []byte) error {
	var res GoogleCloudListOKBodyDataItems0ProjectsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GoogleCloudListUnauthorizedBody google cloud list unauthorized body
swagger:model GoogleCloudListUnauthorizedBody
*/
type GoogleCloudListUnauthorizedBody struct {

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

// Validate validates this google cloud list unauthorized body
func (o *GoogleCloudListUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this google cloud list unauthorized body based on context it is used
func (o *GoogleCloudListUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GoogleCloudListUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GoogleCloudListUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GoogleCloudListUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
