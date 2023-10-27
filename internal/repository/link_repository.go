package repository

import (
	"fmt"
	"github.com/urcop/go-fiber-template/internal/config"
	"github.com/urcop/go-fiber-template/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type LinkRepository interface {
	CreateLink(link *model.Link) error
	GetLink(id string) (*model.Link, error)
	GetAllLinks() ([]*model.Link, error)
	UpdateLink(link *model.Link) error
	DeleteLink(id string) error
}

type LinkRepositoryImpl struct {
	db *gorm.DB
}

func (l *LinkRepositoryImpl) CreateLink(link *model.Link) error {
	result := l.db.Create(&link)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (l *LinkRepositoryImpl) GetLink(id string) (*model.Link, error) {
	var link *model.Link

	result := l.db.Where("id = ?", id).First(&link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (l *LinkRepositoryImpl) GetAllLinks() ([]*model.Link, error) {
	var links []*model.Link

	result := l.db.Find(&links)
	if result.Error != nil {
		return nil, result.Error
	}
	return links, nil
}

func (l *LinkRepositoryImpl) UpdateLink(link *model.Link) error {
	result := l.db.Save(&link)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (l *LinkRepositoryImpl) DeleteLink(id string) error {
	var link *model.Link

	result := l.db.Where("id = ?", id).First(&link)
	link.DeletedAt = gorm.DeletedAt{time.Now().UTC(), true}
	result = l.db.Save(&link)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewLinkRepository() LinkRepository {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot open database")
	}

	pgSvc := &LinkRepositoryImpl{}
	err = db.AutoMigrate(&model.Link{})
	if err != nil {
		panic(err)
	}
	return pgSvc
}
