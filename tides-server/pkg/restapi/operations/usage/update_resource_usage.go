// Code generated by go-swagger; DO NOT EDIT.

package usage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateResourceUsageHandlerFunc turns a function with the right signature into a update resource usage handler
type UpdateResourceUsageHandlerFunc func(UpdateResourceUsageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateResourceUsageHandlerFunc) Handle(params UpdateResourceUsageParams) middleware.Responder {
	return fn(params)
}

// UpdateResourceUsageHandler interface for that can handle valid update resource usage params
type UpdateResourceUsageHandler interface {
	Handle(UpdateResourceUsageParams) middleware.Responder
}

// NewUpdateResourceUsage creates a new http.Handler for the update resource usage operation
func NewUpdateResourceUsage(ctx *middleware.Context, handler UpdateResourceUsageHandler) *UpdateResourceUsage {
	return &UpdateResourceUsage{Context: ctx, Handler: handler}
}

/* UpdateResourceUsage swagger:route PUT /usage/{id} usage updateResourceUsage

update datacenter usage info

*/
type UpdateResourceUsage struct {
	Context *middleware.Context
	Handler UpdateResourceUsageHandler
}

func (o *UpdateResourceUsage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateResourceUsageParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateResourceUsageBody update resource usage body
//
// swagger:model UpdateResourceUsageBody
type UpdateResourceUsageBody struct {

	// current CPU
	CurrentCPU float64 `json:"currentCPU,omitempty"`

	// current disk
	CurrentDisk float64 `json:"currentDisk,omitempty"`

	// current RAM
	CurrentRAM float64 `json:"currentRAM,omitempty"`

	// total CPU
	TotalCPU float64 `json:"totalCPU,omitempty"`

	// total disk
	TotalDisk float64 `json:"totalDisk,omitempty"`

	// total RAM
	TotalRAM float64 `json:"totalRAM,omitempty"`
}

// Validate validates this update resource usage body
func (o *UpdateResourceUsageBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update resource usage body based on context it is used
func (o *UpdateResourceUsageBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateResourceUsageBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateResourceUsageBody) UnmarshalBinary(b []byte) error {
	var res UpdateResourceUsageBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateResourceUsageOKBody update resource usage o k body
//
// swagger:model UpdateResourceUsageOKBody
type UpdateResourceUsageOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this update resource usage o k body
func (o *UpdateResourceUsageOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update resource usage o k body based on context it is used
func (o *UpdateResourceUsageOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateResourceUsageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateResourceUsageOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateResourceUsageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
