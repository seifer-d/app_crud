package main

import (
	"github.com/seifer-d/app_crud/database"
	"github.com/seifer-d/app_crud/models"
)

func main() {
	db := database.GetDatabase()
	alumno := models.Alumno{}
	db.AutoMigrate(&alumno)
}
