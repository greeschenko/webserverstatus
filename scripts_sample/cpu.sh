#!/usr/bin/env bash

MAIN=''
CORES=`grep -c ^processor /proc/cpuinfo`

if [ $# -eq 0 ]; then
    MAIN=`cat /proc/loadavg`
else
    if [[ $1 == status ]]; then
        MAIN=`cat /proc/loadavg | awk '{print($1)}'`
        MAIN=`awk '{print $1*100}' <<<"$MAIN"`
        if [[ $MAIN -gt $(($CORES * 400)) ]]; then
            MAIN="FAIL"
        elif [[ $MAIN -gt $(($CORES * 200)) ]]; then
            MAIN="WARN"
        else
            MAIN="OK"
        fi
    elif [[ $1 == graph ]]; then
        MAIN=`cat /proc/loadavg | awk '{print($1)}'`
    elif [[ $1 == graphmax ]]; then
        MAIN=$(($CORES*4))
    fi
fi

printf "%s" "${MAIN}"
