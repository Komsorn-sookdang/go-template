package repository

import "go.uber.org/fx"

var Module = fx.Module("RepositoryModule",
	fx.Provide(NewAssessmentRepository),
)
