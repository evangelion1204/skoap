package auth

import (
	"github.com/zalando/skipper/filters"
	"net/http"
)

type (
	specReject struct{}

	filterReject struct{}
)

const (
	AuthRejectName = "authReject"
)

func NewAuthReject() filters.Spec {
	s := &specReject{}

	return s
}

func (s *specReject) Name() string {
	return AuthRejectName
}

func (s *specReject) CreateFilter(args []interface{}) (filters.Filter, error) {
	f := &filterReject{}

	return f, nil

}

func (f *filterReject) Request(ctx filters.FilterContext) {
	reason, reasonSet := ctx.StateBag()["auth-reject-reason"].(string)

	if !reasonSet && reason == "" {
		return
	}

	ctx.Serve(&http.Response{StatusCode: http.StatusUnauthorized})
}

func (f *filterReject) Response(ctx filters.FilterContext) {}
