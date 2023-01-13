package usecase

import (
	"peanut/domain"
	"peanut/repository"
)

type ContentUsecase interface {
	GetContents() ([]domain.Content, error)
	CreateContent(domain.Content) (*domain.Content, error)
}

type contentUsecase struct {
	ContentRepo repository.ContentRepo
}

func NewContentUsecase(r repository.ContentRepo) ContentUsecase {
	return &contentUsecase{ContentRepo: r}
}

func (uc *contentUsecase) GetContents() (contents []domain.Content, err error) {
	contents, err = uc.ContentRepo.GetContents()
	if err != nil {
		return nil, err
	}
	return
}

func (uc *contentUsecase) CreateContent(c domain.Content) (content *domain.Content, err error) {
	content, err = uc.ContentRepo.CreateContent(c)
	if err != nil {
		return nil, err
	}
	return
}
