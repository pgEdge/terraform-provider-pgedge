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

// GetBackupStoresIDReader is a Reader for the GetBackupStoresID structure.
type GetBackupStoresIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupStoresIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupStoresIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBackupStoresIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetBackupStoresIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /backup-stores/{id}] GetBackupStoresID", response, response.Code())
	}
}

// NewGetBackupStoresIDOK creates a GetBackupStoresIDOK with default headers values
func NewGetBackupStoresIDOK() *GetBackupStoresIDOK {
	return &GetBackupStoresIDOK{}
}

/*
GetBackupStoresIDOK describes a response with status code 200, with default header values.

Response containing the backup store details.
*/
type GetBackupStoresIDOK struct {
	Payload *models.BackupStore
}

// IsSuccess returns true when this get backup stores Id o k response has a 2xx status code
func (o *GetBackupStoresIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get backup stores Id o k response has a 3xx status code
func (o *GetBackupStoresIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup stores Id o k response has a 4xx status code
func (o *GetBackupStoresIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get backup stores Id o k response has a 5xx status code
func (o *GetBackupStoresIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup stores Id o k response a status code equal to that given
func (o *GetBackupStoresIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get backup stores Id o k response
func (o *GetBackupStoresIDOK) Code() int {
	return 200
}

func (o *GetBackupStoresIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backup-stores/{id}][%d] getBackupStoresIdOK %s", 200, payload)
}

func (o *GetBackupStoresIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backup-stores/{id}][%d] getBackupStoresIdOK %s", 200, payload)
}

func (o *GetBackupStoresIDOK) GetPayload() *models.BackupStore {
	return o.Payload
}

func (o *GetBackupStoresIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupStore)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupStoresIDBadRequest creates a GetBackupStoresIDBadRequest with default headers values
func NewGetBackupStoresIDBadRequest() *GetBackupStoresIDBadRequest {
	return &GetBackupStoresIDBadRequest{}
}

/*
GetBackupStoresIDBadRequest describes a response with status code 400, with default header values.

Bad request.
*/
type GetBackupStoresIDBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get backup stores Id bad request response has a 2xx status code
func (o *GetBackupStoresIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backup stores Id bad request response has a 3xx status code
func (o *GetBackupStoresIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup stores Id bad request response has a 4xx status code
func (o *GetBackupStoresIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get backup stores Id bad request response has a 5xx status code
func (o *GetBackupStoresIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup stores Id bad request response a status code equal to that given
func (o *GetBackupStoresIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get backup stores Id bad request response
func (o *GetBackupStoresIDBadRequest) Code() int {
	return 400
}

func (o *GetBackupStoresIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backup-stores/{id}][%d] getBackupStoresIdBadRequest %s", 400, payload)
}

func (o *GetBackupStoresIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backup-stores/{id}][%d] getBackupStoresIdBadRequest %s", 400, payload)
}

func (o *GetBackupStoresIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackupStoresIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupStoresIDUnauthorized creates a GetBackupStoresIDUnauthorized with default headers values
func NewGetBackupStoresIDUnauthorized() *GetBackupStoresIDUnauthorized {
	return &GetBackupStoresIDUnauthorized{}
}

/*
GetBackupStoresIDUnauthorized describes a response with status code 401, with default header values.

Authorization information is missing or invalid.
*/
type GetBackupStoresIDUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get backup stores Id unauthorized response has a 2xx status code
func (o *GetBackupStoresIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backup stores Id unauthorized response has a 3xx status code
func (o *GetBackupStoresIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup stores Id unauthorized response has a 4xx status code
func (o *GetBackupStoresIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get backup stores Id unauthorized response has a 5xx status code
func (o *GetBackupStoresIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup stores Id unauthorized response a status code equal to that given
func (o *GetBackupStoresIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get backup stores Id unauthorized response
func (o *GetBackupStoresIDUnauthorized) Code() int {
	return 401
}

func (o *GetBackupStoresIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backup-stores/{id}][%d] getBackupStoresIdUnauthorized %s", 401, payload)
}

func (o *GetBackupStoresIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backup-stores/{id}][%d] getBackupStoresIdUnauthorized %s", 401, payload)
}

func (o *GetBackupStoresIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackupStoresIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}