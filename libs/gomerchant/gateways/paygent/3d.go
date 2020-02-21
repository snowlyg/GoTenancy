package paygent

import (
	"errors"
	"net/http"

	"GoTenancy/libs/gomerchant"
)

type SecureCodeParams struct {
	UserAgent  string
	TermURL    string
	HttpAccept string
}

func (paygent *Paygent) SecureCodeAuthorize(amount uint64, secureCodeParams SecureCodeParams, params gomerchant.AuthorizeParams) (gomerchant.AuthorizeResponse, error) {
	if params.Params == nil {
		params.Params = gomerchant.Params{}
	}
	params.Set("Paygent3DMode", true)
	params.Set("Paygent3DParams", secureCodeParams)

	return paygent.Authorize(amount, params)
}

func (paygent *Paygent) CompleteAuthorize(paymentID string, params gomerchant.CompleteAuthorizeParams) (gomerchant.CompleteAuthorizeResponse, error) {
	if req, ok := params.Get("request"); ok {
		if request, ok := req.(*http.Request); ok {
			request.ParseForm()
			response, err := paygent.Request("024", gomerchant.Params{"MD": request.Form.Get("MD"), "PaRes": request.Form.Get("PaRes")})
			return gomerchant.CompleteAuthorizeResponse{Params: response.Params}, err
		}
	}
	return gomerchant.CompleteAuthorizeResponse{}, errors.New("no valid request params found")
}
