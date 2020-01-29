#!/bin/bash

scriptdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

subfolder="${1}"
keep_last="${2}"

if [[ -z "${subfolder}" ]]; then
    subfolder="daily"
fi

if [[ -z "${keep_last}" ]]; then
    keep_last="0"
fi


"${scriptdir}/bkpbot" \
    "${scriptdir}/config.yaml" \
    -l "${scriptdir}/bkpbot.log" \
    -s "${subfolder}" \
    -k ${keep_last}
