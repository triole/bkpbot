#!/bin/bash

set -e

tmpdir="/tmp/bkpbot_test"
scriptdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
basedir=$(echo "${scriptdir}" | grep -Po ".*(?=\/)")
sourcedir="${basedir}/src"

function ff(){
    find "${1}" \
        -mindepth 1 -maxdepth 1 -type ${2} \
        | grep -Po "${3}" \
        | sort
}

gofiles=$(ff "${sourcedir}" f ".*\.go$" | grep -v "_test.go")
configfiles=($(ff "${scriptdir}" f ".*\.yaml$"))

cmd="go run ${gofiles} -k 3"

# main
rm -rf "${tmpdir}/*"
c=0
cd "${sourcedir}"
for cf in "${configfiles[@]}"; do
    sleep 1
    c=$((c+1))

    echo -e "\n\033[0;36mRun ${cf}\033[0m"
    eval ${cmd} "${cf}"

    # generate differently looking folder to test filter
    mkdir -p "${tmpdir}/z_$RANDOM"

    # assert if there are n backup folders
    arr=($(ff "${tmpdir}" d ".*\/[0-9]{8}_[0-9]{6}$"))
    if (( "${#arr[@]}" != 3 && ${c} > 3)); then
        echo -e "\033[0;91mAssertion failed:\033[0m Number of kept folders is wrong. Check: ${tmpdir}"
        exit 1
    fi
done
