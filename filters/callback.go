package filters

import (
    "github.com/zalando/skipper/filters"
    "github.com/zalando-incubator/skoap/strategies"

)

const (
    CallbackName = "callback"
)

type (
	spec struct {
		strategy   strategies.Strategy
	}

	filter struct {
		strategy   strategies.Strategy
	}
)

func NewCallback(strategy strategies.Strategy) filters.Spec {
	s := &spec{strategy: strategy}

	return s
}

func (s *spec) Name() string {
    return CallbackName
}

func (s *spec) CreateFilter(args []interface{}) (filters.Filter, error) {
	f := &filter{strategy: s.strategy}

	return f, nil

}

func (f *filter) Request(ctx filters.FilterContext) {
	// request := ctx.Request()
}

func (f *filter) Response(_ filters.FilterContext) {}
