package mocks

import "github.com/pondara/go_learning/pkg/models"

var Books = []models.Book{
	{
		Id:     1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "A book for Go",
	},
	{
		Id:     2,
		Title:  "Golang 2",
		Author: "Gopher 2",
		Desc:   "A book for Go 2",
	},
}
