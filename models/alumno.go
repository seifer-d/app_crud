package models

import (
	"gorm.io/gorm"
)

type Alumno struct {
	gorm.Model
	Nombre    string `json:"nombre"`
	Documento string `json:"documento"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
}
