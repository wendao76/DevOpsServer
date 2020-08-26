package model

//返回的结果对象
type Result struct{
	success bool
	code uint
	msg string
	data interface{}
}
