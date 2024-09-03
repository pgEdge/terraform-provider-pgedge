// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pgEdge/terraform-provider-pgedge/client/models"
)

// PostDatabasesReader is a Reader for the PostDatabases structure.
type PostDatabasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDatabasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDatabasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDatabasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDatabasesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDatabasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /databases] PostDatabases", response, response.Code())
	}
}

// NewPostDatabasesOK creates a PostDatabasesOK with default headers values
func NewPostDatabasesOK() *PostDatabasesOK {
	return &PostDatabasesOK{}
}

/*
PostDatabasesOK describes a response with status code 200, with default header values.

Successful response
*/
type PostDatabasesOK struct {
	Payload *models.DatabaseCreationResponse
}

// IsSuccess returns true when this post databases o k response has a 2xx status code
func (o *PostDatabasesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post databases o k response has a 3xx status code
func (o *PostDatabasesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post databases o k response has a 4xx status code
func (o *PostDatabasesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post databases o k response has a 5xx status code
func (o *PostDatabasesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post databases o k response a status code equal to that given
func (o *PostDatabasesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post databases o k response
func (o *PostDatabasesOK) Code() int {
	return 200
}

func (o *PostDatabasesOK) Error() string {
	return fmt.Sprintf("[POST /databases][%d] postDatabasesOK  %+v", 200, o.Payload)
}

func (o *PostDatabasesOK) String() string {
	return fmt.Sprintf("[POST /databases][%d] postDatabasesOK  %+v", 200, o.Payload)
}

func (o *PostDatabasesOK) GetPayload() *models.DatabaseCreationResponse {
	return o.Payload
}

func (o *PostDatabasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatabaseCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDatabasesBadRequest creates a PostDatabasesBadRequest with default headers values
func NewPostDatabasesBadRequest() *PostDatabasesBadRequest {
	return &PostDatabasesBadRequest{}
}

/*
PostDatabasesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDatabasesBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post databases bad request response has a 2xx status code
func (o *PostDatabasesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post databases bad request response has a 3xx status code
func (o *PostDatabasesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post databases bad request response has a 4xx status code
func (o *PostDatabasesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post databases bad request response has a 5xx status code
func (o *PostDatabasesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post databases bad request response a status code equal to that given
func (o *PostDatabasesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post databases bad request response
func (o *PostDatabasesBadRequest) Code() int {
	return 400
}

func (o *PostDatabasesBadRequest) Error() string {
	return fmt.Sprintf("[POST /databases][%d] postDatabasesBadRequest  %+v", 400, o.Payload)
}

func (o *PostDatabasesBadRequest) String() string {
	return fmt.Sprintf("[POST /databases][%d] postDatabasesBadRequest  %+v", 400, o.Payload)
}

func (o *PostDatabasesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDatabasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDatabasesUnauthorized creates a PostDatabasesUnauthorized with default headers values
func NewPostDatabasesUnauthorized() *PostDatabasesUnauthorized {
	return &PostDatabasesUnauthorized{}
}

/*
PostDatabasesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDatabasesUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this post databases unauthorized response has a 2xx status code
func (o *PostDatabasesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post databases unauthorized response has a 3xx status code
func (o *PostDatabasesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post databases unauthorized response has a 4xx status code
func (o *PostDatabasesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post databases unauthorized response has a 5xx status code
func (o *PostDatabasesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post databases unauthorized response a status code equal to that given
func (o *PostDatabasesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post databases unauthorized response
func (o *PostDatabasesUnauthorized) Code() int {
	return 401
}

func (o *PostDatabasesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /databases][%d] postDatabasesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostDatabasesUnauthorized) String() string {
	return fmt.Sprintf("[POST /databases][%d] postDatabasesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostDatabasesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDatabasesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDatabasesInternalServerError creates a PostDatabasesInternalServerError with default headers values
func NewPostDatabasesInternalServerError() *PostDatabasesInternalServerError {
	return &PostDatabasesInternalServerError{}
}

/*
PostDatabasesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDatabasesInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this post databases internal server error response has a 2xx status code
func (o *PostDatabasesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post databases internal server error response has a 3xx status code
func (o *PostDatabasesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post databases internal server error response has a 4xx status code
func (o *PostDatabasesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post databases internal server error response has a 5xx status code
func (o *PostDatabasesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post databases internal server error response a status code equal to that given
func (o *PostDatabasesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post databases internal server error response
func (o *PostDatabasesInternalServerError) Code() int {
	return 500
}

func (o *PostDatabasesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /databases][%d] postDatabasesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostDatabasesInternalServerError) String() string {
	return fmt.Sprintf("[POST /databases][%d] postDatabasesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostDatabasesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDatabasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
