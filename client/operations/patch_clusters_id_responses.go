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

// PatchClustersIDReader is a Reader for the PatchClustersID structure.
type PatchClustersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchClustersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchClustersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchClustersIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchClustersIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /clusters/{id}] PatchClustersID", response, response.Code())
	}
}

// NewPatchClustersIDOK creates a PatchClustersIDOK with default headers values
func NewPatchClustersIDOK() *PatchClustersIDOK {
	return &PatchClustersIDOK{}
}

/*
PatchClustersIDOK describes a response with status code 200, with default header values.

Response containing the cluster definition.
*/
type PatchClustersIDOK struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this patch clusters Id o k response has a 2xx status code
func (o *PatchClustersIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch clusters Id o k response has a 3xx status code
func (o *PatchClustersIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch clusters Id o k response has a 4xx status code
func (o *PatchClustersIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch clusters Id o k response has a 5xx status code
func (o *PatchClustersIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch clusters Id o k response a status code equal to that given
func (o *PatchClustersIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch clusters Id o k response
func (o *PatchClustersIDOK) Code() int {
	return 200
}

func (o *PatchClustersIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /clusters/{id}][%d] patchClustersIdOK %s", 200, payload)
}

func (o *PatchClustersIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /clusters/{id}][%d] patchClustersIdOK %s", 200, payload)
}

func (o *PatchClustersIDOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *PatchClustersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchClustersIDBadRequest creates a PatchClustersIDBadRequest with default headers values
func NewPatchClustersIDBadRequest() *PatchClustersIDBadRequest {
	return &PatchClustersIDBadRequest{}
}

/*
PatchClustersIDBadRequest describes a response with status code 400, with default header values.

Bad request.
*/
type PatchClustersIDBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this patch clusters Id bad request response has a 2xx status code
func (o *PatchClustersIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch clusters Id bad request response has a 3xx status code
func (o *PatchClustersIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch clusters Id bad request response has a 4xx status code
func (o *PatchClustersIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch clusters Id bad request response has a 5xx status code
func (o *PatchClustersIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch clusters Id bad request response a status code equal to that given
func (o *PatchClustersIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the patch clusters Id bad request response
func (o *PatchClustersIDBadRequest) Code() int {
	return 400
}

func (o *PatchClustersIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /clusters/{id}][%d] patchClustersIdBadRequest %s", 400, payload)
}

func (o *PatchClustersIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /clusters/{id}][%d] patchClustersIdBadRequest %s", 400, payload)
}

func (o *PatchClustersIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchClustersIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchClustersIDUnauthorized creates a PatchClustersIDUnauthorized with default headers values
func NewPatchClustersIDUnauthorized() *PatchClustersIDUnauthorized {
	return &PatchClustersIDUnauthorized{}
}

/*
PatchClustersIDUnauthorized describes a response with status code 401, with default header values.

Authorization information is missing or invalid.
*/
type PatchClustersIDUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this patch clusters Id unauthorized response has a 2xx status code
func (o *PatchClustersIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch clusters Id unauthorized response has a 3xx status code
func (o *PatchClustersIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch clusters Id unauthorized response has a 4xx status code
func (o *PatchClustersIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch clusters Id unauthorized response has a 5xx status code
func (o *PatchClustersIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch clusters Id unauthorized response a status code equal to that given
func (o *PatchClustersIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch clusters Id unauthorized response
func (o *PatchClustersIDUnauthorized) Code() int {
	return 401
}

func (o *PatchClustersIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /clusters/{id}][%d] patchClustersIdUnauthorized %s", 401, payload)
}

func (o *PatchClustersIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /clusters/{id}][%d] patchClustersIdUnauthorized %s", 401, payload)
}

func (o *PatchClustersIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchClustersIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}