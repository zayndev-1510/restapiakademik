package models

import (
	"gorm.io/gorm"
	"restapi.com/akademik/api/config"
)
var db *gorm.DB
type MahasiswaModel struct{
	Nim string `gorm:"primaryKey" json:"nim"`
	Nama string `gorm:"type:varchar(50)" json:"nama"`
	Kelas string `gorm:"type:varchar(2)" json:"kelas"`
}


func init(){
	db=config.ConnectDB()
	db.Table("tbl_mahasiswa").AutoMigrate(&MahasiswaModel{})
}
func GetDataMahasiswa() ([] MahasiswaModel,*gorm.DB){
	var result []MahasiswaModel
	query:=db.Table("tbl_mahasiswa").Find(&result)
	return  result,query
}

func SaveDataMahasiswa(row MahasiswaModel) *gorm.DB {
	query:=db.Table("tbl_mahasiswa").Create(row)
	return query
}

func UpdateDataMahasiswa(nim string,row MahasiswaModel) *gorm.DB {
	query:=db.Table("tbl_mahasiswa").Where("nim=?",nim).Updates(row)
	return query
}

func DeleteDataMahasiswa(row MahasiswaModel) *gorm.DB {
	query:=db.Table("tbl_mahasiswa").Delete(row)
	return query
}

func SearchDataMahasiswa(nim string) (MahasiswaModel,*gorm.DB){

	var result MahasiswaModel
	query:=db.Table("tbl_mahasiswa").Where("nim=?",nim).Find(&result)
	return result,query
}
