package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

var users = []User{
	{ID: 1,Name: "张三"},
	{ID: 2,Name: "李四"},
	{ID: 3,Name: "王五"},
}

func main() {
	r := gin.Default()
	r.GET("/users", listUser)
	r.GET("/users/:id", getUser)
	r.POST("/users", createUser)
	r.Run()
}

func listUser(c *gin.Context){
	c.JSON(200, users)
}

func getUser(c *gin.Context){
	id := c.Param("id")
	var user User

	found := false

	for _, u := range users {
		if strings.EqualFold(id , strconv.Itoa(u.ID)) {
			user = u
			found = true
			break
		}
	}

	if found{
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{
			"message": "用户不存在",
		})
	}

}

func createUser(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	if name != ""{
		u := User{ID: len(users) + 1,Name: name}
		users = append(users, u)
		c.JSON(http.StatusCreated, u)
	}else {
		c.JSON(http.StatusOK, gin.H{
			"massage": "请输入用户名",
		})
	}
}




func test(){
	http.HandleFunc("/users", GetUsers)
	http.ListenAndServe(":8880", nil)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := json.Marshal(users)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "{\"msg\": \"" + err.Error() + "\"}")
		}else{
			w.WriteHeader(http.StatusOK)
			w.Write(users)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "not found")
	}
}

type User struct {
	ID int
	Name string
}
