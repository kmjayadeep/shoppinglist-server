package main

import (
	"github.com/kmjayadeep/shoppinglist-server/docs"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type ShoppingItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var shoppingItems []ShoppingItem

func main() {
	docs.SwaggerInfo.Title = "Shopping List"
	docs.SwaggerInfo.Description = "Shopping list server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "shoppinglist.cosmos.cboxlab.com"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"https"}

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

// GetShoppingList godoc
//
//	@Summary		Get Shopping List
//	@Description	Get shopping list items
//	@Tags			shopping-list
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]GetShoppingItem
//	@Router			/ [get]
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
