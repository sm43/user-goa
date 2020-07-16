// Code generated by goa v3.2.0, DO NOT EDIT.
//
// user HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/sm43/user-goa/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	user "github.com/sm43/user-goa/gen/user"
	userviews "github.com/sm43/user-goa/gen/user/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "user" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetResponse returns a decoder for responses returned by the user get
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "get", err)
			}
			p := NewGetUserOK(&body)
			view := "default"
			vres := &userviews.User{Projected: p, View: view}
			if err = userviews.ValidateUser(vres); err != nil {
				return nil, goahttp.ErrValidationError("user", "get", err)
			}
			res := user.NewUser(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "get", resp.StatusCode, string(body))
		}
	}
}

// unmarshalCompanyResponseBodyToUserviewsCompanyView builds a value of type
// *userviews.CompanyView from a value of type *CompanyResponseBody.
func unmarshalCompanyResponseBodyToUserviewsCompanyView(v *CompanyResponseBody) *userviews.CompanyView {
	res := &userviews.CompanyView{
		ID:       v.ID,
		Name:     v.Name,
		Location: v.Location,
	}

	return res
}
