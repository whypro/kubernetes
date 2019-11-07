#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

source "$(cd "$(dirname "$0")"; pwd)/init.sh"

BINARY_DIR=${KUBE_ROOT}/_output/bin
GIT_VERSION_FILE=${KUBE_ROOT}/gitversion
TARGETS+=(
  "kubelet"
  "kube-apiserver"
  "kube-controller-manager"
  "kube-scheduler"
  "kube-proxy"
  "kubeadm"
  "kubectl"
)

if [[ -n ${CI_COMMIT_REF_NAME} ]]; then
  echo "Version: ${CI_COMMIT_REF_NAME}"
  
  KUBE_GIT_VERSION=${CI_COMMIT_REF_NAME}
  KUBE_GIT_COMMIT=${CI_COMMIT_SHA-}
  KUBE_GIT_TREE_STATE="clean"

  if [[ "${KUBE_GIT_VERSION}" =~ ^v([0-9]+)\.([0-9]+)(\.[0-9]+)?([-].*)?([+].*)?$ ]]; then
    KUBE_GIT_MAJOR=${BASH_REMATCH[1]}
    KUBE_GIT_MINOR=${BASH_REMATCH[2]}
    if [[ -n "${BASH_REMATCH[4]}" ]]; then
      KUBE_GIT_MINOR+="+"
    fi
  fi

    cat <<EOF >"${GIT_VERSION_FILE}"
KUBE_GIT_COMMIT='${KUBE_GIT_COMMIT-}'
KUBE_GIT_TREE_STATE='${KUBE_GIT_TREE_STATE-}'
KUBE_GIT_VERSION='${KUBE_GIT_VERSION-}'
KUBE_GIT_MAJOR='${KUBE_GIT_MAJOR-}'
KUBE_GIT_MINOR='${KUBE_GIT_MINOR-}'
EOF
  export KUBE_GIT_VERSION_FILE=${GIT_VERSION_FILE}
fi

cd ${KUBE_ROOT}

if [[ ! -d "${ARTIFACT_BIN}" ]]; then
  mkdir ${ARTIFACT_BIN}
fi

for TARGET in "${TARGETS[@]}"
do
  make WHAT=cmd/${TARGET}
  mv ${BINARY_DIR}/${TARGET} ${ARTIFACT_BIN}/${TARGET}
done
