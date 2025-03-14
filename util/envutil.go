package util

import (
	"fmt"
	"os"
)

func GetParamFromEnv(envName string) string {
	// 获取名为 "envName" 的环境变量的值
	envValue := os.Getenv(envName)
	printEnvLog := os.Getenv("PRINT_ENV_LOG")

	if envValue == "" {
		// 环境变量未设置或为空
		if printEnvLog != "" {
			fmt.Println(fmt.Sprintf("环境变量 %s 未设置", envName))
		}
	} else {
		// 打印环境变量的值
		if printEnvLog != "" {
			fmt.Println(fmt.Sprintf("环境变量 %s -> %s", envName, envValue))
		}
	}
	return envValue
}
