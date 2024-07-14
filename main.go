package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ShoppingItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var shoppingItems []ShoppingItem

func main() {
	r := gin.Default()
	r.GET("/shopping-list", GetShoppingList)
	r.POST("/shopping-list", AddToShoppingList)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func GetShoppingList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items": shoppingItems,
	})
}

func AddToShoppingList(c *gin.Context) {
	item := ShoppingItem{}
	c.BindJSON(&item)
	item.ID = uuid.NewString()
	shoppingItems = append(shoppingItems, item)
	c.Status(http.StatusCreated)
}
