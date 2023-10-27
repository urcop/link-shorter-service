package model

import "gorm.io/gorm"

type Link struct {
	*gorm.Model
	ID        string `json:"ID"`
	Link      string `json:"link"`
	ShortLink string `json:"short_link"`
	Clicked   uint32 `json:"clicked"`
	Random    bool   `json:"random"`
}
