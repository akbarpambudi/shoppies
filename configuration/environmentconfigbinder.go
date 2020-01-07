package configuration

import (
	"github.com/spf13/viper"
)

type EnvironmentConfigBinderProperties struct {
	FileName string
	Path     string
}

type EnvironmentConfigBinder struct {
	fileName string
	path     string
	config   AppConfig
}

func (binder *EnvironmentConfigBinder) Bind() (err error) {
	viper.SetConfigName(binder.fileName)
	viper.AddConfigPath(binder.path)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	viper.SetDefault("http.address", ":8080")
	err = viper.Unmarshal(&binder.config)
	return err
}
func (binder *EnvironmentConfigBinder) GetAppConfig() (*AppConfig, error) {
	return &binder.config, nil
}

func NewEnvironmentConfigBinder(properties EnvironmentConfigBinderProperties) *EnvironmentConfigBinder {
	return &EnvironmentConfigBinder{fileName: properties.FileName, path: properties.Path}
}
