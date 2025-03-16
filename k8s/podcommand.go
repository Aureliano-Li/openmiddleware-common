package k8s

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Aureliano-Li/openmiddleware-common/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	olog = util.GetLogger()
)

func ExecCommandInPod(podName, namespace, containerName, command string) (string, string, error) {
	// 获取 Kubernetes 配置
	kubeConfigPath := util.GetParamFromEnv("KUBE_CONFIG_PATH")
	// /Users/li/env/88.conf
	config := ctrl.GetConfigOrDie()
	if kubeConfigPath != "" {
		configPath := kubeConfigPath
		var err error
		config, err = clientcmd.BuildConfigFromFlags("", configPath)
		if err != nil {
			olog.Errorf("Error building kubeconfig: %s", err.Error())
		}
	}

	// 创建 Kubernetes 客户端
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", "", fmt.Errorf("failed to create clientset: %v", err)
	}

	// 定义要执行的命令
	cmd := []string{"sh", "-c", command}

	// 设置 Exec 请求
	req := clientSet.CoreV1().RESTClient().
		Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Command:   cmd,
			Container: containerName,
			Stdin:     false,
			Stdout:    true,
			Stderr:    true,
			TTY:       false,
		}, scheme.ParameterCodec)

	// 执行命令
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		return "", "", fmt.Errorf("failed to create executor: %v", err)
	}

	// 捕获输出
	var stdout, stderr bytes.Buffer
	err = exec.StreamWithContext(context.Background(), remotecommand.StreamOptions{
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to execute command: %v", err)
	}
	olog.Infof("stdout: %s", stdout.String())
	return stdout.String(), stderr.String(), nil
}
