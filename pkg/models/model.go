package models

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (m *Repository) AutoMigrate() {
	m.db.AutoMigrate(&Inventory{})
}
