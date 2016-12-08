package auth

import (
	"github.com/zalando/go-tokens/client"
	"github.com/zalando/skipper/filters"
	"net/http"
	"net/url"
)

type (
	authenticate struct {
		authUrl     string
		realm       string
		redirectUrl string
		provider    client.CredentialsProvider
	}
)

func (strategy authenticate) Unauthorized(ctx filters.FilterContext) {
	credentials, _ := strategy.provider.Get()

	params := url.Values{}

	params.Add("response_type", "code")
	params.Add("client_id", credentials.Id())
	params.Add("realm", strategy.realm)
	params.Add("redirect_uri", strategy.redirectUrl)
	params.Add("state", ctx.OriginalRequest().URL.String())

	ctx.Serve(&http.Response{
		StatusCode: http.StatusFound,
		Header:     http.Header{"Location": []string{strategy.authUrl + "?" + params.Encode()}}})
}

func (strategy authenticate) Authorized(ctx filters.FilterContext) {}

func NewAuthenticate(authUrl string, realm string, redirectUrl string, provider client.CredentialsProvider) Strategy {
	return &authenticate{authUrl: authUrl, realm: realm, redirectUrl: redirectUrl, provider: provider}
}
