package auth

import (
    "net/http"
    "github.com/zalando/skipper/filters"
)

type (
    passthrough struct {}
)

func (strategy passthrough) Unauthorized(ctx filters.FilterContext) {
    ctx.Serve(&http.Response{StatusCode: http.StatusUnauthorized})
}

func (strategy passthrough) Authorized(ctx filters.FilterContext) {}

func NewPassthrough() Strategy {
    return &passthrough{}
}
