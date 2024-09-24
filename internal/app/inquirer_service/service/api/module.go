package api

import "go.uber.org/fx"

var Module = fx.Module("ApiServiceModule",
	fx.Provide(NewAssessmentService),
)
