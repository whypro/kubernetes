package container

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/diff"
	hashutil "k8s.io/kubernetes/pkg/util/hash"
)

func TestHackContainerHash(t *testing.T) {
	testcases := []struct {
		name     string
		original string
		expected string
		hackFunc hashutil.HackFunc
	}{
		{
			"convert 1.9 container go string to 1.7",
			`(v1.Container){Name:(string)kube-proxy Image:(string)index-dev.qiniu.io/kelibrary/kube-proxy-amd64:v1.9.7 Command:([]string)[/usr/local/bin/kube-proxy --config=/var/lib/kube-proxy/config.conf] Args:([]string)<nil> WorkingDir:(string) Ports:([]v1.ContainerPort)<nil> EnvFrom:([]v1.EnvFromSource)<nil> Env:([]v1.EnvVar)<nil> Resources:(v1.ResourceRequirements){Limits:(v1.ResourceList)<nil> Requests:(v1.ResourceList)<nil>} VolumeMounts:([]v1.VolumeMount)[{Name:(string)kube-proxy ReadOnly:(bool)false MountPath:(string)/var/lib/kube-proxy SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)xtables-lock ReadOnly:(bool)false MountPath:(string)/run/xtables.lock SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)lib-modules ReadOnly:(bool)true MountPath:(string)/lib/modules SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)kube-proxy-token-jdlsl ReadOnly:(bool)true MountPath:(string)/var/run/secrets/kubernetes.io/serviceaccount SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>}] VolumeDevices:([]v1.VolumeDevice)<nil> LivenessProbe:(*v1.Probe)<nil> ReadinessProbe:(*v1.Probe)<nil> Lifecycle:(*v1.Lifecycle)<nil> TerminationMessagePath:(string)/dev/termination-log TerminationMessagePolicy:(v1.TerminationMessagePolicy)File ImagePullPolicy:(v1.PullPolicy)IfNotPresent SecurityContext:(*v1.SecurityContext){Capabilities:(*v1.Capabilities)<nil> Privileged:(*bool)true SELinuxOptions:(*v1.SELinuxOptions)<nil> RunAsUser:(*int64)<nil> RunAsNonRoot:(*bool)<nil> ReadOnlyRootFilesystem:(*bool)<nil> AllowPrivilegeEscalation:(*bool)<nil>} Stdin:(bool)false StdinOnce:(bool)false TTY:(bool)false}`,
			`(v1.Container){Name:(string)kube-proxy Image:(string)index-dev.qiniu.io/kelibrary/kube-proxy-amd64:v1.9.7 Command:([]string)[/usr/local/bin/kube-proxy --config=/var/lib/kube-proxy/config.conf] Args:([]string)<nil> WorkingDir:(string) Ports:([]v1.ContainerPort)<nil> EnvFrom:([]v1.EnvFromSource)<nil> Env:([]v1.EnvVar)<nil> Resources:(v1.ResourceRequirements){Limits:(v1.ResourceList)<nil> Requests:(v1.ResourceList)<nil>} VolumeMounts:([]v1.VolumeMount)[{Name:(string)kube-proxy ReadOnly:(bool)false MountPath:(string)/var/lib/kube-proxy SubPath:(string)} {Name:(string)xtables-lock ReadOnly:(bool)false MountPath:(string)/run/xtables.lock SubPath:(string)} {Name:(string)lib-modules ReadOnly:(bool)true MountPath:(string)/lib/modules SubPath:(string)} {Name:(string)kube-proxy-token-jdlsl ReadOnly:(bool)true MountPath:(string)/var/run/secrets/kubernetes.io/serviceaccount SubPath:(string)}] LivenessProbe:(*v1.Probe)<nil> ReadinessProbe:(*v1.Probe)<nil> Lifecycle:(*v1.Lifecycle)<nil> TerminationMessagePath:(string)/dev/termination-log TerminationMessagePolicy:(v1.TerminationMessagePolicy)File ImagePullPolicy:(v1.PullPolicy)IfNotPresent SecurityContext:(*v1.SecurityContext){Capabilities:(*v1.Capabilities)<nil> Privileged:(*bool)true SELinuxOptions:(*v1.SELinuxOptions)<nil> RunAsUser:(*int64)<nil> RunAsNonRoot:(*bool)<nil> ReadOnlyRootFilesystem:(*bool)<nil>} Stdin:(bool)false StdinOnce:(bool)false TTY:(bool)false}`,
			func(s string) string {
				s = hackGoString19to18(s)
				s = hackGoString18to17(s)
				return s
			},
		},
		{
			"convert 1.9 container go string to 1.8",
			`(v1.Container){Name:(string)kube-proxy Image:(string)index-dev.qiniu.io/kelibrary/kube-proxy-amd64:v1.9.7 Command:([]string)[/usr/local/bin/kube-proxy --config=/var/lib/kube-proxy/config.conf] Args:([]string)<nil> WorkingDir:(string) Ports:([]v1.ContainerPort)<nil> EnvFrom:([]v1.EnvFromSource)<nil> Env:([]v1.EnvVar)<nil> Resources:(v1.ResourceRequirements){Limits:(v1.ResourceList)<nil> Requests:(v1.ResourceList)<nil>} VolumeMounts:([]v1.VolumeMount)[{Name:(string)kube-proxy ReadOnly:(bool)false MountPath:(string)/var/lib/kube-proxy SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)xtables-lock ReadOnly:(bool)false MountPath:(string)/run/xtables.lock SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)lib-modules ReadOnly:(bool)true MountPath:(string)/lib/modules SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)kube-proxy-token-jdlsl ReadOnly:(bool)true MountPath:(string)/var/run/secrets/kubernetes.io/serviceaccount SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>}] VolumeDevices:([]v1.VolumeDevice)<nil> LivenessProbe:(*v1.Probe)<nil> ReadinessProbe:(*v1.Probe)<nil> Lifecycle:(*v1.Lifecycle)<nil> TerminationMessagePath:(string)/dev/termination-log TerminationMessagePolicy:(v1.TerminationMessagePolicy)File ImagePullPolicy:(v1.PullPolicy)IfNotPresent SecurityContext:(*v1.SecurityContext){Capabilities:(*v1.Capabilities)<nil> Privileged:(*bool)true SELinuxOptions:(*v1.SELinuxOptions)<nil> RunAsUser:(*int64)<nil> RunAsNonRoot:(*bool)<nil> ReadOnlyRootFilesystem:(*bool)<nil> AllowPrivilegeEscalation:(*bool)<nil>} Stdin:(bool)false StdinOnce:(bool)false TTY:(bool)false}`,
			`(v1.Container){Name:(string)kube-proxy Image:(string)index-dev.qiniu.io/kelibrary/kube-proxy-amd64:v1.9.7 Command:([]string)[/usr/local/bin/kube-proxy --config=/var/lib/kube-proxy/config.conf] Args:([]string)<nil> WorkingDir:(string) Ports:([]v1.ContainerPort)<nil> EnvFrom:([]v1.EnvFromSource)<nil> Env:([]v1.EnvVar)<nil> Resources:(v1.ResourceRequirements){Limits:(v1.ResourceList)<nil> Requests:(v1.ResourceList)<nil>} VolumeMounts:([]v1.VolumeMount)[{Name:(string)kube-proxy ReadOnly:(bool)false MountPath:(string)/var/lib/kube-proxy SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)xtables-lock ReadOnly:(bool)false MountPath:(string)/run/xtables.lock SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)lib-modules ReadOnly:(bool)true MountPath:(string)/lib/modules SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)kube-proxy-token-jdlsl ReadOnly:(bool)true MountPath:(string)/var/run/secrets/kubernetes.io/serviceaccount SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>}] LivenessProbe:(*v1.Probe)<nil> ReadinessProbe:(*v1.Probe)<nil> Lifecycle:(*v1.Lifecycle)<nil> TerminationMessagePath:(string)/dev/termination-log TerminationMessagePolicy:(v1.TerminationMessagePolicy)File ImagePullPolicy:(v1.PullPolicy)IfNotPresent SecurityContext:(*v1.SecurityContext){Capabilities:(*v1.Capabilities)<nil> Privileged:(*bool)true SELinuxOptions:(*v1.SELinuxOptions)<nil> RunAsUser:(*int64)<nil> RunAsNonRoot:(*bool)<nil> ReadOnlyRootFilesystem:(*bool)<nil> AllowPrivilegeEscalation:(*bool)<nil>} Stdin:(bool)false StdinOnce:(bool)false TTY:(bool)false}`,
			func(s string) string {
				s = hackGoString19to18(s)
				return s
			},
		},
	}

	for _, tc := range testcases {
		actual := tc.hackFunc(tc.original)
		if actual != tc.expected {
			t.Errorf("test %q failed, expected: \n %s\nactual: \n%s\n", tc.name, tc.expected, actual)
		}
	}
}

func containerToGoString(container v1.Container) string {
	printer := spew.ConfigState{
		Indent:         " ",
		SortKeys:       true,
		DisableMethods: true,
		SpewKeys:       true,
	}
	return printer.Sprintf("%#v", container)
}

func TestHackLatestContainer(t *testing.T) {
	testcases := []struct {
		name      string
		container v1.Container
		expected  string
		hackFunc  hashutil.HackFunc
	}{
		{
			"convert latest container go string to 1.10",
			v1.Container{
				Name:  "kube-proxy",
				Image: "index-dev.qiniu.io/kelibrary/kube-proxy-amd64:v1.9.7",
				Command: []string{
					"/usr/local/bin/kube-proxy",
					"--config=/var/lib/kube-proxy/config.conf",
				},
				Resources: v1.ResourceRequirements{
					Limits:   nil,
					Requests: nil,
				},
				VolumeMounts: []v1.VolumeMount{
					v1.VolumeMount{
						Name:             "kube-proxy",
						ReadOnly:         false,
						MountPath:        "/var/lib/kube-proxy",
						SubPath:          "",
						MountPropagation: nil,
					},
					v1.VolumeMount{
						Name:             "xtables-lock",
						ReadOnly:         false,
						MountPath:        "/run/xtables.lock",
						SubPath:          "",
						MountPropagation: nil,
					},
					v1.VolumeMount{
						Name:             "lib-modules",
						ReadOnly:         true,
						MountPath:        "/lib/modules",
						SubPath:          "",
						MountPropagation: nil,
					},
					v1.VolumeMount{
						Name:             "kube-proxy-token-jdlsl",
						ReadOnly:         true,
						MountPath:        "/var/run/secrets/kubernetes.io/serviceaccount",
						SubPath:          "",
						MountPropagation: nil,
					},
				},
				TerminationMessagePath:   "/dev/termination-log",
				TerminationMessagePolicy: v1.TerminationMessageReadFile,
				ImagePullPolicy:          v1.PullIfNotPresent,
				SecurityContext: &v1.SecurityContext{
					Capabilities: nil,
					Privileged: func(b bool) *bool {
						return &b
					}(true),
				},
			},
			`(v1.Container){Name:(string)kube-proxy Image:(string)index-dev.qiniu.io/kelibrary/kube-proxy-amd64:v1.9.7 Command:([]string)[/usr/local/bin/kube-proxy --config=/var/lib/kube-proxy/config.conf] Args:([]string)<nil> WorkingDir:(string) Ports:([]v1.ContainerPort)<nil> EnvFrom:([]v1.EnvFromSource)<nil> Env:([]v1.EnvVar)<nil> Resources:(v1.ResourceRequirements){Limits:(v1.ResourceList)<nil> Requests:(v1.ResourceList)<nil>} VolumeMounts:([]v1.VolumeMount)[{Name:(string)kube-proxy ReadOnly:(bool)false MountPath:(string)/var/lib/kube-proxy SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)xtables-lock ReadOnly:(bool)false MountPath:(string)/run/xtables.lock SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)lib-modules ReadOnly:(bool)true MountPath:(string)/lib/modules SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)kube-proxy-token-jdlsl ReadOnly:(bool)true MountPath:(string)/var/run/secrets/kubernetes.io/serviceaccount SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>}] VolumeDevices:([]v1.VolumeDevice)<nil> LivenessProbe:(*v1.Probe)<nil> ReadinessProbe:(*v1.Probe)<nil> Lifecycle:(*v1.Lifecycle)<nil> TerminationMessagePath:(string)/dev/termination-log TerminationMessagePolicy:(v1.TerminationMessagePolicy)File ImagePullPolicy:(v1.PullPolicy)IfNotPresent SecurityContext:(*v1.SecurityContext){Capabilities:(*v1.Capabilities)<nil> Privileged:(*bool)true SELinuxOptions:(*v1.SELinuxOptions)<nil> RunAsUser:(*int64)<nil> RunAsGroup:(*int64)<nil> RunAsNonRoot:(*bool)<nil> ReadOnlyRootFilesystem:(*bool)<nil> AllowPrivilegeEscalation:(*bool)<nil>} Stdin:(bool)false StdinOnce:(bool)false TTY:(bool)false}`,
			hackContainerGoStringTo110,
		},
		{
			"convert latest container go string to 1.9",
			v1.Container{
				Name:  "kube-proxy",
				Image: "index-dev.qiniu.io/kelibrary/kube-proxy-amd64:v1.9.7",
				Command: []string{
					"/usr/local/bin/kube-proxy",
					"--config=/var/lib/kube-proxy/config.conf",
				},
				Resources: v1.ResourceRequirements{
					Limits:   nil,
					Requests: nil,
				},
				VolumeMounts: []v1.VolumeMount{
					v1.VolumeMount{
						Name:             "kube-proxy",
						ReadOnly:         false,
						MountPath:        "/var/lib/kube-proxy",
						SubPath:          "",
						MountPropagation: nil,
					},
					v1.VolumeMount{
						Name:             "xtables-lock",
						ReadOnly:         false,
						MountPath:        "/run/xtables.lock",
						SubPath:          "",
						MountPropagation: nil,
					},
					v1.VolumeMount{
						Name:             "lib-modules",
						ReadOnly:         true,
						MountPath:        "/lib/modules",
						SubPath:          "",
						MountPropagation: nil,
					},
					v1.VolumeMount{
						Name:             "kube-proxy-token-jdlsl",
						ReadOnly:         true,
						MountPath:        "/var/run/secrets/kubernetes.io/serviceaccount",
						SubPath:          "",
						MountPropagation: nil,
					},
				},
				TerminationMessagePath:   "/dev/termination-log",
				TerminationMessagePolicy: v1.TerminationMessageReadFile,
				ImagePullPolicy:          v1.PullIfNotPresent,
				SecurityContext: &v1.SecurityContext{
					Capabilities: nil,
					Privileged: func(b bool) *bool {
						return &b
					}(true),
				},
			},
			`(v1.Container){Name:(string)kube-proxy Image:(string)index-dev.qiniu.io/kelibrary/kube-proxy-amd64:v1.9.7 Command:([]string)[/usr/local/bin/kube-proxy --config=/var/lib/kube-proxy/config.conf] Args:([]string)<nil> WorkingDir:(string) Ports:([]v1.ContainerPort)<nil> EnvFrom:([]v1.EnvFromSource)<nil> Env:([]v1.EnvVar)<nil> Resources:(v1.ResourceRequirements){Limits:(v1.ResourceList)<nil> Requests:(v1.ResourceList)<nil>} VolumeMounts:([]v1.VolumeMount)[{Name:(string)kube-proxy ReadOnly:(bool)false MountPath:(string)/var/lib/kube-proxy SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)xtables-lock ReadOnly:(bool)false MountPath:(string)/run/xtables.lock SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)lib-modules ReadOnly:(bool)true MountPath:(string)/lib/modules SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)kube-proxy-token-jdlsl ReadOnly:(bool)true MountPath:(string)/var/run/secrets/kubernetes.io/serviceaccount SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>}] VolumeDevices:([]v1.VolumeDevice)<nil> LivenessProbe:(*v1.Probe)<nil> ReadinessProbe:(*v1.Probe)<nil> Lifecycle:(*v1.Lifecycle)<nil> TerminationMessagePath:(string)/dev/termination-log TerminationMessagePolicy:(v1.TerminationMessagePolicy)File ImagePullPolicy:(v1.PullPolicy)IfNotPresent SecurityContext:(*v1.SecurityContext){Capabilities:(*v1.Capabilities)<nil> Privileged:(*bool)true SELinuxOptions:(*v1.SELinuxOptions)<nil> RunAsUser:(*int64)<nil> RunAsNonRoot:(*bool)<nil> ReadOnlyRootFilesystem:(*bool)<nil> AllowPrivilegeEscalation:(*bool)<nil>} Stdin:(bool)false StdinOnce:(bool)false TTY:(bool)false}`,
			hackContainerGoStringTo19,
		},
		{
			"convert latest container go string to 1.8",
			v1.Container{
				Name:  "kube-proxy",
				Image: "index-dev.qiniu.io/kelibrary/kube-proxy-amd64:v1.9.7",
				Command: []string{
					"/usr/local/bin/kube-proxy",
					"--config=/var/lib/kube-proxy/config.conf",
				},
				Resources: v1.ResourceRequirements{
					Limits:   nil,
					Requests: nil,
				},
				VolumeMounts: []v1.VolumeMount{
					v1.VolumeMount{
						Name:             "kube-proxy",
						ReadOnly:         false,
						MountPath:        "/var/lib/kube-proxy",
						SubPath:          "",
						MountPropagation: nil,
					},
					v1.VolumeMount{
						Name:             "xtables-lock",
						ReadOnly:         false,
						MountPath:        "/run/xtables.lock",
						SubPath:          "",
						MountPropagation: nil,
					},
					v1.VolumeMount{
						Name:             "lib-modules",
						ReadOnly:         true,
						MountPath:        "/lib/modules",
						SubPath:          "",
						MountPropagation: nil,
					},
					v1.VolumeMount{
						Name:             "kube-proxy-token-jdlsl",
						ReadOnly:         true,
						MountPath:        "/var/run/secrets/kubernetes.io/serviceaccount",
						SubPath:          "",
						MountPropagation: nil,
					},
				},
				TerminationMessagePath:   "/dev/termination-log",
				TerminationMessagePolicy: v1.TerminationMessageReadFile,
				ImagePullPolicy:          v1.PullIfNotPresent,
				SecurityContext: &v1.SecurityContext{
					Capabilities: nil,
					Privileged: func(b bool) *bool {
						return &b
					}(true),
				},
			},
			`(v1.Container){Name:(string)kube-proxy Image:(string)index-dev.qiniu.io/kelibrary/kube-proxy-amd64:v1.9.7 Command:([]string)[/usr/local/bin/kube-proxy --config=/var/lib/kube-proxy/config.conf] Args:([]string)<nil> WorkingDir:(string) Ports:([]v1.ContainerPort)<nil> EnvFrom:([]v1.EnvFromSource)<nil> Env:([]v1.EnvVar)<nil> Resources:(v1.ResourceRequirements){Limits:(v1.ResourceList)<nil> Requests:(v1.ResourceList)<nil>} VolumeMounts:([]v1.VolumeMount)[{Name:(string)kube-proxy ReadOnly:(bool)false MountPath:(string)/var/lib/kube-proxy SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)xtables-lock ReadOnly:(bool)false MountPath:(string)/run/xtables.lock SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)lib-modules ReadOnly:(bool)true MountPath:(string)/lib/modules SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>} {Name:(string)kube-proxy-token-jdlsl ReadOnly:(bool)true MountPath:(string)/var/run/secrets/kubernetes.io/serviceaccount SubPath:(string) MountPropagation:(*v1.MountPropagationMode)<nil>}] LivenessProbe:(*v1.Probe)<nil> ReadinessProbe:(*v1.Probe)<nil> Lifecycle:(*v1.Lifecycle)<nil> TerminationMessagePath:(string)/dev/termination-log TerminationMessagePolicy:(v1.TerminationMessagePolicy)File ImagePullPolicy:(v1.PullPolicy)IfNotPresent SecurityContext:(*v1.SecurityContext){Capabilities:(*v1.Capabilities)<nil> Privileged:(*bool)true SELinuxOptions:(*v1.SELinuxOptions)<nil> RunAsUser:(*int64)<nil> RunAsNonRoot:(*bool)<nil> ReadOnlyRootFilesystem:(*bool)<nil> AllowPrivilegeEscalation:(*bool)<nil>} Stdin:(bool)false StdinOnce:(bool)false TTY:(bool)false}`,
			hackContainerGoStringTo18,
		},
		{
			"convert latest container go string to 1.7",
			v1.Container{
				Name:  "kube-proxy",
				Image: "index-dev.qiniu.io/kelibrary/kube-proxy-amd64:v1.9.7",
				Command: []string{
					"/usr/local/bin/kube-proxy",
					"--config=/var/lib/kube-proxy/config.conf",
				},
				Resources: v1.ResourceRequirements{
					Limits:   nil,
					Requests: nil,
				},
				VolumeMounts: []v1.VolumeMount{
					v1.VolumeMount{
						Name:             "kube-proxy",
						ReadOnly:         false,
						MountPath:        "/var/lib/kube-proxy",
						SubPath:          "",
						MountPropagation: nil,
					},
					v1.VolumeMount{
						Name:             "xtables-lock",
						ReadOnly:         false,
						MountPath:        "/run/xtables.lock",
						SubPath:          "",
						MountPropagation: nil,
					},
					v1.VolumeMount{
						Name:             "lib-modules",
						ReadOnly:         true,
						MountPath:        "/lib/modules",
						SubPath:          "",
						MountPropagation: nil,
					},
					v1.VolumeMount{
						Name:             "kube-proxy-token-jdlsl",
						ReadOnly:         true,
						MountPath:        "/var/run/secrets/kubernetes.io/serviceaccount",
						SubPath:          "",
						MountPropagation: nil,
					},
				},
				TerminationMessagePath:   "/dev/termination-log",
				TerminationMessagePolicy: v1.TerminationMessageReadFile,
				ImagePullPolicy:          v1.PullIfNotPresent,
				SecurityContext: &v1.SecurityContext{
					Capabilities: nil,
					Privileged: func(b bool) *bool {
						return &b
					}(true),
				},
			},
			`(v1.Container){Name:(string)kube-proxy Image:(string)index-dev.qiniu.io/kelibrary/kube-proxy-amd64:v1.9.7 Command:([]string)[/usr/local/bin/kube-proxy --config=/var/lib/kube-proxy/config.conf] Args:([]string)<nil> WorkingDir:(string) Ports:([]v1.ContainerPort)<nil> EnvFrom:([]v1.EnvFromSource)<nil> Env:([]v1.EnvVar)<nil> Resources:(v1.ResourceRequirements){Limits:(v1.ResourceList)<nil> Requests:(v1.ResourceList)<nil>} VolumeMounts:([]v1.VolumeMount)[{Name:(string)kube-proxy ReadOnly:(bool)false MountPath:(string)/var/lib/kube-proxy SubPath:(string)} {Name:(string)xtables-lock ReadOnly:(bool)false MountPath:(string)/run/xtables.lock SubPath:(string)} {Name:(string)lib-modules ReadOnly:(bool)true MountPath:(string)/lib/modules SubPath:(string)} {Name:(string)kube-proxy-token-jdlsl ReadOnly:(bool)true MountPath:(string)/var/run/secrets/kubernetes.io/serviceaccount SubPath:(string)}] LivenessProbe:(*v1.Probe)<nil> ReadinessProbe:(*v1.Probe)<nil> Lifecycle:(*v1.Lifecycle)<nil> TerminationMessagePath:(string)/dev/termination-log TerminationMessagePolicy:(v1.TerminationMessagePolicy)File ImagePullPolicy:(v1.PullPolicy)IfNotPresent SecurityContext:(*v1.SecurityContext){Capabilities:(*v1.Capabilities)<nil> Privileged:(*bool)true SELinuxOptions:(*v1.SELinuxOptions)<nil> RunAsUser:(*int64)<nil> RunAsNonRoot:(*bool)<nil> ReadOnlyRootFilesystem:(*bool)<nil>} Stdin:(bool)false StdinOnce:(bool)false TTY:(bool)false}`,
			hackContainerGoStringTo17,
		},
	}
	for _, tc := range testcases {
		goString := containerToGoString(tc.container)
		actual := tc.hackFunc(goString)
		if actual != tc.expected {
			t.Errorf("test %q failed, diff:\n%s\n", tc.name, diff.StringDiff(tc.expected, actual))
		}
	}
}
