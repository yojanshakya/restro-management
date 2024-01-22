package bootstrap

import (
    "Restro/domain"
    "Restro/pkg"
    "Restro/seeds"

    "go.uber.org/fx"
)

var CommonModules = fx.Module("common",
    fx.Options(
        pkg.Module,
        seeds.Module,
        domain.Module,
    ),
)
