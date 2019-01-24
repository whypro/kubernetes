package api

// LogPolicy events
const (
	LogPolicyConfigUpdateSuccess = "LogPolicyConfigUpdateSuccess"
	LogPolicyConfigUpdateFailed  = "LogPolicyConfigUpdateFailed"
	LogPolicyCreateSuccess       = "LogPolicyCreateSuccess"
	LogPolicyCreateFailed        = "LogPolicyCreateFailed"
	LogPolicyRemoveSuccess       = "LogPolicyRemoveSuccess"
)

const (
	// PodLogPolicyLabelKey is the key of pod log policy
	PodLogPolicyLabelKey = "beta.log.qiniu.com/log-policy"
)

// PodLogPolicy is the log policy definition on a pod
type PodLogPolicy struct {
	// PluginName is log plugin name, eg. logkit, logexporter
	PluginName string `json:"plugin_name"`
	// SafeDeletionEnabled ensure all log been collected before pod are terminated
	// SafeDeletionEnabled == true, pod will keep terminating forever util log plugin says all log are collected.
	// SafeDeletionEnabled == false, pod will terminated before TerminationGracePeriodSeconds in PodSpec.
	SafeDeletionEnabled bool `json:"safe_deletion_enabled"`
	// ContainerLogPolicies is a list of ContainerLogPolicy
	ContainerLogPolicies []ContainerLogPolicy `json:"container_log_policies"`
}

// ContainerLogPolicy is the log policy definition on all containers of a pod
type ContainerLogPolicy struct {
	// ContainerName
	ContainerName string `json:"container_name"`
	// Name is the container log policy name, eg. std(stdout/stderr), app, audit
	Name string `json:"name"`
	// Path is the log volume mount path
	// if path is "-", that means this policy is dedicated for std(stdout/stderr) logs and VolumeName will make no sense.
	Path string `json:"path"`
	// VolumeName is the volume/volumeMount for container file log
	VolumeName string `json:"volume_name"`
	// PluginConfigMap is the configmap name of log plugin configs
	PluginConfigMap string `json:"plugin_configmap"`
}
