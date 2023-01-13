package domain

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	Thumbnail   string
	Name        string
	Media       string
	Description string
	Playtime    int
	Resolution  int
	ARwidth     int
	ARheight    int
	Fever       bool
	Ondemand    bool
}

type CreateContentRequest struct {
	Thumbnail   *multipart.FileHeader `form:"Thumbnail" binding:"required"`
	Media       *multipart.FileHeader `form:"Media" binding:"required"`
	Name        string                `form:"Name" binding:"required"`
	Description string                `form:"Description" binding:"required"`
	Playtime    int                   `form:"Playtime" binding:"required"`
	Resolution  int                   `form:"Resolution" binding:"required"`
	ARwidth     int                   `form:"ARwidth" binding:"required"`
	ARheight    int                   `form:"ARheight" binding:"required"`
	Fever       *bool                 `form:"Fever" binding:"required"`
	Ondemand    *bool                 `form:"Ondemand" binding:"required"`
}
