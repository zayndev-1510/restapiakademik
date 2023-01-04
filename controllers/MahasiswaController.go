package controllers

import (
	"github.com/gin-gonic/gin"
	"restapi.com/akademik/api/models"
)


var mahasiswa models.MahasiswaModel //instalasi model mahasiswa


// function mendapatkan data mahasiswa
var GetDataMahasiswa=func(c *gin.Context){

	data,db:=models.GetDataMahasiswa()
	if db.Error !=nil{
		c.JSON(500,gin.H{"message":"Something error in "+db.Error.Error(),"status":0})
		return
	}
	if db.RowsAffected==0{
		c.JSON(200,gin.H{"message":"Data Is Empty ","status":0,"data":data})
		return
	}
	c.JSON(200,gin.H{"message":"Data not Empty ","status":1,"data":data})
}

// function menyimpan data mahasiswa

var SaveDataMahasiswa=func(c *gin.Context){
	
	errjson:=c.ShouldBindJSON(&mahasiswa)
	if errjson !=nil{
		c.JSON(500,gin.H{"message":"Something error in "+errjson.Error(),"status":0})
		return
	}
	check:=models.SaveDataMahasiswa(mahasiswa)
	if check.Error !=nil{
		c.JSON(500,gin.H{"message":"Something error in "+check.Error.Error(),"status":0})
		return
	}
	if check.RowsAffected==0{
		c.JSON(200,gin.H{"message":"Save Data Failed","status":0})
		return
	}

	c.JSON(200,gin.H{"message":"Save Data Success","status":1})
}
// function memperbarui data mahasiswa
var UpdateDataMahasiswa=func(c *gin.Context){
	errjson:=c.ShouldBindJSON(&mahasiswa)
	if errjson !=nil{
		c.JSON(500,gin.H{"message":"Something error in "+errjson.Error(),"status":0})
		return
	}
	check:=models.UpdateDataMahasiswa(mahasiswa.Nim,mahasiswa)

	if check.Error !=nil{
		c.JSON(500,gin.H{"message":"Something error in "+check.Error.Error(),"status":0})
		return
	}
	if check.RowsAffected==0{
		c.JSON(200,gin.H{"message":"Not data to update","status":0})
		return
	}

	c.JSON(200,gin.H{"message":"Update Data Success","status":1})
}

// function menghapus data mahasiswa
var DeleteDataMahasiswa=func(c *gin.Context){
	mahasiswa.Nim=c.Param("nim")
	check:=models.DeleteDataMahasiswa(mahasiswa)

	if check.Error !=nil{
		c.JSON(500,gin.H{"message":"Something error in "+check.Error.Error(),"status":0})
		return
	}
	if check.RowsAffected==0{
		c.JSON(200,gin.H{"message":"Delete Data Failed","status":0})
		return
	}
	c.JSON(200,gin.H{"message":"Delete Data Success","status":1})
}

// function mencari data mahasiswa

var SearchDataMahasiswa=func(c *gin.Context){
	mahasiswa.Nim=c.Param("nim")
	data,check:=models.SearchDataMahasiswa(mahasiswa.Nim)

	if check.Error !=nil{
		c.JSON(500,gin.H{"message":"Something error in "+check.Error.Error(),"status":0})
		return
	}
	if check.RowsAffected==0{
		c.JSON(200,gin.H{"message":"Data Is Empty","status":0})
		return
	}
	c.JSON(200,gin.H{"message":"Data Is Empty","status":1,"data":data})
}
