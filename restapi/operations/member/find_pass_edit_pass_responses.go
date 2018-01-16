// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injapi/models"
)

// FindPassEditPassOKCode is the HTTP code returned for type FindPassEditPassOK
const FindPassEditPassOKCode int = 200

/*FindPassEditPassOK 登录成功，返回登录信息

swagger:response findPassEditPassOK
*/
type FindPassEditPassOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2003 `json:"body,omitempty"`
}

// NewFindPassEditPassOK creates FindPassEditPassOK with default headers values
func NewFindPassEditPassOK() *FindPassEditPassOK {
	return &FindPassEditPassOK{}
}

// WithPayload adds the payload to the find pass edit pass o k response
func (o *FindPassEditPassOK) WithPayload(payload *models.InlineResponse2003) *FindPassEditPassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find pass edit pass o k response
func (o *FindPassEditPassOK) SetPayload(payload *models.InlineResponse2003) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindPassEditPassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FindPassEditPassDefault error

swagger:response findPassEditPassDefault
*/
type FindPassEditPassDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindPassEditPassDefault creates FindPassEditPassDefault with default headers values
func NewFindPassEditPassDefault(code int) *FindPassEditPassDefault {
	if code <= 0 {
		code = 500
	}

	return &FindPassEditPassDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find pass edit pass default response
func (o *FindPassEditPassDefault) WithStatusCode(code int) *FindPassEditPassDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find pass edit pass default response
func (o *FindPassEditPassDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find pass edit pass default response
func (o *FindPassEditPassDefault) WithPayload(payload *models.Error) *FindPassEditPassDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find pass edit pass default response
func (o *FindPassEditPassDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindPassEditPassDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
