// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pgEdge/terraform-provider-pgedge/client/models"
)

// PostSSHKeysReader is a Reader for the PostSSHKeys structure.
type PostSSHKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSSHKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSSHKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSSHKeysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSSHKeysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /ssh-keys] PostSSHKeys", response, response.Code())
	}
}

// NewPostSSHKeysOK creates a PostSSHKeysOK with default headers values
func NewPostSSHKeysOK() *PostSSHKeysOK {
	return &PostSSHKeysOK{}
}

/*
PostSSHKeysOK describes a response with status code 200, with default header values.

Response containing the SSH key details.
*/
type PostSSHKeysOK struct {
	Payload *models.SSHKey
}

// IsSuccess returns true when this post Ssh keys o k response has a 2xx status code
func (o *PostSSHKeysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post Ssh keys o k response has a 3xx status code
func (o *PostSSHKeysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ssh keys o k response has a 4xx status code
func (o *PostSSHKeysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ssh keys o k response has a 5xx status code
func (o *PostSSHKeysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ssh keys o k response a status code equal to that given
func (o *PostSSHKeysOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post Ssh keys o k response
func (o *PostSSHKeysOK) Code() int {
	return 200
}

func (o *PostSSHKeysOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ssh-keys][%d] postSshKeysOK %s", 200, payload)
}

func (o *PostSSHKeysOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ssh-keys][%d] postSshKeysOK %s", 200, payload)
}

func (o *PostSSHKeysOK) GetPayload() *models.SSHKey {
	return o.Payload
}

func (o *PostSSHKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSSHKeysBadRequest creates a PostSSHKeysBadRequest with default headers values
func NewPostSSHKeysBadRequest() *PostSSHKeysBadRequest {
	return &PostSSHKeysBadRequest{}
}

/*
PostSSHKeysBadRequest describes a response with status code 400, with default header values.

Bad request.
*/
type PostSSHKeysBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post Ssh keys bad request response has a 2xx status code
func (o *PostSSHKeysBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ssh keys bad request response has a 3xx status code
func (o *PostSSHKeysBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ssh keys bad request response has a 4xx status code
func (o *PostSSHKeysBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ssh keys bad request response has a 5xx status code
func (o *PostSSHKeysBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ssh keys bad request response a status code equal to that given
func (o *PostSSHKeysBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post Ssh keys bad request response
func (o *PostSSHKeysBadRequest) Code() int {
	return 400
}

func (o *PostSSHKeysBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ssh-keys][%d] postSshKeysBadRequest %s", 400, payload)
}

func (o *PostSSHKeysBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ssh-keys][%d] postSshKeysBadRequest %s", 400, payload)
}

func (o *PostSSHKeysBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSSHKeysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSSHKeysUnauthorized creates a PostSSHKeysUnauthorized with default headers values
func NewPostSSHKeysUnauthorized() *PostSSHKeysUnauthorized {
	return &PostSSHKeysUnauthorized{}
}

/*
PostSSHKeysUnauthorized describes a response with status code 401, with default header values.

Authorization information is missing or invalid.
*/
type PostSSHKeysUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this post Ssh keys unauthorized response has a 2xx status code
func (o *PostSSHKeysUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ssh keys unauthorized response has a 3xx status code
func (o *PostSSHKeysUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ssh keys unauthorized response has a 4xx status code
func (o *PostSSHKeysUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ssh keys unauthorized response has a 5xx status code
func (o *PostSSHKeysUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ssh keys unauthorized response a status code equal to that given
func (o *PostSSHKeysUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post Ssh keys unauthorized response
func (o *PostSSHKeysUnauthorized) Code() int {
	return 401
}

func (o *PostSSHKeysUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ssh-keys][%d] postSshKeysUnauthorized %s", 401, payload)
}

func (o *PostSSHKeysUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ssh-keys][%d] postSshKeysUnauthorized %s", 401, payload)
}

func (o *PostSSHKeysUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSSHKeysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
