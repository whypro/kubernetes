package v1alpha1

const (
	// PodLogPolicyLabelKey is the key of pod log policy
	PodLogPolicyLabelKey = "alpha.log.qiniu.com/log-policy"
)

// PodLogPolicy is the log policy definition on a pod
type PodLogPolicy struct {
	// LogPlugin is log plugin name, eg. logkit, logexporter
	LogPlugin string `json:"log_plugin"`
	// SafeDeletionEnabled ensure all log been collected before pod are terminated
	// SafeDeletionEnabled == true, pod will keep terminating forever util log plugin says all log are collected.
	// SafeDeletionEnabled == false, pod will terminated before TerminationGracePeriodSeconds in PodSpec.
	SafeDeletionEnabled bool `json:"safe_deletion_enabled"`
	// ContainerLogPolicies is a map of container name -> ContainerLogPolicies
	ContainerLogPolicies map[string]ContainerLogPolicies `json:"container_log_policies"`
}

// ContainerLogPolicies is a list of ContainerLogPolicy
type ContainerLogPolicies []*ContainerLogPolicy

// ContainerLogPolicy is the log policy definition on all containers of a pod
type ContainerLogPolicy struct {
	// Category is log category name, eg. std(stdout/stderr), app, audit
	Category string `json:"category"`
	// Path is the log volume mount path
	// if path is "-", that means this policy is dedicated for std(stdout/stderr) logs and VolumeName will make no sense.
	Path string `json:"path"`
	// VolumeName is the volume/volumeMount for container file log
	VolumeName string `json:"volume_name"`
	// PluginConfigMap is the configmap name of log plugin configs
	PluginConfigMap string `json:"plugin_configmap"`
}
