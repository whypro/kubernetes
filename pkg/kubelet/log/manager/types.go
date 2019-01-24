package manager

import (
	"k8s.io/api/core/v1"
	k8stypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/kubelet/config"
)

// LogManager is a interface of log manager
type LogManager interface {
	// CreateLogPolicy create pod log policy info
	CreateLogPolicy(pod *v1.Pod) error
	// RemoveLogPolicy removes pod log policy info
	RemoveLogPolicy(pod *v1.Pod) error
	// Start non-blocking starts the manager
	Start(sourcesReady config.SourcesReady) error
	// IsCollectFinished check if a pod log collecting is finished by pod uid
	IsCollectFinished(podUID k8stypes.UID) bool
}
