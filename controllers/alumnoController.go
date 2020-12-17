package controllers

import (
	"fmt"
	"net/http"

	"github.com/seifer-d/app_crud/views"
	"github.com/seifer-d/app_crud/database"
	"github.com/seifer-d/app_crud/models"
	"github.com/kataras/iris/v12"
)

func AllAlumnos(ctx iris.Context) {
	alumnos := []models.Alumno{}
	db := database.GetDatabase()
	db.Find(&alumnos)
	ctx.ViewData("alumnos", alumnos)
	ctx.View("list.html")
}

func CreateAlumno(ctx iris.Context) {
	alumno := models.Alumno{}
	err := ctx.ReadForm(&alumno)
	//err := json.NewDecoder(ctx.Request().Body).Decode(&alumno)
	if err != nil {
		// Sí hay algun error al guardar los datos se devolvera un error 500
		fmt.Println(err)
		view.SendErr(ctx, http.StatusInternalServerError)
		return
	}
	db := database.GetDatabase()
	result := db.Create(&alumno)
	// Se codifica el nuevo registro y se devuelve

	fmt.Println(result)
	ctx.View("index.html")
}

func GetAlumno(ctx iris.Context) {
	id := ctx.URLParam("id")

	alumno := models.Alumno{}
	db := database.GetDatabase()
	db.Where("ID = ?", id).First(&alumno)
	ctx.ViewData("alumno", alumno)
	ctx.View("update.html")
}

func UpdateAlumno(ctx iris.Context) {
	id := ctx.URLParam("id")

	alumno := models.Alumno{}
	db := database.GetDatabase()
	db.Where("ID = ?", id).First(&alumno)
	ctx.ViewData("alumno", alumno)
	ctx.View("update.html")
}

func UpdateAlumno2(ctx iris.Context) {
	id := ctx.URLParam("id")
	alumno := models.Alumno{}
	alumnoForm := models.Alumno{}
	db := database.GetDatabase()
	db.First(&alumno, id)
	err := ctx.ReadForm(&alumnoForm)
	//err := json.NewDecoder(ctx.Request().Body).Decode(&alumno)
	if err != nil {
		// Sí hay algun error al guardar los datos se devolvera un error 500
		fmt.Println(err)
		view.SendErr(ctx, http.StatusInternalServerError)
		return
	}
	alumno.Nombre = alumnoForm.Nombre
	alumno.Documento = alumnoForm.Documento
	alumno.Direccion = alumnoForm.Direccion
	alumno.Telefono = alumnoForm.Telefono
	db.Save(&alumno)
	AllAlumnos(ctx)
}

func DeleteAlumno(ctx iris.Context) {
	id := ctx.URLParam("id")
	db := database.GetDatabase()
	db.Delete(&models.Alumno{}, id)
	AllAlumnos(ctx)
}
