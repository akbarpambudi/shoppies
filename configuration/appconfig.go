package configuration

type AppConfig struct {
	Http HttpConfig
}

type HttpConfig struct {
	Address string
}
