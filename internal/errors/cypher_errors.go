package errors

import "errors"

var (
	// LoadConfigError 配置文件相关异常
	LoadConfigError = errors.New("加载配置文件异常！")

	// 获取 AES 密钥异常
	GetAESKeyError = errors.New("获取AES密钥异常！")

	// WriteFileError 文件写入异常
	WriteFileError = errors.New("文件写入异常！")

	// SetKeyError 设置密钥失败
	SetKeyError = errors.New("设置密钥失败！")
)
