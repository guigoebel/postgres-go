package controller

import (
	"net/http"

	"github.com/guigoebel/postgres-go/model"
	"github.com/guigoebel/postgres-go/storage"
	"github.com/labstack/echo/v4"
)

// GetStudents
func GetStudents(c echo.Context) error {
	students, _ := GetRepoStudents()
	return c.JSON(http.StatusOK, students)
}

func GetRepoStudents() ([]model.Students, error) {
	db := storage.GetDBInstance()
	students := []model.Students{}

	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}
