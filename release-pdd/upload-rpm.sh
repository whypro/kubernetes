#!/bin/bash
set -o nounset
set -o pipefail

source "$(cd "$(dirname "$0")"; pwd)/init.sh"

ARCHS=(
  amd64/x86_64
)

upload() {
  sha256sum $1 > $1.txt

  curl -X PUT -i -T $1\
    http://${CI_POS_ENDPOINT}/${CI_KUBE_PREFIX_RPMS}/$1 \
    -H "Host: ${CI_POS_HOST}" \
    -H "X-Pos-Tag: ${CI_POS_TAG}" \
    -H "sod-appid: ${CI_POS_SOD_APPID}" \
    -H "sod-appkey: ${CI_POS_SOD_APPKEY}"

  if [[ $? -ne 0 ]]; then
    echo "Failed to upload $1."
    exit 1
  fi

  curl -X PUT -i -T $1.txt\
    http://${CI_POS_ENDPOINT}/${CI_KUBE_PREFIX_RPMS}/$1.txt \
    -H "Host: ${CI_POS_HOST}" \
    -H "X-Pos-Tag: ${CI_POS_TAG}" \
    -H "sod-appid: ${CI_POS_SOD_APPID}" \
    -H "sod-appkey: ${CI_POS_SOD_APPKEY}"

  if [[ $? -ne 0 ]]; then
    echo "Failed to upload $1.txt."
    exit 1
  fi
}

cd ${ARTIFACT_RPM}

for ARCH in ${ARCHS[@]}; do
  IFS=/ read GOARCH RPMARCH<<< ${ARCH}; unset IFS;
  cd ${RPMARCH}
  FILES=$(ls *.rpm)
  for FILE in ${FILES[@]}; do
    upload $FILE
  done
  cd ..
done
