package Smarthouse_server

import "errors"

type DeviceList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type DeviceItem struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name" binding:"required"`
	Place     string `json:"place" db:"place"`
	Condition bool   `json:"condition" db:"condition"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateItemInput struct {
	Name      *string `json:"name"`
	Place     *string `json:"place"`
	Condition *bool   `json:"condition"`
}

func (i UpdateItemInput) Validate() error {
	if i.Name == nil && i.Place == nil && i.Condition == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
