package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Person struct {
	Age int `form:"age" binding:"required,gt=10"`
	Name string `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"8"`
}

func main() {
	r := gin.Default()
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		//txt := fmt.Printf()
		c.String(200, "id : %s", id)
	})

	r.GET("/users/list", func(c *gin.Context) {
		id := c.Param("id")
		//txt := fmt.Printf()
		c.String(200, "id : %s", id)
	})

	r.GET("/5lmh", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil{
			c.String(500, fmt.Sprint(err))
			return
		}
		c.String(200, fmt.Sprint("%#v", person))
	})

	r.Run(":8080")
}
