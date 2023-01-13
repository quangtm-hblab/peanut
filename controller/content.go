package controller

import (
	"net/http"
	"os"
	"peanut/domain"
	"peanut/repository"
	"peanut/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContentController struct {
	Usecase usecase.ContentUsecase
}

func NewContentController(db *gorm.DB) *ContentController {
	return &ContentController{
		Usecase: usecase.NewContentUsecase(repository.NewContentRepo(db)),
	}
}

// ListContent godoc
//
//	@Summary		content
//	@Description	content
//	@Tags			Content
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]domain.Content
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Router			/contents [get]
func (c *ContentController) GetContents(ctx *gin.Context) {
	contents, err := c.Usecase.GetContents()
	if err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{
			"message": "Not found any record",
		})
		return
	}
	ctx.JSON(http.StatusOK, domain.Response{
		Success: true, Data: contents, Message: "Get data successful",
	})
}

// CreateContent godoc
//
//	@Summary		content
//	@Description	content
//	@Tags			Content
//	@Accept			json
//	@Produce		json
//	@Param			Thumbnail	formData	file	true	"file"
//	@Param			Media		formData	file	true	"file"
//	@Param			Name		formData	string	true	"string"	minlength(1)	maxlength(30)
//	@Param			Description	formData	string	true	"string"	minlength(0)	maxlength(500)
//	@Param			PlayTime	formData	int		true	"int"
//	@Param			Resolution	formData	int		true	"int"
//	@Param			ARheight	formData	int		true	"int"
//	@Param			ARwidth		formData	int		true	"int"
//	@Param			Fever		formData	boolean	false	"boolean"
//	@Param			Ondemand	formData	boolean	false	"boolean"
//	@Success		200			{object}	domain.Content
//	@Failure		400			{object}	domain.ErrorResponse
//	@Failure		500			{object}	domain.ErrorResponse
//	@Router			/contents [post]
func (c *ContentController) CreateContent(ctx *gin.Context) {
	thumbnailDir := os.Getenv("UPLOAD_THUMBNAIL_PATH")
	mediaDir := os.Getenv("UPLOAD_MEDIA_PATH")
	var req domain.CreateContentRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, err.Error()})
	}

	id := uuid.New()
	thumbnailName := id.String() + req.Thumbnail.Filename
	thumbnailDst := thumbnailDir + thumbnailName
	mediaName := id.String() + req.Media.Filename
	mediaDst := mediaDir + mediaName
	var saveErr error
	saveErr = ctx.SaveUploadedFile(req.Media, mediaDst)
	if saveErr != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, saveErr.Error()})
	}
	saveErr = ctx.SaveUploadedFile(req.Thumbnail, thumbnailDst)
	if saveErr != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, saveErr.Error()})
	}

	newContent := domain.Content{
		Thumbnail:   thumbnailDst,
		Media:       mediaDst,
		Name:        req.Name,
		Description: req.Description,
		Playtime:    req.Playtime,
		Resolution:  req.Resolution,
		ARwidth:     req.ARwidth,
		ARheight:    req.ARheight,
		Ondemand:    *req.Ondemand,
		Fever:       *req.Fever,
	}
	content, err := c.Usecase.CreateContent(newContent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.Response{true, content, "create content successfully"})
}
