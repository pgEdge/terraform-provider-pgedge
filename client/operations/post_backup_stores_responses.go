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

// PostBackupStoresReader is a Reader for the PostBackupStores structure.
type PostBackupStoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostBackupStoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostBackupStoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostBackupStoresBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostBackupStoresUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /backup-stores] PostBackupStores", response, response.Code())
	}
}

// NewPostBackupStoresOK creates a PostBackupStoresOK with default headers values
func NewPostBackupStoresOK() *PostBackupStoresOK {
	return &PostBackupStoresOK{}
}

/*
PostBackupStoresOK describes a response with status code 200, with default header values.

Response containing the backup store details.
*/
type PostBackupStoresOK struct {
	Payload *models.BackupStore
}

// IsSuccess returns true when this post backup stores o k response has a 2xx status code
func (o *PostBackupStoresOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post backup stores o k response has a 3xx status code
func (o *PostBackupStoresOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post backup stores o k response has a 4xx status code
func (o *PostBackupStoresOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post backup stores o k response has a 5xx status code
func (o *PostBackupStoresOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post backup stores o k response a status code equal to that given
func (o *PostBackupStoresOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post backup stores o k response
func (o *PostBackupStoresOK) Code() int {
	return 200
}

func (o *PostBackupStoresOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /backup-stores][%d] postBackupStoresOK %s", 200, payload)
}

func (o *PostBackupStoresOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /backup-stores][%d] postBackupStoresOK %s", 200, payload)
}

func (o *PostBackupStoresOK) GetPayload() *models.BackupStore {
	return o.Payload
}

func (o *PostBackupStoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupStore)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostBackupStoresBadRequest creates a PostBackupStoresBadRequest with default headers values
func NewPostBackupStoresBadRequest() *PostBackupStoresBadRequest {
	return &PostBackupStoresBadRequest{}
}

/*
PostBackupStoresBadRequest describes a response with status code 400, with default header values.

Bad request.
*/
type PostBackupStoresBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post backup stores bad request response has a 2xx status code
func (o *PostBackupStoresBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post backup stores bad request response has a 3xx status code
func (o *PostBackupStoresBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post backup stores bad request response has a 4xx status code
func (o *PostBackupStoresBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post backup stores bad request response has a 5xx status code
func (o *PostBackupStoresBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post backup stores bad request response a status code equal to that given
func (o *PostBackupStoresBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post backup stores bad request response
func (o *PostBackupStoresBadRequest) Code() int {
	return 400
}

func (o *PostBackupStoresBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /backup-stores][%d] postBackupStoresBadRequest %s", 400, payload)
}

func (o *PostBackupStoresBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /backup-stores][%d] postBackupStoresBadRequest %s", 400, payload)
}

func (o *PostBackupStoresBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostBackupStoresBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostBackupStoresUnauthorized creates a PostBackupStoresUnauthorized with default headers values
func NewPostBackupStoresUnauthorized() *PostBackupStoresUnauthorized {
	return &PostBackupStoresUnauthorized{}
}

/*
PostBackupStoresUnauthorized describes a response with status code 401, with default header values.

Authorization information is missing or invalid.
*/
type PostBackupStoresUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this post backup stores unauthorized response has a 2xx status code
func (o *PostBackupStoresUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post backup stores unauthorized response has a 3xx status code
func (o *PostBackupStoresUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post backup stores unauthorized response has a 4xx status code
func (o *PostBackupStoresUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post backup stores unauthorized response has a 5xx status code
func (o *PostBackupStoresUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post backup stores unauthorized response a status code equal to that given
func (o *PostBackupStoresUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post backup stores unauthorized response
func (o *PostBackupStoresUnauthorized) Code() int {
	return 401
}

func (o *PostBackupStoresUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /backup-stores][%d] postBackupStoresUnauthorized %s", 401, payload)
}

func (o *PostBackupStoresUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /backup-stores][%d] postBackupStoresUnauthorized %s", 401, payload)
}

func (o *PostBackupStoresUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostBackupStoresUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
