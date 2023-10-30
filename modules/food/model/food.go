package model

import (
	"errors"
	"food-delivery/common"
	"time"
)

const (
	EntityName = "Item"
)

var (
	ErrTitleIsBlank = errors.New("name cannot be blank!")
	ErrItemDeleted  = errors.New("category is deleted!")
)

type Food struct {
	common.SQLModel `json:",inline"`
	Name            string      `json:"name" gorm:"column:name;"`
	Description     string      `json:"description" gorm:"column:description;"`
	Price           int         `json:"price" gorm:"column:price;"`
	Images          *Images     `json:"images" gorm:"column:images;"`
	Status          *FoodStatus `json:"status" gorm:"column:status;"`
	CreatedAt       *time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt       *time.Time  `json:"updated_at" gorm:"column:updated_at;"`
}

func (Food) TableName() string { return "foods" }

type FoodCreation struct {
	common.SQLModel `json:",inline"`
	Name            string      `json:"name" gorm:"column:name;"`
	Description     string      `json:"description" gorm:"column:description;"`
	Price           int         `json:"price" gorm:"column:price;"`
	Images          *Images     `json:"images" gorm:"column:images;"`
	Status          *FoodStatus `json:"status" gorm:"column:status;"`
	CreatedAt       *time.Time  `json:"created_at" gorm:"column:created_at;"`
}

func (FoodCreation) TableName() string { return Food{}.TableName() }

type FoodUpdate struct {
	common.SQLModel `json:",inline"`
	Name            string      `json:"name" gorm:"column:name;"`
	Description     string      `json:"description" gorm:"column:description;"`
	Price           int         `json:"price" gorm:"column:price;"`
	Images          *Images     `json:"images" gorm:"column:images;"`
	Status          *FoodStatus `json:"status" gorm:"column:status;"`
	UpdatedAt       *time.Time  `json:"updated_at" gorm:"column:updated_at;"`
}

func (FoodUpdate) TableName() string { return Food{}.TableName() }
