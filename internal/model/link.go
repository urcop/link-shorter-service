package model

import "gorm.io/gorm"

type Link struct {
	*gorm.Model
	ID        string `json:"ID"`
	Link      string `json:"link"`
	ShortLink string `json:"short_link"`
	Clicked   string `json:"clicked"`
	Random    string `json:"random"`
}
