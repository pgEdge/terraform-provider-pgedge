// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pgEdge/terraform-provider-pgedge/models"
)

// PostDatabasesIDReplicationReader is a Reader for the PostDatabasesIDReplication structure.
type PostDatabasesIDReplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDatabasesIDReplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDatabasesIDReplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostDatabasesIDReplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDatabasesIDReplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDatabasesIDReplicationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /databases/{id}/replication] PostDatabasesIDReplication", response, response.Code())
	}
}

// NewPostDatabasesIDReplicationOK creates a PostDatabasesIDReplicationOK with default headers values
func NewPostDatabasesIDReplicationOK() *PostDatabasesIDReplicationOK {
	return &PostDatabasesIDReplicationOK{}
}

/*
PostDatabasesIDReplicationOK describes a response with status code 200, with default header values.

Successful response
*/
type PostDatabasesIDReplicationOK struct {
	Payload *models.ReplicationResponse
}

// IsSuccess returns true when this post databases Id replication o k response has a 2xx status code
func (o *PostDatabasesIDReplicationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post databases Id replication o k response has a 3xx status code
func (o *PostDatabasesIDReplicationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post databases Id replication o k response has a 4xx status code
func (o *PostDatabasesIDReplicationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post databases Id replication o k response has a 5xx status code
func (o *PostDatabasesIDReplicationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post databases Id replication o k response a status code equal to that given
func (o *PostDatabasesIDReplicationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post databases Id replication o k response
func (o *PostDatabasesIDReplicationOK) Code() int {
	return 200
}

func (o *PostDatabasesIDReplicationOK) Error() string {
	return fmt.Sprintf("[POST /databases/{id}/replication][%d] postDatabasesIdReplicationOK  %+v", 200, o.Payload)
}

func (o *PostDatabasesIDReplicationOK) String() string {
	return fmt.Sprintf("[POST /databases/{id}/replication][%d] postDatabasesIdReplicationOK  %+v", 200, o.Payload)
}

func (o *PostDatabasesIDReplicationOK) GetPayload() *models.ReplicationResponse {
	return o.Payload
}

func (o *PostDatabasesIDReplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReplicationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDatabasesIDReplicationUnauthorized creates a PostDatabasesIDReplicationUnauthorized with default headers values
func NewPostDatabasesIDReplicationUnauthorized() *PostDatabasesIDReplicationUnauthorized {
	return &PostDatabasesIDReplicationUnauthorized{}
}

/*
PostDatabasesIDReplicationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDatabasesIDReplicationUnauthorized struct {
}

// IsSuccess returns true when this post databases Id replication unauthorized response has a 2xx status code
func (o *PostDatabasesIDReplicationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post databases Id replication unauthorized response has a 3xx status code
func (o *PostDatabasesIDReplicationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post databases Id replication unauthorized response has a 4xx status code
func (o *PostDatabasesIDReplicationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post databases Id replication unauthorized response has a 5xx status code
func (o *PostDatabasesIDReplicationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post databases Id replication unauthorized response a status code equal to that given
func (o *PostDatabasesIDReplicationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post databases Id replication unauthorized response
func (o *PostDatabasesIDReplicationUnauthorized) Code() int {
	return 401
}

func (o *PostDatabasesIDReplicationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /databases/{id}/replication][%d] postDatabasesIdReplicationUnauthorized ", 401)
}

func (o *PostDatabasesIDReplicationUnauthorized) String() string {
	return fmt.Sprintf("[POST /databases/{id}/replication][%d] postDatabasesIdReplicationUnauthorized ", 401)
}

func (o *PostDatabasesIDReplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDatabasesIDReplicationNotFound creates a PostDatabasesIDReplicationNotFound with default headers values
func NewPostDatabasesIDReplicationNotFound() *PostDatabasesIDReplicationNotFound {
	return &PostDatabasesIDReplicationNotFound{}
}

/*
PostDatabasesIDReplicationNotFound describes a response with status code 404, with default header values.

Database not found
*/
type PostDatabasesIDReplicationNotFound struct {
}

// IsSuccess returns true when this post databases Id replication not found response has a 2xx status code
func (o *PostDatabasesIDReplicationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post databases Id replication not found response has a 3xx status code
func (o *PostDatabasesIDReplicationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post databases Id replication not found response has a 4xx status code
func (o *PostDatabasesIDReplicationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post databases Id replication not found response has a 5xx status code
func (o *PostDatabasesIDReplicationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post databases Id replication not found response a status code equal to that given
func (o *PostDatabasesIDReplicationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post databases Id replication not found response
func (o *PostDatabasesIDReplicationNotFound) Code() int {
	return 404
}

func (o *PostDatabasesIDReplicationNotFound) Error() string {
	return fmt.Sprintf("[POST /databases/{id}/replication][%d] postDatabasesIdReplicationNotFound ", 404)
}

func (o *PostDatabasesIDReplicationNotFound) String() string {
	return fmt.Sprintf("[POST /databases/{id}/replication][%d] postDatabasesIdReplicationNotFound ", 404)
}

func (o *PostDatabasesIDReplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDatabasesIDReplicationInternalServerError creates a PostDatabasesIDReplicationInternalServerError with default headers values
func NewPostDatabasesIDReplicationInternalServerError() *PostDatabasesIDReplicationInternalServerError {
	return &PostDatabasesIDReplicationInternalServerError{}
}

/*
PostDatabasesIDReplicationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDatabasesIDReplicationInternalServerError struct {
}

// IsSuccess returns true when this post databases Id replication internal server error response has a 2xx status code
func (o *PostDatabasesIDReplicationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post databases Id replication internal server error response has a 3xx status code
func (o *PostDatabasesIDReplicationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post databases Id replication internal server error response has a 4xx status code
func (o *PostDatabasesIDReplicationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post databases Id replication internal server error response has a 5xx status code
func (o *PostDatabasesIDReplicationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post databases Id replication internal server error response a status code equal to that given
func (o *PostDatabasesIDReplicationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post databases Id replication internal server error response
func (o *PostDatabasesIDReplicationInternalServerError) Code() int {
	return 500
}

func (o *PostDatabasesIDReplicationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /databases/{id}/replication][%d] postDatabasesIdReplicationInternalServerError ", 500)
}

func (o *PostDatabasesIDReplicationInternalServerError) String() string {
	return fmt.Sprintf("[POST /databases/{id}/replication][%d] postDatabasesIdReplicationInternalServerError ", 500)
}

func (o *PostDatabasesIDReplicationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
