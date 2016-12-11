package auth

import (
	"github.com/zalando/go-tokens/client"
	"github.com/zalando/skipper/filters"
	"net/http"
	"net/url"
)

type (
	specRedirectProvider struct {
		authUrl     string
		realm       string
		redirectUrl string
		provider    client.CredentialsProvider
	}

	filterRedirectProvider struct {
		authUrl     string
		realm       string
		redirectUrl string
		provider    client.CredentialsProvider
	}
)

const (
	AuthRedirectProviderName = "authRedirectProvider"
)

func NewAuthRedirectProvider(authUrl string, realm string, redirectUrl string, provider client.CredentialsProvider) filters.Spec {
	s := &specRedirectProvider{authUrl: authUrl, realm: realm, redirectUrl: redirectUrl, provider: provider}

	return s
}

func (s *specRedirectProvider) Name() string {
	return AuthRedirectProviderName
}

func (s *specRedirectProvider) CreateFilter(args []interface{}) (filters.Filter, error) {
	f := &filterRedirectProvider{authUrl: s.authUrl, realm: s.realm, redirectUrl: s.redirectUrl, provider: s.provider}

	return f, nil

}

func (f *filterRedirectProvider) Request(ctx filters.FilterContext) {
	reason, reasonSet := ctx.StateBag()["auth-reject-reason"].(string)

	if !reasonSet && reason == "" {
		return
	}

	credentials, _ := f.provider.Get()

	params := url.Values{}

	params.Add("response_type", "code")
	params.Add("client_id", credentials.Id())
	params.Add("realm", f.realm)
	params.Add("redirect_uri", f.redirectUrl)
	params.Add("state", ctx.OriginalRequest().URL.String())

	ctx.Serve(&http.Response{
		StatusCode: http.StatusFound,
		Header:     http.Header{"Location": []string{f.authUrl + "?" + params.Encode()}}})
}

func (f *filterRedirectProvider) Response(ctx filters.FilterContext) {}
