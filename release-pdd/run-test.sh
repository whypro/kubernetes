#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

source "$(cd "$(dirname "$0")"; pwd)/init.sh"
PACKAGES_FILE=${KUBE_ROOT}/release-pdd/packages.txt

cd ${KUBE_ROOT}
echo "Running hack/update-codegen.sh to generate generated files..."
hack/update-codegen.sh > /dev/null 2>&1 || true
echo "Running go list to listing all pkgs..."
go list ./... > ${PACKAGES_FILE} 2> /dev/null || true
PACKAGES=$(grep -F -v -x -f ${KUBE_ROOT}/release-pdd/exclude.txt ${PACKAGES_FILE})
echo "About to run unit tests."
for PACKAGE in "${PACKAGES[@]}"
do
  go test ${PACKAGE} -timeout 600s
done
