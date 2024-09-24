package assessment

import "go.uber.org/fx"

var Module = fx.Module("AssessmentEngineModule",
	fx.Provide(NewAssessmentEngine),
)
