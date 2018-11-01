package logmanager

import (
	"k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/kubelet/config"
)

type Manager interface {
	CreateLogPolicy(pod *v1.Pod) error
	RemoveLogPolicy(pod *v1.Pod) error
	Start(sourcesReady config.SourcesReady) error
}
