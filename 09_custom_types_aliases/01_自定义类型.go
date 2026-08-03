package main

import "fmt"

type Code int

const (
	SuccessCode    Code = 0
	ValidCode      Code = 7
	ServiceErrCode Code = 8
)

func (c Code) GetMsg() string {
	switch c {
	case SuccessCode:
		return "成功"
	case ValidCode:
		return "校验失败"
	case ServiceErrCode:
		return "服务错误"
	default:
		return "未知错误"
	}
}

func main() {
	var code Code = SuccessCode
	fmt.Println(code)
	fmt.Println(code.GetMsg())
}
