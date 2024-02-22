package graph

import (
	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	db *gorm.DB
}

func InitGraphRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}
