// Code generated by go-swagger; DO NOT EDIT.

package unmanaged

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/community-edition/cli/cmd/plugin/ui/server/models"
)

// GetUnmanagedClustersOKCode is the HTTP code returned for type GetUnmanagedClustersOK
const GetUnmanagedClustersOKCode int = 200

/*GetUnmanagedClustersOK Success

swagger:response getUnmanagedClustersOK
*/
type GetUnmanagedClustersOK struct {

	/*List of unmanaged clusters
	  In: Body
	*/
	Payload []*models.UnmanagedCluster `json:"body,omitempty"`
}

// NewGetUnmanagedClustersOK creates GetUnmanagedClustersOK with default headers values
func NewGetUnmanagedClustersOK() *GetUnmanagedClustersOK {

	return &GetUnmanagedClustersOK{}
}

// WithPayload adds the payload to the get unmanaged clusters o k response
func (o *GetUnmanagedClustersOK) WithPayload(payload []*models.UnmanagedCluster) *GetUnmanagedClustersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get unmanaged clusters o k response
func (o *GetUnmanagedClustersOK) SetPayload(payload []*models.UnmanagedCluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUnmanagedClustersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.UnmanagedCluster, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetUnmanagedClustersBadRequestCode is the HTTP code returned for type GetUnmanagedClustersBadRequest
const GetUnmanagedClustersBadRequestCode int = 400

/*GetUnmanagedClustersBadRequest Bad request

swagger:response getUnmanagedClustersBadRequest
*/
type GetUnmanagedClustersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUnmanagedClustersBadRequest creates GetUnmanagedClustersBadRequest with default headers values
func NewGetUnmanagedClustersBadRequest() *GetUnmanagedClustersBadRequest {

	return &GetUnmanagedClustersBadRequest{}
}

// WithPayload adds the payload to the get unmanaged clusters bad request response
func (o *GetUnmanagedClustersBadRequest) WithPayload(payload *models.Error) *GetUnmanagedClustersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get unmanaged clusters bad request response
func (o *GetUnmanagedClustersBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUnmanagedClustersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUnmanagedClustersInternalServerErrorCode is the HTTP code returned for type GetUnmanagedClustersInternalServerError
const GetUnmanagedClustersInternalServerErrorCode int = 500

/*GetUnmanagedClustersInternalServerError Internal server error

swagger:response getUnmanagedClustersInternalServerError
*/
type GetUnmanagedClustersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUnmanagedClustersInternalServerError creates GetUnmanagedClustersInternalServerError with default headers values
func NewGetUnmanagedClustersInternalServerError() *GetUnmanagedClustersInternalServerError {

	return &GetUnmanagedClustersInternalServerError{}
}

// WithPayload adds the payload to the get unmanaged clusters internal server error response
func (o *GetUnmanagedClustersInternalServerError) WithPayload(payload *models.Error) *GetUnmanagedClustersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get unmanaged clusters internal server error response
func (o *GetUnmanagedClustersInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUnmanagedClustersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}