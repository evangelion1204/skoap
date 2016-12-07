package auth

import (
	"github.com/zalando/skipper/filters"
	"net/http"
)

const (
	AuthStorageName = "authStorage"
	AuthStorageCookieName = "auth"
	AuthorizationPrefix = "Bearer "
)

type (
	spec struct {
	}

	filter struct {
	}
)

func NewAuthStorage() filters.Spec {
	s := &spec{}

	return s
}

func (s *spec) Name() string {
	return AuthStorageName
}

func (s *spec) CreateFilter(args []interface{}) (filters.Filter, error) {
	f := &filter{}

	return f, nil

}

func (f *filter) Request(ctx filters.FilterContext) {
	request := ctx.Request()
	authCookie, err := request.Cookie(AuthStorageCookieName)

	if err == nil {
		request.Header.Set("Authorization", AuthorizationPrefix + authCookie.Value)
	}
}

func (f *filter) Response(ctx filters.FilterContext) {
	newAuth := ctx.Response().Header.Get("Authorization")

	if newAuth != ""  {
		extractedToken := newAuth[len(AuthorizationPrefix):]
		cookie := &http.Cookie{
			Name:     AuthStorageCookieName,
			Value:    extractedToken,
			HttpOnly: true,
			Secure:   true,
			Domain:   "",
			Path:     "/",
			MaxAge:   0}

		ctx.Response().Header.Add("Set-Cookie", cookie.String())
	}

}
