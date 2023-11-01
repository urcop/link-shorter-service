package model

import "gorm.io/gorm"

type Link struct {
	*gorm.Model
	Link      string `json:"link"`
	ShortLink string `gorm:"primarykey;unique:true",json:"short_link"`
	Clicked   uint32 `json:"clicked"`
	Random    bool   `json:"random"`
}
