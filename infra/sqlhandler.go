package infra

import (
	"github.com/dionomusuko/todos-go/interfaces/database"
	"github.com/jinzhu/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	conn, err := gorm.Open("mysql", "root")
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
