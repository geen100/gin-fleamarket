package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
		{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	}

	ItemRepository := repositories.NewItemMemoryRepository(items)
	ItemService := services.NewItemService(ItemRepository)
	ItemController := controllers.NewItemController(ItemService)

	r := gin.Default()
	r.GET("/items", ItemController.FindAll)
	r.GET("/items/:id", ItemController.FindById)
	r.Run("localhost:8080")
}
