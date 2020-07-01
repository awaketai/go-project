package model

import "errors"

// 自定义错误
var (
	ERROR_USER_NOT_EXISTS = errors.New("user not exists")
	ERROR_USER_EXISTS = errors.New("user already exists")
	ERROR_USER_WRONG_PWD = errors.New("user wrong password")
)
