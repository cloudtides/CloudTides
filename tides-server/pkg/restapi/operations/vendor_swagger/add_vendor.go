// Code generated by go-swagger; DO NOT EDIT.

package vendor_swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddVendorHandlerFunc turns a function with the right signature into a add vendor handler
type AddVendorHandlerFunc func(AddVendorParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddVendorHandlerFunc) Handle(params AddVendorParams) middleware.Responder {
	return fn(params)
}

// AddVendorHandler interface for that can handle valid add vendor params
type AddVendorHandler interface {
	Handle(AddVendorParams) middleware.Responder
}

// NewAddVendor creates a new http.Handler for the add vendor operation
func NewAddVendor(ctx *middleware.Context, handler AddVendorHandler) *AddVendor {
	return &AddVendor{Context: ctx, Handler: handler}
}

/* AddVendor swagger:route POST /vendors vendor addVendor

add vendor

*/
type AddVendor struct {
	Context *middleware.Context
	Handler AddVendorHandler
}

func (o *AddVendor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddVendorParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddVendorBody add vendor body
//
// swagger:model AddVendorBody
type AddVendorBody struct {

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// vendor type
	VendorType string `json:"vendorType,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this add vendor body
func (o *AddVendorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add vendor body based on context it is used
func (o *AddVendorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddVendorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddVendorBody) UnmarshalBinary(b []byte) error {
	var res AddVendorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddVendorOKBody add vendor o k body
//
// swagger:model AddVendorOKBody
type AddVendorOKBody struct {

	// id
	ID int64 `json:"id,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add vendor o k body
func (o *AddVendorOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add vendor o k body based on context it is used
func (o *AddVendorOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddVendorOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddVendorOKBody) UnmarshalBinary(b []byte) error {
	var res AddVendorOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}