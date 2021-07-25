package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func main() {
	id, err := DB()
	if errors.Is(err, sql.ErrNoRows) {
		//TODO
	}
	fmt.Printf("%+v \n", err)
	fmt.Println(id)
}

func Conn() (*sql.DB, error) {

	return sql.Open("mysql",
		"root:root.@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")

}

func DB() (id int, err error) {

	db, err := Conn()

	if err != nil {
		return 0, fmt.Errorf("db conn err :%w", err)
	}
	defer db.Close()
	err = db.QueryRow("select id from admin where id = ? ", 6).Scan(&id)
	//经常数据库查询只是用于判断查询是否存在，此时 sql.ErrNoRows 并不认为是一个错误信息
	//此时将 sql.ErrNoRows wrap处理可能会影响错误日志查看
	//我这里的处理方式是：将 sql.ErrNoRows 的原始错误直接返回给调用者，由调用者进行空处理
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, errors.Wrap(err, "db query err ")
	}
	return id, err

}
