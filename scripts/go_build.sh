#!/bin/bash -e
watch="servicces core"
. shouldIBuild.sh
shouldIBuild
if [[ $SHOULD_BUILD = 0 ]]; then
    exit 0
fi
make linux-binaries-api
BUILD_EXIT_STATUS=$?
exit $BUILD_EXIT_STATUS