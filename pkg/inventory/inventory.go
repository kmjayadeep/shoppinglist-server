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

type ItemRequest struct {
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

// Edit modify item in inventory
//
//	@Summary		Modify inventory item
//	@Description	Modify inventory item
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Param			Item	body		ItemRequest	true	"Edit inventory item"
//	@Param			id	path		string	true	"inventory item id"
//	@Success		200
//	@Router			/inventory/{id} [post]
func Edit(c *gin.Context) {
	id := c.Param("id")
	req := ItemRequest{}
	if err := c.BindJSON(&req); err != nil {
		return
	}

	for k, item := range inventoryItems {
		if id == item.ID {
			inventoryItems[k].Name = req.Name
			inventoryItems[k].Expiry = req.Expiry
			inventoryItems[k].Quantity = req.Quantity
			inventoryItems[k].StorageLocation = req.StorageLocation
			inventoryItems[k].Unit = req.Unit
			c.Status(http.StatusOK)
			return
		}
	}

	c.Status(http.StatusNotFound)
}

// Add Add item to inventory list
//
//	@Summary		Add to inventory list
//	@Description	Add item to inventory list
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Param			Item	body		ItemRequest	true	"Add inventory item"
//	@Success		201
//	@Router			/inventory [post]
func Add(c *gin.Context) {
	item := ItemRequest{}
	c.BindJSON(&item)
	inventoryItems = append(inventoryItems, Item{
		ID:       uuid.NewString(),
		Name:     item.Name,
		Expiry:   item.Expiry,
		Quantity: item.Quantity,
		Unit:     item.Unit,
	})
	c.Status(http.StatusCreated)
}

// Delee Delete item from inventory list
//
//		@Summary		Delete from inventory list
//		@Description	Delete item from inventory list
//		@Tags			inventory
//		@Accept			json
//		@Produce		json
//		@Param			id	path		string	true	"inventory item id"
//		@Success		200
//	 @Failure     404
//		@Router			/inventory [delete]
func Delete(c *gin.Context) {
	id := c.Param("id")

	newItems := []Item{}
	found := false

	for _, item := range inventoryItems {
		if item.ID == id {
			found = true
		} else {
			newItems = append(newItems, item)
		}
	}

	if found {
		inventoryItems = newItems
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusNotFound)
}
