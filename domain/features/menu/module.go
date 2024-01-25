package menu

import "go.uber.org/fx"

var Module = fx.Module("menu",
	fx.Options(
		fx.Provide(NewMenuController),
		fx.Invoke(NewRoute),
	),
)