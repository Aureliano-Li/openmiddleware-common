package middleware

import (
	"context"
	"fmt"
	"github.com/Aureliano-Li/openmiddleware-common/util"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strconv"
)

var (
	mlog = util.GetLogger()
)

func GetInitConfigMapName(middlewareName string) string {
	return fmt.Sprintf("%s-init-shell", middlewareName)
}

// GetHeadlessServiceName 获取中间件headless svc name
func GetHeadlessServiceName(middlewareName string) string {
	return middlewareName + "-headless"
}

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

func EnsureConfigMapExistsFromFile(client client.Client, ctx context.Context, namespace, configMapName, filePath, fileKey string) error {
	// 尝试获取 ConfigMap
	cm := &corev1.ConfigMap{}
	err := client.Get(ctx, types.NamespacedName{Name: configMapName, Namespace: namespace}, cm)
	if err == nil {
		// ConfigMap 已存在
		mlog.Infof("ConfigMap %s already exists", configMapName)
		return nil
	}

	// 读取文件内容
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取文件 %s 失败: %v", filePath, err)
	}

	err = client.Create(ctx, &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapName,
			Namespace: namespace,
		},
		Data: map[string]string{
			fileKey: string(fileContent),
		},
	})
	if err != nil {
		mlog.Errorf("创建 ConfigMap %s 失败: %v", configMapName, err)
		return fmt.Errorf("创建 ConfigMap %s 失败: %v", configMapName, err)
	}
	mlog.Infof("ConfigMap %s create success\n", configMapName)
	return nil
}
