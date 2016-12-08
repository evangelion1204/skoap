package callback

import (
	"github.com/zalando-incubator/skoap/filters/auth"
	"github.com/zalando-incubator/skoap/oauth"
	"github.com/zalando/skipper/filters"
	"net/http"
)

const (
	CallbackName = "callback"
)

type (
	spec struct {
		authUrlBase string
		strategy    auth.Strategy
		client      oauth.Client
	}

	filter struct {
		authUrlBase string
		strategy    auth.Strategy
		client      oauth.Client
	}
)

func New(authUrlBase string, strategy auth.Strategy, client oauth.Client) filters.Spec {
	s := &spec{authUrlBase: authUrlBase, strategy: strategy, client: client}

	return s
}

func (s *spec) Name() string {
	return CallbackName
}

func (s *spec) CreateFilter(args []interface{}) (filters.Filter, error) {
	f := &filter{authUrlBase: s.authUrlBase, strategy: s.strategy, client: s.client}

	return f, nil

}

func (f *filter) Request(ctx filters.FilterContext) {
	request := ctx.Request()
	params := request.URL.Query()

	newToken, _ := f.client.GetAccessTokenByCode(params.Get("code"))

	var stateUrl string
	if stateUrl = params.Get("state"); stateUrl == "" {
		stateUrl = "/"
	}

	ctx.Serve(&http.Response{
		StatusCode: http.StatusFound,
		Header:     http.Header{"Location": []string{stateUrl}}})

	ctx.Response().Header.Set("Authorization", "Bearer "+newToken)
}

func (f *filter) Response(_ filters.FilterContext) {}
