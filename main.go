package main

import (
	"github.com/gin-gonic/gin"
	"restapi.com/akademik/api/routers"
)

func main(){
	c:=gin.Default()
	routers.RouterMahasiswa(c)
	c.Run()

}