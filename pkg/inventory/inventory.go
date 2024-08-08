package inventory

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Item struct {
	ID              string    `json:"id"  swaggertype:"string"`
	Name            string    `json:"name"`
	Expiry          time.Time `json:"expiry"`
	Quantity        int       `json:"quntity"`
	StorageLocation string    `json:"storageLocation"`
	Unit            string    `json:"unit"`
}

var inventoryItems []Item

func init() {
	inventoryItems = append(inventoryItems, Item{
		Name:            "mayo",
		ID:              uuid.NewString(),
		Expiry:          time.Now().Add(7 * 24 * 60 * time.Minute),
		StorageLocation: "fridge",
		Unit:            "KG",
		Quantity:        1,
	})

}

// Get return inventory items
//
//	@Summary		Get Inventory
//	@Description	Return items in inventory
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	Item
//	@Router			/inventory [get]
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items": inventoryItems,
	})
}
