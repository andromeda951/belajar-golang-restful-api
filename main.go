package main

import (
	"andromeda/belajar-golang-restful-api/app"
	"andromeda/belajar-golang-restful-api/controller"
	"andromeda/belajar-golang-restful-api/repository"
	"andromeda/belajar-golang-restful-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories", categoryController.Update)
	router.DELETE("/api/categories", categoryController.Delete)
}