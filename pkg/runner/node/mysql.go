package node

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	Host     string
	Port     uint
	Username string
	Password string
	Database string
}

type Mysql struct {
	Config MysqlConfig
	*sql.DB
}

func NewMysql(config MysqlConfig) *Mysql {
	return &Mysql{Config: config}
}

func (m *Mysql) Name() string {
	return "mysql"
}

func (m *Mysql) Run() error {
	db, err := sql.Open("mysql", makeDSN(m.Config))

	if err != nil {
		return err
	}

	m.DB = db

	return nil
}

func (m *Mysql) Close() error {
	return m.DB.Close()
}

func makeDSN(config MysqlConfig) string {
	return fmt.Sprintf(
		"%s:%s@%s:%d/%s?parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}
