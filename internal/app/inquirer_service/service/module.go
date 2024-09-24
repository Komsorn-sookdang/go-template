package service

import (
	"github.com/Komsorn-sookdang/go-template/internal/app/inquirer_service/service/api"
	"go.uber.org/fx"
)

var Module = fx.Module("ServiceModule",
	api.Module,
)
