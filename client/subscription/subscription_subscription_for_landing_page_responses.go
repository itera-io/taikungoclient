// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// SubscriptionSubscriptionForLandingPageReader is a Reader for the SubscriptionSubscriptionForLandingPage structure.
type SubscriptionSubscriptionForLandingPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionSubscriptionForLandingPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionSubscriptionForLandingPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubscriptionSubscriptionForLandingPageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSubscriptionSubscriptionForLandingPageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubscriptionSubscriptionForLandingPageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubscriptionSubscriptionForLandingPageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubscriptionSubscriptionForLandingPageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubscriptionSubscriptionForLandingPageOK creates a SubscriptionSubscriptionForLandingPageOK with default headers values
func NewSubscriptionSubscriptionForLandingPageOK() *SubscriptionSubscriptionForLandingPageOK {
	return &SubscriptionSubscriptionForLandingPageOK{}
}

/*
SubscriptionSubscriptionForLandingPageOK describes a response with status code 200, with default header values.

Success
*/
type SubscriptionSubscriptionForLandingPageOK struct {
	Payload []*SubscriptionSubscriptionForLandingPageOKBodyItems0
}

// IsSuccess returns true when this subscription subscription for landing page o k response has a 2xx status code
func (o *SubscriptionSubscriptionForLandingPageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription subscription for landing page o k response has a 3xx status code
func (o *SubscriptionSubscriptionForLandingPageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for landing page o k response has a 4xx status code
func (o *SubscriptionSubscriptionForLandingPageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription subscription for landing page o k response has a 5xx status code
func (o *SubscriptionSubscriptionForLandingPageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for landing page o k response a status code equal to that given
func (o *SubscriptionSubscriptionForLandingPageOK) IsCode(code int) bool {
	return code == 200
}

func (o *SubscriptionSubscriptionForLandingPageOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageOK  %+v", 200, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageOK  %+v", 200, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageOK) GetPayload() []*SubscriptionSubscriptionForLandingPageOKBodyItems0 {
	return o.Payload
}

func (o *SubscriptionSubscriptionForLandingPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForLandingPageBadRequest creates a SubscriptionSubscriptionForLandingPageBadRequest with default headers values
func NewSubscriptionSubscriptionForLandingPageBadRequest() *SubscriptionSubscriptionForLandingPageBadRequest {
	return &SubscriptionSubscriptionForLandingPageBadRequest{}
}

/*
SubscriptionSubscriptionForLandingPageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SubscriptionSubscriptionForLandingPageBadRequest struct {
	Payload []*SubscriptionSubscriptionForLandingPageBadRequestBodyItems0
}

// IsSuccess returns true when this subscription subscription for landing page bad request response has a 2xx status code
func (o *SubscriptionSubscriptionForLandingPageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for landing page bad request response has a 3xx status code
func (o *SubscriptionSubscriptionForLandingPageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for landing page bad request response has a 4xx status code
func (o *SubscriptionSubscriptionForLandingPageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription subscription for landing page bad request response has a 5xx status code
func (o *SubscriptionSubscriptionForLandingPageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for landing page bad request response a status code equal to that given
func (o *SubscriptionSubscriptionForLandingPageBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SubscriptionSubscriptionForLandingPageBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageBadRequest  %+v", 400, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageBadRequest  %+v", 400, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageBadRequest) GetPayload() []*SubscriptionSubscriptionForLandingPageBadRequestBodyItems0 {
	return o.Payload
}

func (o *SubscriptionSubscriptionForLandingPageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForLandingPageUnauthorized creates a SubscriptionSubscriptionForLandingPageUnauthorized with default headers values
func NewSubscriptionSubscriptionForLandingPageUnauthorized() *SubscriptionSubscriptionForLandingPageUnauthorized {
	return &SubscriptionSubscriptionForLandingPageUnauthorized{}
}

/*
SubscriptionSubscriptionForLandingPageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SubscriptionSubscriptionForLandingPageUnauthorized struct {
	Payload *SubscriptionSubscriptionForLandingPageUnauthorizedBody
}

// IsSuccess returns true when this subscription subscription for landing page unauthorized response has a 2xx status code
func (o *SubscriptionSubscriptionForLandingPageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for landing page unauthorized response has a 3xx status code
func (o *SubscriptionSubscriptionForLandingPageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for landing page unauthorized response has a 4xx status code
func (o *SubscriptionSubscriptionForLandingPageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription subscription for landing page unauthorized response has a 5xx status code
func (o *SubscriptionSubscriptionForLandingPageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for landing page unauthorized response a status code equal to that given
func (o *SubscriptionSubscriptionForLandingPageUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SubscriptionSubscriptionForLandingPageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageUnauthorized  %+v", 401, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageUnauthorized  %+v", 401, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageUnauthorized) GetPayload() *SubscriptionSubscriptionForLandingPageUnauthorizedBody {
	return o.Payload
}

func (o *SubscriptionSubscriptionForLandingPageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SubscriptionSubscriptionForLandingPageUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForLandingPageForbidden creates a SubscriptionSubscriptionForLandingPageForbidden with default headers values
func NewSubscriptionSubscriptionForLandingPageForbidden() *SubscriptionSubscriptionForLandingPageForbidden {
	return &SubscriptionSubscriptionForLandingPageForbidden{}
}

/*
SubscriptionSubscriptionForLandingPageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SubscriptionSubscriptionForLandingPageForbidden struct {
	Payload *SubscriptionSubscriptionForLandingPageForbiddenBody
}

// IsSuccess returns true when this subscription subscription for landing page forbidden response has a 2xx status code
func (o *SubscriptionSubscriptionForLandingPageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for landing page forbidden response has a 3xx status code
func (o *SubscriptionSubscriptionForLandingPageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for landing page forbidden response has a 4xx status code
func (o *SubscriptionSubscriptionForLandingPageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription subscription for landing page forbidden response has a 5xx status code
func (o *SubscriptionSubscriptionForLandingPageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for landing page forbidden response a status code equal to that given
func (o *SubscriptionSubscriptionForLandingPageForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SubscriptionSubscriptionForLandingPageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageForbidden  %+v", 403, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageForbidden  %+v", 403, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageForbidden) GetPayload() *SubscriptionSubscriptionForLandingPageForbiddenBody {
	return o.Payload
}

func (o *SubscriptionSubscriptionForLandingPageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SubscriptionSubscriptionForLandingPageForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForLandingPageNotFound creates a SubscriptionSubscriptionForLandingPageNotFound with default headers values
func NewSubscriptionSubscriptionForLandingPageNotFound() *SubscriptionSubscriptionForLandingPageNotFound {
	return &SubscriptionSubscriptionForLandingPageNotFound{}
}

/*
SubscriptionSubscriptionForLandingPageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SubscriptionSubscriptionForLandingPageNotFound struct {
	Payload *SubscriptionSubscriptionForLandingPageNotFoundBody
}

// IsSuccess returns true when this subscription subscription for landing page not found response has a 2xx status code
func (o *SubscriptionSubscriptionForLandingPageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for landing page not found response has a 3xx status code
func (o *SubscriptionSubscriptionForLandingPageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for landing page not found response has a 4xx status code
func (o *SubscriptionSubscriptionForLandingPageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription subscription for landing page not found response has a 5xx status code
func (o *SubscriptionSubscriptionForLandingPageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for landing page not found response a status code equal to that given
func (o *SubscriptionSubscriptionForLandingPageNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SubscriptionSubscriptionForLandingPageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageNotFound  %+v", 404, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageNotFound  %+v", 404, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageNotFound) GetPayload() *SubscriptionSubscriptionForLandingPageNotFoundBody {
	return o.Payload
}

func (o *SubscriptionSubscriptionForLandingPageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SubscriptionSubscriptionForLandingPageNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForLandingPageInternalServerError creates a SubscriptionSubscriptionForLandingPageInternalServerError with default headers values
func NewSubscriptionSubscriptionForLandingPageInternalServerError() *SubscriptionSubscriptionForLandingPageInternalServerError {
	return &SubscriptionSubscriptionForLandingPageInternalServerError{}
}

/*
SubscriptionSubscriptionForLandingPageInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SubscriptionSubscriptionForLandingPageInternalServerError struct {
}

// IsSuccess returns true when this subscription subscription for landing page internal server error response has a 2xx status code
func (o *SubscriptionSubscriptionForLandingPageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for landing page internal server error response has a 3xx status code
func (o *SubscriptionSubscriptionForLandingPageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for landing page internal server error response has a 4xx status code
func (o *SubscriptionSubscriptionForLandingPageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription subscription for landing page internal server error response has a 5xx status code
func (o *SubscriptionSubscriptionForLandingPageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this subscription subscription for landing page internal server error response a status code equal to that given
func (o *SubscriptionSubscriptionForLandingPageInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SubscriptionSubscriptionForLandingPageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageInternalServerError ", 500)
}

func (o *SubscriptionSubscriptionForLandingPageInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageInternalServerError ", 500)
}

func (o *SubscriptionSubscriptionForLandingPageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
SubscriptionSubscriptionForLandingPageBadRequestBodyItems0 subscription subscription for landing page bad request body items0
swagger:model SubscriptionSubscriptionForLandingPageBadRequestBodyItems0
*/
type SubscriptionSubscriptionForLandingPageBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this subscription subscription for landing page bad request body items0
func (o *SubscriptionSubscriptionForLandingPageBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this subscription subscription for landing page bad request body items0 based on context it is used
func (o *SubscriptionSubscriptionForLandingPageBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SubscriptionSubscriptionForLandingPageBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SubscriptionSubscriptionForLandingPageBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res SubscriptionSubscriptionForLandingPageBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SubscriptionSubscriptionForLandingPageForbiddenBody subscription subscription for landing page forbidden body
swagger:model SubscriptionSubscriptionForLandingPageForbiddenBody
*/
type SubscriptionSubscriptionForLandingPageForbiddenBody struct {

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

// Validate validates this subscription subscription for landing page forbidden body
func (o *SubscriptionSubscriptionForLandingPageForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this subscription subscription for landing page forbidden body based on context it is used
func (o *SubscriptionSubscriptionForLandingPageForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SubscriptionSubscriptionForLandingPageForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SubscriptionSubscriptionForLandingPageForbiddenBody) UnmarshalBinary(b []byte) error {
	var res SubscriptionSubscriptionForLandingPageForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SubscriptionSubscriptionForLandingPageNotFoundBody subscription subscription for landing page not found body
swagger:model SubscriptionSubscriptionForLandingPageNotFoundBody
*/
type SubscriptionSubscriptionForLandingPageNotFoundBody struct {

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

// Validate validates this subscription subscription for landing page not found body
func (o *SubscriptionSubscriptionForLandingPageNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this subscription subscription for landing page not found body based on context it is used
func (o *SubscriptionSubscriptionForLandingPageNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SubscriptionSubscriptionForLandingPageNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SubscriptionSubscriptionForLandingPageNotFoundBody) UnmarshalBinary(b []byte) error {
	var res SubscriptionSubscriptionForLandingPageNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SubscriptionSubscriptionForLandingPageOKBodyItems0 subscription subscription for landing page o k body items0
swagger:model SubscriptionSubscriptionForLandingPageOKBodyItems0
*/
type SubscriptionSubscriptionForLandingPageOKBodyItems0 struct {

	// cloud credential limit
	CloudCredentialLimit int32 `json:"cloudCredentialLimit,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is demo
	IsDemo bool `json:"isDemo"`

	// is deprecated
	IsDeprecated bool `json:"isDeprecated"`

	// monthly price
	MonthlyPrice float64 `json:"monthlyPrice,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// partner Id
	PartnerID int32 `json:"partnerId,omitempty"`

	// project limit
	ProjectLimit int32 `json:"projectLimit,omitempty"`

	// server limit
	ServerLimit int32 `json:"serverLimit,omitempty"`

	// tcu price
	TcuPrice float64 `json:"tcuPrice,omitempty"`

	// trial days
	TrialDays int32 `json:"trialDays,omitempty"`

	// user limit
	UserLimit int32 `json:"userLimit,omitempty"`

	// yearly price
	YearlyPrice float64 `json:"yearlyPrice,omitempty"`
}

// Validate validates this subscription subscription for landing page o k body items0
func (o *SubscriptionSubscriptionForLandingPageOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this subscription subscription for landing page o k body items0 based on context it is used
func (o *SubscriptionSubscriptionForLandingPageOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SubscriptionSubscriptionForLandingPageOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SubscriptionSubscriptionForLandingPageOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res SubscriptionSubscriptionForLandingPageOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SubscriptionSubscriptionForLandingPageUnauthorizedBody subscription subscription for landing page unauthorized body
swagger:model SubscriptionSubscriptionForLandingPageUnauthorizedBody
*/
type SubscriptionSubscriptionForLandingPageUnauthorizedBody struct {

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

// Validate validates this subscription subscription for landing page unauthorized body
func (o *SubscriptionSubscriptionForLandingPageUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this subscription subscription for landing page unauthorized body based on context it is used
func (o *SubscriptionSubscriptionForLandingPageUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SubscriptionSubscriptionForLandingPageUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SubscriptionSubscriptionForLandingPageUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res SubscriptionSubscriptionForLandingPageUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
