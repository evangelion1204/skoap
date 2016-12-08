package auth

import (
    "net/http"
    "github.com/zalando/skipper/filters"
)

type (
    authenticate struct {
        authUrl string
    }
)

func (strategy authenticate) Unauthorized(ctx filters.FilterContext) {
    ctx.Serve(&http.Response{
		StatusCode: http.StatusFound,
		Header:     http.Header{"Location": []string{strategy.authUrl}}})
}

func (strategy authenticate) Authorized(ctx filters.FilterContext) {}

func NewAuthenticate(authUrl string) Strategy {
    return &authenticate{authUrl: authUrl}
}
