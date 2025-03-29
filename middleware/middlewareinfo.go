package middleware

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RoleStatusInfo struct {
	PodName        string      `json:"podName,omitempty"`
	Role           string      `json:"role,omitempty"`
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`
	PodStatus      string      `json:"podStatus,omitempty"`
}
