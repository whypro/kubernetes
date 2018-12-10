#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

LOGFILE=/tmp/make.$TARGET.log

rm -f $LOGFILE

# We need to redirect logs to files to work around traivs 4MB log limit, see
# https://docs.travis-ci.com/user/common-build-problems/#Log-Length-exceeded.
dump_output() {
    test -f $LOGFILE && tail -500 $LOGFILE
}

trap 'dump_output' EXIT

if [[ "$TARGET" == "test-integration" ]]; then
    ./hack/install-etcd.sh
fi

args=(
    $TARGET
    # CPU resources in travis is very low, slow down to make sure tests will
    # not timeout.
    # See https://docs.travis-ci.com/user/reference/overview/.
    PARALLEL=1
)

# pass all KUBE_ environments
while IFS='=' read -r -d '' n v; do
    if [[ $n =~ ^KUBE_ ]]; then
        args+=($n="$v")
    fi
done < <(env -0)

./build/run.sh make "${args[@]}" &> $LOGFILE
