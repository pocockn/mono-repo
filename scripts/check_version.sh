#!/bin/bash -e

CHANGED=$(git diff master VERSION)
if [ ! "$CHANGED" ];
    then echo "VERSION file must be changed between PRs for modified services"
fi
