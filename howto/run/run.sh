#!/bin/bash

IFS=$'\n'
scriptdir="$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)"
basedir="$(echo "${scriptdir}" | grep -Po ".*(?=\/)")"

bindir="${basedir}/bin"
bkpdir="${basedir}/bkp"
moddir="${basedir}/run/mod"
logdir="${basedir}/log"
logfile="${logdir}/bkpbot.log"

bot_config="${basedir}/bot_config.yaml"
run_config="${basedir}/run_config.toml"

mods_to_run_before=($(stoml "${run_config}" "mods_to_run_before"))
mods_to_run_after=($(stoml "${run_config}" "mods_to_run_after"))

# exported variables used to mods
# export MY_USER=$(stoml "${run_config}" "my_user")
# export MY_HOME=$(stoml "${run_config}" "my_home")
# export RELEVANT_DOCKERS=($(stoml "${run_config}" "relevant_dockers"))

subfolder="${1}"
keep_last="${2}"

if [[ -z "${subfolder}" ]]; then
    subfolder="daily"
fi

if [[ -z "${keep_last}" ]]; then
    keep_last="0"
fi

export DEBUG=""
for val in "$@"; do
    if [[ "${val}" =~ debug ]]; then
        export DEBUG="--debug"
    fi
done

if ( [[ -n "${debug}" ]] && (( ${#@} < 3 )) ); then
    echo "Debug must be third arg. Provide 'subfolder' and 'keep last'."
    exit 1
fi


# functions
function ifmodrun(){
    mods=($(echo "${1}" | tr " " "\n"))
    for mod in "${mods[@]}"; do
        modfile="${moddir}/${mod}.sh"
        if [[ -f "${modfile}" ]]; then
            echo "Run ${modfile}"
            eval "${modfile}"
        fi
    done
}


# main
echo -e "\nLaunch mods before"
ifmodrun "${mods_to_run_before}"

echo ""
"${bindir}/bkpbot" "${bot_config}" \
    -l "${logfile}" \
    -s "${subfolder}" \
    -k ${keep_last} ${DEBUG}

echo -e "\nLaunch mods after"
ifmodrun ${mods_to_run_after}
