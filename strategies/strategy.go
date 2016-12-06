package strategies

import (
    "github.com/zalando/skipper/filters"
)

type Strategy interface {
    Authorized(filters.FilterContext)
    Unauthorized(filters.FilterContext)
}
