package configuration

type AppConfig struct {
	Http HttpConfig
	Db   DbConfig
}

type HttpConfig struct {
	Address string
}

type DbConfig struct {
	Dialect    string
	Connection string
}
