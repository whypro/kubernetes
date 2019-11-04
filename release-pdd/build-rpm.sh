#!/bin/bash
set -o nounset
set -o pipefail

if [[ -z ${CI_COMMIT_REF_NAME} ]]; then
  echo "No CI_COMMIT_REF_NAME is provided."
  exit 1
fi

source "$(cd "$(dirname "$0")"; pwd)/init.sh"

SPECS=${KUBE_ROOT}/release-pdd/spec
CRI_TOOLS_VERSION=v1.14.0
CNI_VERSION=v0.7.5
KUBE_VERSION=$(echo ${CI_COMMIT_REF_NAME} | cut -d ' ' -f 2 | tr -s '-' '_' | sed 's/^v//')

BINS+=(
  "kubelet"
  "kubeadm"
  "kubectl"
)

ARCHS=(
  amd64/x86_64
)

download_and_check() {
  curl -X GET \
    http://${CI_POS_ENDPOINT}/${CI_POS_TAG}/${CI_KUBE_PREFIX_TOOLS}/$1 \
      -H "Host: ${CI_POS_HOST}" \
      -H "sod-appid: ${CI_POS_SOD_APPID}" \
      -H "sod-appkey: ${CI_POS_SOD_APPKEY}" -O

  if [[ $? -ne 0 ]]; then
    echo "Cannot download $1."
    exit 1
  fi

  curl -X GET \
    http://${CI_POS_ENDPOINT}/${CI_POS_TAG}/${CI_KUBE_PREFIX_TOOLS}/$1.txt \
      -H "Host: ${CI_POS_HOST}" \
      -H "sod-appid: ${CI_POS_SOD_APPID}" \
      -H "sod-appkey: ${CI_POS_SOD_APPKEY}" -O
  
  if [[ $? -ne 0 ]]; then
    echo "Cannot download $1.txt."
    exit 1
  fi

  cat $1.txt | sha256sum --check --status
  if [[ $? -ne 0 ]]; then
    echo "$1 SHA256 is invalid."
    exit 1
  fi
}

echo "KUBE_VERSION: ${KUBE_VERSION}"

sed -i "s/{{KUBE_VERSION}}/${KUBE_VERSION}/g" ${SPECS}/rpm.spec
sed -i "s/{{CRI_TOOLS_VERSION}}/${CRI_TOOLS_VERSION}/g" ${SPECS}/rpm.spec
sed -i "s/{{CNI_VERSION}}/${CNI_VERSION}/g" ${SPECS}/rpm.spec

for BIN in ${BINS[@]}; do
  cp ${ARTIFACT_BIN}/${BIN} ${SPECS}
done

for ARCH in ${ARCHS[@]}; do
  IFS=/ read GOARCH RPMARCH<<< ${ARCH}; unset IFS;
  SRC_PATH="/root/rpmbuild/SOURCES/${RPMARCH}"
  mkdir -p ${SRC_PATH}
  cp -r ${SPECS}/* ${SRC_PATH}
  echo "Building RPM's for ${GOARCH}....."
  sed -i "s/\%global ARCH.*/\%global ARCH ${GOARCH}/" ${SRC_PATH}/rpm.spec
  cd ${SRC_PATH}
  download_and_check crictl-${CRI_TOOLS_VERSION}-linux-${GOARCH}.tar.gz
  download_and_check cni-plugins-${GOARCH}-${CNI_VERSION}.tgz
  /usr/bin/rpmbuild --target ${RPMARCH} --define "_sourcedir ${SRC_PATH}" -bb ${SRC_PATH}/rpm.spec
  mkdir -p ${ARTIFACT_RPM}/${RPMARCH}
  createrepo -o ${ARTIFACT_RPM}/${RPMARCH}/ ${ARTIFACT_RPM}/${RPMARCH}
  mv /root/rpmbuild/RPMS/${RPMARCH}/* ${ARTIFACT_RPM}/${RPMARCH}
  rm -fr ${ARTIFACT_RPM}/${RPMARCH}/repodata
done
