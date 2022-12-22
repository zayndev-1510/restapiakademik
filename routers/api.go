package routers

import (
	"github.com/gin-gonic/gin"
	"restapi.com/akademik/api/controllers"
)

func RouterMahasiswa(router *gin.Engine){

	grup:=router.Group("mahasiswa")
	grup.GET("get-data",controllers.GetDataMahasiswa)
	grup.POST("save-data",controllers.SaveDataMahasiswa)
	grup.POST("update-data",controllers.UpdateDataMahasiswa)
	grup.GET("search-data/:nim",controllers.SearchDataMahasiswa)
	grup.DELETE("delete-data/:nim",controllers.DeleteDataMahasiswa)
}