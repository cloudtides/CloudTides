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

// GetResourceUsageHandlerFunc turns a function with the right signature into a get resource usage handler
type GetResourceUsageHandlerFunc func(GetResourceUsageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetResourceUsageHandlerFunc) Handle(params GetResourceUsageParams) middleware.Responder {
	return fn(params)
}

// GetResourceUsageHandler interface for that can handle valid get resource usage params
type GetResourceUsageHandler interface {
	Handle(GetResourceUsageParams) middleware.Responder
}

// NewGetResourceUsage creates a new http.Handler for the get resource usage operation
func NewGetResourceUsage(ctx *middleware.Context, handler GetResourceUsageHandler) *GetResourceUsage {
	return &GetResourceUsage{Context: ctx, Handler: handler}
}

/* GetResourceUsage swagger:route GET /usage/{id} usage getResourceUsage

get resource usage

*/
type GetResourceUsage struct {
	Context *middleware.Context
	Handler GetResourceUsageHandler
}

func (o *GetResourceUsage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetResourceUsageParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetResourceUsageNotFoundBody get resource usage not found body
//
// swagger:model GetResourceUsageNotFoundBody
type GetResourceUsageNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get resource usage not found body
func (o *GetResourceUsageNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get resource usage not found body based on context it is used
func (o *GetResourceUsageNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetResourceUsageNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourceUsageNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetResourceUsageNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetResourceUsageOKBody get resource usage o k body
//
// swagger:model GetResourceUsageOKBody
type GetResourceUsageOKBody struct {

	// current CPU
	CurrentCPU float64 `json:"currentCPU,omitempty"`

	// current disk
	CurrentDisk float64 `json:"currentDisk,omitempty"`

	// current RAM
	CurrentRAM float64 `json:"currentRAM,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// percent CPU
	PercentCPU float64 `json:"percentCPU,omitempty"`

	// percent disk
	PercentDisk float64 `json:"percentDisk,omitempty"`

	// percent RAM
	PercentRAM float64 `json:"percentRAM,omitempty"`

	// total CPU
	TotalCPU float64 `json:"totalCPU,omitempty"`

	// total disk
	TotalDisk float64 `json:"totalDisk,omitempty"`

	// total RAM
	TotalRAM float64 `json:"totalRAM,omitempty"`
}

// Validate validates this get resource usage o k body
func (o *GetResourceUsageOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get resource usage o k body based on context it is used
func (o *GetResourceUsageOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetResourceUsageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourceUsageOKBody) UnmarshalBinary(b []byte) error {
	var res GetResourceUsageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
