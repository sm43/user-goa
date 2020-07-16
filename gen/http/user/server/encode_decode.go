// Code generated by goa v3.2.0, DO NOT EDIT.
//
// user HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/sm43/user-goa/design

package server

import (
	"context"
	"net/http"

	userviews "github.com/sm43/user-goa/gen/user/views"
	goahttp "goa.design/goa/v3/http"
)

// EncodeGetResponse returns an encoder for responses returned by the user get
// endpoint.
func EncodeGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*userviews.User)
		enc := encoder(ctx, w)
		body := NewGetResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// marshalUserviewsCompanyViewToCompanyResponseBodyTiny builds a value of type
// *CompanyResponseBodyTiny from a value of type *userviews.CompanyView.
func marshalUserviewsCompanyViewToCompanyResponseBodyTiny(v *userviews.CompanyView) *CompanyResponseBodyTiny {
	res := &CompanyResponseBodyTiny{
		Name: *v.Name,
	}

	return res
}
