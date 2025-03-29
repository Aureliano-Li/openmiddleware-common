package middleware

import (
	corev1 "k8s.io/api/core/v1"
	"strconv"
)

// GetMiddlewareLabels 获取中间件label
func GetMiddlewareLabels(middlewareName, middlewareType string) map[string]string {
	return map[string]string{
		KeyLabelOpenMiddlewareName: middlewareName,
		KeyLabelOpenMiddlewareType: middlewareType,
	}
}

// GetMiddlewareBasicEnv 获取基础环境变量
func GetMiddlewareBasicEnv(replicas int32, middlewareName, middlewareType string) []corev1.EnvVar {
	envs := []corev1.EnvVar{
		{
			Name: "POD_NAMESPACE",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "metadata.namespace",
				},
			},
		},
		{
			Name: "POD_NAME",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "metadata.name",
				},
			},
		},
		{
			Name:  "REPLICAS",
			Value: strconv.Itoa(int(replicas)),
		},
		{
			Name:  "OPEN_MIDDLEWARE_CLUSTER_NAME",
			Value: middlewareName,
		},
		{
			Name:  "OPEN_MIDDLEWARE_CLUSTER_TYPE",
			Value: middlewareType,
		},
	}
	return envs
}
