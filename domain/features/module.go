package features

import (
  "Restro/domain/features/hello"

  "go.uber.org/fx"
)

var Module = fx.Module("features",
  fx.Options(hello.Module),
)
