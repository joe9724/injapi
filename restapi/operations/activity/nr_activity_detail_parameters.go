// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNrActivityDetailParams creates a new NrActivityDetailParams object
// with the default values initialized.
func NewNrActivityDetailParams() NrActivityDetailParams {
	var ()
	return NrActivityDetailParams{}
}

// NrActivityDetailParams contains all the bound params for the activity detail operation
// typically these are obtained from a http.Request
//
// swagger:parameters /activity/detail
type NrActivityDetailParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*活动关键词
	  In: query
	*/
	Key *string
	/*会员id
	  In: query
	*/
	MemberID *string
	/*服务状态(付费0，免费1,全部2)
	  In: query
	*/
	ServiceType *int64
	/*状态 0=正常 1锁定
	  In: query
	*/
	Status *int64
	/*title
	  In: query
	*/
	Title *string
	/*活动详情
	  In: query
	*/
	Type *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrActivityDetailParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qKey, qhkKey, _ := qs.GetOK("key")
	if err := o.bindKey(qKey, qhkKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qMemberID, qhkMemberID, _ := qs.GetOK("memberId")
	if err := o.bindMemberID(qMemberID, qhkMemberID, route.Formats); err != nil {
		res = append(res, err)
	}

	qServiceType, qhkServiceType, _ := qs.GetOK("serviceType")
	if err := o.bindServiceType(qServiceType, qhkServiceType, route.Formats); err != nil {
		res = append(res, err)
	}

	qStatus, qhkStatus, _ := qs.GetOK("status")
	if err := o.bindStatus(qStatus, qhkStatus, route.Formats); err != nil {
		res = append(res, err)
	}

	qTitle, qhkTitle, _ := qs.GetOK("title")
	if err := o.bindTitle(qTitle, qhkTitle, route.Formats); err != nil {
		res = append(res, err)
	}

	qType, qhkType, _ := qs.GetOK("type")
	if err := o.bindType(qType, qhkType, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NrActivityDetailParams) bindKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Key = &raw

	return nil
}

func (o *NrActivityDetailParams) bindMemberID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.MemberID = &raw

	return nil
}

func (o *NrActivityDetailParams) bindServiceType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("serviceType", "query", "int64", raw)
	}
	o.ServiceType = &value

	return nil
}

func (o *NrActivityDetailParams) bindStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("status", "query", "int64", raw)
	}
	o.Status = &value

	return nil
}

func (o *NrActivityDetailParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Title = &raw

	return nil
}

func (o *NrActivityDetailParams) bindType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("type", "query", "int64", raw)
	}
	o.Type = &value

	return nil
}
