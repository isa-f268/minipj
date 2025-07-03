package helper

import "main/dto"

type respHelper[T interface{}] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type RegisterResp struct {
	respHelper[dto.RegisterResp]
}

type LoginResp struct {
	respHelper[string]
}

func RespHelper[T interface{}](msg string, d T) respHelper[T] {
	return respHelper[T]{Message: msg, Data: d}
}
