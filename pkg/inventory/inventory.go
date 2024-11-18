package inventory

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kmjayadeep/shoppinglist-server/pkg/models"
)

type InventoryService struct {
	repo *models.Repository
}

func NewService(repo *models.Repository) *InventoryService {
	return &InventoryService{
		repo: repo,
	}
}

type ItemRequest struct {
	Name            string    `json:"name"`
	Expiry          time.Time `json:"expiry"`
	Quantity        int       `json:"quntity"`
	StorageLocation string    `json:"storageLocation"`
	Unit            string    `json:"unit"`
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
func (s *InventoryService) Get(c *gin.Context) {
	items, _ := s.repo.GetInventory()

	c.JSON(http.StatusOK, gin.H{
		"items": items,
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
func (s *InventoryService) Edit(c *gin.Context) {
	// id := c.Param("id")
	req := ItemRequest{}
	if err := c.BindJSON(&req); err != nil {
		return
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
func (s *InventoryService) Add(c *gin.Context) {
	item := ItemRequest{}
	c.BindJSON(&item)
	inv := models.Inventory{
		Name:            item.Name,
		Expiry:          item.Expiry,
		Quantity:        item.Quantity,
		StorageLocation: item.StorageLocation,
		Unit:            item.Unit,
	}
	_ = s.repo.CreateInventory(&inv)
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
func (s *InventoryService) Delete(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)

	err := s.repo.DeleteInventoryByID(uint(i))
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusOK)
	}
}
