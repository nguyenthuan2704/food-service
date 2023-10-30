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

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string            `json:"name" gorm:"column:name;"`
	Addr            string            `json:"addr" gorm:"column:addr;"`
	Logo            *Logo             `json:"logo" gorm:"column:logo;"`
	Status          *RestaurantStatus `json:"status" gorm:"column:status;"`
	CreatedAt       *time.Time        `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt       *time.Time        `json:"updated_at" gorm:"column:updated_at;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreation struct {
	common.SQLModel `json:",inline"`
	Name            string            `json:"name" gorm:"column:name;"`
	Addr            string            `json:"addr" gorm:"column:addr;"`
	Logo            *Logo             `json:"logo" gorm:"column:logo;"`
	Status          *RestaurantStatus `json:"status" gorm:"column:status;"`
	CreatedAt       *time.Time        `json:"created_at" gorm:"column:created_at;"`
}

func (RestaurantCreation) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	common.SQLModel `json:",inline"`
	Name            string            `json:"name" gorm:"column:name;"`
	Addr            string            `json:"addr" gorm:"column:addr;"`
	Logo            *Logo             `json:"logo" gorm:"column:logo;"`
	Status          *RestaurantStatus `json:"status" gorm:"column:status;"`
	UpdatedAt       *time.Time        `json:"updated_at" gorm:"column:updated_at;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
