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

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string          `json:"name" gorm:"column:name;"`
	Description     string          `json:"description" gorm:"column:description;"`
	Icon            *Icon           `json:"icon" gorm:"column:icon;"`
	Status          *CategoryStatus `json:"status" gorm:"column:status;"`
	CreatedAt       *time.Time      `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt       *time.Time      `json:"updated_at" gorm:"column:updated_at;"`
}

func (Category) TableName() string { return "categories" }

type CategoryCreation struct {
	common.SQLModel `json:",inline"`
	/*	Ví dụ về show lỗi hiển thị
		Name            string          `json:"name" gorm:"column:namem;"`*/
	Name        string          `json:"name" gorm:"column:name;"`
	Description string          `json:"description" gorm:"column:description;"`
	Icon        *Icon           `json:"icon" gorm:"column:icon;"`
	Status      *CategoryStatus `json:"status" gorm:"column:status;"`
	CreatedAt   *time.Time      `json:"created_at" gorm:"column:created_at;"`
}

func (CategoryCreation) TableName() string { return Category{}.TableName() }

type CategoryUpdate struct {
	common.SQLModel `json:",inline"`
	Name            *string         `json:"name" gorm:"column:name;"`
	Description     *string         `json:"description" gorm:"column:description;"`
	Icon            *Icon           `json:"icon" gorm:"column:icon;"`
	Status          *CategoryStatus `json:"status" gorm:"column:status;"`
	UpdatedAt       *time.Time      `json:"updated_at" gorm:"column:updated_at;"`
}

func (CategoryUpdate) TableName() string { return Category{}.TableName() }
