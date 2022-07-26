package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type User struct {
	ID int `gorm:"primaryKey column:id"`
	CreatedAt  time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Created int64 `gorm:"autoCreateTime"`
	Name string `gorm:"column:name"`
}

type Matomo_log_action struct {
	Idaction  int    `gorm:"idaction" json:"idaction"`
	Name      string `gorm:"name" json:"name"`
	Hash      int    `gorm:"hash" json:"hash"`
	Type      int    `gorm:"type" json:"type"`
	UrlPrefix int    `gorm:"url_prefix" json:"url_prefix"`
}

type APIUser struct {
	ID int
	Name string
}

func main() {



}

func conn(){
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3308)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase: true, // skip the snake_casing of names
		},
	})
	if err != nil {
		fmt.Println("创建数据库连接错误:",err)
		return
	}

	//db.Use(
	//	)

	dbSelect(db)
}

func dbSelect(db *gorm.DB){
	//shar
	//var user []User
	//res := db.Find(&user)
	//fmt.Println(res)
	//
	//result := map[string]interface{}{}
	//db.Model(&User{}).First(&result)
	//
	//result1 := map[string]interface{}{}
	//db.Table("test.user").Take(&result1)
	//fmt.Println(result)
	//
	result2 := map[string]interface{}{}
	db.Model(&User{}).First(&result2, 4)
	fmt.Println(result2)
	//
	//var result3  = User{ID: 5}
	//db.First(&result3)
	//fmt.Println(result)
	//
	//var result4 User
	//db.Model(&User{ID: 5}).First(&result4)
	//fmt.Println(result)

	var users []User
	db.Where("name like ?", "%rl%").Find(&users)
	fmt.Println(users)

	var users1 []User
	db.Where(&User{Name: "rl"}).Find(&users1)
	fmt.Println(users)

	//var users2 []interface{}
	row, err :=db.Table("test.User").Select("count(distinct name)").Rows()
	if err != nil {
		fmt.Println("查询出错：",err)
	}
	fmt.Println(row)

	var apiUser []APIUser
	db.Model(&User{}).Limit(5).Find(&apiUser).Table("User")
	fmt.Println(apiUser)

	stmt := db.Session(&gorm.Session{DryRun: true}).Model(&User{}).Limit(5).Find(&apiUser).Statement
	a := stmt.SQL.String()
	fmt.Println("sql: ", stmt.SQL.String(), a)
	fmt.Println(stmt)

	var usersAttr User
	db.Where(User{Name: "rlg"}).Assign("Created", "123").FirstOrInit(&usersAttr)
	fmt.Println(apiUser)

}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if u.Name == "rl" || u.Name == "rlg" {
		u.Name = "钩子函数"
	}
	return
}


func creat(db *gorm.DB){
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	//sqlDB, err := db.DB()
	var users = []User{{Name: "rll1"},{Name: "rll2"},{Name: "rll3"}}

	db.Create(&users)
	//db.Model(&User{}).Create([]map[string]interface{}{
	//	{"Name": "rl"},
	//	{"Name": "rl1"},
	//	{"Name": "rl2"},
	//})
	fmt.Println("user插入数据成功")
}
