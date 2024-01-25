package models

import (
	"time"

	"qna-management/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityUsers struct {
	ID        string `gorm:"primaryKey;"`
	FullName  string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:boolean;default:false"`
	Role      string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *EntityUsers) BeforeCreate(db *gorm.DB) error {
	e.ID = uuid.New().String()
	e.Password = utils.HashPassword(e.Password)
	e.CreatedAt = time.Now().Local()
	return nil
}

func (e *EntityUsers) BeforeUpdate(db *gorm.DB) error {
	e.UpdatedAt = time.Now().Local()
	return nil
}
