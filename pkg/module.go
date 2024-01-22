package pkg

import (
	"Restro/pkg/framework"
	"Restro/pkg/infrastructure"
	"Restro/pkg/middlewares"
	"Restro/pkg/services"

	"go.uber.org/fx"
)

var Module = fx.Module("pkg",
	framework.Module,
	services.Module,
	middlewares.Module,
	infrastructure.Module,
)
