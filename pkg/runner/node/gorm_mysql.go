package node

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormMysqlConfig struct {
	Host     string
	Port     uint
	Username string
	Password string
	Database string
}

type GormMysql struct {
	*gorm.DB
	Config GormMysqlConfig
	Models []interface{}
}

func NewGormMysql(config GormMysqlConfig) *GormMysql {
	return &GormMysql{
		Config: config,
	}
}

func (g *GormMysql) Name() string {
	return "gorm.mysql"
}

func (g *GormMysql) Run() error {
	dsn := gormMysqlMakeDSN(g.Config)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	g.DB = db

	return g.DB.AutoMigrate(g.Models...)

}

func (g *GormMysql) Close() error {
	sqlDB, err := g.DB.DB()

	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (g *GormMysql) AddModelForAutoMigration(model interface{}) {
	g.Models = append(g.Models, model)
}

func gormMysqlMakeDSN(config GormMysqlConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}
