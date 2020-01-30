#!/bin/bash

scriptdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
basedir="$(echo "${scriptdir}" | grep -Po ".*(?=\/)" | grep -Po ".*(?=\/)")"
config="${basedir}/run_config.toml"

relevant_dockers=($(stoml "${config}" "relevant_dockers"))

echo "Stop dockers"
for c in "${relevant_dockers[@]}"; do
    ids=($(docker ps -aq --filter "name=${c}"))
    if (( ${#ids} > 0 )); then
        for id in "${ids[@]}"; do
            cmd="docker stop ${id}"
            if [[ -z "${DEBUG}" ]]; then
                eval "${cmd}"
            fi
        done
    else
        echo "No container found for ${c}"
    fi
done
