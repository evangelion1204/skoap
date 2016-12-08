package auth

import (
	"github.com/zalando/skipper/filters"
	"net/http"
)

type (
	passthrough struct{}
)

func (strategy passthrough) Unauthorized(ctx filters.FilterContext) {
	ctx.Serve(&http.Response{StatusCode: http.StatusUnauthorized})
}

func (strategy passthrough) Authorized(ctx filters.FilterContext) {}

func NewPassthrough() Strategy {
	return &passthrough{}
}
