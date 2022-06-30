#!/bin/bash -e

if [[ -f VERSION ]]; then
    cat VERSION > NEW_VERSION
elif [[ -f package.json ]]; then
    echo $(grep '"version":' package.json | cut -d\" -f4) > VERSION
else
    echo "Could not find a suitable version file"
    exit 1
fi

git checkout master
if [[ -f VERSION ]]; then
    cat VERSION > OLD_VERSION
elif [[ -f package.json ]]; then
    echo $(grep '"version":' package.json | cut -d\" -f4) > OLD_VERSION
else
    echo "Could not find a suitable version file"
    exit 1
fi
git checkout -

new_version=$(cat ./NEW_VERSION)
old_version=$(cat ./OLD_VERSION)

echo "New version: $new_version"
echo "Old version: $old_version"

if [[ "$new_version" == "NONE" && "$old_version" == "NONE" ]]; then
   echo "No version check required"
   exit 0
fi

 if [[ "$new_version" == "$old_version" ]]; then
    echo "Version has not changed, version must change if the service has been modified"
    exit 1
fi

major_current=$(echo $old_version | cut -d. -f1)
minor_current=$(echo $old_version | cut -d. -f2)
patch_current=$(echo $old_version | cut -d. -f3)

valid_major="$(($major_current + 1)).0.0"
valid_minor="$major_current.$(($minor_current + 1)).0"
valid_patch="$major_current.$minor_current.$(($patch_current + 1))"

if [[ "$new_version" != "$valid_major" && "$new_version" != "$valid_minor" && "$new_version" != "$valid_patch" ]]; then
     echo "Invalid version bump"
     exit 1
fi
