package domain

import (
    "Restro/domain/features"
    "Restro/domain/middlewares"

    "go.uber.org/fx"
)

var Module = fx.Options(
	middlewares.Module,
	features.Module,
)
