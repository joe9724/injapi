// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injapi/models"
)

// OrderListOKCode is the HTTP code returned for type OrderListOK
const OrderListOKCode int = 200

/*OrderListOK 获取反馈列表

swagger:response orderListOK
*/
type OrderListOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2005 `json:"body,omitempty"`
}

// NewOrderListOK creates OrderListOK with default headers values
func NewOrderListOK() *OrderListOK {
	return &OrderListOK{}
}

// WithPayload adds the payload to the order list o k response
func (o *OrderListOK) WithPayload(payload *models.InlineResponse2005) *OrderListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the order list o k response
func (o *OrderListOK) SetPayload(payload *models.InlineResponse2005) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *OrderListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*OrderListDefault error

swagger:response orderListDefault
*/
type OrderListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewOrderListDefault creates OrderListDefault with default headers values
func NewOrderListDefault(code int) *OrderListDefault {
	if code <= 0 {
		code = 500
	}

	return &OrderListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the order list default response
func (o *OrderListDefault) WithStatusCode(code int) *OrderListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the order list default response
func (o *OrderListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the order list default response
func (o *OrderListDefault) WithPayload(payload *models.Error) *OrderListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the order list default response
func (o *OrderListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *OrderListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}