// Code generated by go-swagger; DO NOT EDIT.

package home

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	middleware "github.com/go-openapi/runtime/middleware"
	"injapi/models"
	"fmt"
	"injapi/var"
)

// NrHomeHandlerFunc turns a function with the right signature into a home handler
type NrHomeHandlerFunc func(NrHomeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrHomeHandlerFunc) Handle(params NrHomeParams) middleware.Responder {
	return fn(params)
}

// NrHomeHandler interface for that can handle valid home params
type NrHomeHandler interface {
	Handle(NrHomeParams) middleware.Responder
}

// NewNrHome creates a new http.Handler for the home operation
func NewNrHome(ctx *middleware.Context, handler NrHomeHandler) *NrHome {
	return &NrHome{Context: ctx, Handler: handler}
}

/*NrHome swagger:route GET /home Home home

大首页接口

大首页接口

*/
type NrHome struct {
	Context *middleware.Context
	Handler NrHomeHandler
}

func (o *NrHome) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrHomeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok HomeOK
	var response models.InlineResponse200
	/*var banners models.InlineResponse200Banners
	var icons models.InlineResponse200Icons
	var news models.InlineResponse200News
	var people models.InlineResponse200People*/
	var status models.Response

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}

	var Banners models.InlineResponse200Banners
	db.Table("banners").Where(map[string]interface{}{"status":0}).Limit(5).Find(&Banners)
	response.Banners = Banners

	var Icons models.InlineResponse200Icons
	db.Table("icons").Where(map[string]interface{}{"status":0}).Limit(5).Find(&Icons)
	response.Icons = Icons

	//status

	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
