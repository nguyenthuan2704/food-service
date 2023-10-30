package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Logo struct {
	Url    string `json:"url" gorm:"column:url;"`
	Name   string `json:"name" gorm:"column:name;"`
	Width  int    `json:"width" gorm:"column:width;"`
	Height int    `json:"height" gorm:"column:height;"`
}

func (j *Logo) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img Logo
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *Logo) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
