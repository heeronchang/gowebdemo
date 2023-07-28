package controllers

import (
	"gowebdemo/internal/app/appone/models/request"
	"gowebdemo/internal/app/appone/models/response"
	"gowebdemo/internal/app/appone/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Auth struct {
}

func (a *Auth) Login(ctx *gin.Context) {
	l := &request.Login{}
	if err := ctx.ShouldBindJSON(l); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Fail(ctx, err.Error())
			return
		}

		response.Fail(ctx, errs.Translate(GetTrans()))
		return
	}

	res, err := services.Login(l)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	response.OK(ctx, res, nil)
}
