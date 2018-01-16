// Code generated by go-swagger; DO NOT EDIT.

package banner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrBannerListHandlerFunc turns a function with the right signature into a banner list handler
type NrBannerListHandlerFunc func(NrBannerListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrBannerListHandlerFunc) Handle(params NrBannerListParams) middleware.Responder {
	return fn(params)
}

// NrBannerListHandler interface for that can handle valid banner list params
type NrBannerListHandler interface {
	Handle(NrBannerListParams) middleware.Responder
}

// NewNrBannerList creates a new http.Handler for the banner list operation
func NewNrBannerList(ctx *middleware.Context, handler NrBannerListHandler) *NrBannerList {
	return &NrBannerList{Context: ctx, Handler: handler}
}

/*NrBannerList swagger:route GET /banner/list Banner bannerList

Banner列表

Banner列表

*/
type NrBannerList struct {
	Context *middleware.Context
	Handler NrBannerListHandler
}

func (o *NrBannerList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrBannerListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}