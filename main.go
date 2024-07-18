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
		Name: "eggs",
		ID:   uuid.NewString(),
	}, ShoppingItem{
		Name: "milk",
		ID:   uuid.NewString(),
	})

	r := gin.Default()
	r.Use(cors.Default()) // Allow all origins

	r.GET("/api/v1/shopping-list", GetShoppingList)
	r.POST("/api/v1/shopping-list", AddToShoppingList)
	r.POST("/api/v1/shopping-list/:id", EditShoppingList)
	r.DELETE("/api/v1/shopping-list/:id", DeleteFromShoppingList)
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

// EditShoppingList Add item to shopping list
//
//	@Summary		Edit shopping list
//	@Description	Edit item in shopping list
//	@Tags			shopping-list
//	@Accept			json
//	@Produce		json
//	@Param			shoppingItem	body		ShoppingItemRequest	true	"Edit shopping item"
//	@Param			id	path		string	true	"shopping item id"
//	@Success		200
//	@Router			/shopping-list/{id} [post]
func EditShoppingList(c *gin.Context) {
	id := c.Param("id")
	req := ShoppingItemRequest{}
	if err := c.BindJSON(&req); err != nil {
		return
	}

	for k, item := range shoppingItems {
		if id == item.ID {
			shoppingItems[k].Name = req.Name
			c.Status(http.StatusOK)
			return
		}
	}

	c.Status(http.StatusNotFound)
}

// DeleteFromShoppingList Delete item from shopping list
//
//		@Summary		Delete from shopping list
//		@Description	Delete item from shopping list
//		@Tags			shopping-list
//		@Accept			json
//		@Produce		json
//		@Param			id	path		string	true	"shopping item id"
//		@Success		200
//	 @Failure     404
//		@Router			/shopping-list [delete]
func DeleteFromShoppingList(c *gin.Context) {
	id := c.Param("id")

	newItems := []ShoppingItem{}
	found := false

	for _, item := range shoppingItems {
		if item.ID == id {
			found = true
		} else {
			newItems = append(newItems, item)
		}
	}

	if found {
		shoppingItems = newItems
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusNotFound)
}
