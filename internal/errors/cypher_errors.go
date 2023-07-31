package errors

import "errors"

var (
	// 配置文件相关异常
	LoadConfigError = errors.New("加载配置文件异常！")
	GetAESKeyError  = errors.New("获取AES密钥异常！")

	// 文件写入异常
	WriteFileError = errors.New("文件写入异常！")
)
