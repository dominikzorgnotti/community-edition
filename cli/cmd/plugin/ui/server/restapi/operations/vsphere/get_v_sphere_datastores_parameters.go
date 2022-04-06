// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetVSphereDatastoresParams creates a new GetVSphereDatastoresParams object
//
// There are no default values defined in the spec.
func NewGetVSphereDatastoresParams() GetVSphereDatastoresParams {

	return GetVSphereDatastoresParams{}
}

// GetVSphereDatastoresParams contains all the bound params for the get v sphere datastores operation
// typically these are obtained from a http.Request
//
// swagger:parameters getVSphereDatastores
type GetVSphereDatastoresParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*datacenter managed object Id, e.g. datacenter-2
	  Required: true
	  In: query
	*/
	Dc string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetVSphereDatastoresParams() beforehand.
func (o *GetVSphereDatastoresParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDc, qhkDc, _ := qs.GetOK("dc")
	if err := o.bindDc(qDc, qhkDc, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDc binds and validates parameter Dc from query.
func (o *GetVSphereDatastoresParams) bindDc(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("dc", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("dc", "query", raw); err != nil {
		return err
	}
	o.Dc = raw

	return nil
}