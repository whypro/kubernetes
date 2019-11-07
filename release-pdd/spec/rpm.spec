%global KUBE_VERSION {{KUBE_VERSION}}
%global RPM_RELEASE 0
%global ARCH amd64

%global CNI_VERSION {{CNI_VERSION}}
%global CRI_TOOLS_VERSION {{CRI_TOOLS_VERSION}}

Name: kubelet
Version: %{KUBE_VERSION}
Release: %{RPM_RELEASE}
Summary: Container cluster management
License: ASL 2.0

URL: https://kubernetes.io
Source0: kubelet
Source1: kubelet.service
Source2: kubectl
Source3: kubeadm
Source4: 10-kubeadm.conf
Source5: cni-plugins-%{ARCH}-%{CNI_VERSION}.tgz
Source6: kubelet.env
Source7: crictl-%{CRI_TOOLS_VERSION}-linux-%{ARCH}.tar.gz

BuildRequires: systemd
BuildRequires: curl
Requires: iptables >= 1.4.21
Requires: kubernetes-cni >= %{CNI_VERSION}
Requires: socat
Requires: util-linux
Requires: ethtool
Requires: iproute
Requires: ebtables
Requires: conntrack


%description
The node agent of Kubernetes, the container cluster manager.

%package -n kubernetes-cni

Version: %{CNI_VERSION}
Release: %{RPM_RELEASE}
Summary: Binaries required to provision kubernetes container networking
Requires: kubelet

%description -n kubernetes-cni
Binaries required to provision container networking.

%package -n kubectl

Version: %{KUBE_VERSION}
Release: %{RPM_RELEASE}
Summary: Command-line utility for interacting with a Kubernetes cluster.

%description -n kubectl
Command-line utility for interacting with a Kubernetes cluster.

%package -n kubeadm

Version: %{KUBE_VERSION}
Release: %{RPM_RELEASE}
Summary: Command-line utility for administering a Kubernetes cluster.
Requires: kubelet >= 1.13.0
Requires: kubectl >= 1.13.0
Requires: kubernetes-cni >= 0.7.5
Requires: cri-tools >= 1.13.0

%description -n kubeadm
Command-line utility for administering a Kubernetes cluster.

%package -n cri-tools

Version: %{CRI_TOOLS_VERSION}
Release: %{RPM_RELEASE}
Summary: Command-line utility for interacting with a container runtime.

%description -n cri-tools
Command-line utility for interacting with a container runtime.

%prep

cp -p %SOURCE0 %{_builddir}/
cp -p %SOURCE1 %{_builddir}/
cp -p %SOURCE2 %{_builddir}/
cp -p %SOURCE3 %{_builddir}/
cp -p %SOURCE4 %{_builddir}/
cp -p %SOURCE6 %{_builddir}/
%setup -c -D -T -a 5 -n cni-plugins
%setup -c -a 7 -T -n cri-tools

%install

cd %{_builddir}
install -m 755 -d %{buildroot}%{_unitdir}
install -m 755 -d %{buildroot}%{_unitdir}/kubelet.service.d/
install -m 755 -d %{buildroot}%{_bindir}
install -m 755 -d %{buildroot}%{_sysconfdir}/cni/net.d/
install -m 755 -d %{buildroot}%{_sysconfdir}/kubernetes/manifests/
install -m 755 -d %{buildroot}/var/lib/kubelet/
install -p -m 755 -t %{buildroot}%{_bindir}/ kubelet
install -p -m 755 -t %{buildroot}%{_bindir}/ kubectl
install -p -m 755 -t %{buildroot}%{_bindir}/ kubeadm
install -p -m 644 -t %{buildroot}%{_unitdir}/ kubelet.service
install -p -m 644 -t %{buildroot}%{_unitdir}/kubelet.service.d/ 10-kubeadm.conf
install -p -m 755 -t %{buildroot}%{_bindir}/ cri-tools/crictl

install -m 755 -d %{buildroot}%{_sysconfdir}/sysconfig/
install -p -m 644 -T kubelet.env %{buildroot}%{_sysconfdir}/sysconfig/kubelet

install -m 755 -d %{buildroot}/opt/cni/bin
mv cni-plugins/* %{buildroot}/opt/cni/bin/

%files
%{_bindir}/kubelet
%{_unitdir}/kubelet.service
%{_sysconfdir}/kubernetes/manifests/

%config(noreplace) %{_sysconfdir}/sysconfig/kubelet

%files -n kubernetes-cni
/opt/cni

%files -n kubectl
%{_bindir}/kubectl

%files -n kubeadm
%{_bindir}/kubeadm
%{_unitdir}/kubelet.service.d/10-kubeadm.conf

%files -n cri-tools
%{_bindir}/crictl
