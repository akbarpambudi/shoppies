//+build wireinject

package injector

import (
	"github.com/akbarpambudi/shoppies/configuration"
	"github.com/google/wire"
)

func InitializeAppConfigBinder() configuration.ConfigBinder {
	panic(wire.Build(wire.Value(configuration.EnvironmentConfigBinderProperties{
		FileName: "app-config",
		Path:     "./env",
	}),
		configuration.NewEnvironmentConfigBinder,
		wire.Bind(new(configuration.ConfigBinder),
			new(*configuration.EnvironmentConfigBinder))),
	)
}
