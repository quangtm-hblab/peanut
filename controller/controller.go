package controller

import (
	"net/http"
	"peanut/domain"
	"peanut/pkg/apierrors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func bindJSON(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBindJSON(obj)
	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}

func bindForm(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBind(obj)
	// check required params
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, err.Error()})
		return false
	}

	// check if params is null/""
	form, _ := ctx.MultipartForm()
	values := form.Value

	for _, value := range values {
		if value[0] == "" {
			ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, "Required params shouldn't be null"})
			return false
		}
	}
	return true
}

func bindQueryParams(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBindQuery(obj)

	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}

func checkError(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	_ = ctx.Error(err).SetType(gin.ErrorTypePublic)
	return true
}
