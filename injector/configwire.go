//+build wireinject

package injector

import (
	"github.com/akbarpambudi/shoppies/configuration"
	"github.com/google/wire"
)

var (
	configModuleSets = wire.NewSet(wire.Value(configuration.EnvironmentConfigBinderProperties{
		FileName: "app-config",
		Path:     "./env",
	}),
		provideAppConfig)
)

func provideAppConfig(properties configuration.EnvironmentConfigBinderProperties) (*configuration.AppConfig, error) {
	environmentConfigBinder := configuration.NewEnvironmentConfigBinder(properties)
	environmentConfigBinder.Bind()
	return environmentConfigBinder.GetAppConfig()

}

func InitializeAppConfig() (*configuration.AppConfig, error) {
	panic(wire.Build(configModuleSets))
}
