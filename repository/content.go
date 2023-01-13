package repository

import (
	"peanut/domain"

	"gorm.io/gorm"
)

type ContentRepo interface {
	GetContents() ([]domain.Content, error)
	CreateContent(domain.Content) (*domain.Content, error)
}

type contentRepo struct {
	DB *gorm.DB
}

func NewContentRepo(db *gorm.DB) ContentRepo {
	return &contentRepo{DB: db}
}

func (r *contentRepo) GetContents() (contents []domain.Content, err error) {
	result := r.DB.Find(&contents)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}

func (r *contentRepo) CreateContent(c domain.Content) (newcontent *domain.Content, err error) {
	newcontent = &domain.Content{
		Model:       gorm.Model{},
		Thumbnail:   c.Thumbnail,
		Name:        c.Name,
		Media:       c.Media,
		Description: c.Description,
		Playtime:    c.Playtime,
		Resolution:  c.Resolution,
		ARwidth:     c.ARwidth,
		ARheight:    c.ARheight,
		Fever:       c.Fever,
		Ondemand:    c.Ondemand,
	}
	result := r.DB.Create(newcontent)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
