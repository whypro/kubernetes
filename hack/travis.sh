#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# build linux/amd64 only
export KUBE_FASTBUILD=true

LOGFILE=/tmp/make.$TARGET.log

if [[ "$TARGET" == "test-integration" ]]; then
    ./hack/install-etcd.sh
fi

dump_output() {
    tail -500 $LOGFILE
}

trap 'dump_output' EXIT

# We need to redirect logs to files to work around traivs 4MB log limit, see
# https://docs.travis-ci.com/user/common-build-problems/#Log-Length-exceeded.
./build/run.sh make $TARGET &> $LOGFILE
