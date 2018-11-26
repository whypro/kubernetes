package logmanager

import (
	"sync"

	k8stypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/logplugin/v1alpha1"
	"k8s.io/kubernetes/pkg/kubelet/log/logmanager/api"
)

type logVolume struct {
	volumeName string
	// path in container
	// eg. /var/log/<category>
	path string
	// real mount path in host
	// eg. /var/lib/kubelet/pods/<pod-uid>/volumes/kubernetes.io~<volume-type>/<volume-name>
	hostPath string
	// pod logs symlink path
	// eg. /var/log/pods/<pod-uid>/<container-name>/<category>
	logDirPath string
}

// volumeName -> logVolume
type logVolumesMap map[string]*logVolume

// configName -> config
type logConfigsMap map[string]*pluginapi.Config

type podStateManager struct {
	mutex sync.RWMutex
	// pod uid -> podLogPolicy
	// desired state
	podLogPolicies map[k8stypes.UID]*api.PodLogPolicy
	// pod uid -> podLogVolume
	// desired state
	podLogVolumes map[k8stypes.UID]logVolumesMap
	// pod uid -> configmap key set
	// update from pod policy
	// desired state
	podConfigMaps map[k8stypes.UID]sets.String
	// configmap key -> pod uid set
	// configmap key := <namespace>/<name>
	// update from pod policy
	// desired state
	configMapPodUIDs map[string]sets.String
}

type pluginStateManager struct {
	mutex sync.RWMutex
	// pod uid -> config name set
	// update from log plugins
	// current state
	podLogConfigNames    map[k8stypes.UID]sets.String
	podLogPluginEndpoint map[k8stypes.UID]pluginEndpoint
}

func newPodStateManager() *podStateManager {
	return &podStateManager{
		podLogPolicies:   make(map[k8stypes.UID]*api.PodLogPolicy),
		podLogVolumes:    make(map[k8stypes.UID]logVolumesMap),
		podConfigMaps:    make(map[k8stypes.UID]sets.String),
		configMapPodUIDs: make(map[string]sets.String),
	}
}

func (m *podStateManager) updateConfigMapKeys(podUID k8stypes.UID, configMapKeys sets.String) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.podConfigMaps[podUID] = configMapKeys
	for key := range configMapKeys {
		podUIDs, exists := m.configMapPodUIDs[key]
		if !exists {
			podUIDs = sets.NewString()
			m.configMapPodUIDs[key] = podUIDs
		}
		podUIDs.Insert(string(podUID))
	}
}

func (m *podStateManager) removeConfigMapKeys(podUID k8stypes.UID) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	configMapKeys, uidExists := m.podConfigMaps[podUID]
	if uidExists {
		for key := range configMapKeys {
			podUIDs, keyExists := m.configMapPodUIDs[key]
			if keyExists {
				podUIDs.Delete(string(podUID))
				if podUIDs.Len() == 0 {
					delete(m.configMapPodUIDs, key)
				}
			}
		}
	}
	delete(m.podConfigMaps, podUID)
}

func (m *podStateManager) getAllConfigMapKeys() sets.String {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	configMapKeys := sets.NewString()
	for key := range m.configMapPodUIDs {
		configMapKeys.Insert(key)
	}
	return configMapKeys
}

func (m *podStateManager) getPodUIDs(configMapKey string) sets.String {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	podUIDs, exists := m.configMapPodUIDs[configMapKey]
	if !exists {
		podUIDs = sets.NewString()
	}
	return podUIDs
}

func (m *podStateManager) updateLogPolicy(podUID k8stypes.UID, podLogPolicy *api.PodLogPolicy) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.podLogPolicies[podUID] = podLogPolicy
}

func (m *podStateManager) removeLogPolicy(podUID k8stypes.UID) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.podLogPolicies, podUID)
}

func (m *podStateManager) getLogPolicy(podUID k8stypes.UID) (*api.PodLogPolicy, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	podLogPolicy, exists := m.podLogPolicies[podUID]
	return podLogPolicy, exists
}

func (m *podStateManager) updateLogVolumes(podUID k8stypes.UID, logVolumes logVolumesMap) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.podLogVolumes[podUID] = logVolumes
}

func (m *podStateManager) removeLogVolumes(podUID k8stypes.UID) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.podLogVolumes, podUID)
}

func (m *podStateManager) getLogVolumes(podUID k8stypes.UID) (logVolumesMap, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	podLogVolume, exists := m.podLogVolumes[podUID]
	return podLogVolume, exists
}

func newPluginStateManager() *pluginStateManager {
	return &pluginStateManager{
		podLogConfigNames: make(map[k8stypes.UID]sets.String),
	}
}

func (m *pluginStateManager) getLogConfigNames(podUID k8stypes.UID) sets.String {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	configNames, exists := m.podLogConfigNames[podUID]
	if !exists {
		return sets.NewString()
	}
	return configNames
}

func (m *pluginStateManager) updateAllLogConfigs(configs []*pluginapi.Config, endpoint pluginEndpoint) {
	if configs == nil {
		return
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	podLogConfigNames := make(map[k8stypes.UID]sets.String)
	podLogPluginEndpoint := make(map[k8stypes.UID]pluginEndpoint)
	for _, config := range configs {
		// update podLogConfigNames map
		configNames, exists := podLogConfigNames[k8stypes.UID(config.Metadata.PodUID)]
		if !exists {
			configNames = sets.NewString()
			podLogConfigNames[k8stypes.UID(config.Metadata.PodUID)] = configNames
		}
		configNames.Insert(config.Metadata.Name)
		// update podLogPluginName map
		_, exists = podLogPluginEndpoint[k8stypes.UID(config.Metadata.PodUID)]
		if !exists {
			podLogPluginEndpoint[k8stypes.UID(config.Metadata.PodUID)] = endpoint
		}
	}
	m.podLogConfigNames = podLogConfigNames
	m.podLogPluginEndpoint = podLogPluginEndpoint

}

func (m *pluginStateManager) getAllPodUIDs() sets.String {
	podUIDs := sets.NewString()
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	for uid := range m.podLogConfigNames {
		podUIDs.Insert(string(uid))
	}
	return podUIDs
}

func (m *pluginStateManager) getLogPluginEndpoint(podUID k8stypes.UID) (pluginEndpoint, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	endpoint, exists := m.podLogPluginEndpoint[podUID]
	return endpoint, exists
}
