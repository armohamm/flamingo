package templatefunctions

import (
	"context"

	"flamingo.me/flamingo/v3/core/csrfPreventionFilter/application"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/session"
)

type (
	CsrfTokenFunc struct {
		service application.Service
		logger  flamingo.Logger
	}
)

func (f *CsrfTokenFunc) Inject(s application.Service, l flamingo.Logger) {
	f.service = s
	f.logger = l
}

func (f *CsrfTokenFunc) Func(ctx context.Context) interface{} {
	return func() interface{} {
		s, ok := session.FromContext(ctx)
		if !ok {
			f.logger.WithField("csrf", "templateFunc").Error("can't find session")
			return ""
		}

		return f.service.Generate(s)
	}
}
