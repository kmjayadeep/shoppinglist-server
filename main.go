package main

import (
	_ "github.com/kmjayadeep/shoppinglist-server/docs"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type ShoppingItem struct {
	ID   string `json:"id"  swaggertype:"string"`
	Name string `json:"name"`
}

type ShoppingItemRequest struct {
	Name string `json:"name"`
}

var shoppingItems []ShoppingItem

//	@title			Shopping List
//	@version		1.0
//	@description	Shopping list manager

//	@host		shoppinglist.cosmos.cboxlab.com
//	@BasePath	/api/v1
//	@Schemes	https

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	shoppingItems = append(shoppingItems, ShoppingItem{
		Name: "chilly powder",
		ID:   uuid.NewString(),
	}, ShoppingItem{
		Name: "garam masala",
		ID:   uuid.NewString(),
	}, ShoppingItem{
		Name: "milk",
		ID:   uuid.NewString(),
	})

	r := gin.Default()
	r.Use(cors.Default()) // Allow all origins

	r.GET("/api/v1/shopping-list", GetShoppingList)
	r.POST("/api/v1/shopping-list", AddToShoppingList)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

// GetShoppingList return shopping list items
//
//	@Summary		Get Shopping List
//	@Description	Return shopping list items
//	@Tags			shopping-list
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	ShoppingItem
//	@Router			/shopping-list [get]
func GetShoppingList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items": shoppingItems,
	})
}

// AddToShoppingList Add item to shopping list
//
//	@Summary		Add to shopping list
//	@Description	Add item to shopping list
//	@Tags			shopping-list
//	@Accept			json
//	@Produce		json
//	@Param			shoppingItem	body		ShoppingItemRequest	true	"Add shopping item"
//	@Success		201
//	@Router			/shopping-list [post]
func AddToShoppingList(c *gin.Context) {
	item := ShoppingItemRequest{}
	c.BindJSON(&item)
	shoppingItems = append(shoppingItems, ShoppingItem{
		ID:   uuid.NewString(),
		Name: item.Name,
	})
	c.Status(http.StatusCreated)
}
