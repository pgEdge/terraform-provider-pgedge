// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostOauthTokenReader is a Reader for the PostOauthToken structure.
type PostOauthTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOauthTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOauthTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOauthTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOauthTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /oauth/token] PostOauthToken", response, response.Code())
	}
}

// NewPostOauthTokenOK creates a PostOauthTokenOK with default headers values
func NewPostOauthTokenOK() *PostOauthTokenOK {
	return &PostOauthTokenOK{}
}

/*
PostOauthTokenOK describes a response with status code 200, with default header values.

Successful response
*/
type PostOauthTokenOK struct {
	Payload *PostOauthTokenOKBody
}

// IsSuccess returns true when this post oauth token o k response has a 2xx status code
func (o *PostOauthTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post oauth token o k response has a 3xx status code
func (o *PostOauthTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth token o k response has a 4xx status code
func (o *PostOauthTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post oauth token o k response has a 5xx status code
func (o *PostOauthTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post oauth token o k response a status code equal to that given
func (o *PostOauthTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post oauth token o k response
func (o *PostOauthTokenOK) Code() int {
	return 200
}

func (o *PostOauthTokenOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /oauth/token][%d] postOauthTokenOK %s", 200, payload)
}

func (o *PostOauthTokenOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /oauth/token][%d] postOauthTokenOK %s", 200, payload)
}

func (o *PostOauthTokenOK) GetPayload() *PostOauthTokenOKBody {
	return o.Payload
}

func (o *PostOauthTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostOauthTokenOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthTokenBadRequest creates a PostOauthTokenBadRequest with default headers values
func NewPostOauthTokenBadRequest() *PostOauthTokenBadRequest {
	return &PostOauthTokenBadRequest{}
}

/*
PostOauthTokenBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostOauthTokenBadRequest struct {
	Payload *PostOauthTokenBadRequestBody
}

// IsSuccess returns true when this post oauth token bad request response has a 2xx status code
func (o *PostOauthTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post oauth token bad request response has a 3xx status code
func (o *PostOauthTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth token bad request response has a 4xx status code
func (o *PostOauthTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post oauth token bad request response has a 5xx status code
func (o *PostOauthTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post oauth token bad request response a status code equal to that given
func (o *PostOauthTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post oauth token bad request response
func (o *PostOauthTokenBadRequest) Code() int {
	return 400
}

func (o *PostOauthTokenBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /oauth/token][%d] postOauthTokenBadRequest %s", 400, payload)
}

func (o *PostOauthTokenBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /oauth/token][%d] postOauthTokenBadRequest %s", 400, payload)
}

func (o *PostOauthTokenBadRequest) GetPayload() *PostOauthTokenBadRequestBody {
	return o.Payload
}

func (o *PostOauthTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostOauthTokenBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthTokenInternalServerError creates a PostOauthTokenInternalServerError with default headers values
func NewPostOauthTokenInternalServerError() *PostOauthTokenInternalServerError {
	return &PostOauthTokenInternalServerError{}
}

/*
PostOauthTokenInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostOauthTokenInternalServerError struct {
	Payload *PostOauthTokenInternalServerErrorBody
}

// IsSuccess returns true when this post oauth token internal server error response has a 2xx status code
func (o *PostOauthTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post oauth token internal server error response has a 3xx status code
func (o *PostOauthTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth token internal server error response has a 4xx status code
func (o *PostOauthTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post oauth token internal server error response has a 5xx status code
func (o *PostOauthTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post oauth token internal server error response a status code equal to that given
func (o *PostOauthTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post oauth token internal server error response
func (o *PostOauthTokenInternalServerError) Code() int {
	return 500
}

func (o *PostOauthTokenInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /oauth/token][%d] postOauthTokenInternalServerError %s", 500, payload)
}

func (o *PostOauthTokenInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /oauth/token][%d] postOauthTokenInternalServerError %s", 500, payload)
}

func (o *PostOauthTokenInternalServerError) GetPayload() *PostOauthTokenInternalServerErrorBody {
	return o.Payload
}

func (o *PostOauthTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostOauthTokenInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostOauthTokenBadRequestBody post oauth token bad request body
swagger:model PostOauthTokenBadRequestBody
*/
type PostOauthTokenBadRequestBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this post oauth token bad request body
func (o *PostOauthTokenBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post oauth token bad request body based on context it is used
func (o *PostOauthTokenBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOauthTokenBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOauthTokenBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostOauthTokenBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostOauthTokenBody post oauth token body
swagger:model PostOauthTokenBody
*/
type PostOauthTokenBody struct {

	// client id
	ClientID string `json:"client_id,omitempty"`

	// client secret
	ClientSecret string `json:"client_secret,omitempty"`

	// grant type
	GrantType string `json:"grant_type,omitempty"`
}

// Validate validates this post oauth token body
func (o *PostOauthTokenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post oauth token body based on context it is used
func (o *PostOauthTokenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOauthTokenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOauthTokenBody) UnmarshalBinary(b []byte) error {
	var res PostOauthTokenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostOauthTokenInternalServerErrorBody post oauth token internal server error body
swagger:model PostOauthTokenInternalServerErrorBody
*/
type PostOauthTokenInternalServerErrorBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this post oauth token internal server error body
func (o *PostOauthTokenInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post oauth token internal server error body based on context it is used
func (o *PostOauthTokenInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOauthTokenInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOauthTokenInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostOauthTokenInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostOauthTokenOKBody post oauth token o k body
swagger:model PostOauthTokenOKBody
*/
type PostOauthTokenOKBody struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// expires in
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// token type
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this post oauth token o k body
func (o *PostOauthTokenOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post oauth token o k body based on context it is used
func (o *PostOauthTokenOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOauthTokenOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOauthTokenOKBody) UnmarshalBinary(b []byte) error {
	var res PostOauthTokenOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}