/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package decorategpupod

import (
	"testing"
	"strconv"

	"k8s.io/apiserver/pkg/admission"
	"k8s.io/kubernetes/pkg/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)



func getPod(name string, numContainers int, needGPU bool) *api.Pod {
	res := api.ResourceRequirements{}
	if needGPU {
		res.Requests = api.ResourceList{api.ResourceNvidiaGPU: resource.MustParse("1")}
		res.Limits = api.ResourceList{api.ResourceNvidiaGPU: resource.MustParse("1")}
	}

	pod := &api.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: "test"},
		Spec:       api.PodSpec{},
	}
	pod.Spec.Containers = make([]api.Container, 0, numContainers)
	for i := 0; i < numContainers; i++ {
		pod.Spec.Containers = append(pod.Spec.Containers, api.Container{
			Image:     "foo:V" + strconv.Itoa(i),
			Resources: res,
		})
	}

	return pod
}

func getToleration() *api.Toleration {
	tolerationSeconds := int64(3600)
	toleration := &api.Toleration{
	Key: "dedicated",
	Operator: api.TolerationOpEqual,
	Value: "gpu",
	Effect: api.TaintEffectNoExecute,
	TolerationSeconds: &tolerationSeconds,
	}
	return toleration
}

// tolerationEqual checks if two provided tolerations are equal or not.
// will move to "k8s.io/kubernetes/pkg/util/tolerations" in future
func tolerationEqual(first, second api.Toleration) bool {
	if first.Key == second.Key &&
		first.Operator == second.Operator &&
		first.Value == second.Value &&
		first.Effect == second.Effect &&
		tolerationSecondsEqual(first.TolerationSeconds, second.TolerationSeconds) {
		return true
	}
	return false
}

// tolerationSecondsEqual checks if two provided TolerationSeconds are equal or not.
// will move to "k8s.io/kubernetes/pkg/util/tolerations" in future
func tolerationSecondsEqual(ts1, ts2 *int64) bool {
	if ts1 == ts2 {
		return true
	}
	if ts1 != nil && ts2 != nil && *ts1 == *ts2 {
		return true
	}
	return false
}

func tolerationContains(toleration *api.Toleration, tolerationSlice []api.Toleration) bool {
	for _, t := range tolerationSlice {
		if tolerationEqual(t, *toleration) {
			return true
		}
	}
	return false
}

func TestAdmitNeedGPUOnCreateShouldAddToleration(t *testing.T) {
	handler := NewDecorateGPUPodPlugin()
	newPod := getPod("test", 2, true)
	err := handler.Admit(admission.NewAttributesRecord(newPod, nil, api.Kind("Pod").WithVersion("version"), newPod.Namespace, newPod.Name, api.Resource("pods").WithVersion("version"), "", admission.Create, nil))
	if err != nil {
		t.Errorf("Unexpected error returned from admission handler")
	}

	// check toleration
	if !tolerationContains(getToleration(), newPod.Spec.Tolerations) {
		t.Errorf("Check toleration failed")
	}
}

func TestAdmitNeedGPUOnCreateShouldNotAddToleration(t *testing.T) {
	handler := NewDecorateGPUPodPlugin()
	newPod := getPod("test", 2, false)
	err := handler.Admit(admission.NewAttributesRecord(newPod, nil, api.Kind("Pod").WithVersion("version"), newPod.Namespace, newPod.Name, api.Resource("pods").WithVersion("version"), "", admission.Create, nil))
	if err != nil {
		t.Errorf("Unexpected error returned from admission handler")
	}

	// check toleration
	if tolerationContains(getToleration(), newPod.Spec.Tolerations) {
		t.Errorf("Check toleration failed")
	}
}

func TestHandles(t *testing.T) {
	for op, shouldHandle := range map[admission.Operation]bool{
		admission.Create:  true,
		admission.Update:  false,
		admission.Connect: false,
		admission.Delete:  false,
	} {
		handler := NewDecorateGPUPodPlugin()
		if e, a := shouldHandle, handler.Handles(op); e != a {
			t.Errorf("%v: shouldHandle=%t, handles=%t", op, e, a)
		}
	}
}
