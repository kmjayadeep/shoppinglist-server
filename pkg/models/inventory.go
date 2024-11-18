package models

import (
	"gorm.io/gorm"
	"time"
)

type Inventory struct {
	gorm.Model
	Name            string    `json:"name"`
	Expiry          time.Time `json:"expiry"`
	Quantity        int       `json:"quntity"`
	StorageLocation string    `json:"storageLocation"`
	Unit            string    `json:"unit"`
}

func (r *Repository) CreateInventory(inv *Inventory) error {
	return r.db.Create(inv).Error
}

func (r *Repository) GetInventoryByID(id uint) (*Inventory, error) {
	var inv Inventory
	err := r.db.First(&inv, id).Error
	return &inv, err
}

func (r *Repository) DeleteInventoryByID(id uint) error {
	var inv Inventory
	err := r.db.Delete(&inv, id).Error
	return err
}

func (r *Repository) GetInventory() ([]Inventory, error) {
	inv := []Inventory{}
	res := r.db.Find(&inv)
	return inv, res.Error
}
