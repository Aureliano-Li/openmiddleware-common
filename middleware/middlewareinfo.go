package middleware

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BasicStatus struct {
	Status     string           `json:"status,omitempty"`
	Phase      string           `json:"phase,omitempty"`
	RoleStatus []RoleStatusInfo `json:"roleStatus,omitempty"`
}

type RoleStatusInfo struct {
	PodName        string      `json:"podName,omitempty"`
	Role           string      `json:"role,omitempty"`
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`
	PodStatus      string      `json:"podStatus,omitempty"`
}
