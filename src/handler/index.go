package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Hello(c echo.Context) error {
	return c.HTML(http.StatusOK, "<html><head><title>Just hi!</title></head><body><h1>Hello from echo, alright!</h1><p>haha hihi API</p></body></html>")
}

type Item struct {
	Id   int    `json:"id" xml:"id"`
	Text string `json:"text" xml:"text"`
}

type ResItem struct {
	Desc  string  `json:"desc" xml:"desc"`
	Items []*Item `json:"items" xml:"items"`
}

var items []*Item
var db *gorm.DB

func InitItems() {
	var item1 = new(Item)
	item1.Id = 0
	item1.Text = "abdi"
	fmt.Println(item1)
	items = append(items, item1)
	items = append(items, &Item{Id: 1, Text: "waw"})
}

func InitDB(gormdb *gorm.DB) {
	db = gormdb
}

func GetAll(c echo.Context) error {
	db.Find(&items)
	res := &ResItem{Items: items, Desc: "testing"}
	return c.JSON(http.StatusOK, res)
}

func GetOne(c echo.Context) error {
	selectedId, convertErr := strconv.Atoi(c.Param("id"))
	if convertErr != nil {
		panic(convertErr)
	}
	var selectedItem *Item
	for i := 0; i < len(items); i++ {
		if items[i].Id == selectedId {
			selectedItem = items[i]
		}
	}

	return c.JSON(http.StatusOK, selectedItem)
}

func CreateItem(c echo.Context) error {
	item := new(Item)
	if err := c.Bind(item); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	result := db.Create(item)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	return c.JSON(http.StatusCreated, &ResItem{Desc: "New item is successfully created."})
}

func UpdateItem(c echo.Context) error {
	selectedId, convertErr := strconv.Atoi(c.Param("id"))
	if convertErr != nil {
		panic(convertErr)
	}
	body := new(Item)
	if err := c.Bind(body); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}
	db.Save(&Item{Id: selectedId, Text: body.Text})
	return c.JSON(http.StatusOK, &ResItem{Desc: "Updated"})
}

func RemoveItem(c echo.Context) error {
	selectedId, convertErr := strconv.Atoi(c.Param("id"))
	if convertErr != nil {
		panic(convertErr)
	}
	db.Delete(&Item{}, selectedId)
	return c.JSON(http.StatusOK, &ResItem{Desc: "Deleted"})
}
