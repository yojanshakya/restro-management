package features

import (
	"Restro/domain/features/hello"
	"Restro/domain/features/menu"

	"go.uber.org/fx"
)

var Module = fx.Module("features",
  fx.Options(hello.Module),
  fx.Options(menu.Module),
)
