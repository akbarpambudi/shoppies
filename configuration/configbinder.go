package configuration

type ConfigBinder interface {
	Bind() (err error)
	GetAppConfig() (*AppConfig, error)
}
