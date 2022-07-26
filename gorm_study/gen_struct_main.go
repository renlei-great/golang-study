package main

import (
	"fmt"
	"github.com/gohouse/converter"
)

func main() {
	err := converter.NewTable2Struct().
		SavePath("./model.go").
		Dsn("gaea:T49QbHnncFJ2AI8@tcp(35.222.2.35:3306)/matomo?charset=utf8").
		TagKey("gorm").
		EnableJsonTag(true).
		Table("matomo_log_action").
		Run()
	fmt.Println(err)
}
