#!/bin/bash

scriptdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
basedir="$(echo "${scriptdir}" | grep -Po ".*(?=\/)" | grep -Po ".*(?=\/)")"
config="${basedir}/run_config.toml"

my_user=($(stoml "${config}" "my_user"))
my_home=($(stoml "${config}" "my_home"))

cmd="chown -R ${my_user}.${my_user} \"${my_home}/.backup\""
echo "${cmd}"
if [[ -z "${DEBUG}" ]]; then
    eval "${cmd}"
else
    echo "Debug mode. Did not fix."
fi
