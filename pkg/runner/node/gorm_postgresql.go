package node

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormPostgresqlConfig struct {
	Host     string
	Port     uint
	Username string
	Password string
	Database string
}

type GormPostgresql struct {
	*gorm.DB
	Config GormPostgresqlConfig
	Models []interface{}
}

func NewGormPostgresql(config GormPostgresqlConfig) *GormPostgresql {
	return &GormPostgresql{
		Config: config,
	}
}

func (g *GormPostgresql) Name() string {
	return "gorm.mysql"
}

func (g *GormPostgresql) Run() error {
	dsn := gormPostgresqlMakeDSN(g.Config)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	g.DB = db

	return g.DB.AutoMigrate(g.Models...)

}

func (g *GormPostgresql) Close() error {
	sqlDB, err := g.DB.DB()

	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (g *GormPostgresql) AddModelForAutoMigration(model interface{}) {
	g.Models = append(g.Models, model)
}

func gormPostgresqlMakeDSN(config GormPostgresqlConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		config.Host,
		config.Username,
		config.Password,
		config.Database,
		config.Port,
	)
}
