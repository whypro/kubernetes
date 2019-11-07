#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

source "$(cd "$(dirname "$0")"; pwd)/init.sh"

BASEIMAGE=harbor-htj.srv.yiran.com/k8s/debian-base-amd64:v1.0.0 
BASEIMAGE_IPTABLES=harbor-htj.srv.yiran.com/k8s/debian-iptables-amd64:v11.0.2
DOCKERFILE_DIR=${KUBE_ROOT}/dockerfiles
DOCKERFILE=${KUBE_ROOT}/dockerfiles/Dockerfile
BINARY=$1

if [[ -d "${DOCKERFILE_DIR}" ]]; then
  echo "Remove existing dockerfiles dir."
  rm -fr ${DOCKERFILE_DIR}
fi

echo "Create dockerfiles dir."
mkdir -p ${DOCKERFILE_DIR}

if [[ ! -f "${ARTIFACT_BIN}/${BINARY}" ]]; then
  echo "Binary ${BINARY} does not exist in ${ARTIFACT_BIN}"
  exit 1
else
  cp ${ARTIFACT_BIN}/${BINARY} ${DOCKERFILE_DIR}
fi

if [[ ${BINARY} =~ "kube-proxy" ]]; then
  cat <<EOF > "${DOCKERFILE}"
FROM ${BASEIMAGE_IPTABLES}
COPY ${BINARY} /usr/local/bin/${BINARY}
EOF
else
  cat <<EOF > "${DOCKERFILE}"
FROM ${BASEIMAGE}
COPY ${BINARY} /usr/local/bin/${BINARY}
EOF
fi
