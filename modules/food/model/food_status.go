package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type FoodStatus int

const (
	FoodDeleted FoodStatus = iota
	FoodDoing
)

var allFoodStatuses = [2]string{"Deleted", "Doing"}

func (item *FoodStatus) String() string {
	return allFoodStatuses[*item]
}

func parseStr2ItemStatus(s string) (FoodStatus, error) {
	for i := range allFoodStatuses {
		if allFoodStatuses[i] == s {
			return FoodStatus(i), nil
		}
	}
	return FoodStatus(0), errors.New("invalid status string")
}

/* Scan đọc dữ liệu từ Mysql để đem lên CategoryStatus */

func (item *FoodStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	v, err := parseStr2ItemStatus(string(bytes))

	if err != nil {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	*item = v

	return nil

}

/* Ngược lại với hàm Scan */

func (item *FoodStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}

/* Json Encoding -  từ data structure từ CategoryStatus sang json value */

func (item *FoodStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

/* Json Decoding -  ngược lại với MarshalJSON */

func (item *FoodStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	itemValue, err := parseStr2ItemStatus(str)
	if err != nil {
		return err
	}

	*item = itemValue

	return nil
}
