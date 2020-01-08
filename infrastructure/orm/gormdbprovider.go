package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type OrmConfiguration struct {
	ConnectionString string
	Dialect          string
}

func ProvideGormDB(config OrmConfiguration) (*gorm.DB, error) {
	return gorm.Open(config.Dialect, config.ConnectionString)
}
