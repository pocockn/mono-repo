#!/bin/bash -e
SHOULD_BUILD=0
shouldIBuild() {
    if [[ "${DRONE_DEPLOY_TO}" ]]; then
        SHOULD_BUILD=1
    else
        . detectChangedFolders.sh
        detect_changed_folders
        toW=($(echo "$watch" | tr ' ' '\n'))
        changed=($(echo "$changed_components"))
        for i in "${toW[@]}"
        do
            for j in "${changed[@]}"
            do
                if [[ $i = $j ]]; then
                SHOULD_BUILD=1
                fi
            done
        done
    fi
}