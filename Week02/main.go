package main

import (
	"Go-000/Week02/service"
	"database/sql"
	"github.com/pkg/errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var db *sql.DB

	defer func() {
		db.Close()

		r := recover()

		if r != nil {
			var rerr error
			switch e := r.(type) {
			case string:
				rerr = errors.New(e)
			case error:
				rerr = e
			default:
				rerr = errors.New(fmt.Sprintf("%v", e))
			}

			logErr(rerr)
		}
	} ()

	db, err :=sql.Open("mysql","x:x@tcp(x.x.x.x:3306)/test?charset=utf8")

	if err != nil {
		logErr(err)
		return
	}

	userService := service.NewUserService(db)
	user, err := userService.GetByName("name11")

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			//todo 没有找到的处理
		} else {
			//todo 其他错误的处理
		}

		logErr(err)
		return
	}

	fmt.Println("查找成功")
	fmt.Println(user)
}


func logErr(err error) {
	fmt.Printf("%+v\n", err)
}
