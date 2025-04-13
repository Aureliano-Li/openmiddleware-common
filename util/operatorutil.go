package util

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"strings"
)

func GetNamespaceFilter() predicate.Funcs {
	// 定义 Predicate，过滤命名空间
	watchNamespace := GetParamFromEnv("WATCH_NAMESPACE")
	nsMap := StringToMap(watchNamespace)
	namespacePredicate := predicate.Funcs{
		CreateFunc: func(e event.CreateEvent) bool {
			return checkNamespaceExists(nsMap, e.Object.GetNamespace())
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			return checkNamespaceExists(nsMap, e.ObjectOld.GetNamespace())
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			return checkNamespaceExists(nsMap, e.Object.GetNamespace())
		},
		GenericFunc: func(e event.GenericEvent) bool {
			return checkNamespaceExists(nsMap, e.Object.GetNamespace())
		},
	}
	return namespacePredicate
}

func checkNamespaceExists(nsMap map[string]struct{}, ns string) bool {
	if len(nsMap) == 0 {
		return true
	}
	_, exists := nsMap[ns]
	return exists
}

func StringToMap(input string) map[string]struct{} {
	result := make(map[string]struct{})

	// 如果输入为空，返回空map
	if input == "" {
		return result
	}

	// 按逗号分割字符串
	pairs := strings.Split(input, ",")

	for _, pair := range pairs {
		// 去除空格
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}

		result[pair] = struct{}{}
	}

	return result
}
