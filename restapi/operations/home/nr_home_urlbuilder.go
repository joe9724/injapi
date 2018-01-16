// Code generated by go-swagger; DO NOT EDIT.

package home

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// NrHomeURL generates an URL for the home operation
type NrHomeURL struct {
	Client      *string
	Imei        *string
	IsRecommend *int64
	MemberID    *string
	PosID       *string
	Ts          *int64
	Type        *int64
	Val         *string
	Version     *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrHomeURL) WithBasePath(bp string) *NrHomeURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrHomeURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *NrHomeURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/home"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/joe9724/jinlinonline/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var client string
	if o.Client != nil {
		client = *o.Client
	}
	if client != "" {
		qs.Set("client", client)
	}

	var imei string
	if o.Imei != nil {
		imei = *o.Imei
	}
	if imei != "" {
		qs.Set("imei", imei)
	}

	var isRecommend string
	if o.IsRecommend != nil {
		isRecommend = swag.FormatInt64(*o.IsRecommend)
	}
	if isRecommend != "" {
		qs.Set("isRecommend", isRecommend)
	}

	var memberID string
	if o.MemberID != nil {
		memberID = *o.MemberID
	}
	if memberID != "" {
		qs.Set("memberId", memberID)
	}

	var posID string
	if o.PosID != nil {
		posID = *o.PosID
	}
	if posID != "" {
		qs.Set("posId", posID)
	}

	var ts string
	if o.Ts != nil {
		ts = swag.FormatInt64(*o.Ts)
	}
	if ts != "" {
		qs.Set("ts", ts)
	}

	var typeVar string
	if o.Type != nil {
		typeVar = swag.FormatInt64(*o.Type)
	}
	if typeVar != "" {
		qs.Set("type", typeVar)
	}

	var val string
	if o.Val != nil {
		val = *o.Val
	}
	if val != "" {
		qs.Set("val", val)
	}

	var version string
	if o.Version != nil {
		version = *o.Version
	}
	if version != "" {
		qs.Set("version", version)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *NrHomeURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *NrHomeURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *NrHomeURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on NrHomeURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on NrHomeURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *NrHomeURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}