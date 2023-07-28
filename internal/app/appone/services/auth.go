package services

import (
	"gowebdemo/internal/app/appone/models/request"
	"gowebdemo/internal/app/appone/models/response"
	"gowebdemo/internal/pkg/jwt"
)

func Login(l *request.Login) (*response.Login, error) {
	// 根据 req.Login 查询用户
	res := &response.Login{}

	// 根据用户ID 生成token
	claims := map[string]any{
		"user_id":   12345,
		"user_name": "heeron",
	}
	tokenStr, err := jwt.Token(claims)
	if err != nil {
		return nil, err
	}

	res.Token = tokenStr

	return res, nil
}
