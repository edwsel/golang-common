package node

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormConfig struct {
	Host     string
	Port     uint
	Username string
	Password string
	Database string
}

type Gorm struct {
	*gorm.DB
	Config GormConfig
	Models []interface{}
}

func NewGorm(config GormConfig) *Gorm {
	return &Gorm{
		Config: config,
	}
}

func (g *Gorm) Name() string {
	return "gorm.mysql"
}

func (g *Gorm) Run() error {
	dsn := gormMakeDSN(g.Config)

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

func (g *Gorm) Close() error {
	sqlDB, err := g.DB.DB()

	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (g *Gorm) AddModelForAutoMigration(model interface{}) {
	g.Models = append(g.Models, model)
}

func gormMakeDSN(config GormConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}
