package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dictuantranit/go-ecommerce-backend-api/global"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/errors"
)

func InitMySqlC() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := sql.Open("mysql", s)
	errors.Must(global.Logger, err, "InitMysql initialization error")
	global.Logger.Info("Initializing MySQL Successfully sql")
	global.Mdbc = db

	// set Pool
	SetPool()
}

func SetPoolC() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error: %s::", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}
 
