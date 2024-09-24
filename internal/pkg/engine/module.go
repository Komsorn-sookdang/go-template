package engine

import (
	"github.com/Komsorn-sookdang/go-template/internal/pkg/engine/assessment"
	"go.uber.org/fx"
)

var Module = fx.Module("EngineModule",
	assessment.Module,
)
